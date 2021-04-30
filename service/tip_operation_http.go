package service

import (
	"encoding/json"
	"mylab/cpagent/db"
	"mylab/cpagent/responses"
	"net/http"

	"github.com/gorilla/mux"
	logger "github.com/sirupsen/logrus"
)

func createTipOperationHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		var tipOpr db.TipOperation
		err := json.NewDecoder(req.Body).Decode(&tipOpr)
		if err != nil {
			logger.WithField("err", err.Error()).Errorln(responses.TipOperationDecodeError)
			responseCodeAndMsg(rw, http.StatusBadRequest, ErrObj{Err: responses.TipOperationDecodeError.Error()})
			return
		}

		valid, respBytes := validate(tipOpr)
		if !valid {
			logger.WithField("err", responses.TipOperationValidationError)
			responseBadRequest(rw, respBytes)
			return
		}

		var tipOperation db.TipOperation
		tipOperation, err = deps.Store.CreateTipOperation(req.Context(), tipOpr)
		if err != nil {
			logger.WithField("err", err.Error()).Errorln(responses.TipOperationCreateError)
			responseCodeAndMsg(rw, http.StatusInternalServerError, ErrObj{Err: responses.TipOperationCreateError.Error()})
			return
		}
		logger.Infoln(responses.TipOperationCreateSuccess)
		responseCodeAndMsg(rw, http.StatusCreated, tipOperation)
	})
}

func showTipOperationHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)

		id, err := parseUUID(vars["id"])
		if err != nil {
			responseCodeAndMsg(rw, http.StatusBadRequest, ErrObj{Err: responses.UUIDParseError.Error()})
			return
		}

		var tipOperation db.TipOperation

		tipOperation, err = deps.Store.ShowTipOperation(req.Context(), id)
		if err != nil {
			responseCodeAndMsg(rw, http.StatusInternalServerError, ErrObj{Err: responses.TipOperationFetchError.Error()})
			logger.WithField("err", err.Error()).Error(responses.TipOperationFetchError)
			return
		}

		logger.Infoln(responses.TipOperationFetchSuccess)
		responseCodeAndMsg(rw, http.StatusOK, tipOperation)
	})
}

func updateTipOperationHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		id, err := parseUUID(vars["id"])
		if err != nil {
			responseCodeAndMsg(rw, http.StatusBadRequest, ErrObj{Err: responses.UUIDParseError.Error()})
			return
		}

		var tipOpr db.TipOperation

		err = json.NewDecoder(req.Body).Decode(&tipOpr)
		if err != nil {
			logger.WithField("err", err.Error()).Errorln(responses.TipOperationDecodeError)
			responseCodeAndMsg(rw, http.StatusBadRequest, ErrObj{Err: responses.TipOperationDecodeError.Error()})
			return
		}

		valid, respBytes := validate(tipOpr)
		if !valid {
			logger.WithField("err", responses.TipOperationValidationError)
			responseBadRequest(rw, respBytes)
			return
		}

		tipOpr.ProcessID = id
		err = deps.Store.UpdateTipOperation(req.Context(), tipOpr)
		if err != nil {
			logger.WithField("err", err.Error()).Error(responses.TipOperationUpdateError)
			responseCodeAndMsg(rw, http.StatusInternalServerError, ErrObj{Err: responses.TipOperationUpdateError.Error()})
			return
		}

		logger.Infoln(responses.TipOperationUpdateSuccess)
		responseCodeAndMsg(rw, http.StatusOK, MsgObj{Msg: responses.TipOperationUpdateSuccess})
	})
}
