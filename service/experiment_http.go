package service

import (
	"encoding/json"
	"fmt"
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
		plcStage = makePLCStage(ss)
		fmt.Printf("plcStage: %+v ", plcStage)
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

		err = deps.Store.UpdateStartTimeExperiments(req.Context(), time.Now(), expID)
		if err != nil {
			logger.WithField("err", err.Error()).Error("Error fetching data")
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}

		// experimentID set
		experimentID = expID

		rw.Header().Add("Content-Type", "application/json")
		rw.WriteHeader(http.StatusOK)
		rw.Write([]byte(`{"msg":"experiment started"}`))
	})
}

func monitorExperimentHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		// if origin not allowed it returns 403
		upgrader.CheckOrigin = func(r *http.Request) bool { return true }

		c, err := upgrader.Upgrade(rw, req, nil)
		if err != nil {
			logger.WithField("err", err.Error()).Error("Websocket upgrader failed")
			return
		}
		defer c.Close()

		// retruns all targets configured for experiment
		targetDetails, err := deps.Store.ListConfTargets(req.Context(), experimentID)
		if err != nil {
			logger.WithField("err", err.Error()).Error("Error fetching target data")
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}

		activeWells := config.ActiveWells("activeWells")

		var cycle uint16
		var previousCycle uint16

		cycle = 0
		fmt.Println("Monitor Invoked with cycle:", cycle)
		for {
			scan, err := deps.Plc.Monitor(cycle)
			if err != nil {
				logger.WithField("err", err.Error()).Error("Error in plc monitor")
				return
			}
			// scan.CycleComplete returns value for same cycle even when read ones, so using previousCycle to not collect already read cycle data
			if scan.CycleComplete && scan.Cycle != previousCycle {

				// write to db & serve
				fmt.Println("cycle: ", cycle)
				fmt.Println("previousCycle: ", previousCycle)
				fmt.Println("scan cycle:", scan.Cycle)
				fmt.Printf("scan: %+v ", scan)

				// makeResult returns data in DB result format
				result := makeResult(activeWells, scan, targetDetails, experimentID)
				// fmt.Printf("%+v",result)

				// insert current cycle result into Database
				DBResult, err := deps.Store.InsertResult(req.Context(), result)
				if err != nil {
					logger.WithField("err", err.Error()).Error("Error inserting result data")
					rw.WriteHeader(http.StatusInternalServerError)
					return
				}

				fmt.Printf("call analyseResult ac %v tr %v db %v\n", len(activeWells), len(targetDetails), len(DBResult))

				// analyseResult returns data required for ploting graph
				Finalresult := analyseResult(activeWells, targetDetails, DBResult, plcStage.CycleCount)

				var Result db.FinalResult
				Result.MaxThreshold = maxThreshold
				Result.Data = append(Result.Data, Finalresult...)

				respBytes, err := json.Marshal(Result)
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
					fmt.Println("socket closed")
					break
				}
				// if scan.Cycle == 2 {
				//			break
				//		}
				cycle++
				previousCycle++
			}
		}
	})
}
