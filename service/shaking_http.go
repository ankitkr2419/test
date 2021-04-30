package service

import (
	"encoding/json"
	"mylab/cpagent/db"
	"mylab/cpagent/responses"
	"net/http"

	"github.com/gorilla/mux"
	logger "github.com/sirupsen/logrus"
)

func createShakingHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		var shaObj db.Shaker
		err := json.NewDecoder(req.Body).Decode(&shaObj)
		if err != nil {
			logger.WithField("err", err.Error()).Errorln(responses.ShakingDecodeError)
			responseCodeAndMsg(rw, http.StatusBadRequest, ErrObj{Err: responses.ShakingDecodeError.Error()})
			return
		}

		valid, respBytes := validate(shaObj)
		if !valid {
			logger.WithField("err", responses.ShakingValidationError)
			responseBadRequest(rw, respBytes)
			return
		}

		var shaker db.Shaker
		shaker, err = deps.Store.CreateShaking(req.Context(), shaObj)
		if err != nil {
			logger.WithField("err", err.Error()).Errorln(responses.ShakingCreateError)
			responseCodeAndMsg(rw, http.StatusInternalServerError, ErrObj{Err: responses.ShakingCreateError.Error()})
			return
		}
		logger.Infoln(responses.ShakingCreateSuccess)
		responseCodeAndMsg(rw, http.StatusCreated, shaker)
	})
}

func showShakingHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		id, err := parseUUID(vars["id"])
		if err != nil {
			responseCodeAndMsg(rw, http.StatusBadRequest, ErrObj{Err: responses.UUIDParseError.Error()})
			return
		}

		var shaking db.Shaker
		shaking, err = deps.Store.ShowShaking(req.Context(), id)
		if err != nil {
			responseCodeAndMsg(rw, http.StatusInternalServerError, ErrObj{Err: responses.ShakingFetchError.Error()})
			logger.WithField("err", err.Error()).Error(responses.ShakingFetchError)
			return
		}

		logger.Infoln(responses.ShakingFetchSuccess)
		responseCodeAndMsg(rw, http.StatusOK, shaking)
	})
}

func updateShakingHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		id, err := parseUUID(vars["id"])
		if err != nil {
			responseCodeAndMsg(rw, http.StatusBadRequest, ErrObj{Err: responses.UUIDParseError.Error()})
			return
		}
		var shObj db.Shaker
		err = json.NewDecoder(req.Body).Decode(&shObj)
		if err != nil {
			logger.WithField("err", err.Error()).Errorln(responses.ShakingDecodeError)
			responseCodeAndMsg(rw, http.StatusBadRequest, ErrObj{Err: responses.ShakingDecodeError.Error()})
			return
		}
		valid, respBytes := validate(shObj)
		if !valid {
			logger.WithField("err", responses.ShakingValidationError)
			responseBadRequest(rw, respBytes)
			return
		}
		shObj.ProcessID = id
		err = deps.Store.UpdateShaking(req.Context(), shObj)
		if err != nil {
			logger.WithField("err", err.Error()).Error(responses.ShakingUpdateError)
			responseCodeAndMsg(rw, http.StatusInternalServerError, ErrObj{Err: responses.ShakingUpdateError.Error()})
			return
		}

		logger.Infoln(responses.ShakingUpdateSuccess)
		responseCodeAndMsg(rw, http.StatusOK, MsgObj{Msg: responses.ShakingUpdateSuccess})
	})
}
