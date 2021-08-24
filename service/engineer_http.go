package service

import (
	"encoding/json"
	"mylab/cpagent/db"
	"mylab/cpagent/plc"
	"mylab/cpagent/responses"
	"net/http"

	"github.com/gorilla/mux"
	logger "github.com/sirupsen/logrus"
)

func pidCalibrationHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		// TODO: Logging this API
		vars := mux.Vars(req)
		deck := vars["deck"]
		var err error

		if deps.PlcDeck[deck].IsRunInProgress() {
			logger.WithField("err", err.Error()).Error(responses.PreviousRunInProgressError)
			return
		}

		deps.PlcDeck[deck].SetRunInProgress()
		defer deps.PlcDeck[deck].ResetRunInProgress()

		go deps.PlcDeck[deck].PIDCalibration(req.Context())
		if err != nil {
			logger.WithField("err", err.Error()).Error(responses.PIDCalibrationError)
			responseCodeAndMsg(rw, http.StatusInternalServerError, ErrObj{Err: responses.PIDCalibrationError.Error(), Deck: deck})
			return
		}

		logger.Infoln(responses.PIDCalibrationStarted)
		responseCodeAndMsg(rw, http.StatusOK, MsgObj{Msg: responses.PIDCalibrationStarted})
	})
}

func lidPIDCalibrationStartHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		// TODO: Logging this API
		go deps.Plc.LidPIDCalibration()

		logger.Infoln(responses.PIDCalibrationStarted)
		responseCodeAndMsg(rw, http.StatusOK, MsgObj{Msg: responses.PIDCalibrationStarted})
	})
}

func lidPIDCalibrationStopHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		if !plc.LidPidTuningInProgress {
			logger.Errorln(responses.LidPidTuningPresentError)
			responseCodeAndMsg(rw, http.StatusBadRequest, ErrObj{Err: responses.LidPidTuningPresentError.Error()})
		}

		// Stop the lid PID Tuning
		err := deps.Plc.Stop()
		if err != nil {
			logger.WithField("err", err.Error()).Error("Error in plc stop")
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}

		responseCodeAndMsg(rw, http.StatusOK, MsgObj{Msg: responses.LidPIDTuningStopped})
	})
}

func shakerHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		// TODO: Logging this API
		vars := mux.Vars(req)
		deck := vars["deck"]
		var err error

		var shObj db.Shaker
		err = json.NewDecoder(req.Body).Decode(&shObj)
		if err != nil {
			logger.WithField("err", err.Error()).Errorln(responses.ShakingDecodeError)
			responseCodeAndMsg(rw, http.StatusBadRequest, ErrObj{Err: responses.ShakingDecodeError.Error()})
			return
		}

		valid, respBytes := validate(shObj)
		if !valid {
			logger.WithField("err", "Validation Error").Errorln(responses.ShakingValidationError)
			responseBadRequest(rw, respBytes)
			return
		}

		if deps.PlcDeck[deck].IsRunInProgress() {
			logger.WithField("err", err.Error()).Error(responses.PreviousRunInProgressError)
			return
		}

		deps.PlcDeck[deck].SetRunInProgress()
		defer deps.PlcDeck[deck].ResetRunInProgress()

		go deps.PlcDeck[deck].Shaking(shObj)
		if err != nil {
			logger.WithField("err", err.Error()).Error(responses.ShakingError)
			responseCodeAndMsg(rw, http.StatusInternalServerError, ErrObj{Err: responses.ShakingError.Error(), Deck: deck})
			return
		}

		logger.Infoln(responses.ShakingSuccess)
		responseCodeAndMsg(rw, http.StatusAccepted, MsgObj{Msg: responses.ShakingSuccess})
	})
}

func heaterHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		// TODO: Logging this API
		vars := mux.Vars(req)
		deck := vars["deck"]
		var err error

		var hObj db.Heating
		err = json.NewDecoder(req.Body).Decode(&hObj)
		if err != nil {
			logger.WithField("err", err.Error()).Errorln(responses.HeatingDecodeError)
			responseCodeAndMsg(rw, http.StatusBadRequest, ErrObj{Err: responses.HeatingDecodeError.Error()})
			return
		}

		valid, respBytes := validate(hObj)
		if !valid {
			logger.WithField("err", "Validation Error").Errorln(responses.HeatingValidationError)
			responseBadRequest(rw, respBytes)
			return
		}

		if deps.PlcDeck[deck].IsRunInProgress() {
			logger.WithField("err", err.Error()).Error(responses.PreviousRunInProgressError)
			return
		}

		deps.PlcDeck[deck].SetRunInProgress()
		defer deps.PlcDeck[deck].ResetRunInProgress()

		go deps.PlcDeck[deck].Heating(hObj)
		if err != nil {
			logger.WithField("err", err.Error()).Error(responses.HeatingError)
			responseCodeAndMsg(rw, http.StatusInternalServerError, ErrObj{Err: responses.HeatingError.Error(), Deck: deck})
			return
		}

		logger.Infoln(responses.HeatingSuccess)
		responseCodeAndMsg(rw, http.StatusAccepted, MsgObj{Msg: responses.HeatingSuccess})
	})
}
