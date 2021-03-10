package service

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Manual struct {
	Deck      string `json:"deck"`
	MotorNum  int    `json:"motor_number"`
	Pulses    int    `json:"pulses"`
	Direction int    `json:"direction"`
}

func manualHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		var response string
		var err error

		var m Manual
		err = json.NewDecoder(req.Body).Decode(&m)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			fmt.Println("Error decoding manual data")
			return
		}

		switch {
		case m.Deck != "A" && m.Deck != "B":
			rw.WriteHeader(http.StatusBadRequest)
			err = fmt.Errorf("Use A or B deck only")
			fmt.Println(err)
			return
		case m.MotorNum <= 4 || m.MotorNum > 10:
			rw.WriteHeader(http.StatusBadRequest)
			err = fmt.Errorf("Select motor num in only in between 5-10")
			fmt.Println(err)
			return
		case m.Direction != 0 && m.Direction != 1:
			rw.WriteHeader(http.StatusBadRequest)
			err = fmt.Errorf("Select motor direction in only as 0 or 1")
			fmt.Println(err)
			return
		case m.Pulses > 10000:
			rw.WriteHeader(http.StatusBadRequest)
			err = fmt.Errorf("Consider pulses only less than or equal to 10000")
			fmt.Println(err)
			return
		}

		switch m.Deck {
		case "A", "B":
			response, err = deps.PlcDeck[m.Deck].ManualMovement(uint16(m.MotorNum), uint16(m.Direction), uint16(m.Pulses))
		default:
			err = fmt.Errorf("Please check your deck")
		}

		if err != nil {
			fmt.Fprintf(rw, err.Error())
			fmt.Println(err.Error())
			rw.WriteHeader(http.StatusBadRequest)
		} else {
			fmt.Fprintf(rw, response+" Manual Movements in Progress/Done")
			rw.WriteHeader(http.StatusOK)
		}
	})
}

func pauseHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		var response string
		var err error

		vars := mux.Vars(req)
		deck := vars["deck"]
		switch deck {
		case "A", "B":
			response, err = singleDeckOperation(deps, deck, "Pause")
		default:
			err = fmt.Errorf("Check your deck name")
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

func resumeHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		var response string
		var err error

		vars := mux.Vars(req)
		deck := vars["deck"]
		switch deck {
		case "":
		case "A", "B":
			response, err = singleDeckOperation(deps, deck, "Resume")
		default:
			err = fmt.Errorf("Check your deck name")
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

func abortHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		var response string
		var err error

		vars := mux.Vars(req)
		deck := vars["deck"]

		fmt.Println("Inside ABORT... value of deck:", deck, len(deck))
		switch deck {
		case "A", "B":
			response, err = singleDeckOperation(deps, deck, "Abort")
		default:
			err = fmt.Errorf("Check your deck name")
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
