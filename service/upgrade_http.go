package service

import (
	"mylab/cpagent/plc"
	"mylab/cpagent/responses"
	"net/http"
	"time"

	logger "github.com/sirupsen/logrus"
)

func safeToUpgradeHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		if deps.PlcDeck[plc.DeckA].IsRunInProgress() || deps.PlcDeck[plc.DeckB].IsRunInProgress() {
			logger.Errorln(responses.RunInProgressForSomeDeckError)
			responseCodeAndMsg(rw, http.StatusBadRequest, ErrObj{Err: responses.RunInProgressForSomeDeckError.Error()})
			return
		}

		// set Run in Progress for 10 sec for both the decks
		// this gives script ample amount to shut down cpagent
		go temporarySetRunInProgress(deps, 10)

		logger.Infoln(responses.SafeToUpgrade)
		responseCodeAndMsg(rw, http.StatusOK, MsgObj{Msg: responses.SafeToUpgrade})
	})
}

// set Run in Progress for dur seconds for both the decks
func temporarySetRunInProgress(deps Dependencies, dur time.Duration) {
	logger.Infoln(responses.TempSettingBothDeckRun)

	deps.PlcDeck[plc.DeckA].SetRunInProgress()
	defer deps.PlcDeck[plc.DeckA].ResetRunInProgress()
	deps.PlcDeck[plc.DeckB].SetRunInProgress()
	defer deps.PlcDeck[plc.DeckB].ResetRunInProgress()

	time.Sleep(dur * time.Second)

	logger.Infoln(responses.ResettingBothDeckRun)
}
