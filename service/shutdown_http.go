package service

import (
	"mylab/cpagent/plc"
	"mylab/cpagent/responses"
	"net/http"
	"os/exec"

	logger "github.com/sirupsen/logrus"
)

func shutDownHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		//writing to off register for station1 as asked by sanket
		err := deps.PlcDeck[plc.DeckA].ShutDown()
		if err != nil {
			logger.WithField("err", err.Error()).Errorln(responses.ShutdownError)
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}
		go func() {
			cmd := exec.Command("bash", "-c", "sleep 8 && shutdown now")
			err := cmd.Run()
			if err != nil {
				logger.Errorf("There is an error shutdown the system")
				return
			}
		}()
		logger.Infoln(responses.ShutDownSucess)
		responseCodeAndMsg(rw, http.StatusCreated, "shutdown")
	})
}
