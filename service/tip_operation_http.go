package service

import (
	"encoding/json"
	"mylab/cpagent/db"
	"net/http"

	"github.com/gorilla/mux"
	logger "github.com/sirupsen/logrus"
)

func createTipOperationHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		var tipOpr db.TipOperation
		err := json.NewDecoder(req.Body).Decode(&tipOpr)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			logger.WithField("err", err.Error()).Error("Error while decoding TipOperation data")
			return
		}

		valid, respBytes := validate(tipOpr)
		if !valid {
			responseBadRequest(rw, respBytes)
			return
		}

		var createdTemp db.TipOperation
		createdTemp, err = deps.Store.CreateTipOperation(req.Context(), tipOpr)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			logger.WithField("err", err.Error()).Error("Error create TipOperation")
			return
		}

		respBytes, err = json.Marshal(createdTemp)
		if err != nil {
			logger.WithField("err", err.Error()).Error("Error marshaling TipOperation data")
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}

		rw.Header().Add("Content-Type", "application/json")
		rw.WriteHeader(http.StatusCreated)
		rw.Write(respBytes)
	})
}

func showTipOperationHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)

		id, err := parseUUID(vars["id"])
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			return
		}

		var TipOperation db.TipOperation

		TipOperation, err = deps.Store.ShowTipOperation(req.Context(), id)
		if err != nil {
			rw.WriteHeader(http.StatusNotFound)
			logger.WithField("err", err.Error()).Error("Error show TipOperation")
			return
		}

		respBytes, err := json.Marshal(TipOperation)
		if err != nil {
			logger.WithField("err", err.Error()).Error("Error marshaling TipOperation data")
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}

		rw.Header().Add("Content-Type", "application/json")
		rw.WriteHeader(http.StatusOK)
		rw.Write(respBytes)
	})
}

func updateTipOperationHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		id, err := parseUUID(vars["id"])
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			return
		}

		var tipOpr db.TipOperation

		err = json.NewDecoder(req.Body).Decode(&tipOpr)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			logger.WithField("err", err.Error()).Error("Error while decoding piercing data")
			return
		}

		valid, respBytes := validate(tipOpr)
		if !valid {
			responseBadRequest(rw, respBytes)
			return
		}

		tipOpr.ProcessID = id
		err = deps.Store.UpdateTipOperation(req.Context(), tipOpr)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			logger.WithField("err", err.Error()).Error("Error update piercing")
			return
		}

		rw.WriteHeader(http.StatusOK)
		rw.Header().Add("Content-Type", "application/json")
		rw.Write([]byte(`{"msg":"TipOperation record updated successfully"}`))
	})
}
