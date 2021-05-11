package service

import (
	"encoding/json"
	"mylab/cpagent/db"
	"mylab/cpagent/responses"
	"net/http"

	"github.com/gorilla/mux"
	logger "github.com/sirupsen/logrus"
)

func createDelayHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		username := req.Context().Value(contextKeyUsername).(string)
		go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.InitialisedState, db.CreateOperation, "", responses.DelayInitialisedState, username)

		var delay db.Delay
		err := json.NewDecoder(req.Body).Decode(&delay)

		// for logging error if there is any otherwise logging success
		defer func() {
			if err != nil {
				go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.ErrorState, db.CreateOperation, "", err.Error(), username)

			} else {
				go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.CompletedState, db.CreateOperation, "", responses.DelayCompletedState, username)

			}

		}()

		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			logger.WithField("err", err.Error()).Error("Error while decoding Delay data")
			return
		}

		valid, respBytes := validate(delay)
		if !valid {
			responseBadRequest(rw, respBytes)
			return
		}

		var createdTemp db.Delay
		createdTemp, err = deps.Store.CreateDelay(req.Context(), delay)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			logger.WithField("err", err.Error()).Error("Error create Delay")
			return
		}

		respBytes, err = json.Marshal(createdTemp)
		if err != nil {
			logger.WithField("err", err.Error()).Error("Error marshaling Delay data")
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}

		rw.Header().Add("Content-Type", "application/json")
		rw.WriteHeader(http.StatusCreated)
		rw.Write(respBytes)
	})
}

func showDelayHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		//logging when the api is initialised
		username := req.Context().Value(contextKeyUsername).(string)
		go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.InitialisedState, db.ShowOperation, "", responses.DelayInitialisedState, username)

		vars := mux.Vars(req)

		id, err := parseUUID(vars["id"])

		// for logging error if there is any otherwise logging success
		defer func() {
			if err != nil {
				go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.ErrorState, db.ShowOperation, "", err.Error(), username)

			} else {
				go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.CompletedState, db.ShowOperation, "", responses.DelayCompletedState, username)

			}

		}()

		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			return
		}

		var Delay db.Delay

		Delay, err = deps.Store.ShowDelay(req.Context(), id)
		if err != nil {
			rw.WriteHeader(http.StatusNotFound)
			logger.WithField("err", err.Error()).Error("Error show Delay")
			return
		}

		respBytes, err := json.Marshal(Delay)
		if err != nil {
			logger.WithField("err", err.Error()).Error("Error marshaling Delay data")
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}

		rw.Header().Add("Content-Type", "application/json")
		rw.WriteHeader(http.StatusOK)
		rw.Write(respBytes)
	})
}

func updateDelayHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		//logging when the api is initialised
		username := req.Context().Value(contextKeyUsername).(string)
		go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.InitialisedState, db.UpdateOperation, "", responses.DelayInitialisedState, username)

		vars := mux.Vars(req)
		id, err := parseUUID(vars["id"])
		// for logging error if there is any otherwise logging success
		defer func() {
			if err != nil {
				go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.ErrorState, db.UpdateOperation, "", err.Error(), username)

			} else {
				go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.CompletedState, db.UpdateOperation, "", responses.DelayCompletedState, username)

			}

		}()

		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			return
		}

		var delay db.Delay

		err = json.NewDecoder(req.Body).Decode(&delay)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			logger.WithField("err", err.Error()).Error("Error while decoding delay data")
			return
		}

		valid, respBytes := validate(delay)
		if !valid {
			responseBadRequest(rw, respBytes)
			return
		}

		delay.ProcessID = id
		err = deps.Store.UpdateDelay(req.Context(), delay)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			logger.WithField("err", err.Error()).Error("Error update delay")
			return
		}

		rw.WriteHeader(http.StatusOK)
		rw.Header().Add("Content-Type", "application/json")
		rw.Write([]byte(`{"msg":"Delay record updated successfully"}`))
	})
}
