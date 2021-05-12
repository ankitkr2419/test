package service

import (
	"encoding/json"
	"mylab/cpagent/db"
	"mylab/cpagent/responses"
	"net/http"

	"github.com/gorilla/mux"
	logger "github.com/sirupsen/logrus"
)

func createTipDockHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		username := req.Context().Value(contextKeyUsername).(string)
		go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.InitialisedState, db.CreateOperation, "", responses.TipDockingInitialisedState, username)

		var tdObj db.TipDock
		err := json.NewDecoder(req.Body).Decode(&tdObj)

		// for logging error if there is any otherwise logging success
		defer func() {
			if err != nil {
				go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.ErrorState, db.CreateOperation, "", err.Error(), username)

			} else {
				go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.CompletedState, db.CreateOperation, "", responses.TipDockingCompletedState, username)

			}

		}()

		if err != nil {
			logger.WithField("err", err.Error()).Errorln(responses.TipDockingDecodeError)
			responseCodeAndMsg(rw, http.StatusBadRequest, ErrObj{Err: responses.TipDockingDecodeError.Error()})
			return
		}

		valid, respBytes := validate(tdObj)
		if !valid {
			logger.WithField("err", "Validation Error").Errorln( responses.TipDockingValidationError)
			responseBadRequest(rw, respBytes)
			return
		}

		var tipDock db.TipDock
		tipDock, err = deps.Store.CreateTipDocking(req.Context(), tdObj)
		if err != nil {
			logger.WithField("err", err.Error()).Errorln(responses.TipDockingCreateError)
			responseCodeAndMsg(rw, http.StatusInternalServerError, ErrObj{Err: responses.TipDockingCreateError.Error()})
			return
		}
		logger.Infoln(responses.TipDockingCreateSuccess)
		responseCodeAndMsg(rw, http.StatusCreated, tipDock)
	})
}

func showTipDockHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		//logging when the api is initialised
		username := req.Context().Value(contextKeyUsername).(string)
		go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.InitialisedState, db.ShowOperation, "", responses.TipDockingInitialisedState, username)

		vars := mux.Vars(req)
		id, err := parseUUID(vars["id"])

		// for logging error if there is any otherwise logging success
		defer func() {
			if err != nil {
				go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.ErrorState, db.ShowOperation, "", err.Error(), username)

			} else {
				go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.CompletedState, db.ShowOperation, "", responses.TipDockingCompletedState, username)

			}

		}()

		if err != nil {
			responseCodeAndMsg(rw, http.StatusBadRequest, ErrObj{Err: responses.UUIDParseError.Error()})
			return
		}

		var tipDock db.TipDock
		tipDock, err = deps.Store.ShowTipDocking(req.Context(), id)
		if err != nil {
			responseCodeAndMsg(rw, http.StatusInternalServerError, ErrObj{Err: responses.TipDockingFetchError.Error()})
			logger.WithField("err", err.Error()).Error(responses.TipDockingFetchError)
			return
		}

		logger.Infoln(responses.TipDockingFetchSuccess)
		responseCodeAndMsg(rw, http.StatusOK, tipDock)
	})
}

func updateTipDockHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		//logging when the api is initialised
		username := req.Context().Value(contextKeyUsername).(string)
		go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.InitialisedState, db.UpdateOperation, "", responses.TipDockingInitialisedState, username)

		vars := mux.Vars(req)
		id, err := parseUUID(vars["id"])
		// for logging error if there is any otherwise logging success
		defer func() {
			if err != nil {
				go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.ErrorState, db.UpdateOperation, "", err.Error(), username)

			} else {
				go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.CompletedState, db.UpdateOperation, "", responses.TipDockingCompletedState, username)

			}

		}()
		if err != nil {
			responseCodeAndMsg(rw, http.StatusBadRequest, ErrObj{Err: responses.UUIDParseError.Error()})
			return
		}
		var tdObj db.TipDock
		err = json.NewDecoder(req.Body).Decode(&tdObj)
		if err != nil {
			logger.WithField("err", err.Error()).Errorln(responses.TipDockingDecodeError)
			responseCodeAndMsg(rw, http.StatusBadRequest, ErrObj{Err: responses.TipDockingDecodeError.Error()})
			return
		}
		valid, respBytes := validate(tdObj)
		if !valid {
			logger.WithField("err", "Validation Error").Errorln( responses.TipDockingValidationError)
			responseBadRequest(rw, respBytes)
			return
		}
		tdObj.ProcessID = id
		err = deps.Store.UpdateTipDock(req.Context(), tdObj)
		if err != nil {
			logger.WithField("err", err.Error()).Error(responses.TipDockingUpdateError)
			responseCodeAndMsg(rw, http.StatusInternalServerError, ErrObj{Err: responses.TipDockingUpdateError.Error()})
			return
		}

		logger.Infoln(responses.TipDockingUpdateSuccess)
		responseCodeAndMsg(rw, http.StatusOK, MsgObj{Msg: responses.TipDockingUpdateSuccess})
	})
}
