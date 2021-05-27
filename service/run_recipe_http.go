package service

import (
	"context"
	"encoding/json"
	"fmt"
	"mylab/cpagent/db"
	"mylab/cpagent/plc"
	"mylab/cpagent/responses"
	"net/http"
	"time"

	"github.com/google/uuid"
	logger "github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
)

func runRecipeHandler(deps Dependencies, runStepWise bool) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.InitialisedState, db.ExecuteOperation, "", responses.RunRecipeInitialisedState)

		var err error

		// for logging error if there is any otherwise logging success
		defer func() {
			if err != nil {
				go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.ErrorState, db.ExecuteOperation, "", err.Error())

			} else {
				go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.CompletedState, db.ExecuteOperation, "", responses.DelayCompletedState)

			}

		}()

		vars := mux.Vars(req)
		deck := vars["deck"]

		recipeID, err := parseUUID(vars["id"])
		if err != nil {
			logger.Errorln(err)
			responseCodeAndMsg(rw, http.StatusBadRequest, ErrObj{Err: err.Error(), Deck: deck})
			return
		}

		go runRecipe(req.Context(), deps, deck, runStepWise, recipeID)
		logger.Infoln(responses.RecipeRunInProgress)
		responseCodeAndMsg(rw, http.StatusOK, MsgObj{Msg: responses.RecipeRunInProgress, Deck: deck})
		return

	})
}

func runNextStepHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		var err error

		vars := mux.Vars(req)
		deck := vars["deck"]

		// If runNext is set means this API is called at wrong time
		if runNext[deck] {
			err = responses.StepRunNotInProgressError
			logger.Errorln(err)
			responseCodeAndMsg(rw, http.StatusInternalServerError, ErrObj{Err: err.Error(), Deck: deck})
			return
		}

		populateNextStepChan(deck)
		logger.Infoln(responses.NextStepRunInProgress)
		responseCodeAndMsg(rw, http.StatusOK, MsgObj{Msg: responses.NextStepRunInProgress, Deck: deck})
		return

	})
}

func runRecipe(ctx context.Context, deps Dependencies, deck string, runStepWise bool, recipeID uuid.UUID) (response string, err error) {

	defer func() {
		if err != nil {
			logger.Errorln(err.Error())
			deps.WsErrCh <- fmt.Errorf("%v_%v_%v", plc.ErrorExtractionMonitor, deck, err.Error())
			go deps.Store.AddAuditLog(ctx, db.MachineOperation, db.ErrorState, db.ExecuteOperation, deck, err.Error())
		}
		resetStepRunInProgress(deck)
	}()

	if !deps.PlcDeck[deck].IsMachineHomed() {
		err = responses.PleaseHomeMachineError
		return
	}

	if deps.PlcDeck[deck].IsRunInProgress() {
		err = responses.PreviousRunInProgressError
		return
	}

	deps.PlcDeck[deck].SetRunInProgress()
	defer deps.PlcDeck[deck].ResetRunInProgress()

	// Get the recipe
	recipe, err := deps.Store.ShowRecipe(ctx, recipeID)
	if err != nil {
		return
	}

	// Get Processes associated with recipe
	processes, err := deps.Store.ListProcesses(ctx, recipe.ID)
	if err != nil {
		return
	}

	var currentCartridgeID int64
	// No cartridge selected so cartridge_id by default is 0
	// Depending on cartridge_1 or cartridge_2 type we shall
	//  select cartridge_id from recipe field

	var currentTip db.TipsTubes
	//  No tip selected
	//  This field will be set when a tip is picked up
	//  We will get its id from recipe and its details from tipsTubes

	if runStepWise {
		setStepRunInProgress(deck)
		populateNextStepChan(deck)
	}

	for i, p := range processes {

		// TODO : percentage calculation from inside the process.
		sendWSData(deps, deck, recipeID, len(processes), i+1, p.Name, p.Type)

		if runStepWise {

			logger.Infoln(responses.WaitingRunNextProcess)
			resetRunNext(deck)
			// To resume the next step admin needs to hits the run-next-step API only
			err = checkForAbortOrNext(deck)
			if err != nil {
				return
			}
			setRunNext(deck)
			logger.Infoln(responses.NextProcessInProgress)
		}
		go deps.Store.AddAuditLog(ctx, db.MachineOperation, db.InitialisedState, db.ExecuteOperation, deck, responses.GetMachineOperationMessage(string(p.Type), string(db.InitialisedState)))

		switch p.Type {
		case db.AspireDispenseProcess:
			ad, err := deps.Store.ShowAspireDispense(ctx, p.ID)
			if err != nil {
				return "", err
			}
			fmt.Println(ad)

			if ad.CartridgeType == db.Cartridge1 {
				currentCartridgeID = recipe.Cartridge1Position
			} else {
				currentCartridgeID = recipe.Cartridge2Position
			}
			// TODO: Pass the complete Tip rather than just name for volume validations
			response, err = deps.PlcDeck[deck].AspireDispense(ad, currentCartridgeID, currentTip.Name)
			if err != nil {
				return "", err
			}

		case db.HeatingProcess:
			heat, err := deps.Store.ShowHeating(ctx, p.ID)
			fmt.Printf("heat object %v", heat)
			ht, err := deps.PlcDeck[deck].Heating(heat)
			if err != nil {

				return "", err
			}
			fmt.Println(ht)

		case db.ShakingProcess:
			shaker, err := deps.Store.ShowShaking(ctx, p.ID)
			if err != nil {
				return "", err
			}
			fmt.Printf("shaker object %v", shaker)

			sha, err := deps.PlcDeck[deck].Shaking(shaker)
			if err != nil {
				return "", err
			}
			fmt.Println(sha)

		case db.PiercingProcess:
			pi, err := deps.Store.ShowPiercing(ctx, p.ID)
			if err != nil {
				return "", err
			}
			fmt.Println(pi)

			if pi.Type == db.Cartridge1 {
				currentCartridgeID = recipe.Cartridge1Position
			} else {
				currentCartridgeID = recipe.Cartridge2Position
			}

			response, err = deps.PlcDeck[deck].Piercing(pi, currentCartridgeID)
			if err != nil {
				return "", err
			}

		case db.AttachDetachProcess:
			ad, err := deps.Store.ShowAttachDetach(ctx, p.ID)
			fmt.Printf("attach detach record %v \n", ad)
			if err != nil {
				return "", err
			}
			response, err = deps.PlcDeck[deck].AttachDetach(ad)
			if err != nil {
				return "", err
			}

		case db.TipDiscardProcess, db.TipPickupProcess:
			to, err := deps.Store.ShowTipOperation(ctx, p.ID)
			if err != nil {
				return "", err
			}
			fmt.Println(to)

			response, err = deps.PlcDeck[deck].TipOperation(to)
			if err != nil {
				return "", err
			}

			switch to.Type {
			case db.PickupTip:
				// Store Current Tip here
				tipID, err := getTipIDFromRecipePosition(recipe, to.Position)
				if err != nil {
					return "", err
				}
				currentTip, err = deps.Store.ShowTip(tipID)
				if err != nil {
					return "", err
				}
			case db.DiscardTip:
				currentTip = db.TipsTubes{}

			}
		case db.TipDockingProcess:
			td, err := deps.Store.ShowTipDocking(ctx, p.ID)
			if err != nil {
				return "", err
			}
			fmt.Println(td)
			if td.Type == string(db.Cartridge1) {
				currentCartridgeID = recipe.Cartridge1Position
			} else {
				currentCartridgeID = recipe.Cartridge2Position
			}
			response, err = deps.PlcDeck[deck].TipDocking(td, currentCartridgeID)
			if err != nil {
				return "", err
			}
		case db.DelayProcess:
			delay, err := deps.Store.ShowDelay(ctx, p.ID)
			if err != nil {
				return "", err
			}
			fmt.Print(delay)
			response, err = deps.PlcDeck[deck].AddDelay(delay)
			if err != nil {
				return "", err
			}

		}
		go deps.Store.AddAuditLog(ctx, db.MachineOperation, db.CompletedState, db.ExecuteOperation, deck, responses.GetMachineOperationMessage(string(p.Type), string(db.CompletedState)))

	}

	plength := len(processes)
	sendWSData(deps, deck, recipeID, plength, plength+1, processes[plength-1].Name, processes[plength-1].Type)

	// send websocket success data
	successWsData := plc.WSData{
		Progress: 100,
		Deck:     deck,
		Status:   "SUCCESS_RECIPE",
		OperationDetails: plc.OperationDetails{
			Message:        fmt.Sprintf("successfully completed recipe %v for deck %v", recipeID, deck),
			CurrentStep:    len(processes),
			RecipeID:       recipeID,
			TotalProcesses: len(processes),
		},
	}
	wsData, err := json.Marshal(successWsData)
	if err != nil {
		logger.WithField("err", err.Error()).Errorln(responses.WebsocketMarshallingError)
		return
	}
	deps.WsMsgCh <- fmt.Sprintf("success_recipe_%v", string(wsData))

	// Home the machine
	deps.PlcDeck[deck].ResetRunInProgress()
	response, err = deps.PlcDeck[deck].Homing()
	if err != nil {
		return
	}

	return "SUCCESS", nil
}

