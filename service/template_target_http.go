package service

import (
	"encoding/json"
	"mylab/cpagent/db"
	"net/http"
	"fmt"
	"github.com/gorilla/mux"
	logger "github.com/sirupsen/logrus"
)

func listTempTargetsHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		tempID, err := parseUUID(vars["template_id"])
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			return
		}
		t, err := deps.Store.ListTemplateTargets(req.Context(), tempID)
		if err != nil {
			logger.WithField("err", err.Error()).Error("Error fetching data")
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}

		respBytes, err := json.Marshal(t)
		if err != nil {
			logger.WithField("err", err.Error()).Error("Error marshaling template targets data")
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}

		rw.Header().Add("Content-Type", "application/json")
		rw.WriteHeader(http.StatusOK)
		rw.Write(respBytes)
	})
}

// @Title createTemplateHandler
// @Description Create createTemplateHandler
// @Router /template [post]
// @Accept  json
// @Success 200 {object}
// @Failure 400 {object}
func updateTempTargetsHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		tempID, err := parseUUID(vars["template_id"])
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			return
		}
		var t []db.TemplateTarget
		err = json.NewDecoder(req.Body).Decode(&t)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			logger.WithField("err", err.Error()).Error("Error while decoding template target data")
			return
		}
		fmt.Println(t)
		for _, tt := range t {
			valid, respBytes := validate(tt)
			if !valid {
				rw.Header().Add("Content-Type", "application/json")
				rw.WriteHeader(http.StatusBadRequest)
				rw.Write(respBytes)
				return
			}
		}

		var createdTemp []db.TemplateTarget
		createdTemp, err = deps.Store.UpsertTemplateTarget(req.Context(), t, tempID)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			logger.WithField("err", err.Error()).Error("Error create target")
			return
		}

		respBytes, err := json.Marshal(createdTemp)
		if err != nil {
			logger.WithField("err", err.Error()).Error("Error marshaling targets data")
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}

		rw.WriteHeader(http.StatusCreated)
		rw.Write(respBytes)
		rw.Header().Add("Content-Type", "application/json")
	})
}
