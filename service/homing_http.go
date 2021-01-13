package service

import (
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/gorilla/mux"
)

func deckHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		var response string
		var err error

		vars := mux.Vars(req)
		deck := vars["deck"]
		switch deck {
		case "":
			response, err = bothDeckOperation(deps, "DeckHoming")
		case "A", "B":
			response, err = singleDeckOperation(deps, deck, "DeckHoming")
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

func syringeHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		var response string
		var err error

		vars := mux.Vars(req)
		deck := vars["deck"]
		switch deck {
		case "":
			response, err = bothDeckOperation(deps, "SyringeHoming")
		case "A", "B":
			response, err = singleDeckOperation(deps, deck, "SyringeHoming")
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

func syringeModuleHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		var response string
		var err error

		vars := mux.Vars(req)
		deck := vars["deck"]
		switch deck {
		case "":
			response, err = bothDeckOperation(deps, "SyringeModuleHoming")
		case "A", "B":
			response, err = singleDeckOperation(deps, deck, "SyringeModuleHoming")
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

func magnetHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		var response string
		var err error

		vars := mux.Vars(req)
		deck := vars["deck"]
		switch deck {
		case "":
			response, err = bothDeckOperation(deps, "MagnetHoming")
		case "A", "B":
			response, err = singleDeckOperation(deps, deck, "MagnetHoming")
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

func homingHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		var response string
		var err error

		vars := mux.Vars(req)
		deck := vars["deck"]
		switch deck {
		case "":
			response, err = bothDeckOperation(deps, "Homing")
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
			fmt.Printf("Error %s deck A", operation, deckAErr)
			return "", deckAErr
		case deckBErr != nil:
			fmt.Printf("Error %s deck B", operation, deckBErr)
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
