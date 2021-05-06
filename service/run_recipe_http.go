package service

import (
	"context"
	"encoding/json"
	"fmt"
	"mylab/cpagent/db"
	"mylab/cpagent/plc"
	"net/http"
	"time"

	"github.com/google/uuid"
	logger "github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
)

func runRecipeHandler(deps Dependencies, runStepWise bool) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		var err error

		vars := mux.Vars(req)
		deck := vars["deck"]

		recipeID, err := parseUUID(vars["id"])
		if err != nil {
			fmt.Fprintf(rw, err.Error())
			rw.WriteHeader(http.StatusBadRequest)
			return
		}

		switch deck {
		case "A", "B":
			go runRecipe(req.Context(), deps, deck, runStepWise, recipeID)
			rw.WriteHeader(http.StatusOK)
			rw.Header().Add("Content-Type", "application/json")
			rw.Write([]byte(fmt.Sprintf(`{"msg":"recipe run is in progress", "deck": "%v"}`, deck)))

		default:
			err = fmt.Errorf("Check your deck name")
		}

		if err != nil {
			rw.Header().Add("Content-Type", "application/json")
			rw.WriteHeader(http.StatusBadRequest)
			rw.Write([]byte(`{"msg":"check your deck name"}`))
			logger.Errorln(err.Error())
		}
	})
}

func runNextStepHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		var err error

		vars := mux.Vars(req)
		deck := vars["deck"]

		switch deck {
		case "A", "B":
			// If runNext is set means this API is called at wrong time
			if runNext[deck] {
				rw.WriteHeader(http.StatusBadRequest)
				rw.Header().Add("Content-Type", "application/json")
				rw.Write([]byte(fmt.Sprintf(`{"msg":"check if the step-run is in progress", "deck": "%v"}`, deck)))
				return
			}

			logger.Infoln("Populating the nextStep channel for deck", deck)
			nextStep[deck] <- struct{}{}

			rw.WriteHeader(http.StatusOK)
			rw.Header().Add("Content-Type", "application/json")
			rw.Write([]byte(fmt.Sprintf(`{"msg":"next step run is in progress", "deck":"%v"}`, deck)))
			return

		default:
			err = fmt.Errorf("Check your deck name")
		}

		if err != nil {
			rw.Header().Add("Content-Type", "application/json")
			rw.WriteHeader(http.StatusBadRequest)
			rw.Write([]byte(`{"msg":"check your deck name"}`))

			logger.Errorln(err.Error())
		}
		return
	})
}

func runRecipe(ctx context.Context, deps Dependencies, deck string, runStepWise bool, recipeID uuid.UUID) (response string, err error) {

	defer func() {
		if err != nil {
			logger.Errorln(err.Error())
			deps.WsErrCh <- fmt.Errorf("%v_%v_%v", plc.ErrorExtractionMonitor, deck, err.Error())
		}
		resetStepRunInProgress(deck)
	}()

	if !deps.PlcDeck[deck].IsMachineHomed() {
		err = fmt.Errorf("Please home the machine first!")
		return
	}

	if deps.PlcDeck[deck].IsRunInProgress() {
		err = fmt.Errorf("previous run already in progress... wait or abort it")

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

		return "", err
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
		logger.Infoln("Populating the nextStep channel for 1st process for deck", deck)
		nextStep[deck] <- struct{}{}
	}

	for i, p := range processes {

		// TODO : percentage calculation from inside the process.
		sendWSData(deps, deck, recipeID, len(processes), i+1)

		if runStepWise {

			logger.Infoln("Waiting to run next process")
			resetRunNext(deck)
			// To resume the next step admin needs to hits the run-next-step API only
			err = checkForAbortOrNext(deck)
			if err != nil {
				return
			}
			setRunNext(deck)
			logger.Infoln("Next process is in progress")
		}

		switch p.Type {
		case "AspireDispense":
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

		case "Heating":
			heat, err := deps.Store.ShowHeating(ctx, p.ID)
			fmt.Printf("heat object %v", heat)
			ht, err := deps.PlcDeck[deck].Heating(heat)
			if err != nil {

				return "", err
			}
			fmt.Println(ht)

		case "Shaking":
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

		case "Piercing":
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

		case "AttachDetach":
			ad, err := deps.Store.ShowAttachDetach(ctx, p.ID)
			fmt.Printf("attach detach record %v \n", ad)
			if err != nil {
				return "", err
			}
			response, err = deps.PlcDeck[deck].AttachDetach(ad)
			if err != nil {
				return "", err
			}

		case "TipOperation":
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
		case "TipDocking":
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
		case "Delay":
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
	}
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
		logger.Errorf("error in marshalling web socket data %v", err.Error())
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
	err = fmt.Errorf("position is invalid to pickup the tip")
	return 0, err
}

func sendWSData(deps Dependencies, deck string, recipeID uuid.UUID, processLength, currentStep int) (response string, err error) {
	// percentage calculation for each process

	progress := float64((currentStep * 100) / processLength)

	wsProgressOperation := plc.WSData{
		Progress: progress,
		Deck:     deck,
		Status:   "PROGRESS_RECIPE",
		OperationDetails: plc.OperationDetails{
			Message:        fmt.Sprintf("process %v for deck %v in progress", currentStep, deck),
			CurrentStep:    currentStep,
			RecipeID:       recipeID,
			TotalProcesses: processLength,
		},
	}

	wsData, err := json.Marshal(wsProgressOperation)
	if err != nil {
		logger.Errorf("error in marshalling web socket data %v", err.Error())
	}
	deps.WsMsgCh <- fmt.Sprintf("progress_recipe_%v", string(wsData))

	return
}

func checkForAbortOrNext(deck string) (err error) {
	for {
		time.Sleep(200 * time.Millisecond)
		select {
		case <-nextStep[deck]:
			logger.Infoln("Next Step will be Run")
			return nil
		case <-abortStepRun[deck]:
			logger.Infoln("Step Run will be Aborted")
			return fmt.Errorf("step run aborted")
		}
	}
}
