package service

import (
	"encoding/json"
	"mylab/cpagent/db"
	"mylab/cpagent/responses"
	"net/http"

	"github.com/gorilla/mux"
	logger "github.com/sirupsen/logrus"
)

func createAttachDetachHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		username := req.Context().Value(contextKeyUsername).(string)
		go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.InitialisedState, db.CreateOperation, "", responses.AttachDetachInitialisedState, username)

		var adObj db.AttachDetach
		err := json.NewDecoder(req.Body).Decode(&adObj)

		// for logging error if there is any otherwise logging success
		defer func() {
			if err != nil {
				go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.ErrorState, db.CreateOperation, "", err.Error(), username)

			} else {
				go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.CompletedState, db.CreateOperation, "", responses.AttachDetachCompletedState, username)

			}

		}()

		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			logger.WithField("err", err.Error()).Errorln(responses.AttachDetachDecodeError)
			responseCodeAndMsg(rw, http.StatusBadRequest, ErrObj{Err: responses.AttachDetachDecodeError.Error()})
			return
		}

		valid, respBytes := validate(adObj)
		if !valid {
			logger.WithField("err", "Validation Error").Errorln( responses.AttachDetachValidationError)
			responseBadRequest(rw, respBytes)
			return
		}

		var createdAtDt db.AttachDetach
		createdAtDt, err = deps.Store.CreateAttachDetach(req.Context(), adObj)
		if err != nil {
			logger.WithField("err", err.Error()).Errorln(responses.AttachDetachCreateError)
			responseCodeAndMsg(rw, http.StatusInternalServerError, ErrObj{Err: responses.AttachDetachCreateError.Error()})
			return
		}

		logger.Infoln(responses.AttachDetachCreateSuccess)
		responseCodeAndMsg(rw, http.StatusCreated, createdAtDt)
	})
}

func showAttachDetachHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		//logging when the api is initialised
		username := req.Context().Value(contextKeyUsername).(string)
		go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.InitialisedState, db.ShowOperation, "", responses.AttachDetachInitialisedState, username)

		vars := mux.Vars(req)
		id, err := parseUUID(vars["id"])

		// for logging error if there is any otherwise logging success
		defer func() {
			if err != nil {
				go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.ErrorState, db.ShowOperation, "", err.Error(), username)

			} else {
				go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.CompletedState, db.ShowOperation, "", responses.AttachDetachCompletedState, username)

			}

		}()

		if err != nil {
			responseCodeAndMsg(rw, http.StatusBadRequest, ErrObj{Err: responses.UUIDParseError.Error()})
			return
		}

		var attachDetach db.AttachDetach
		attachDetach, err = deps.Store.ShowAttachDetach(req.Context(), id)
		if err != nil {
			responseCodeAndMsg(rw, http.StatusInternalServerError, ErrObj{Err: responses.AttachDetachFetchError.Error()})
			logger.WithField("err", err.Error()).Error(responses.AttachDetachFetchError)
			return
		}

		logger.Infoln(responses.AttachDetachFetchSuccess)
		responseCodeAndMsg(rw, http.StatusOK, attachDetach)
	})
}

func updateAttachDetachHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		//logging when the api is initialised
		username := req.Context().Value(contextKeyUsername).(string)
		go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.InitialisedState, db.UpdateOperation, "", responses.AttachDetachInitialisedState, username)

		vars := mux.Vars(req)
		id, err := parseUUID(vars["id"])

		// for logging error if there is any otherwise logging success
		defer func() {
			if err != nil {
				go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.ErrorState, db.UpdateOperation, "", err.Error(), username)

			} else {
				go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.CompletedState, db.UpdateOperation, "", responses.AttachDetachCompletedState, username)

			}

		}()

		if err != nil {
			responseCodeAndMsg(rw, http.StatusBadRequest, ErrObj{Err: responses.UUIDParseError.Error()})
			return
		}

		var adObj db.AttachDetach
		err = json.NewDecoder(req.Body).Decode(&adObj)
		if err != nil {
			logger.WithField("err", err.Error()).Errorln(responses.AttachDetachDecodeError)
			responseCodeAndMsg(rw, http.StatusBadRequest, ErrObj{Err: responses.AttachDetachDecodeError.Error()})
			return
		}
		valid, respBytes := validate(adObj)
		if !valid {
			logger.WithField("err", "Validation Error").Errorln( responses.AttachDetachValidationError)
			responseBadRequest(rw, respBytes)
			return
		}
		adObj.ProcessID = id
		err = deps.Store.UpdateAttachDetach(req.Context(), adObj)
		if err != nil {
			logger.WithField("err", err.Error()).Error(responses.AttachDetachUpdateError)
			responseCodeAndMsg(rw, http.StatusInternalServerError, ErrObj{Err: responses.AttachDetachUpdateError.Error()})
			return
		}
		logger.Infoln(responses.AttachDetachUpdateSuccess)
		responseCodeAndMsg(rw, http.StatusOK, MsgObj{Msg: responses.AttachDetachUpdateSuccess})
	})
}
