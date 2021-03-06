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

type Threshold struct {
	TargetID      uuid.UUID `json:"target_id"`
	AutoThreshold bool      `json:"auto_threshold"`
	Threshold     float32   `json:"threshold"`
	StartCycle    uint16    `json:"start_cycle" validate:"gte=0"`
	EndCycle      uint16    `json:"end_cycle" validate:"gte=0"`
}
type Baseline struct {
	AutoBaseline bool   `json:"auto_baseline"`
	StartCycle   uint16 `json:"start_cycle" validate:"gte=0"`
	EndCycle     uint16 `json:"end_cycle" validate:"gte=0"`
}
type TargetWell struct {
	Target uuid.UUID
	Well   int32
}

func setThresholdHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		var tc Threshold

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
		valid, validationResp := Validate(tc)
		if !valid {
			logger.WithField("err", "Validation Error").Errorln(responses.ThresholdValidationError)
			responseBadRequest(rw, validationResp)
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

		e, err := deps.Store.ShowExperiment(req.Context(), expID)
		if err != nil {
			logger.WithField("err", err.Error()).Error("Error fetching experiment data")
			responseCodeAndMsg(rw, http.StatusInternalServerError, ErrObj{Err: err.Error()})

			return
		}

		respBytes, err := getWellsDataByThreshold(deps, expID, wellPositions, targets, wells, e.RepeatCycle, tc)
		if err != nil {
			logger.WithField("err", err.Error()).Error("Error marshaling Result data")
			responseCodeAndMsg(rw, http.StatusInternalServerError, ErrObj{Err: err.Error()})
			return
		}

		logger.Infoln("Successfully fetched the threshold values")
		responseCodeAndMsg(rw, http.StatusOK, respBytes)
	})
}

func getBaselineNormalisedValuesHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		var bl Baseline
		rw.Header().Add("Content-Type", "application/json")

		vars := mux.Vars(req)
		expID, err := parseUUID(vars["experiment_id"])
		if err != nil {
			logger.Errorln("Invalid Experiment ID: ", expID)
			responseCodeAndMsg(rw, http.StatusBadRequest, ErrObj{Err: responses.InvalidExperimentID.Error()})
			return
		}

		err = json.NewDecoder(req.Body).Decode(&bl)
		if err != nil {
			logger.WithField("err", err.Error()).Error("Error while decoding Threshold Settings data")
			responseCodeAndMsg(rw, http.StatusBadRequest, ErrObj{Err: err.Error()})
			return
		}
		valid, validationResp := Validate(bl)
		if !valid {
			logger.WithField("err", "Validation Error").Errorln(responses.BaselineValidationError)
			responseBadRequest(rw, validationResp)
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
			responseCodeAndMsg(rw, http.StatusExpectationFailed, ErrObj{Err: err.Error()})
			return
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

		e, err := deps.Store.ShowExperiment(req.Context(), expID)
		if err != nil {
			logger.WithField("err", err.Error()).Error("Error fetching experiment data")
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}

		respBytes, err := getBaselineData(deps, expID, wellPositions, targets, wells, e.RepeatCycle, bl)
		if err != nil {
			logger.WithField("err", err.Error()).Error("Error marshaling Result data")
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}

		logger.Infoln("Successfully fetched the baseline values")
		responseCodeAndMsg(rw, http.StatusOK, respBytes)
	})
}
