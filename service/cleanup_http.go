package service

import (
	"fmt"
	"mylab/cpagent/db"
	"mylab/cpagent/responses"
	"net/http"

	"github.com/gorilla/mux"
	logger "github.com/sirupsen/logrus"
)

func discardBoxCleanupHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.InitialisedState, db.ExecuteOperation, "", responses.DiscardBoxInitialisedState)

		var response string
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
		switch deck {
		case "A", "B":
			response, err = singleDeckOperation(deps, deck, "DiscardBoxCleanup")
		default:
			err = fmt.Errorf("Check your deck name")
		}

		if err != nil {
			responseCodeAndMsg(rw, http.StatusInternalServerError, ErrObj{Err: err.Error()})
			logger.WithField("err", err.Error()).Error(responses.WrongDeckError)
		} else {
			logger.Infoln(response)
			responseCodeAndMsg(rw, http.StatusOK, MsgObj{Msg: response, Deck: deck})
		}
	})
}

func restoreDeckHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.InitialisedState, db.ExecuteOperation, "", responses.RestoreDeckInitialisedState)

		var response string
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
		switch deck {
		case "A", "B":
			response, err = singleDeckOperation(deps, deck, "RestoreDeck")
		default:
			err = fmt.Errorf("Check your deck name")
		}

		if err != nil {
			responseCodeAndMsg(rw, http.StatusInternalServerError, ErrObj{Err: err.Error()})
			logger.WithField("err", err.Error()).Error(responses.WrongDeckError)
		} else {
			logger.Infoln(response)
			responseCodeAndMsg(rw, http.StatusOK, MsgObj{Msg: response, Deck: deck})
		}
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

		switch deck {
		case "A", "B":
			logger.Infoln(responses.UvCleanUpSuccess)
			responseCodeAndMsg(rw, http.StatusOK, MsgObj{Msg: responses.UvCleanUpSuccess, Deck: deck})
			go deps.PlcDeck[deck].UVLight(uvTime)
		default:
			responseCodeAndMsg(rw, http.StatusInternalServerError, ErrObj{Err: err.Error()})
			logger.WithField("err", err.Error()).Error(err)
			deps.WsErrCh <- err
		}

	})
}
