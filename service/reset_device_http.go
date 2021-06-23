package service

import (
	"mylab/cpagent/tec"
	"net/http"
)

func resetDeviceHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		var err error

		err = tec.ResetDevice()
		if err != nil{
			responseCodeAndMsg(rw, http.StatusInternalServerError, ErrObj{Err: err.Error()})
			return
		}

		responseCodeAndMsg(rw, http.StatusOK, MsgObj{Msg: "Reset Device success"} )
	})
}