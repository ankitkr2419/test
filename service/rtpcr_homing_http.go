package service

import (
	"fmt"
	"mylab/cpagent/responses"
	"net/http"

	logger "github.com/sirupsen/logrus"
)

func rtpcrHomingHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		err := deps.Plc.HomingRTPCR()
		if err != nil {
			responseCodeAndMsg(rw, http.StatusInternalServerError, ErrObj{Err: responses.RTPCRHomingError.Error()})
			logger.WithField("err", err.Error()).Error(responses.RTPCRHomingError)
			return
		}

		logger.Infoln(responses.RTPCRHomingSuccess)
		responseCodeAndMsg(rw, http.StatusOK, responses.RTPCRHomingSuccess)
		return
	})
}

func rtpcrStartCycleHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		err := deps.Plc.Cycle()
		if err != nil {
			err = fmt.Errorf("%v", "error in starting cycle rt-pcr")
			rw.Write([]byte(err.Error()))
			rw.WriteHeader(http.StatusBadRequest)
			logger.Errorln(err)
			return
		}
		return

	})
}

func rtpcrMonitorHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		err := deps.Plc.Start()
		if err != nil {
			err = fmt.Errorf("%v", "error in monitoring rt-pcr")
			rw.Write([]byte(err.Error()))
			rw.WriteHeader(http.StatusBadRequest)
			logger.Errorln(err)
			return
		}
		return

	})
}
