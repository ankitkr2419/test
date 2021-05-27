package service

import (
	"mylab/cpagent/responses"
	"net/http"
	"time"

	logger "github.com/sirupsen/logrus"
)

func safeToUpgradeHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		if deps.PlcDeck[deckA].IsRunInProgress() || deps.PlcDeck[deckB].IsRunInProgress() {
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
	
	deps.PlcDeck[deckA].SetRunInProgress()
	defer deps.PlcDeck[deckA].ResetRunInProgress()
	deps.PlcDeck[deckB].SetRunInProgress()
	defer deps.PlcDeck[deckB].ResetRunInProgress()

	time.Sleep(dur * time.Second)
	
	logger.Infoln(responses.ResettingBothDeckRun)
}
