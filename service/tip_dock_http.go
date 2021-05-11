package service

import (
	"encoding/json"
	"mylab/cpagent/db"
	"mylab/cpagent/responses"
	"net/http"

	"github.com/gorilla/mux"
	logger "github.com/sirupsen/logrus"
)

func createTipDockHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		username := req.Context().Value(contextKeyUsername).(string)
		go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.InitialisedState, db.CreateOperation, "", responses.TipDockingInitialisedState, username)

		var tdObj db.TipDock
		err := json.NewDecoder(req.Body).Decode(&tdObj)

		// for logging error if there is any otherwise logging success
		defer func() {
			if err != nil {
				go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.ErrorState, db.CreateOperation, "", err.Error(), username)

			} else {
				go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.CompletedState, db.CreateOperation, "", responses.TipDockingCompletedState, username)

			}

		}()

		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			logger.WithField("err", err.Error()).Error("Error while decoding Tip  Docking data")
			return
		}

		valid, respBytes := validate(tdObj)
		if !valid {
			responseBadRequest(rw, respBytes)
			return
		}

		var createdAtDt db.TipDock
		createdAtDt, err = deps.Store.CreateTipDocking(req.Context(), tdObj)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			logger.WithField("err", err.Error()).Error("Error create Tip Docking")
			return
		}

		respBytes, err = json.Marshal(createdAtDt)
		if err != nil {
			logger.WithField("err", err.Error()).Error("Error marshaling Tip Docking data")
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}

		rw.Header().Add("Content-Type", "application/json")
		rw.WriteHeader(http.StatusCreated)
		rw.Write(respBytes)
	})
}

func showTipDockHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		//logging when the api is initialised
		username := req.Context().Value(contextKeyUsername).(string)
		go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.InitialisedState, db.ShowOperation, "", responses.TipDockingInitialisedState, username)

		vars := mux.Vars(req)
		id, err := parseUUID(vars["id"])

		// for logging error if there is any otherwise logging success
		defer func() {
			if err != nil {
				go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.ErrorState, db.ShowOperation, "", err.Error(), username)

			} else {
				go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.CompletedState, db.ShowOperation, "", responses.TipDockingCompletedState, username)

			}

		}()

		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			return
		}

		var TipDock db.TipDock
		TipDock, err = deps.Store.ShowTipDocking(req.Context(), id)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			logger.WithField("err", err.Error()).Error("Error show Tip Docking")
			return
		}

		respBytes, err := json.Marshal(TipDock)
		if err != nil {
			logger.WithField("err", err.Error()).Error("Error marshaling Tip Docking data")
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}
		rw.Header().Add("Content-Type", "application/json")
		rw.WriteHeader(http.StatusOK)
		rw.Write(respBytes)
	})
}

func updateTipDockHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		//logging when the api is initialised
		username := req.Context().Value(contextKeyUsername).(string)
		go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.InitialisedState, db.UpdateOperation, "", responses.TipDockingInitialisedState, username)

		vars := mux.Vars(req)
		id, err := parseUUID(vars["id"])
		// for logging error if there is any otherwise logging success
		defer func() {
			if err != nil {
				go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.ErrorState, db.UpdateOperation, "", err.Error(), username)

			} else {
				go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.CompletedState, db.UpdateOperation, "", responses.TipDockingCompletedState, username)

			}

		}()
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			return
		}
		var tdObj db.TipDock
		err = json.NewDecoder(req.Body).Decode(&tdObj)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			logger.WithField("err", err.Error()).Error("Error while decoding tip docking data")
			return
		}
		valid, respBytes := validate(tdObj)
		if !valid {
			responseBadRequest(rw, respBytes)
			return
		}
		tdObj.ProcessID = id
		err = deps.Store.UpdateTipDock(req.Context(), tdObj)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			logger.WithField("err", err.Error()).Error("Error update tip docking")
			return
		}
		rw.WriteHeader(http.StatusOK)
		rw.Header().Add("Content-Type", "application/json")
		rw.Write([]byte(`{"msg":"Tip Dock record updated successfully"}`))
	})
}
