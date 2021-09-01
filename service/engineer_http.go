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

		if deps.PlcDeck[deck].IsRunInProgress() {
			logger.Errorln(responses.PreviousRunInProgressError)
			responseCodeAndMsg(rw, http.StatusBadRequest, ErrObj{Err: responses.PreviousRunInProgressError.Error(), Deck: deck})
			return
		}

		deps.PlcDeck[deck].SetRunInProgress()
		defer deps.PlcDeck[deck].ResetRunInProgress()

		go deps.PlcDeck[deck].PIDCalibration(req.Context())

		logger.Infoln(responses.PIDCalibrationStarted)
		responseCodeAndMsg(rw, http.StatusOK, MsgObj{Msg: responses.PIDCalibrationStarted})
	})
}

func lidPIDCalibrationStartHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		if plc.LidPidTuningInProgress {
			logger.Errorln(responses.LidPIDTuningPresentError)
			responseCodeAndMsg(rw, http.StatusBadRequest, ErrObj{Err: responses.LidPIDTuningPresentError.Error()})
			return
		}

		// TODO: Logging this API
		go deps.Plc.LidPIDCalibration()

		logger.Infoln(responses.PIDCalibrationStarted)
		responseCodeAndMsg(rw, http.StatusOK, MsgObj{Msg: responses.PIDCalibrationStarted})
	})
}

func lidPIDCalibrationStopHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		if !plc.LidPidTuningInProgress {
			logger.Errorln(responses.LidPidTuningStartError)
			responseCodeAndMsg(rw, http.StatusBadRequest, ErrObj{Err: responses.LidPidTuningStartError.Error()})
			return
		}

		// Stop the lid PID Tuning
		err := deps.Plc.Stop()
		if err != nil {
			logger.WithField("err", err.Error()).Error(responses.LidPidTuningNotOffError)
			responseCodeAndMsg(rw, http.StatusInternalServerError, MsgObj{Msg: responses.LidPidTuningNotOffError.Error()})
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
			logger.Errorln(responses.PreviousRunInProgressError)
			responseCodeAndMsg(rw, http.StatusBadRequest, ErrObj{Err: responses.PreviousRunInProgressError.Error(), Deck: deck})
			return
		}

		deps.PlcDeck[deck].SetRunInProgress()
		defer deps.PlcDeck[deck].ResetRunInProgress()

		go deps.PlcDeck[deck].Shaking(shObj)

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
			logger.Errorln(responses.PreviousRunInProgressError)
			responseCodeAndMsg(rw, http.StatusBadRequest, ErrObj{Err: responses.PreviousRunInProgressError.Error(), Deck: deck})
			return
		}

		deps.PlcDeck[deck].SetRunInProgress()
		defer deps.PlcDeck[deck].ResetRunInProgress()

		go deps.PlcDeck[deck].Heating(hObj)

		logger.Infoln(responses.HeatingSuccess)
		responseCodeAndMsg(rw, http.StatusAccepted, MsgObj{Msg: responses.HeatingSuccess})
	})
}
