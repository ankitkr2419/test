package service

import (
	"encoding/json"
	"fmt"

	// "mylab/cpagent/db"
	"mylab/cpagent/responses"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"

	logger "github.com/sirupsen/logrus"
)

type ThresholdCals struct {
	Dye           string  `json:"dye" validate:"required"`
	AutoThreshold bool    `json:"auto_threshold"`
	Threshold     float64 `json:"threshold"`
	AutoBaseline  bool    `json:"auto_baseline"`
	StartCycle    int64   `json:"start_cycle"`
	EndCycle      int64   `json:"end_cycle"`
}
type TargetCycleWell struct {
	Target uuid.UUID
	Well   int32
}

func setThresholdHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		var tc ThresholdCals

		vars := mux.Vars(req)
		expID, err := parseUUID(vars["experiment_id"])
		if err != nil {
			logger.Errorln("Invalid Experiment ID: ", expID)
			responseCodeAndMsg(rw, http.StatusBadRequest, ErrObj{Err: responses.InvalidExperimentID.Error()})
			return
		}

		err = json.NewDecoder(req.Body).Decode(&tc)
		if err != nil {
			logger.WithField("err", err.Error()).Error("Error while decoding Threshold Settings data")
			responseCodeAndMsg(rw, http.StatusBadRequest, ErrObj{Err: err.Error()})
			return
		}

		wells, err := deps.Store.ListWells(req.Context(), expID)
		if err != nil {
			logger.WithField("err", err.Error()).Error("Error fetching wells data")
			responseCodeAndMsg(rw, http.StatusInternalServerError, ErrObj{Err: "Error fetching wells data"})
			return
		}

		if len(wells) == 0 {
			err = fmt.Errorf("No wells configured")
			logger.Errorln(err)
			responseCodeAndMsg(rw, http.StatusBadRequest, ErrObj{Err: err.Error()})
		}

		wellPositions := make([]int32, 0)
		for _, w := range wells {
			wellPositions = append(wellPositions, int32(w.Position))
		}

		targets, err := deps.Store.ListConfTargets(req.Context(), expID)
		if err != nil {
			logger.WithField("err", err.Error()).Error("Error fetching targets data")
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}
		logger.Infoln("targetssssss", targets)
		e, err := deps.Store.ShowExperiment(req.Context(), expID)
		if err != nil {
			logger.WithField("err", err.Error()).Error("Error fetching experiment data")
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}

		respBytes, err := getWellsDataByThreshold(deps, expID, wellPositions, targets, e.RepeatCycle, tc)
		if err != nil {
			logger.WithField("err", err.Error()).Error("Error marshaling Result data")
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}

		rw.Write(respBytes)
		rw.Header().Add("Content-Type", "application/json")
		rw.WriteHeader(http.StatusAccepted)
	})
}
