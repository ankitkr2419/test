package service

import (
	"net/http"

	"github.com/gorilla/mux"
	logger "github.com/sirupsen/logrus"
)

func getResultHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)

		experimentID, err := parseUUID(vars["id"])
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			return
		}

		wells, err := deps.Store.ListWells(req.Context(), experimentID)
		if err != nil {
			logger.WithField("err", err.Error()).Error("Error fetching data")
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}

		if len(wells) > 0 {
			wellPositions := make([]int32, 0)
			for _, w := range wells {
				wellPositions = append(wellPositions, int32(w.Position))
			}

			targets, err := deps.Store.ListConfTargets(req.Context(), experimentID)
			if err != nil {
				logger.WithField("err", err.Error()).Error("Error fetching targets data")
				rw.WriteHeader(http.StatusInternalServerError)
				return
			}

			respBytes, err := getGraph(deps, experimentID, wellPositions, targets)
			if err != nil {
				logger.WithField("err", err.Error()).Error("Error marshaling Result data")
				rw.WriteHeader(http.StatusInternalServerError)
				return
			}
			rw.Header().Add("Content-Type", "application/json")
			rw.WriteHeader(http.StatusOK)
			rw.Write(respBytes)
			return
		}

		rw.Header().Add("Content-Type", "application/json")
		rw.WriteHeader(http.StatusExpectationFailed)
		rw.Write([]byte(`{"msg":"wells not configured"}`))
		return
	})
}
