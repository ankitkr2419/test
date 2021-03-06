package service

import (
	"encoding/json"
	"fmt"
	"mylab/cpagent/plc"
	"net/http"

	"github.com/gorilla/mux"
	logger "github.com/sirupsen/logrus"
)

type Manual struct {
	Deck      string  `json:"deck"`
	MotorNum  int     `json:"motor_number"`
	Distance  float32 `json:"distance"`
	Direction uint16  `json:"direction"`
}

func manualHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		var response string
		var err error

		var m Manual
		err = json.NewDecoder(req.Body).Decode(&m)
		if err != nil {
			responseCodeAndMsg(rw, http.StatusBadRequest, ErrObj{Err: err.Error()})
			return
		}

		switch {
		case m.Deck != plc.DeckA && m.Deck != plc.DeckB:
			err = fmt.Errorf("Use A or B deck only")
		case m.MotorNum <= 4 || m.MotorNum > 10 || m.MotorNum == 8:
			err = fmt.Errorf("Select motor num in only in between 5-10. Shaker Motor not allowed")
		case m.Direction != plc.TowardsSensor && m.Direction != plc.AgainstSensor:
			err = fmt.Errorf("Select motor direction in only as 0 or 1")
		case m.Distance > 100 || m.Distance <= 0:
			err = fmt.Errorf("Consider distance only less than or equal to 100 mm")
		}

		if err != nil {
			responseCodeAndMsg(rw, http.StatusBadRequest, ErrObj{Err: err.Error(), Deck: m.Deck})
			return
		}

		response, err = deps.PlcDeck[m.Deck].ManualMovement(uint16(m.MotorNum), m.Direction, m.Distance)

		if err != nil {
			responseCodeAndMsg(rw, http.StatusBadRequest, ErrObj{Err: err.Error(), Deck: m.Deck})
		} else {
			responseCodeAndMsg(rw, http.StatusOK, MsgObj{Msg: response + " Manual Movements in Progress/Done", Deck: m.Deck})
			logger.Infoln(response)
		}
	})
}

func pauseHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		var response string
		var err error

		vars := mux.Vars(req)
		deck := vars["deck"]

		response, err = singleDeckOperation(req.Context(), deps, deck, "Pause")

		if err != nil {
			responseCodeAndMsg(rw, http.StatusInternalServerError, ErrObj{Err: err.Error(), Deck: deck})
		} else {
			responseCodeAndMsg(rw, http.StatusOK, MsgObj{Msg: response, Deck: deck})
			logger.Infoln(response)
		}
	})
}

func resumeHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		var response string
		var err error

		vars := mux.Vars(req)
		deck := vars["deck"]

		response, err = singleDeckOperation(req.Context(), deps, deck, "Resume")

		if err != nil {
			responseCodeAndMsg(rw, http.StatusInternalServerError, ErrObj{Err: err.Error(), Deck: deck})
		} else {
			responseCodeAndMsg(rw, http.StatusOK, MsgObj{Msg: response, Deck: deck})
			logger.Infoln(response)
		}
	})
}

func abortHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		var response string
		var err error

		vars := mux.Vars(req)
		deck := vars["deck"]

		logger.Infoln("Inside ABORT... value of deck:", deck, len(deck))

		response, err = singleDeckOperation(req.Context(), deps, deck, "Abort")

		if err != nil {
			responseCodeAndMsg(rw, http.StatusInternalServerError, ErrObj{Err: err.Error(), Deck: deck})
		} else {
			responseCodeAndMsg(rw, http.StatusOK, MsgObj{Msg: response, Deck: deck})
			logger.Infoln(response)
		}
	})
}
