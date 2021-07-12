package service

import (
	"context"
	"encoding/json"
	"fmt"
	"mylab/cpagent/config"
	"mylab/cpagent/db"
	"mylab/cpagent/plc"
	"mylab/cpagent/tec"
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
		file := plc.GetExcelFile(tec.LogsPath, fmt.Sprintf("output_%v", expID))

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

		t, err := deps.Store.ShowTemplate(req.Context(), e.TemplateID)
		if err != nil {
			logger.WithField("err", err.Error()).Error("Error fetching template data")
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}

		// Lid Temp is multiplied by 10 for PLC can't handle floats
		plcStage.IdealLidTemp = uint16(t.LidTemp * 10)
		e.Result = "running"
		err = deps.Store.UpdateStartTimeExperiments(req.Context(), time.Now(), expID, plcStage.CycleCount, e.Result)
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

		plc.AddMergeRowToExcel(file, plc.RTPCRSheet, heading, len(config.ActiveWells("activeWells")))

		row := []interface{}{"well positions"}
		for _, v := range config.ActiveWells("activeWells") {
			row = append(row, v)
		}
		plc.AddRowToExcel(file, plc.RTPCRSheet, row)

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

	// Home the TEC
	// Reset is implicit in Homing
	err = deps.Plc.HomingRTPCR()
	if err != nil {
		return
	}

	// TODO: Lid Temp reaching
	err = deps.Plc.SetLidTemp(p.IdealLidTemp)
	if err != nil {
		return
	}

	//invoke monitor
	go monitorExperiment(deps, file)

	// Start line
	headers := []interface{}{"Description", "Time Taken", "Expected Time", "Initial Temp", "Final Temp", "Ramp"}
	plc.AddRowToExcel(file, plc.TECSheet, headers)

	timeStarted := time.Now()
	row := []interface{}{"Experiment Started at: ", timeStarted.String()}
	plc.AddRowToExcel(file, plc.TempLogs, row)

	row = []interface{}{"Holding Stage About to start"}
	plc.AddRowToExcel(file, plc.TECSheet, row)

	row = []interface{}{"timestamp", "current Temperature", "lid Temperature"}
	plc.AddRowToExcel(file, plc.TempLogs, row)

	// Run Holding Stage
	logger.Infoln("Holding Stage Started")
	err = deps.Tec.RunStage(p.Holding, deps.Plc, file, 0)
	if err != nil {
		return
	}

	// Run Cycle Stage
	row = []interface{}{"Cycle Stage About to start"}
	plc.AddRowToExcel(file, plc.TECSheet, row)

	for i := uint16(1); i <= p.CycleCount; i++ {
		logger.Infoln("Started Cycle->", i)
		err = deps.Tec.RunStage(p.Cycle, deps.Plc, file, i)
		if err != nil {
			return
		}
		logger.Infoln("Cycle Completed -> ", i)
	}

	row = []interface{}{"Experiment Completed at: ", time.Now().String()}
	plc.AddRowToExcel(file, plc.TECSheet, row)

	row = []interface{}{"Total Time Taken by Experiment: ", time.Now().Sub(timeStarted).String()}
	plc.AddRowToExcel(file, plc.TECSheet, row)

	return
}
