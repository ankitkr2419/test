package service

import (
	"context"
	"encoding/json"
	"fmt"
	"mylab/cpagent/db"
	"mylab/cpagent/plc"
	"mylab/cpagent/responses"
	"net/http"
	"reflect"
	"time"

	"github.com/gorilla/mux"
	logger "github.com/sirupsen/logrus"
)

func homingHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		var err error
		var msg string
		vars := mux.Vars(req)
		deck := vars["deck"]

		ctx := context.WithValue(req.Context(), contextKeyUsername, "main")

		switch deck {
		case "":
			fmt.Println("At both deck!!!")
			msg = "homing in progress for both decks"
			plc.SetBothDeckHomingInProgress()
			go bothDeckOperation(ctx, deps, "Homing")
		case "A", "B":
			msg = "homing in progress for single deck"
			go singleDeckOperation(ctx, deps, deck, "Homing")
		default:
			err = fmt.Errorf("Check your deck name")
		}

		if err != nil {
			rw.Write([]byte(err.Error()))
			rw.WriteHeader(http.StatusBadRequest)
			logger.Errorln(err)
		} else {
			rw.Header().Add("Content-Type", "application/json")
			rw.WriteHeader(http.StatusOK)
			rw.Write([]byte(fmt.Sprintf(`{"msg":"%v","deck":"%v"}`, msg, deck)))
			logger.Infoln(msg)
		}

	})
}

func bothDeckOperation(ctx context.Context, deps Dependencies, operation string) (response string, err error) {
	defer plc.ResetBothDeckHomingInProgress()

	var deckAResponse, deckBResponse string
	var deckAErr, deckBErr error

	go func() {
		deckAResponse, deckAErr = singleDeckOperation(ctx, deps, "A", operation)
	}()
	go func() {
		deckBResponse, deckBErr = singleDeckOperation(ctx, deps, "B", operation)
	}()

	for {
		switch {
		case deckAErr != nil:
			fmt.Printf("Error %s deck A", operation)
			return "", deckAErr
		case deckBErr != nil:
			fmt.Printf("Error %s deck B", operation)
			return "", deckBErr
		case deckAResponse != "" && deckBResponse != "":
			operationSuccessMsg := fmt.Sprintf("%s Success for both Decks!", operation)
			successWsData := plc.WSData{
				Progress: 100,
				Deck:     "",
				Status:   "SUCCESS_HOMING",
				OperationDetails: plc.OperationDetails{
					Message: operationSuccessMsg,
				},
			}
			wsData, err := json.Marshal(successWsData)
			if err != nil {
				logger.Errorf("error in marshalling web socket data %v", err.Error())
				deps.WsErrCh <- fmt.Errorf("%v_%v_%v", plc.ErrorExtractionMonitor, "", err.Error())
				return "", err
			}
			deps.WsMsgCh <- fmt.Sprintf("success_homing_%v", string(wsData))
			fmt.Println(operationSuccessMsg)
			return operationSuccessMsg, nil
		default:
			// Only check every 400 milli second
			time.Sleep(400 * time.Millisecond)
		}
	}

}

func singleDeckOperation(ctx context.Context, deps Dependencies, deck, operation string) (response string, err error) {

	go deps.Store.AddAuditLog(ctx, db.MachineOperation, db.InitialisedState, db.ExecuteOperation, deck, responses.GetMachineOperationMessage(operation, string(db.InitialisedState)))

	defer func() {
		if err != nil {
			logger.Errorln(err.Error())
			deps.WsErrCh <- fmt.Errorf("%v_%v_%v", plc.ErrorExtractionMonitor, deck, err.Error())
			go deps.Store.AddAuditLog(ctx, db.MachineOperation, db.ErrorState, db.ExecuteOperation, deck, err.Error())
		} else {
			go deps.Store.AddAuditLog(ctx, db.MachineOperation, db.CompletedState, db.ExecuteOperation, deck, responses.GetMachineOperationMessage(operation, string(db.CompletedState)))
		}
	}()

	// Compact32Deck is the type of deps.PlcDeck[deck]
	result := reflect.ValueOf(deps.PlcDeck[deck]).MethodByName(operation).Call([]reflect.Value{})
	// TODO : variadic parameters ought to be handled as well
	//  this will make Call to any method generic
	// TODO : Handle Panics with recover()

	fmt.Println("Result from Operation ", operation, result)

	if len(result) != 2 {
		fmt.Println("result is different in this reflect Call !", result)
		err := fmt.Errorf("unexpected length result")

		return "", err
	}

	response = result[0].String()
	if len(result) == 2 {
		errRes := result[1].Interface()
		if errRes != nil {
			fmt.Println(errRes)
			err := fmt.Errorf("%v", errRes)
			return "", err
		}
		if operation == "Abort" && stepRunInProgress[deck] && !runNext[deck] {
			// handle the abortStepRun channel
			abortStepRun[deck] <- struct{}{}
		}
	}
	return
}
