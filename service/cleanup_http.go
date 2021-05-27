package service

import (
	"mylab/cpagent/db"
	"mylab/cpagent/responses"
	"net/http"

	"github.com/gorilla/mux"
	logger "github.com/sirupsen/logrus"
)

func discardBoxCleanupHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.InitialisedState, db.ExecuteOperation, "", responses.DiscardBoxInitialisedState)

		var err error

		// for logging error if there is any otherwise logging success
		defer func() {
			if err != nil {
				go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.ErrorState, db.ExecuteOperation, "", err.Error())
			} else {
				go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.CompletedState, db.ExecuteOperation, "", responses.DiscardBoxCompletedState)
			}
		}()

		vars := mux.Vars(req)
		deck := vars["deck"]

		_, err = singleDeckOperation(req.Context(), deps, deck, "DiscardBoxCleanup")
		if err != nil {
			logger.Errorln(err.Error())
			responseCodeAndMsg(rw, http.StatusInternalServerError, ErrObj{Err: responses.DiscardBoxMoveError.Error(), Deck: deck})
			return
		}

		logger.Infoln(responses.DiscardBoxMovedSuccess)
		responseCodeAndMsg(rw, http.StatusOK, MsgObj{Msg: responses.DiscardBoxMovedSuccess, Deck: deck})
		return
	})
}

func restoreDeckHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.InitialisedState, db.ExecuteOperation, "", responses.RestoreDeckInitialisedState)

		var err error
		// for logging error if there is any otherwise logging success
		defer func() {
			if err != nil {
				go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.ErrorState, db.ExecuteOperation, "", err.Error())
			} else {
				go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.CompletedState, db.ExecuteOperation, "", responses.RestoreDeckCompletedState)
			}
		}()
		vars := mux.Vars(req)
		deck := vars["deck"]

		_, err = singleDeckOperation(req.Context(), deps, deck, "RestoreDeck")
		if err != nil {
			logger.Errorln(err.Error())
			responseCodeAndMsg(rw, http.StatusInternalServerError, ErrObj{Err: responses.RestoreDeckError.Error(), Deck: deck})
			return
		}

		logger.Infoln(responses.RestoreDeckSuccess)
		responseCodeAndMsg(rw, http.StatusOK, MsgObj{Msg: responses.RestoreDeckSuccess, Deck: deck})
		return
	})
}

func uvLightHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.InitialisedState, db.ExecuteOperation, "", responses.UvLightInitialisedState)

		vars := mux.Vars(req)
		deck := vars["deck"]

		uvTime := vars["time"]
		var err error
		// for logging error if there is any otherwise logging success
		defer func() {
			if err != nil {
				go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.ErrorState, db.ExecuteOperation, "", err.Error())
			} else {
				go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.CompletedState, db.ExecuteOperation, "", responses.UvLightCompletedState)
			}
		}()

		go deps.PlcDeck[deck].UVLight(uvTime)
		logger.Infoln(responses.UVCleanupProgress)
		responseCodeAndMsg(rw, http.StatusOK, MsgObj{Msg: responses.UVCleanupProgress, Deck: deck})
		return
	})
}
