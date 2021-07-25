package service

import (
	"context"
	"encoding/json"
	"mylab/cpagent/config"
	"mylab/cpagent/db"
	"net/http"

	"github.com/gorilla/mux"
	logger "github.com/sirupsen/logrus"
)

type Scale struct {
	XAxisMin int     `json:"x_axis_min"`
	XAxisMax int     `json:"x_axis_max"`
	YAxisMin float32 `json:"y_axis_min"`
	YAxisMax float32 `json:"y_axis_max"`
}

func updateScaleHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		vars := mux.Vars(req)
		expId, err := parseUUID(vars["id"])
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			return
		}
		var t Scale
		err = json.NewDecoder(req.Body).Decode(&t)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			logger.WithField("err", err.Error()).Error("Error while decoding scale data")
			return
		}

		e, err := deps.Store.ShowExperiment(req.Context(), expId)
		if err != nil {
			logger.WithField("err", err.Error()).Error("Error fetching experiment data")
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}

		// returns all targets configured for experiment
		targetDetails, err := deps.Store.ListConfTargets(req.Context(), expId)
		if err != nil {
			logger.WithField("err", err.Error()).Error("Error fetching target data")
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}

		DBResult, err := deps.Store.GetResult(context.Background(), expId)
		if err != nil {
			logger.WithField("err", err.Error()).Error("Error fetching result data")
			return
		}
		result := UpdateScale(DBResult, config.ActiveWells("activeWells"), targetDetails, e.RepeatCycle, t)

		responseCodeAndMsg(rw, http.StatusOK, result)
	})
}

func UpdateScale(result []db.Result, wells []int32, targets []db.TargetDetails, cycles uint16, s Scale) (finalResult []graph) {

	// ex: for 8 active wells * 6 targets * no of cycle
	for _, aw := range wells {
		var wellResult graph
		wellResult.WellPosition = aw

		for _, t := range targets {
			wellResult.TargetID = t.TargetID
			for _, r := range result {
				if r.WellPosition == wellResult.WellPosition && r.TargetID == wellResult.TargetID {
					if r.Cycle <= uint16(s.XAxisMax) && r.Cycle >= uint16(s.XAxisMin) {
						wellResult.ExperimentID = r.ExperimentID
						wellResult.TargetID = r.TargetID
						wellResult.Threshold = r.Threshold
						wellResult.TotalCycles = cycles

						// if cycle found do not add again!
						if !found(r.Cycle, wellResult.Cycle) {
							wellResult.Cycle = append(wellResult.Cycle, r.Cycle)
							wellResult.FValue = append(wellResult.FValue, scaleValue(float32(r.FValue), s.YAxisMin, s.YAxisMax))
						}
					}

				}
			}
			finalResult = append(finalResult, wellResult)
			wellResult.Cycle = []uint16{}
			wellResult.FValue = []float32{}
		}

	}
	return
}

func scaleValue(value float32, minScale, maxScale float32) (retValue float32) {

	return ((value-pcrMin)/(pcrMax-pcrMin))*(maxScale-minScale) + minScale
}
