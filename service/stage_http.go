package service

import (
	"encoding/json"
	"mylab/cpagent/db"
	"net/http"

	"github.com/gorilla/mux"
	logger "github.com/sirupsen/logrus"
)

func listStagesHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		template_id, err := parseUUID(vars["template_id"])
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			return
		}
		t, err := deps.Store.ListStages(req.Context(), template_id)
		if err != nil {
			logger.WithField("err", err.Error()).Error("Error fetching data")
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}

		respBytes, err := json.Marshal(t)
		if err != nil {
			logger.WithField("err", err.Error()).Error("Error marshaling stages data")
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}

		rw.Header().Add("Content-Type", "application/json")
		rw.WriteHeader(http.StatusOK)
		rw.Write(respBytes)
	})
}

// @Title createStageHandler
// @Description Create createStageHandler
// @Router /stage [post]
// @Accept  json
// @Success 200 {object}
// @Failure 400 {object}
func createStageHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		var t db.Stage
		err := json.NewDecoder(req.Body).Decode(&t)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			logger.WithField("err", err.Error()).Error("Error while decoding stage data")
			return
		}

		valid, respBytes := validate(t)
		if !valid {
			rw.Header().Add("Content-Type", "application/json")
			rw.WriteHeader(http.StatusBadRequest)
			rw.Write(respBytes)
			return
		}

		var createdTemp db.Stage
		createdTemp, err = deps.Store.CreateStage(req.Context(), t)
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

func updateStageHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		id, err := parseUUID(vars["id"])
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			return
		}

		var t db.Stage

		err = json.NewDecoder(req.Body).Decode(&t)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			logger.WithField("err", err.Error()).Error("Error while decoding stage data")
			return
		}

		valid, respBytes := validate(t)
		if !valid {
			rw.Header().Add("Content-Type", "application/json")
			rw.WriteHeader(http.StatusBadRequest)
			rw.Write(respBytes)
			return
		}

		t.ID = id

		err = deps.Store.UpdateStage(req.Context(), t)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			logger.WithField("err", err.Error()).Error("Error update stage")
			return
		}

		rw.WriteHeader(http.StatusOK)
		rw.Header().Add("Content-Type", "application/json")
		rw.Write([]byte("stage updated successfully"))
	})
}

func showStageHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)

		id, err := parseUUID(vars["id"])
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			return
		}

		var latestT db.Stage

		latestT, err = deps.Store.ShowStage(req.Context(), id)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			logger.WithField("err", err.Error()).Error("Error show stage")
			return
		}

		respBytes, err := json.Marshal(latestT)
		if err != nil {
			logger.WithField("err", err.Error()).Error("Error marshaling stage data")
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}

		rw.WriteHeader(http.StatusOK)
		rw.Write(respBytes)
		rw.Header().Add("Content-Type", "application/json")
	})
}

func deleteStageHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		id, err := parseUUID(vars["id"])
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			return
		}

		err = deps.Store.DeleteStage(req.Context(), id)
		if err != nil {
			logger.WithField("err", err.Error()).Error("Error while deleting stage")
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}

		rw.WriteHeader(http.StatusOK)
		rw.Header().Add("Content-Type", "application/json")
	})
}
