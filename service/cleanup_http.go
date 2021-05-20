package service

import (
	"net/http"

	"github.com/gorilla/mux"
	"mylab/cpagent/responses"

	logger "github.com/sirupsen/logrus"
)

func discardBoxCleanupHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		var err error

		vars := mux.Vars(req)
		deck := vars["deck"]

		_, err = singleDeckOperation(deps, deck, "DiscardBoxCleanup")
		if err != nil {
			logger.Errorln(err.Error())
			responseCodeAndMsg(rw, http.StatusInternalServerError, ErrObj{Err: err.Error(), Deck: deck})
			return
		}

		logger.Infoln(responses.DiscardBoxMovedSuccess)
		responseCodeAndMsg(rw, http.StatusInternalServerError, MsgObj{Msg: responses.DiscardBoxMovedSuccess, Deck: deck})
		return
	})
}

func restoreDeckHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		var err error

		vars := mux.Vars(req)
		deck := vars["deck"]

		_, err = singleDeckOperation(deps, deck, "RestoreDeck")
		if err != nil {
			logger.Errorln(err.Error())
			responseCodeAndMsg(rw, http.StatusInternalServerError, ErrObj{Err: err.Error(), Deck: deck})
			return
		}

		logger.Infoln(responses.RestoreDeckSuccess)
		responseCodeAndMsg(rw, http.StatusInternalServerError, MsgObj{Msg: responses.RestoreDeckSuccess, Deck: deck})
		return
	})
}

func uvLightHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		vars := mux.Vars(req)
		deck := vars["deck"]

		uvTime := vars["time"]

		go deps.PlcDeck[deck].UVLight(uvTime)
		logger.Infoln(responses.UVCleanupProgress)
		responseCodeAndMsg(rw, http.StatusOK, MsgObj{Msg: responses.UVCleanupProgress, Deck: deck})
		return
	})
}
