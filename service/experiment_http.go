package service

import (
	"context"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"mylab/cpagent/config"
	"mylab/cpagent/db"
	"mylab/cpagent/plc"
	"mylab/cpagent/tec"
	"mylab/cpagent/responses"
	"net/http"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	logger "github.com/sirupsen/logrus"
)

func listExperimentHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		e, err := deps.Store.ListExperiments(req.Context())
		if err != nil {
			logger.WithField("err", err.Error()).Error("Error fetching experiment data")
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}

		respBytes, err := json.Marshal(e)
		if err != nil {
			logger.WithField("err", err.Error()).Error("Error marshaling experiments data")
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}

		rw.Header().Add("Content-Type", "application/json")
		rw.WriteHeader(http.StatusOK)
		rw.Write(respBytes)
	})
}

func createExperimentHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		var e db.Experiment
		err := json.NewDecoder(req.Body).Decode(&e)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			logger.WithField("err", err.Error()).Error("Error while decoding experiment data")
			return
		}

		valid, respBytes := validate(e)
		if !valid {
			responseBadRequest(rw, respBytes)
			return
		}

		// create new experiment
		var createdExp db.Experiment
		createdExp, err = deps.Store.CreateExperiment(req.Context(), e)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			logger.WithField("err", err.Error()).Error("Error create experiment")
			return
		}

		// get targets with template

		tt, err := deps.Store.ListTemplateTargets(req.Context(), createdExp.TemplateID)
		if err != nil {
			logger.WithField("err", err.Error()).Error("Error fetching template target data")
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}

		// add targets to experiment template

		exptargets := make([]db.ExpTemplateTarget, 0)

		for _, t := range tt {
			ett := db.ExpTemplateTarget{}

			ett.ExperimentID = createdExp.ID
			ett.TemplateID = t.TemplateID
			ett.TargetID = t.TargetID
			ett.Threshold = t.Threshold

			exptargets = append(exptargets, ett)
		}

		// insert in exp template

		_, err = deps.Store.UpsertExpTemplateTarget(req.Context(), exptargets, createdExp.ID)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			logger.WithField("err", err.Error()).Error("Error create target")
			return
		}

		respBytes, err = json.Marshal(createdExp)
		if err != nil {
			logger.WithField("err", err.Error()).Error("Error marshaling templates data")
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}

		rw.WriteHeader(http.StatusCreated)
		rw.Write(respBytes)
		rw.Header().Add("Content-Type", "application/json")
	})
}

func showExperimentHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)

		id, err := parseUUID(vars["id"])
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			return
		}

		var latestE db.Experiment

		latestE, err = deps.Store.ShowExperiment(req.Context(), id)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			logger.WithField("err", err.Error()).Error("Error show experiment")
			return
		}

		respBytes, err := json.Marshal(latestE)
		if err != nil {
			logger.WithField("err", err.Error()).Error("Error marshaling experiment data")
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}

		rw.WriteHeader(http.StatusOK)
		rw.Write(respBytes)
		rw.Header().Add("Content-Type", "application/json")
	})
}

func runExperimentHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		expID, err := parseUUID(vars["experiment_id"])
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			return
		}

		wells, err := deps.Store.ListWells(req.Context(), expID)
		if err != nil {
			logger.WithField("err", err.Error()).Error("Error fetching wells data")
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}

		// validate of NC,PC or NTC set for wells

		isValid, response := db.ValidateExperiment(wells)

		e, err := deps.Store.ShowExperiment(req.Context(), expID)
		if err != nil {
			logger.WithField("err", err.Error()).Error("Error fetching experiment data")
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}

		ss, err := deps.Store.ListStageSteps(req.Context(), e.TemplateID)
		if err != nil {
			logger.WithField("err", err.Error()).Error("Error fetching data")
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}
		plcStage := makePLCStage(ss)

		// err = deps.Plc.ConfigureRun(plcStage)
		// if err != nil {
		// 	logger.WithField("err", err.Error()).Error("Error in ConfigureRun")
		// 	rw.WriteHeader(http.StatusInternalServerError)
		// 	return
		// }

		// err = tec.Run(plcStage)

		// err = deps.Plc.Start()
		// if err != nil {
		// 	logger.WithField("err", err.Error()).Error("Error in plc start")
		// 	rw.WriteHeader(http.StatusInternalServerError)
		// 	return
		// }

		err = deps.Store.UpdateStartTimeExperiments(req.Context(), time.Now(), expID, plcStage.CycleCount)
		if err != nil {
			logger.WithField("err", err.Error()).Error("Error fetching data")
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}

		// retruns all targets configured for experiment
		targetDetails, err := deps.Store.ListConfTargets(req.Context(), expID)
		if err != nil {
			logger.WithField("err", err.Error()).Error("Error fetching target data")
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}

		var ICTargetID uuid.UUID

		for _, t := range targetDetails {
			if t.DyePosition == int32(config.GetICPosition()) {
				ICTargetID = t.TargetID
			}

		}

		setExperimentValues(config.ActiveWells("activeWells"), ICTargetID, targetDetails, expID, plcStage)

		WellTargets := initializeWellTargets()

		// update well targets value in DB
		_, err = deps.Store.UpsertWellTargets(context.Background(), WellTargets, experimentValues.experimentID, false)
		if err != nil {
			// send error
			logger.WithField("err", err.Error()).Error("Error upsert wells")
			return
		}
		//experimentRunning set true
		experimentRunning = true
		go startExp(deps, plcStage)

		//invoke monitor
		// ASK: Do we need this?
		go monitorExperiment(deps)

		if !isValid {
			respBytes, err := json.Marshal(response)
			if err != nil {
				logger.WithField("err", err.Error()).Error("Error marshaling experiments data")
				rw.WriteHeader(http.StatusInternalServerError)
				return
			}
			rw.Header().Add("Content-Type", "application/json")
			rw.WriteHeader(http.StatusAccepted)
			rw.Write(respBytes)
			return
		}

		rw.Header().Add("Content-Type", "application/json")
		rw.WriteHeader(http.StatusOK)
		rw.Write([]byte(`{"msg":"experiment started"}`))
		return
	})
}

func stopExperimentHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		vars := mux.Vars(req)
		expID, err := parseUUID(vars["experiment_id"])
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			return
		}
		// instruct plc to stop the experiment: stops if experiment is already running else returns error
		err = deps.Plc.Stop()
		if err != nil {
			logger.WithField("err", err.Error()).Error("Error in plc stop")
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}

		err = deps.Store.UpdateStopTimeExperiments(req.Context(), time.Now(), expID, "aborted")
		if err != nil {
			logger.WithField("err", err.Error()).Error("Error fetching data")
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}

		rw.Header().Add("Content-Type", "application/json")
		rw.WriteHeader(http.StatusOK)
		rw.Write([]byte(`{"msg":"experiment stopped"}`))

	})
}

func startExp(deps Dependencies, p plc.Stage) (err error) {
	// logging output to file and console
	if _, err := os.Stat(tec.LogsPath); os.IsNotExist(err) {
		os.MkdirAll(tec.LogsPath, 0755)
		// ignore error and try creating log output file
	}

	file, err := os.Create(fmt.Sprintf("%v/output_%v.csv", tec.LogsPath, time.Now().Unix()))
	if err != nil {
		logger.WithField("Err", err).Errorln(responses.FileCreationError)
		return	
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Home the TEC
	// Reset is implicit in Homing
	deps.Plc.HomingRTPCR()

	// Start line
	err = writer.Write([]string{"Description", "Time Taken", "Expected Time", "Initial Temp", "Final Temp", "Ramp"})
	if err != nil {
		return
	}
	
	timeStarted := time.Now()
	writer.Write([]string{"Experiment Started at: ", timeStarted.String()})

	writer.Write([]string{"Holding Stage About to start"})

	tec.TempMonStarted = true

	//Go back to Room Temp at the end
	defer func() {
		tec.TempMonStarted = false
		experimentRunning = false
		if err != nil {
			return
		}
		err = deps.Tec.ReachRoomTemp()
		return
	}()
	// Run Holding Stage
	logger.Infoln("Holding Stage Started")
	deps.Tec.RunStage(p.Holding, writer, 0)
	writer.Flush()

	// Cycle in Plc
	deps.Plc.Cycle()

	// Run Cycle Stage
	err = writer.Write([]string{"Cycle Stage About to start"})
	if err != nil {
		return
	}

	for i := uint16(1); i <= p.CycleCount; i++ {
		logger.Infoln("Started Cycle->", i)
		deps.Tec.RunStage(p.Cycle, writer, i)
		writer.Flush()
		logger.Infoln("Cycle Completed -> ", i)
		// Cycle in Plc
		deps.Plc.Cycle()
	}

	writer.Write([]string{"Experiment Completed at: ", time.Now().String()})
	writer.Write([]string{"Total Time Taken by Experiment: ", time.Now().Sub(timeStarted).String()})

	return
}
