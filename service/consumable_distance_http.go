package service

import (
	"encoding/json"
	"mylab/cpagent/db"
	"net/http"

	logger "github.com/sirupsen/logrus"
)

func createConsumableDistanceHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		var cd db.ConsumableDistance
		err := json.NewDecoder(req.Body).Decode(&cd)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			logger.WithField("err", err.Error()).Error("Error while decoding Consumable Distance data")
			return
		}

		valid, respBytes := validate(cd)
		if !valid {
			responseBadRequest(rw, respBytes)
			return
		}

		err = deps.Store.InsertConsumableDistance(req.Context(), []db.ConsumableDistance{cd})
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			logger.WithField("err", err.Error()).Error("Error while inserting comsumable distance")
			return
		}

		respBytes, err = json.Marshal(cd)
		if err != nil {
			logger.WithField("err", err.Error()).Error("Error marshaling Consumable Distance data")
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}
		rw.Header().Add("Content-Type", "application/json")
		rw.WriteHeader(http.StatusCreated)
		rw.Write(respBytes)
	})
}
