package service

import (
	logger "github.com/sirupsen/logrus"
	"mylab/cpagent/tec"

	"encoding/json"
	"net/http"
)

func runProfileHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		var err error

		var t tec.TempProfile
		err = json.NewDecoder(req.Body).Decode(&t)
		if err != nil {
			responseCodeAndMsg(rw, http.StatusBadRequest, ErrObj{Err: err.Error()})
			return
		}

		logger.Infoln("TempProfile: ", t)

		err = deps.Tec.RunProfile(deps.Plc, t)
		if err != nil {
			responseCodeAndMsg(rw, http.StatusInternalServerError, ErrObj{Err: err.Error()})
			return
		}

		responseCodeAndMsg(rw, http.StatusOK, MsgObj{Msg: "Profile Run success"})
	})
}
