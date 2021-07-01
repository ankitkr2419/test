package service

import (
	"mylab/cpagent/tec"

	"encoding/json"
	"net/http"
)

func setTempAndRampHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		var err error

		var t tec.TECTempSet
		err = json.NewDecoder(req.Body).Decode(&t)
		if err != nil {
			responseCodeAndMsg(rw, http.StatusBadRequest, ErrObj{Err: err.Error()})
			return
		}

		go deps.Tec.ConnectTEC(t)

		responseCodeAndMsg(rw, http.StatusOK, MsgObj{Msg: "Temp and Ramp set success"} )
	})
}