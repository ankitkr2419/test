package service

import (
	"encoding/json"
	"mylab/cpagent/db"
	"mylab/cpagent/responses"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	logger "github.com/sirupsen/logrus"
)

func createTipTubeHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.InitialisedState, db.CreateOperation, "", responses.TipTubeInitialisedState)

		var tt db.TipsTubes
		err := json.NewDecoder(req.Body).Decode(&tt)

		// for logging error if there is any otherwise logging success
		defer func() {
			if err != nil {
				go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.ErrorState, db.CreateOperation, "", err.Error())

			} else {
				go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.CompletedState, db.CreateOperation, "", responses.TipTubeCompletedState)

			}

		}()

		if err != nil {
			logger.WithField("err", err.Error()).Errorln(responses.TipTubeDecodeError)
			responseCodeAndMsg(rw, http.StatusBadRequest, ErrObj{Err: responses.TipTubeDecodeError.Error()})
			return
		}

		valid, respBytes := Validate(tt)
		if !valid {
			responseBadRequest(rw, respBytes)
			return
		}

		go db.SetTipsTubesValues([]db.TipsTubes{tt})
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
		go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.InitialisedState, db.ShowOperation, "", responses.TipTubeInitialisedState)

		vars := mux.Vars(req)
		tipTubeType := vars["tiptube"]

		var tipsTubes []db.TipsTubes
		var err error

		// for logging error if there is any otherwise logging success
		defer func() {
			if err != nil {
				go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.ErrorState, db.ShowOperation, "", err.Error())

			} else {
				go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.CompletedState, db.ShowOperation, "", responses.TipTubeCompletedState)

			}

		}()

		switch tipTubeType {
		case "tip", "tube", "":
			tipsTubes, err = deps.Store.ListTipsTubes(tipTubeType)
			if err != nil {
				responseCodeAndMsg(rw, http.StatusInternalServerError, ErrObj{Err: responses.TipTubeFetchError.Error()})
				logger.WithField("err", err.Error()).Error(responses.TipTubeFetchError)
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

func listTipsTubesPositionHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		//logging when the api is initialised
		go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.InitialisedState, db.ShowOperation, "", responses.TipTubeInitialisedState)

		vars := mux.Vars(req)
		tipTubeType := vars["tiptube"]
		position, err := strconv.ParseInt(vars["position"], 10, 64)
		var tipsTubes []db.TipsTubes

		// for logging error if there is any otherwise logging success
		defer func() {
			if err != nil {
				go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.ErrorState, db.ShowOperation, "", err.Error())

			} else {
				go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.CompletedState, db.ShowOperation, "", responses.TipTubeCompletedState)

			}

		}()

		switch tipTubeType {
		case "tip", "tube", "":
			tipsTubes, err = deps.Store.ListTipsTubesByPosition(req.Context(), tipTubeType, position)
			if err != nil {
				responseCodeAndMsg(rw, http.StatusInternalServerError, ErrObj{Err: responses.TipTubeFetchError.Error()})
				logger.WithField("err", err.Error()).Error(responses.TipTubeFetchError)
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

func deleteTipTubeHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		id, err := strconv.ParseInt(vars["id"], 10, 64)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			return
		}

		err = deps.Store.DeleteTipTube(req.Context(), id)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			logger.WithField("err", err.Error()).Error("Error while deleting TipTube")
			return
		}
		response := MsgObj{
			Msg: "TipTube deleted successfully",
		}
		responseCodeAndMsg(rw, http.StatusOK, response)

	})
}
