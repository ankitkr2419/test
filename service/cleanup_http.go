package service

import (
	"fmt"
	"mylab/cpagent/db"
	"mylab/cpagent/plc"
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
		// totalTime is UVLight timer time in Seconds
		var totalTime int64

		// for logging error if there is any otherwise logging success
		defer func() {
			if err != nil {
				go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.ErrorState, db.ExecuteOperation, "", err.Error())
			} else {
				go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.CompletedState, db.ExecuteOperation, "", responses.UvLightCompletedState)
			}
		}()

		totalTime, err = db.CalculateTimeInSeconds(uvTime)
		if err != nil {
			logger.WithField("Decode Error", err).Errorln(responses.UVTimeFormatDecodeError)
			responseCodeAndMsg(rw, http.StatusBadRequest, ErrObj{Err: responses.UVTimeFormatDecodeError.Error(), Deck: deck})
			return
		}

		if totalTime < plc.MinimumUVLightOnTime {
			err = fmt.Errorf("please check your time. minimum time is : %v seconds", plc.MinimumUVLightOnTime)
			logger.WithField("Validation Error", err).Errorln(responses.UVMinimumTimeError)
			responseCodeAndMsg(rw, http.StatusBadRequest, ErrObj{Err: responses.UVMinimumTimeError.Error(), Deck: deck})
			return
		}

		logger.Infoln("totalTime", totalTime)

		go deps.PlcDeck[deck].UVLight(totalTime)
		logger.Infoln(responses.UVCleanupProgress)
		responseCodeAndMsg(rw, http.StatusOK, MsgObj{Msg: responses.UVCleanupProgress, Deck: deck})
		return
	})
}
