package service

import (
	"encoding/json"
	"mylab/cpagent/db"
	"net/http"

	"github.com/gorilla/mux"
	logger "github.com/sirupsen/logrus"
)

func createAttachDetachHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		var adObj db.AttachDetach
		err := json.NewDecoder(req.Body).Decode(&adObj)
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
		vars := mux.Vars(req)
		id, err := parseUUID(vars["id"])
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
		vars := mux.Vars(req)
		id, err := parseUUID(vars["id"])
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
		adObj.ID = id
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
