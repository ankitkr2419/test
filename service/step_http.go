package service

import (
	"encoding/json"
	"mylab/cpagent/db"
	"net/http"

	"github.com/gorilla/mux"
	logger "github.com/sirupsen/logrus"
)

func listStepsHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		stage_id, err := parseUUID(vars["stage_id"])
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			return
		}
		t, err := deps.Store.ListSteps(req.Context(), stage_id)
		if err != nil {
			logger.WithField("err", err.Error()).Error("Error fetching data")
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}

		respBytes, err := json.Marshal(t)
		if err != nil {
			logger.WithField("err", err.Error()).Error("Error marshaling steps data")
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}

		rw.Header().Add("Content-Type", "application/json")
		rw.WriteHeader(http.StatusOK)
		rw.Write(respBytes)
	})
}

// @Title createStepHandler
// @Description Create createStepHandler
// @Router /step [post]
// @Accept  json
// @Success 200 {object}
// @Failure 400 {object}
func createStepHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		var t db.Step
		err := json.NewDecoder(req.Body).Decode(&t)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			logger.WithField("err", err.Error()).Error("Error while decoding step data")
			return
		}

		valid, respBytes := validate(t)
		if !valid {
			responseBadRequest(rw, respBytes)
			return
		}

		var createdTemp db.Step
		createdTemp, err = deps.Store.CreateStep(req.Context(), t)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			logger.WithField("err", err.Error()).Error("Error create target")
			return
		}

		respBytes, err = json.Marshal(createdTemp)
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

func updateStepHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		id, err := parseUUID(vars["id"])
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			return
		}

		var t db.Step

		err = json.NewDecoder(req.Body).Decode(&t)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			logger.WithField("err", err.Error()).Error("Error while decoding step data")
			return
		}

		valid, respBytes := validate(t)
		if !valid {
			responseBadRequest(rw, respBytes)
			return
		}

		t.ID = id

		err = deps.Store.UpdateStep(req.Context(), t)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			logger.WithField("err", err.Error()).Error("Error update step")
			return
		}

		rw.WriteHeader(http.StatusOK)
		rw.Header().Add("Content-Type", "application/json")
		rw.Write([]byte(`{"msg":"step updated successfully"}`))
	})
}

func showStepHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)

		id, err := parseUUID(vars["id"])
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			return
		}

		var latestT db.Step

		latestT, err = deps.Store.ShowStep(req.Context(), id)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			logger.WithField("err", err.Error()).Error("Error show step")
			return
		}

		respBytes, err := json.Marshal(latestT)
		if err != nil {
			logger.WithField("err", err.Error()).Error("Error marshaling step data")
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}

		rw.WriteHeader(http.StatusOK)
		rw.Write(respBytes)
		rw.Header().Add("Content-Type", "application/json")
	})
}

func deleteStepHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		id, err := parseUUID(vars["id"])
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			return
		}

		err = deps.Store.DeleteStep(req.Context(), id)
		if err != nil {
			logger.WithField("err", err.Error()).Error("Error while deleting step")
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}

		rw.WriteHeader(http.StatusOK)
		rw.Header().Add("Content-Type", "application/json")
		rw.Write([]byte(`{"msg":"step deleted successfully"}`))
	})
}
