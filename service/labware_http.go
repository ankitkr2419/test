package service

import (
	"encoding/json"
	"mylab/cpagent/db"
	"net/http"

	logger "github.com/sirupsen/logrus"
)

func createLabwareHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		var l db.Labware
		err := json.NewDecoder(req.Body).Decode(&l)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			logger.WithField("err", err.Error()).Error("Error while decoding Labware data")
			return
		}

		valid, respBytes := validate(l)
		if !valid {
			responseBadRequest(rw, respBytes)
			return
		}

		err = deps.Store.InsertLabware(req.Context(), []db.Labware{l})
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			logger.WithField("err", err.Error()).Error("Error while inserting labware")
			return
		}

		respBytes, err = json.Marshal(l)
		if err != nil {
			logger.WithField("err", err.Error()).Error("Error marshaling Labware data")
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}

		rw.WriteHeader(http.StatusCreated)
		rw.Write(respBytes)
		rw.Header().Add("Content-Type", "application/json")
	})
}
