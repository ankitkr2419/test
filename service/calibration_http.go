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

		err := deps.PlcDeck[deck].PIDCalibration(req.Context())
		if err != nil {
			logger.WithField("err", err.Error()).Error(responses.PIDCalibrationError)
			// TODO: Add Deck below whenever Deck PR is merged
			responseCodeAndMsg(rw, http.StatusInternalServerError, ErrObj{Err: responses.PIDCalibrationError.Error()}) //, Deck: deck})
			return
		}

		logger.Infoln(responses.PIDCalibrationSuccess)
		responseCodeAndMsg(rw, http.StatusOK, MsgObj{Msg: responses.PIDCalibrationSuccess})
	})
}
