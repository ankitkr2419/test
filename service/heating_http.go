package service

import (
	"encoding/json"
	"mylab/cpagent/db"
	"mylab/cpagent/responses"
	"net/http"

	"github.com/gorilla/mux"
	logger "github.com/sirupsen/logrus"
)

func createHeatingHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		username := req.Context().Value(contextKeyUsername).(string)
		go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.InitialisedState, db.CreateOperation, "", responses.HeatingInitialisedState, username)

		var htObj db.Heating
		err := json.NewDecoder(req.Body).Decode(&htObj)

		// for logging error if there is any otherwise logging success
		defer func() {
			if err != nil {
				go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.ErrorState, db.CreateOperation, "", err.Error(), username)

			} else {
				go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.CompletedState, db.CreateOperation, "", responses.HeatingCompletedState, username)

			}

		}()

		if err != nil {
			logger.WithField("err", err.Error()).Errorln(responses.HeatingDecodeError)
			responseCodeAndMsg(rw, http.StatusBadRequest, ErrObj{Err: responses.HeatingDecodeError.Error()})
			return
		}

		valid, respBytes := validate(htObj)
		if !valid {
			logger.WithField("err", "Validation Error").Errorln(responses.HeatingValidationError)
			responseBadRequest(rw, respBytes)
			return
		}

		var createdTemp db.Heating
		createdTemp, err = deps.Store.CreateHeating(req.Context(), htObj)
		if err != nil {
			logger.WithField("err", err.Error()).Errorln(responses.HeatingCreateError)
			responseCodeAndMsg(rw, http.StatusInternalServerError, ErrObj{Err: responses.HeatingCreateError.Error()})
			return
		}
		logger.Infoln(responses.HeatingCreateSuccess)
		responseCodeAndMsg(rw, http.StatusCreated, createdTemp)
	})
}

func showHeatingHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		//logging when the api is initialised
		username := req.Context().Value(contextKeyUsername).(string)
		go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.InitialisedState, db.ShowOperation, "", responses.HeatingInitialisedState, username)

		vars := mux.Vars(req)

		id, err := parseUUID(vars["id"])

		// for logging error if there is any otherwise logging success
		defer func() {
			if err != nil {
				go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.ErrorState, db.ShowOperation, "", err.Error(), username)

			} else {
				go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.CompletedState, db.ShowOperation, "", responses.HeatingCompletedState, username)

			}

		}()

		if err != nil {
			responseCodeAndMsg(rw, http.StatusBadRequest, ErrObj{Err: responses.UUIDParseError.Error()})
			return
		}

		var heating db.Heating

		heating, err = deps.Store.ShowHeating(req.Context(), id)
		if err != nil {
			responseCodeAndMsg(rw, http.StatusInternalServerError, ErrObj{Err: responses.HeatingFetchError.Error()})
			logger.WithField("err", err.Error()).Error(responses.HeatingFetchError)
			return
		}

		logger.Infoln(responses.HeatingFetchSuccess)
		responseCodeAndMsg(rw, http.StatusOK, heating)
	})
}

func updateHeatingHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		username := req.Context().Value(contextKeyUsername).(string)
		go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.InitialisedState, db.UpdateOperation, "", responses.HeatingInitialisedState, username)

		vars := mux.Vars(req)
		id, err := parseUUID(vars["id"])

		// for logging error if there is any otherwise logging success
		defer func() {
			if err != nil {
				go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.ErrorState, db.UpdateOperation, "", err.Error(), username)

			} else {
				go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.CompletedState, db.UpdateOperation, "", responses.HeatingCompletedState, username)

			}

		}()

		if err != nil {
			responseCodeAndMsg(rw, http.StatusBadRequest, ErrObj{Err: responses.UUIDParseError.Error()})
			return
		}

		var htObj db.Heating

		err = json.NewDecoder(req.Body).Decode(&htObj)
		if err != nil {
			logger.WithField("err", err.Error()).Errorln(responses.HeatingDecodeError)
			responseCodeAndMsg(rw, http.StatusBadRequest, ErrObj{Err: responses.HeatingDecodeError.Error()})
			return
		}

		valid, respBytes := validate(htObj)
		if !valid {
			logger.WithField("err", "Validation Error").Errorln(responses.HeatingValidationError)
			responseBadRequest(rw, respBytes)
			return
		}

		htObj.ProcessID = id
		err = deps.Store.UpdateHeating(req.Context(), htObj)
		if err != nil {
			logger.WithField("err", err.Error()).Error(responses.HeatingUpdateError)
			responseCodeAndMsg(rw, http.StatusInternalServerError, ErrObj{Err: responses.HeatingUpdateError.Error()})
			return
		}

		logger.Infoln(responses.HeatingUpdateSuccess)
		responseCodeAndMsg(rw, http.StatusOK, MsgObj{Msg: responses.HeatingUpdateSuccess})
	})
}
