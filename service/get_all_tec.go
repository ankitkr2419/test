package service

import (
	"net/http"	
)

func getAllTECHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

	err := deps.Tec.GetAllTEC()
	if err != nil{
		responseCodeAndMsg(rw, http.StatusInternalServerError, MsgObj{Msg: "Couldn't get the TEC values"} )

	}
	responseCodeAndMsg(rw, http.StatusOK, MsgObj{Msg: "Get All Values Success"} )
	})
}