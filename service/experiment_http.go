package service

import (
	"encoding/json"
	"mylab/cpagent/config"
	"mylab/cpagent/db"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	logger "github.com/sirupsen/logrus"
)

var upgrader = websocket.Upgrader{} // use default options

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
			logger.WithField("err", err.Error()).Error("Error create template")
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
		experimentID, err := parseUUID(vars["experiment_id"])
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			return
		}

		e, err := deps.Store.ShowExperiment(req.Context(), experimentID)
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
		plcStage = makePLCStage(ss)

		err = deps.Plc.ConfigureRun(plcStage)
		if err != nil {
			logger.WithField("err", err.Error()).Error("Error in ConfigureRun")
			return
		}

		err = deps.Plc.Start()
		if err != nil {
			logger.WithField("err", err.Error()).Error("Error in plc start")
			return
		}

		err = deps.Store.UpdateStartTimeExperiments(req.Context(), time.Now(), experimentID)
		if err != nil {
			logger.WithField("err", err.Error()).Error("Error fetching data")
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}

		rw.Header().Add("Content-Type", "application/json")
		rw.WriteHeader(http.StatusOK)
		rw.Write([]byte(`{"msg":"experiment started"}`))
	})
}

func monitorExperimentHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		c, err := upgrader.Upgrade(rw, req, nil)
		if err != nil {
			logger.WithField("err", err.Error()).Error("Websocket upgrader failed")
			return
		}
		defer c.Close()

		vars := mux.Vars(req)
		experimentID, err := parseUUID(vars["experiment_id"])
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			return
		}

		targetDetails, err := deps.Store.ListConfTargets(req.Context(), experimentID)
		if err != nil {
			logger.WithField("err", err.Error()).Error("Error fetching target data")
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}

		activeWells := config.ActiveWells("activeWells")

		var cycle uint16

		for {
			scan, err := deps.Plc.Monitor(cycle)
			if err != nil {
				logger.WithField("err", err.Error()).Error("Error in plc monitor")
				return
			}

			if scan.CycleComplete {
				// write to db & serve
				result := makeResult(activeWells, scan, targetDetails, experimentID)
				err := deps.Store.InsertResult(req.Context(), result)
				if err != nil {
					logger.WithField("err", err.Error()).Error("Error inserting result data")
					rw.WriteHeader(http.StatusInternalServerError)
					return
				}

				respBytes, err := json.Marshal(result)
				if err != nil {
					logger.WithField("err", err.Error()).Error("Error marshaling experiment data")
					rw.WriteHeader(http.StatusInternalServerError)
					return
				}
				err = c.WriteMessage(1, respBytes)
				if err != nil {
					logger.WithField("err", err.Error()).Error("Websocket failed to write")
					break
				}

				if scan.Cycle == plcStage.CycleCount {
					// last cycle
					break
				}

				cycle++
			}
		}
	})
}
