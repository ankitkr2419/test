package service

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func deckHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		response, err := deps.PlcDeckA.DeckHoming()
		if err != nil {
			fmt.Fprintf(rw, err.Error())
		} else {
			fmt.Fprintf(rw, response)
		}
	})
}

func syringeHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		response, err := deps.PlcDeckA.SyringeHoming()
		if err != nil {
			fmt.Fprintf(rw, err.Error())
		} else {
			fmt.Fprintf(rw, response)
		}
	})
}

func syringeModuleHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		response, err := deps.PlcDeckA.SyringeModuleHoming()
		if err != nil {
			fmt.Fprintf(rw, err.Error())
		} else {
			fmt.Fprintf(rw, response)
		}
	})
}

func magnetHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		response, err := deps.PlcDeckA.MagnetHoming()
		if err != nil {
			fmt.Fprintf(rw, err.Error())
		} else {
			fmt.Fprintf(rw, response)
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
			response, err = homeBothDecks(deps)
		case "A", "B":
			response, err = homeDeck(deps, deck)
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

func homeBothDecks(deps Dependencies) (response string, err error) {

	var deckAResponse, deckBResponse string
	var deckAErr, deckBErr error

	go func() {
		deckAResponse, deckAErr = homeDeck(deps, "A")
	}()
	go func() {
		deckBResponse, deckBErr = homeDeck(deps, "B")
	}()

	for {
		switch {
		case deckAErr != nil:
			fmt.Println("Error homing deck A")
			return "", deckAErr
		case deckBErr != nil:
			fmt.Println("Error homing deck B")
			return "", deckBErr
		case deckAResponse != "" && deckBResponse != "":
			fmt.Println("Homing Success for both decks")
			return "Homing Success for both Decks!", nil
		default:
			// Only check every 400 milli second
			time.Sleep(400 * time.Millisecond)
		}
	}
}

func homeDeck(deps Dependencies, deck string) (response string, err error) {
	switch deck {
	case "A":
		response, err = deps.PlcDeckA.Homing()
	case "B":
		response, err = deps.PlcDeckB.Homing()
	default:
		err = fmt.Errorf("Check your deck name")
	}
	return
}
