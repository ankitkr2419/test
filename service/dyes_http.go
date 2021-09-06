package service

import (
	"encoding/json"
	"mylab/cpagent/db"
	"mylab/cpagent/responses"
	"net/http"

	logger "github.com/sirupsen/logrus"
)

func updateDyeToleranceHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		var cd []db.Dye
		err := json.NewDecoder(req.Body).Decode(&cd)
		if err != nil {
			responseCodeAndMsg(rw, http.StatusBadRequest, ErrObj{Err: responses.DyeDecodeError.Error()})
			logger.WithField("err", err.Error()).Error(responses.DyeDecodeError)
			return
		}

		go db.UpdateDyesTolerance(cd)

		dyes, err := deps.Store.InsertDyes(req.Context(), cd)
		if err != nil {
			responseCodeAndMsg(rw, http.StatusInternalServerError, ErrObj{Err: responses.DyeInsertError.Error()})
			logger.WithField("err", err.Error()).Error(responses.DyeInsertError)
			return
		}

		logger.Infoln(responses.DyeCreateSuccess)
		responseCodeAndMsg(rw, http.StatusCreated, dyes)
	})
}

func listDyesHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		e, err := deps.Store.ListDyes(req.Context())
		if err != nil {
			responseCodeAndMsg(rw, http.StatusInternalServerError, ErrObj{Err: responses.DyeFetchError.Error()})
			logger.WithField("err", err.Error()).Error(responses.DyeFetchError)
			return
		}

		logger.Infoln(responses.DyeCreateSuccess)
		responseCodeAndMsg(rw, http.StatusCreated, e)
	})
}
