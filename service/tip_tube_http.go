package service

import (
	"encoding/json"
	"mylab/cpagent/db"
	"mylab/cpagent/responses"
	"net/http"

	"github.com/gorilla/mux"
	logger "github.com/sirupsen/logrus"
)

func createTipTubeHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		username := req.Context().Value(contextKeyUsername).(string)
		go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.InitialisedState, db.CreateOperation, "", responses.TipTubeInitialisedState, username)

		var tt db.TipsTubes
		err := json.NewDecoder(req.Body).Decode(&tt)

		// for logging error if there is any otherwise logging success
		defer func() {
			if err != nil {
				go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.ErrorState, db.CreateOperation, "", err.Error(), username)

			} else {
				go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.CompletedState, db.CreateOperation, "", responses.TipTubeCompletedState, username)

			}

		}()

		if err != nil {
			logger.WithField("err", err.Error()).Errorln(responses.TipTubeDecodeError)
			responseCodeAndMsg(rw, http.StatusBadRequest, ErrObj{Err: responses.TipTubeDecodeError.Error()})
			return
		}

		valid, respBytes := validate(tt)
		if !valid {
			responseBadRequest(rw, respBytes)
			return
		}

		err = deps.Store.InsertTipsTubes(req.Context(), []db.TipsTubes{tt})
		if err != nil {
			logger.WithField("err", err.Error()).Errorln(responses.TipTubeCreateError)
			responseCodeAndMsg(rw, http.StatusInternalServerError, ErrObj{Err: responses.TipTubeCreateError.Error()})
			return
		}

		logger.Infoln(responses.TipTubeCreateSuccess)
		responseCodeAndMsg(rw, http.StatusCreated, tt)
	})
}

func listTipsTubesHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		//logging when the api is initialised
		username := req.Context().Value(contextKeyUsername).(string)
		go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.InitialisedState, db.ShowOperation, "", responses.TipTubeInitialisedState, username)

		vars := mux.Vars(req)
		tipTubeType := vars["tiptube"]

		var tipsTubes []db.TipsTubes
		var err error

		// for logging error if there is any otherwise logging success
		defer func() {
			if err != nil {
				go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.ErrorState, db.ShowOperation, "", err.Error(), username)

			} else {
				go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.CompletedState, db.ShowOperation, "", responses.TipTubeCompletedState, username)

			}

		}()

		switch tipTubeType {
		case "tip", "tube", "":
			tipsTubes, err = deps.Store.ListTipsTubes(tipTubeType)
			if err != nil {
				responseCodeAndMsg(rw, http.StatusInternalServerError, ErrObj{Err: responses.TipTubeFetchError.Error()})
				logger.WithField("err", err.Error()).Error(responses.TipTubeFetchError)
				return
				return
			}
		default:
			responseCodeAndMsg(rw, http.StatusBadRequest, ErrObj{Err: responses.TipTubeArgumentsError.Error()})
			return
		}

		logger.Infoln(responses.TipTubeFetchSuccess)
		responseCodeAndMsg(rw, http.StatusOK, tipsTubes)
	})
}
