package service

import (
	"encoding/json"
	"mylab/cpagent/db"
	"mylab/cpagent/responses"
	"net/http"

	"github.com/gorilla/mux"
	logger "github.com/sirupsen/logrus"
)

func createHeatingHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		username := req.Context().Value(contextKeyUsername).(string)
		go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.InitialisedState, db.CreateOperation, "", responses.HeatingInitialisedState, username)

		var htObj db.Heating
		err := json.NewDecoder(req.Body).Decode(&htObj)

		// for logging error if there is any otherwise logging success
		defer func() {
			if err != nil {
				go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.ErrorState, db.CreateOperation, "", err.Error(), username)

			} else {
				go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.CompletedState, db.CreateOperation, "", responses.HeatingCompletedState, username)

			}

		}()

		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			logger.WithField("err", err.Error()).Error("Error while decoding heating data")
			return
		}

		valid, respBytes := validate(htObj)
		if !valid {
			responseBadRequest(rw, respBytes)
			return
		}

		var createdTemp db.Heating
		createdTemp, err = deps.Store.CreateHeating(req.Context(), htObj)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			logger.WithField("err", err.Error()).Error("Error create Heating")
			return
		}

		respBytes, err = json.Marshal(createdTemp)
		if err != nil {
			logger.WithField("err", err.Error()).Error("Error marshaling heating data")
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}

		rw.Header().Add("Content-Type", "application/json")
		rw.WriteHeader(http.StatusCreated)
		rw.Write(respBytes)
	})
}

func showHeatingHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		//logging when the api is initialised
		username := req.Context().Value(contextKeyUsername).(string)
		go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.InitialisedState, db.ShowOperation, "", responses.HeatingInitialisedState, username)

		vars := mux.Vars(req)

		id, err := parseUUID(vars["id"])

		// for logging error if there is any otherwise logging success
		defer func() {
			if err != nil {
				go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.ErrorState, db.ShowOperation, "", err.Error(), username)

			} else {
				go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.CompletedState, db.ShowOperation, "", responses.HeatingCompletedState, username)

			}

		}()

		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			return
		}

		var heating db.Heating

		heating, err = deps.Store.ShowHeating(req.Context(), id)
		if err != nil {
			rw.WriteHeader(http.StatusNotFound)
			logger.WithField("err", err.Error()).Error("Error show heating")
			return
		}

		respBytes, err := json.Marshal(heating)
		if err != nil {
			logger.WithField("err", err.Error()).Error("Error marshaling heating data")
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}

		rw.Header().Add("Content-Type", "application/json")
		rw.WriteHeader(http.StatusOK)
		rw.Write(respBytes)
	})
}

func updateHeatingHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		username := req.Context().Value(contextKeyUsername).(string)
		go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.InitialisedState, db.UpdateOperation, "", responses.HeatingInitialisedState, username)

		vars := mux.Vars(req)
		id, err := parseUUID(vars["id"])

		// for logging error if there is any otherwise logging success
		defer func() {
			if err != nil {
				go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.ErrorState, db.UpdateOperation, "", err.Error(), username)

			} else {
				go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.CompletedState, db.UpdateOperation, "", responses.HeatingCompletedState, username)

			}

		}()

		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			return
		}

		var htObj db.Heating

		err = json.NewDecoder(req.Body).Decode(&htObj)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			logger.WithField("err", err.Error()).Error("Error while decoding piercing data")
			return
		}

		valid, respBytes := validate(htObj)
		if !valid {
			responseBadRequest(rw, respBytes)
			return
		}

		htObj.ProcessID = id
		err = deps.Store.UpdateHeating(req.Context(), htObj)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			logger.WithField("err", err.Error()).Error("Error update piercing")
			return
		}

		rw.WriteHeader(http.StatusOK)
		rw.Header().Add("Content-Type", "application/json")
		rw.Write([]byte(`{"msg":"heating record updated successfully"}`))
	})
}
