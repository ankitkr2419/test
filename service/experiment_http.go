package service

import (
	"context"
	"encoding/json"
	"fmt"
	"math"
	"mylab/cpagent/config"
	"mylab/cpagent/db"
	"mylab/cpagent/plc"
	"net/http"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
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
		// create  new file for each experiment with experiment id in file name.
		file := db.GetExcelFile(ExpOutputPath, fmt.Sprintf("output_%v", expID))

		db.SetExperimentExcelFile(file)
		deps.Store.SetExcelHeadings(file, expID)

		wells, err := deps.Store.ListWells(req.Context(), expID)
		if err != nil {
			logger.WithField("err", err.Error()).Error("Error fetching wells data")
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}

		for _, v := range wells {
			db.AddRowToExcel(file, db.WellSample, []interface{}{v.ID, v.Position, v.ExperimentID, v.SampleID, v.Task, v.ColorCode, v.SampleName})
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

		for _, v := range ss {
			db.AddRowToExcel(file, db.StepsStageSheet, []interface{}{
				v.Stage.ID,
				v.Stage.Type,
				v.Stage.RepeatCount,
				v.Stage.TemplateID,
				v.Stage.StepCount,
				v.Stage.CreatedAt.String(),
				v.Stage.UpdatedAt.String(),
				v.Step.ID,
				v.Step.StageID,
				v.Step.RampRate,
				v.Step.TargetTemperature,
				v.Step.HoldTime,
				v.Step.DataCapture,
				v.Step.CreatedAt.String(),
				v.Step.UpdatedAt.String()})
		}

		plcStage := makePLCStage(ss)

		t, err := deps.Store.ShowTemplate(req.Context(), e.TemplateID)
		if err != nil {
			logger.WithField("err", err.Error()).Error("Error fetching template data")
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}

		db.AddRowToExcel(file, db.TemplateSheet, []interface{}{
			t.ID,
			t.Name,
			t.Description,
			t.Publish,
			t.CreatedAt.String(),
			t.UpdatedAt.String(),
			t.Volume,
			t.LidTemp,
			t.EstimatedTime,
			t.Finished})
		// Set currentExpTemplate to current running Template
		currentExpTemplate = t

		// Lid Temp is multiplied by 10 for PLC can't handle floats
		plcStage.IdealLidTemp = uint16(t.LidTemp * 10)
		e.Result = "running"

		// Set expStartTime to current Time
		expStartTime = time.Now()

		err = deps.Store.UpdateStartTimeExperiments(req.Context(), expStartTime, expID, plcStage.CycleCount, e.Result)
		if err != nil {
			logger.WithField("err", err.Error()).Error("Error fetching data")
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}

		// returns all targets configured for experiment
		targetDetails, err := deps.Store.ListConfTargets(req.Context(), expID)
		if err != nil {
			logger.WithField("err", err.Error()).Error("Error fetching target data")
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}

		for _, v := range targetDetails {
			db.AddRowToExcel(file, db.TargetSheet, []interface{}{
				v.ExperimentID,
				v.TemplateID,
				v.TargetID,
				v.Threshold,
				v.DyePosition,
				v.TargetName,
			})
		}

		var ICTargetID uuid.UUID
		var dyePositions []int32

		for _, t := range targetDetails {
			if t.DyePosition == int32(config.GetICPosition()) {
				ICTargetID = t.TargetID
			}
			dyePositions = append(dyePositions, t.DyePosition)

		}
		heading := []interface{}{"Dye Position"}
		for _, v := range dyePositions {
			heading = append(heading, v)
		}

		db.AddMergeRowToExcel(file, db.RTPCRSheet, heading, len(config.ActiveWells("activeWells")))

		row := []interface{}{"well positions"}
		for range dyePositions {
			for _, v := range config.ActiveWells("activeWells") {
				row = append(row, v)
			}
		}
		db.AddRowToExcel(file, db.RTPCRSheet, row)

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
		plc.ExperimentRunning = true
		go startExp(deps, plcStage, file)

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

		if plc.ExperimentRunning {
			e, err := deps.Store.ShowExperiment(req.Context(), expID)
			if err != nil {
				logger.WithField("err", err.Error()).Error("Error fetching experiment data")
				rw.WriteHeader(http.StatusInternalServerError)
				return
			}
			if e.Result != "running" {
				logger.WithField("err", "invalid running experiment").Error("this experiment id not running")
				responseCodeAndMsg(rw, http.StatusInternalServerError, ErrObj{Err: "this experiment is not running"})
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
				responseCodeAndMsg(rw, http.StatusInternalServerError, ErrObj{Err: "error fetching data"})
				return
			}
		} else {
			logger.WithField("err", "experiment not running").Error("No experiment running")
			responseCodeAndMsg(rw, http.StatusInternalServerError, ErrObj{Err: "no experiment is running"})
			return
		}
		responseCodeAndMsg(rw, http.StatusOK, MsgObj{Msg: "experiment stopped"})
	})
}

