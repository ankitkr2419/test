package service

import (
	"encoding/json"
	"fmt"
	"mylab/cpagent/plc"
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

		switch deck {
		case "":
			fmt.Println("At both deck!!!")
			msg = "homing in progress for both decks"
			plc.SetBothDeckHomingInProgress()
			go bothDeckOperation(deps, "Homing")
		case "A", "B":
			msg = "homing in progress for single deck"
			go singleDeckOperation(deps, deck, "Homing")
		default:
			err = fmt.Errorf("Check your deck name")
		}

		if err != nil {
			rw.Write([]byte(err.Error()))
			rw.WriteHeader(http.StatusBadRequest)
			logger.Errorln(err)
			deps.WsErrCh <- err
		} else {
			rw.Header().Add("Content-Type", "application/json")
			rw.WriteHeader(http.StatusOK)
			rw.Write([]byte(fmt.Sprintf(`{"msg":"%v","deck":"%v"}`, msg, deck)))
			logger.Infoln(msg)
		}

	})
}

func bothDeckOperation(deps Dependencies, operation string) (response string, err error) {
	defer plc.ResetBothDeckHomingInProgress()

	var deckAResponse, deckBResponse string
	var deckAErr, deckBErr error

	go func() {
		deckAResponse, deckAErr = singleDeckOperation(deps, "A", operation)
	}()
	go func() {
		deckBResponse, deckBErr = singleDeckOperation(deps, "B", operation)
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
				deps.WsErrCh <- err
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

func singleDeckOperation(deps Dependencies, deck, operation string) (response string, err error) {

	// Compact32Deck is the type of deps.PlcDeck[deck]
	result := reflect.ValueOf(deps.PlcDeck[deck]).MethodByName(operation).Call([]reflect.Value{})
	// TODO : variadic parameters ought to be handled as well
	//  this will make Call to any method generic
	// TODO : Handle Panics with recover()

	if len(result) != 2 {
		fmt.Println("result is different in this reflect Call !", result)
		err := fmt.Errorf("unexpected length result")
		deps.WsErrCh <- err
		return "", err
	}

	fmt.Println("Correct Result: ", result)
	response = result[0].String()
	if len(response) > 0 {
		errRes := result[1].Interface()
		if errRes != nil {
			fmt.Println(errRes)
			err := fmt.Errorf("%v", errRes)
			deps.WsErrCh <- err
			return "", err
		}
		if operation == "Homing" {
			successWsData := plc.WSData{
				Progress: 100,
				Deck:     deck,
				Status:   "SUCCESS_HOMING",
				OperationDetails: plc.OperationDetails{
					Message: fmt.Sprintf("successfully homed for deck %v", deck),
				},
			}
			wsData, err := json.Marshal(successWsData)
			if err != nil {
				logger.Errorf("error in marshalling web socket data %v", err.Error())
				deps.WsErrCh <- err
			}
			deps.WsMsgCh <- fmt.Sprintf("success_homing_%v_%v", deck, string(wsData))
		}
	}

	return
}
