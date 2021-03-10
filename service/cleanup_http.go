package service

import (
	"net/http"
	"fmt"

	"github.com/gorilla/mux"
)

func discardBoxCleanupHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		var response string
		var err error

		vars := mux.Vars(req)
		deck := vars["deck"]
		switch deck {
		case "A", "B":
			response, err = singleDeckOperation(deps, deck, "DiscardBoxCleanup")
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

func restoreDeckHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		var response string
		var err error

		vars := mux.Vars(req)
		deck := vars["deck"]
		switch deck {
		case "A", "B":
			response, err = singleDeckOperation(deps, deck, "RestoreDeck")
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