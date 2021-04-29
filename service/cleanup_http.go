package service

import (
	"fmt"
	"net/http"

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
			err = fmt.Errorf("Check your deck name")
		}

		if err != nil {
			fmt.Fprintf(rw, err.Error())
			fmt.Println(err.Error())
			rw.WriteHeader(http.StatusInternalServerError)
		} else {
			rw.Header().Add("Content-Type", "application/json")
			rw.Write([]byte(fmt.Sprintf(`{"msg":%v,"deck":"%v"}`,response, deck)))
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
			err = fmt.Errorf("Check your deck name")
		}

		if err != nil {
			fmt.Fprintf(rw, err.Error())
			fmt.Println(err.Error())
			rw.WriteHeader(http.StatusInternalServerError)
		} else {
			rw.Header().Add("Content-Type", "application/json")
			rw.Write([]byte(fmt.Sprintf(`{"msg":%v,"deck":"%v"}`,response, deck)))
			rw.WriteHeader(http.StatusOK)
		}
	})
}

func uvLightHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		vars := mux.Vars(req)
		deck := vars["deck"]

		uvTime := vars["time"]

		switch deck {
		case "A", "B":
			rw.Header().Add("Content-Type", "application/json")
			rw.WriteHeader(http.StatusOK)
			rw.Write([]byte(fmt.Sprintf(`{"msg":"uv light clean up in progress","deck":"%v"}`, deck)))
			go deps.PlcDeck[deck].UVLight(uvTime)
		default:
			err := fmt.Errorf("Check your deck name")
			deps.WsErrCh <- err
		}

	})
}
