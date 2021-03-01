package service

import (
	"fmt"
	"net/http"
	"reflect"

	"github.com/gorilla/mux"
)

func homingHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		var response string
		var err error

		vars := mux.Vars(req)
		deck := vars["deck"]
		switch deck {
		case "A", "B":
			response, err = singleDeckOperation(deps, deck, "Homing")
		default:
			err = fmt.Errorf("Check you deck name")
		}

		if err != nil {
			fmt.Fprintf(rw, err.Error())
			fmt.Println(err.Error())
			rw.WriteHeader(http.StatusInternalServerError)
		} else {
			fmt.Fprintf(rw, response)
			rw.WriteHeader(http.StatusOK)
		}
	})
}

func singleDeckOperation(deps Dependencies, deck, operation string) (response string, err error) {
	switch deck {
	case "A", "B":

		// Compact32Deck is the type of deps.PlcDeck[deck]
		result := reflect.ValueOf(deps.PlcDeck[deck]).MethodByName(operation).Call([]reflect.Value{})
		// TODO : variadic parameters ought to be handled as well
		//  this will make Call to any method generic
		// TODO : Handle Panics with recover()

		if len(result) != 2 {
			fmt.Println("result is different in this reflect Call !", result)
			return "", fmt.Errorf("unexpected length result")
		}

		fmt.Println("Correct Result: ", result)
		response = result[0].String()
		if len(response) > 0 {
			return
		}
		errRes := result[1].Interface()
		if errRes != nil {
			fmt.Println(errRes)
			return "", fmt.Errorf("%v", errRes)
		}

	default:
		err = fmt.Errorf("Check your deck name")
	}
	return
}
