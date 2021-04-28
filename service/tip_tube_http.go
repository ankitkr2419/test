package service

import (
	"encoding/json"
	"mylab/cpagent/db"
	"net/http"

	"github.com/gorilla/mux"
	logger "github.com/sirupsen/logrus"
)

func createTipTubeHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		var tt db.TipsTubes
		err := json.NewDecoder(req.Body).Decode(&tt)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			logger.WithField("err", err.Error()).Error("Error while decoding Tip or Tube data")
			return
		}

		valid, respBytes := validate(tt)
		if !valid {
			responseBadRequest(rw, respBytes)
			return
		}

		err = deps.Store.InsertTipsTubes(req.Context(), []db.TipsTubes{tt})
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			logger.WithField("err", err.Error()).Error("Error while inserting Tip or Tube")
			return
		}

		respBytes, err = json.Marshal(tt)
		if err != nil {
			logger.WithField("err", err.Error()).Error("Error marshalling Tip or Tube data")
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}

		rw.WriteHeader(http.StatusCreated)
		rw.Write(respBytes)
		rw.Header().Add("Content-Type", "application/json")
	})
}

func listTipsTubesHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		tipTubeType := vars["tiptube"]

		var tipsTubes []db.TipsTubes
		var err error

		switch tipTubeType {
		case "tip", "tube", "":
			tipsTubes, err = deps.Store.ListTipsTubes(tipTubeType)
			if err != nil {
				rw.WriteHeader(http.StatusInternalServerError)
				logger.WithField("err", err.Error()).Error("Error showing Tip tubes")
				return
			}
		default:
			rw.WriteHeader(http.StatusBadRequest)
			logger.WithField("err", "invalid argument").Error("Invalid Argument")
			rw.Write([]byte(`{"error":"invalid arguments"}`))
			return
		}

		respBytes, err := json.Marshal(tipsTubes)
		if err != nil {
			logger.WithField("err", err.Error()).Error("Error marshaling Tip tubes data")
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}
		rw.Header().Add("Content-Type", "application/json")
		rw.WriteHeader(http.StatusOK)
		rw.Write(respBytes)
	})
}
