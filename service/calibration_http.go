package service

import (
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

		logger.Infoln(responses.PIDCalibrationSuccess)
		responseCodeAndMsg(rw, http.StatusOK, MsgObj{Msg: responses.PIDCalibrationSuccess})
	})
}
