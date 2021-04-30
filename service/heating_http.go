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
		var htObj db.Heating
		err := json.NewDecoder(req.Body).Decode(&htObj)
		if err != nil {
			logger.WithField("err", err.Error()).Errorln(responses.HeatingDecodeError)
			responseCodeAndMsg(rw, http.StatusBadRequest, ErrObj{Err: responses.HeatingDecodeError.Error()})
			return
		}

		valid, respBytes := validate(htObj)
		if !valid {
			logger.WithField("err", responses.HeatingValidationError)
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
		vars := mux.Vars(req)

		id, err := parseUUID(vars["id"])
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
		vars := mux.Vars(req)
		id, err := parseUUID(vars["id"])
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
			logger.WithField("err", responses.HeatingValidationError)
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
