package service

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func deckHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		var response string
		var err error

		vars := mux.Vars(req)
		deck := vars["id"]
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
		deck := vars["id"]
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
		deck := vars["id"]
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
		deck := vars["id"]
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
		deck := vars["id"]
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
			fmt.Printf("Error %s deck A", operation)
			// abort Deck B Operation as Well
			response, err = deps.PlcDeck["B"].Abort()
			if err != nil {
				return
			}
			return "", deckAErr
		case deckBErr != nil:
			fmt.Printf("Error %s deck B", operation)
			// Abort Deck A Operation as well
			response, err = deps.PlcDeck["A"].Abort()
			if err != nil {
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
	switch deck {
	case "A", "B":
		switch operation {
		case "Homing":
			response, err = deps.PlcDeck[deck].Homing()
		case "SyringeHoming":
			response, err = deps.PlcDeck[deck].SyringeHoming()
		case "SyringeModuleHoming":
			response, err = deps.PlcDeck[deck].SyringeModuleHoming()
		case "DeckHoming":
			response, err = deps.PlcDeck[deck].DeckHoming()
		case "MagnetHoming":
			response, err = deps.PlcDeck[deck].MagnetHoming()
		case "Pause":
			response, err = deps.PlcDeck[deck].Pause()
		case "Resume":
			response, err = deps.PlcDeck[deck].Resume()
		case "Abort":
			response, err = deps.PlcDeck[deck].Abort()

		default:
			err = fmt.Errorf("Invalid Operation")
		}
	default:
		err = fmt.Errorf("Check your deck name")
	}
	return
}
