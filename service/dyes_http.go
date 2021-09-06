package service

import (
	"encoding/json"
	"mylab/cpagent/db"
	"net/http"

	logger "github.com/sirupsen/logrus"
)

func updateDyeToleranceHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		var cd []db.Dye
		err := json.NewDecoder(req.Body).Decode(&cd)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			logger.WithField("err", err.Error()).Error("Error while decoding dyes data")
			return
		}

		go db.UpdateDyesTolerance(cd)

		dyes, err := deps.Store.InsertDyes(req.Context(), cd)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			logger.WithField("err", err.Error()).Error("Error while inserting dyes")
			return
		}

		respBytes, err := json.Marshal(dyes)
		if err != nil {
			logger.WithField("err", err.Error()).Error("Error marshaling dyes data")
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}
		rw.Header().Add("Content-Type", "application/json")
		rw.WriteHeader(http.StatusCreated)
		rw.Write(respBytes)
	})
}
func listDyesHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		e, err := deps.Store.ListDyes(req.Context())
		if err != nil {
			logger.WithField("err", err.Error()).Error("Error fetching Dyes data")
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}

		respBytes, err := json.Marshal(e)
		if err != nil {
			logger.WithField("err", err.Error()).Error("Error marshaling Dyes data")
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}

		rw.Header().Add("Content-Type", "application/json")
		rw.WriteHeader(http.StatusOK)
		rw.Write(respBytes)
	})
}
