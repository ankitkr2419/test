package service

import (
	"mylab/cpagent/responses"
	"net/http"

	logger "github.com/sirupsen/logrus"
)

func rtpcrResetHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		err := deps.Plc.Reset()
		if err != nil {
			responseCodeAndMsg(rw, http.StatusInternalServerError, ErrObj{Err: responses.RTPCRResetError.Error()})
			logger.WithField("err", err.Error()).Error(responses.RTPCRResetError)
			return
		}
		logger.Infoln(responses.RTPCRResetSuccess)
		responseCodeAndMsg(rw, http.StatusOK, responses.RTPCRResetSuccess)
		return

	})
}