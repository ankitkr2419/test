package service

import (
	"encoding/json"

	"mylab/cpagent/tec"
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

		err = tec.ConnectTEC(t)
		if err != nil{
			responseCodeAndMsg(rw, http.StatusInternalServerError, ErrObj{Err: err.Error()})
			return
		}

		responseCodeAndMsg(rw, http.StatusOK, MsgObj{Msg: "Temp and Ramp set success"} )
	})
}