func startExp(deps Dependencies, p plc.Stage, file *excelize.File) (err error) {
	//Go back to Room Temp at the end

	// Experiment Running should stop at the end
	// And then Homing should happen
	defer func() {
		plc.ExperimentRunning = false
		err = deps.Plc.HomingRTPCR()
		if err != nil {
			deps.WsErrCh <- err
			return
		}

	}()

	// Send error on websocket
	// Reach Room Temp and SwitchOffLid
	defer func() {

		if err != nil {
			deps.WsErrCh <- err
		} else {
			deps.WsMsgCh <- "stop"
		}
		err = deps.Plc.SwitchOffLidTemp()
		if err != nil {
			deps.WsErrCh <- err
		}
		err = deps.Tec.ReachRoomTemp()
		if err != nil {
			deps.WsErrCh <- err
		}
		return
	}()

	err = deps.Tec.ReachRoomTemp()
	if err != nil {
		return
	}

	// Home the PLC
	// Reset is implicit in Homing
	err = deps.Plc.HomingRTPCR()
	if err != nil {
		return
	}

	// templateRunSuccess has to happen before monitor is called
	templateRunSuccess = false

	// invoke monitor after 2 secs
	go func() {
		time.Sleep(2 * time.Second)
		go monitorExperiment(deps, file)
	}()

	lidTempStartTime := time.Now()
	err = deps.Plc.SetLidTemp(p.IdealLidTemp)
	if err != nil {
		return
	}
	lidTempSecs := time.Now().Sub(lidTempStartTime).Seconds()

	// Setting currentExpTemplate Estimated Time more accurately.
	if currentExpTemplate.EstimatedTime != 0 {
		// Below formula is copied from estimated_time.go
		// We are removing variable LidTempTime
		estimatedLidTime := int64(math.Abs(float64(p.IdealLidTemp/10)-config.GetRoomTemp()) / 0.5)
		currentExpTemplate.EstimatedTime = currentExpTemplate.EstimatedTime - estimatedLidTime + int64(lidTempSecs)
	}

	// Start line
	headers := []interface{}{"Description", "Time Taken", "Expected Time", "Initial Temp", "Final Temp", "Ramp"}
	db.AddRowToExcel(file, db.TECSheet, headers)

	timeStarted := time.Now()
	row := []interface{}{"Experiment Started at: ", timeStarted.String()}
	db.AddRowToExcel(file, db.TempLogs, row)

	row = []interface{}{"Holding Stage About to start"}
	db.AddRowToExcel(file, db.TECSheet, row)

	row = []interface{}{"timestamp", "current Temperature", "lid Temperature"}
	db.AddRowToExcel(file, db.TempLogs, row)

	// Run Holding Stage
	logger.Infoln("Holding Stage Started")
	err = deps.Tec.RunStage(p.Holding, deps.Plc, file, 0)
	if err != nil {
		return
	}

	// Run Cycle Stage
	row = []interface{}{"Cycle Stage About to start"}
	db.AddRowToExcel(file, db.TECSheet, row)

	for i := uint16(1); i <= p.CycleCount; i++ {
		logger.Infoln("Started Cycle->", i)
		err = deps.Tec.RunStage(p.Cycle, deps.Plc, file, i)
		if err != nil {
			return
		}
		logger.Infoln("Cycle Completed -> ", i)
		// Home every n number of cycles
		if (config.GetNumHomingCycles() != 0) &&
			(int(i)%config.GetNumHomingCycles() == 0) {
			err = deps.Plc.HomingRTPCR()
			if err != nil {
				return
			}
		}
	}

	templateRunSuccess = true

	row = []interface{}{"Experiment Completed at: ", time.Now().String()}
	db.AddRowToExcel(file, db.TECSheet, row)

	row = []interface{}{"Total Time Taken by Experiment: ", time.Now().Sub(timeStarted).String()}
	db.AddRowToExcel(file, db.TECSheet, row)

	return
}
