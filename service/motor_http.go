package service

import (
	"encoding/json"
	"mylab/cpagent/db"
	"net/http"

	logger "github.com/sirupsen/logrus"
)

func createMotorHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		var m db.Motor
		err := json.NewDecoder(req.Body).Decode(&m)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			logger.WithField("err", err.Error()).Error("Error while decoding motor data")
			return
		}

		valid, respBytes := validate(m)
		if !valid {
			responseBadRequest(rw, respBytes)
			return
		}

		err = deps.Store.InsertMotor(req.Context(), []db.Motor{m})
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			logger.WithField("err", err.Error()).Error("Error while inserting motor")
			return
		}

		respBytes, err = json.Marshal(m)
		if err != nil {
			logger.WithField("err", err.Error()).Error("Error marshaling Motor data")
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}

		rw.WriteHeader(http.StatusCreated)
		rw.Write(respBytes)
		rw.Header().Add("Content-Type", "application/json")
	})
}
