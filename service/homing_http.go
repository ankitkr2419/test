package service

import (
	"fmt"
	"mylab/cpagent/plc/compact32"
	"net/http"
	"reflect"
	"time"

	"github.com/gorilla/mux"
	logger "github.com/sirupsen/logrus"
)

func homingHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		var response string
		var err error
		vars := mux.Vars(req)
		deck := vars["deck"]

		switch deck {
		case "":
			fmt.Println("At both deck!!!")
			rw.Write([]byte(`Operation in progress for both decks`))
			rw.WriteHeader(http.StatusOK)
			compact32.SetBothDeckHomingInProgress()
			defer compact32.ResetBothDeckHomingInProgress()
			go bothDeckOperation(deps, "Homing")
		case "A", "B":
			rw.Write([]byte(`Operation in progress for single deck`))
			rw.WriteHeader(http.StatusOK)
			go singleDeckOperation(deps, deck, "Homing")
		default:
			err = fmt.Errorf("Check your deck name")
			rw.Write([]byte(err.Error()))
			rw.WriteHeader(http.StatusBadRequest)
		}

		if err != nil {
			logger.Errorln(err)
			deps.WsErrCh <- err
		} else {
			logger.Infoln(response)
			deps.WsMsgCh <- fmt.Sprintf("success_homing%v_successfully homed", deck)
		}

	})
}

func bothDeckOperation(deps Dependencies, operation string) (response string, err error) {

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
			// abort Deck B Operation as Well
			response, err = deps.PlcDeck["B"].Abort()
			if err != nil {
				deps.WsErrCh <- err
				return
			}
			return "", deckAErr
		case deckBErr != nil:
			fmt.Printf("Error %s deck B", operation)
			// Abort Deck A Operation as well
			response, err = deps.PlcDeck["A"].Abort()
			if err != nil {
				deps.WsErrCh <- err
				return
			}
			return "", deckBErr
		case deckAResponse != "" && deckBResponse != "":

			operationSuccessMsg := fmt.Sprintf("%s Success for both Decks!", operation)
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
		return
	}
	errRes := result[1].Interface()
	if errRes != nil {
		fmt.Println(errRes)
		err := fmt.Errorf("%v", errRes)
		deps.WsErrCh <- err
		return "", err
	}

	return
}
