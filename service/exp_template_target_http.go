package service

import (
	"encoding/json"
	"mylab/cpagent/db"
	"net/http"

	"github.com/gorilla/mux"
	logger "github.com/sirupsen/logrus"
)

func listExpTempTargetsHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)

		//TBD find alternative option to parsing UUID
		expID, err := parseUUID(vars["experiment_id"])
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			return
		}
		t, err := deps.Store.ListExpTemplateTargets(req.Context(), expID)
		if err != nil {
			logger.WithField("err", err.Error()).Error("Error fetching data")
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}

		respBytes, err := json.Marshal(t)
		if err != nil {
			logger.WithField("err", err.Error()).Error("Error marshaling experiment template targets data")
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}

		rw.Header().Add("Content-Type", "application/json")
		rw.WriteHeader(http.StatusOK)
		rw.Write(respBytes)
	})
}

func updateExpTempTargetsHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		expID, err := parseUUID(vars["experiment_id"])
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			return
		}
		var t []db.ExpTemplateTarget
		err = json.NewDecoder(req.Body).Decode(&t)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			logger.WithField("err", err.Error()).Error("Error while decoding template target data")
			return
		}

		for _, tt := range t {
			valid, respBytes := Validate(tt)
			if !valid {
				responseBadRequest(rw, respBytes)
				return
			}
		}

		var createdET []db.ExpTempTargeTDye
		createdET, err = deps.Store.UpsertExpTemplateTarget(req.Context(), t, expID)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			logger.WithField("err", err.Error()).Error("Error create target")
			return
		}

		respBytes, err := json.Marshal(createdET)
		if err != nil {
			logger.WithField("err", err.Error()).Error("Error marshaling experiment targets data")
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}
		rw.WriteHeader(http.StatusCreated)
		rw.Write(respBytes)
		rw.Header().Add("Content-Type", "application/json")
	})
}
