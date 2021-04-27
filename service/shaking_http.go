package service

import (
	"encoding/json"
	"mylab/cpagent/db"
	"net/http"

	"github.com/gorilla/mux"
	logger "github.com/sirupsen/logrus"
)

func createShakingHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		var shaObj db.Shaker
		err := json.NewDecoder(req.Body).Decode(&shaObj)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			logger.WithField("err", err.Error()).Error("Error while decoding shaking data")
			return
		}

		valid, respBytes := validate(shaObj)
		if !valid {
			responseBadRequest(rw, respBytes)
			return
		}

		err = updateProcessName(deps, shaObj.ProcessID, "Shaking", shaObj)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			logger.WithField("err", err.Error()).Error("Error updating process name")
			return
		}

		var createdShaker db.Shaker
		createdShaker, err = deps.Store.CreateShaking(req.Context(), shaObj)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			logger.WithField("err", err.Error()).Error("Error create shaking")
			return
		}

		respBytes, err = json.Marshal(createdShaker)
		if err != nil {
			logger.WithField("err", err.Error()).Error("Error marshaling shaking data")
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}

		rw.Header().Add("Content-Type", "application/json")
		rw.WriteHeader(http.StatusCreated)
		rw.Write(respBytes)
	})
}

func showShakingHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		id, err := parseUUID(vars["id"])
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			return
		}

		var shaking db.Shaker
		shaking, err = deps.Store.ShowShaking(req.Context(), id)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			logger.WithField("err", err.Error()).Error("Error show shaking")
			return
		}

		respBytes, err := json.Marshal(shaking)
		if err != nil {
			logger.WithField("err", err.Error()).Error("Error marshaling shaking data")
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}
		rw.Header().Add("Content-Type", "application/json")
		rw.WriteHeader(http.StatusOK)
		rw.Write(respBytes)
	})
}

func updateShakingHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		id, err := parseUUID(vars["id"])
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			return
		}
		var shObj db.Shaker
		err = json.NewDecoder(req.Body).Decode(&shObj)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			logger.WithField("err", err.Error()).Error("Error while decoding shaking data")
			return
		}
		valid, respBytes := validate(shObj)
		if !valid {
			responseBadRequest(rw, respBytes)
			return
		}
		shObj.ProcessID = id
		err = deps.Store.UpdateShaking(req.Context(), shObj)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			logger.WithField("err", err.Error()).Error("Error update shaking")
			return
		}
		rw.WriteHeader(http.StatusOK)
		rw.Header().Add("Content-Type", "application/json")
		rw.Write([]byte(`{"msg":"shaking record updated successfully"}`))
	})
}
