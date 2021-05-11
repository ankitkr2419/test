package service

import (
	"encoding/json"
	"mylab/cpagent/db"
	"mylab/cpagent/responses"
	"net/http"

	"github.com/gorilla/mux"
	logger "github.com/sirupsen/logrus"
)

func createAttachDetachHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		username := req.Context().Value(contextKeyUsername).(string)
		go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.InitialisedState, db.CreateOperation, "", responses.AttachDetachInitialisedState, username)

		var adObj db.AttachDetach
		err := json.NewDecoder(req.Body).Decode(&adObj)

		// for logging error if there is any otherwise logging success
		defer func() {
			if err != nil {
				go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.ErrorState, db.CreateOperation, "", err.Error(), username)

			} else {
				go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.CompletedState, db.CreateOperation, "", responses.AttachDetachCompletedState, username)

			}

		}()

		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			logger.WithField("err", err.Error()).Error("Error while decoding attach detach data")
			return
		}

		valid, respBytes := validate(adObj)
		if !valid {
			responseBadRequest(rw, respBytes)
			return
		}

		var createdAtDt db.AttachDetach
		createdAtDt, err = deps.Store.CreateAttachDetach(req.Context(), adObj)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			logger.WithField("err", err.Error()).Error("Error create attach detach")
			return
		}

		respBytes, err = json.Marshal(createdAtDt)
		if err != nil {
			logger.WithField("err", err.Error()).Error("Error marshaling attach detach data")
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}

		rw.Header().Add("Content-Type", "application/json")
		rw.WriteHeader(http.StatusCreated)
		rw.Write(respBytes)
	})
}

func showAttachDetachHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		//logging when the api is initialised
		username := req.Context().Value(contextKeyUsername).(string)
		go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.InitialisedState, db.ShowOperation, "", responses.AttachDetachInitialisedState, username)

		vars := mux.Vars(req)
		id, err := parseUUID(vars["id"])

		// for logging error if there is any otherwise logging success
		defer func() {
			if err != nil {
				go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.ErrorState, db.ShowOperation, "", err.Error(), username)

			} else {
				go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.CompletedState, db.ShowOperation, "", responses.AttachDetachCompletedState, username)

			}

		}()

		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			return
		}

		var AttachDetach db.AttachDetach
		AttachDetach, err = deps.Store.ShowAttachDetach(req.Context(), id)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			logger.WithField("err", err.Error()).Error("Error show attach detach")
			return
		}

		respBytes, err := json.Marshal(AttachDetach)
		if err != nil {
			logger.WithField("err", err.Error()).Error("Error marshaling attach detach data")
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}
		rw.Header().Add("Content-Type", "application/json")
		rw.WriteHeader(http.StatusOK)
		rw.Write(respBytes)
	})
}

func updateAttachDetachHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		//logging when the api is initialised
		username := req.Context().Value(contextKeyUsername).(string)
		go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.InitialisedState, db.UpdateOperation, "", responses.AttachDetachInitialisedState, username)

		vars := mux.Vars(req)
		id, err := parseUUID(vars["id"])

		// for logging error if there is any otherwise logging success
		defer func() {
			if err != nil {
				go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.ErrorState, db.UpdateOperation, "", err.Error(), username)

			} else {
				go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.CompletedState, db.UpdateOperation, "", responses.AttachDetachCompletedState, username)

			}

		}()

		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			return
		}
		var adObj db.AttachDetach
		err = json.NewDecoder(req.Body).Decode(&adObj)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			logger.WithField("err", err.Error()).Error("Error while decoding attach detach data")
			return
		}
		valid, respBytes := validate(adObj)
		if !valid {
			responseBadRequest(rw, respBytes)
			return
		}
		adObj.ProcessID = id
		err = deps.Store.UpdateAttachDetach(req.Context(), adObj)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			logger.WithField("err", err.Error()).Error("Error update attach detach")
			return
		}
		rw.WriteHeader(http.StatusOK)
		rw.Header().Add("Content-Type", "application/json")
		rw.Write([]byte(`{"msg":"Attach Detach record updated successfully"}`))
	})
}
