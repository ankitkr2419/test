package service

import (
	"mylab/cpagent/responses"
	"net/http"

	logger "github.com/sirupsen/logrus"
)

func pingHandler(rw http.ResponseWriter, req *http.Request) {
	logger.Infoln("Server was pinged")
	responseCodeAndMsg(rw, http.StatusOK, MsgObj{Msg: responses.Pong})
	return
}