func getTipIDFromRecipePosition(recipe db.Recipe, position int64) (id int64, err error) {
	// Currently only 3 positions are allowed for tips deck version 1.2
	// TODO: Change this for version 1.3
	switch position {
	case 1:
		return recipe.Position1, nil
	case 2:
		return recipe.Position2, nil
	case 3:
		return recipe.Position3, nil
	}
	err = responses.PickupPositionInvalid
	return 0, err
}

func populateNextStepChan(deck string) {
	logger.Infoln("Populating the nextStep channel for deck", deck)
	nextStep[deck] <- struct{}{}
}

func sendWSData(deps Dependencies, deck string, recipeID uuid.UUID, processLength, currentStep int, processName string, processType db.ProcessType) {

	// percentage calculation for each process

	progress := float64(((currentStep - 1) * 100) / processLength)

	wsProgressOperation := plc.WSData{
		Progress: progress,
		Deck:     deck,
		Status:   "PROGRESS_RECIPE",
		OperationDetails: plc.OperationDetails{
			Message:        fmt.Sprintf("process %v for deck %v in progress", currentStep, deck),
			CurrentStep:    currentStep,
			RecipeID:       recipeID,
			TotalProcesses: processLength,
			ProcessName:    processName,
			ProcessType:    processType,
		},
	}

	if processLength < currentStep {
		wsProgressOperation.OperationDetails.Message = fmt.Sprintf("process %v for deck %v completed", processLength, deck)
		wsProgressOperation.OperationDetails.CurrentStep = processLength
	}

	wsData, err := json.Marshal(wsProgressOperation)
	if err != nil {
		logger.WithField("err", err.Error()).Errorln(responses.WebsocketMarshallingError)
	}
	deps.WsMsgCh <- fmt.Sprintf("progress_recipe_%v", string(wsData))

	return
}

func checkForAbortOrNext(deck string) (err error) {
	for {
		time.Sleep(200 * time.Millisecond)
		select {
		case <-nextStep[deck]:
			logger.Infoln(responses.NextStepWillRun)
			return nil
		case <-abortStepRun[deck]:
			logger.Infoln(responses.StepRunWillAbort)
			return responses.StepRunAborted
		}
	}
}
