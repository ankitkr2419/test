package service

import (
	"encoding/json"
	"mylab/cpagent/db"
	"mylab/cpagent/responses"
	"net/http"

	"github.com/gorilla/mux"
	logger "github.com/sirupsen/logrus"
)

func createDelayHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		var delay db.Delay
		err := json.NewDecoder(req.Body).Decode(&delay)
		if err != nil {
			logger.WithField("err", err.Error()).Errorln(responses.DelayDecodeError)
			responseCodeAndMsg(rw, http.StatusBadRequest, ErrObj{Err: responses.DelayDecodeError.Error()})
			return
		}

		valid, respBytes := validate(delay)
		if !valid {
			logger.WithField("err", err.Error()).Errorln(responses.DelayValidationError)
			responseBadRequest(rw, respBytes)
			return
		}

		var createdTemp db.Delay
		createdTemp, err = deps.Store.CreateDelay(req.Context(), delay)
		if err != nil {
			logger.WithField("err", err.Error()).Errorln(responses.DelayCreateError)
			responseCodeAndMsg(rw, http.StatusInternalServerError, ErrObj{Err: responses.DelayCreateError.Error()})
			return
		}

		logger.Infoln(responses.DelayCreateSuccess)
		responseCodeAndMsg(rw, http.StatusCreated, createdTemp)
	})
}

func showDelayHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)

		id, err := parseUUID(vars["id"])
		if err != nil {
			// This error is already logged
			responseCodeAndMsg(rw, http.StatusBadRequest, ErrObj{Err: responses.UUIDParseError.Error()})
			return
		}

		var delay db.Delay

		delay, err = deps.Store.ShowDelay(req.Context(), id)
		if err != nil {
			responseCodeAndMsg(rw, http.StatusInternalServerError, ErrObj{Err: responses.DelayFetchError.Error()})
			logger.WithField("err", err.Error()).Error(responses.DelayFetchError)
			return
		}

		logger.Infoln(responses.DelayFetchSuccess)
		responseCodeAndMsg(rw, http.StatusOK, delay)
	})
}

func updateDelayHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		id, err := parseUUID(vars["id"])
		if err != nil {
			responseCodeAndMsg(rw, http.StatusBadRequest, ErrObj{Err: responses.UUIDParseError.Error()})
			return
		}

		var delay db.Delay

		err = json.NewDecoder(req.Body).Decode(&delay)
		if err != nil {
			logger.WithField("err", err.Error()).Errorln(responses.DelayDecodeError)
			responseCodeAndMsg(rw, http.StatusBadRequest, ErrObj{Err: responses.DelayDecodeError.Error()})
			return
		}

		valid, respBytes := validate(delay)
		if !valid {
			logger.WithField("err", err.Error()).Errorln(responses.DelayValidationError)
			responseBadRequest(rw, respBytes)
			return
		}

		delay.ProcessID = id
		err = deps.Store.UpdateDelay(req.Context(), delay)
		if err != nil {
			logger.WithField("err", err.Error()).Error(responses.DelayUpdateError)
			responseCodeAndMsg(rw, http.StatusInternalServerError, ErrObj{Err: responses.DelayUpdateError.Error()})
			return
		}

		logger.Infoln(responses.DelayUpdateSuccess)
		responseCodeAndMsg(rw, http.StatusOK, MsgObj{Msg: responses.DelayUpdateSuccess})
	})
}
