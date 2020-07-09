package service

import (
	"context"
	"encoding/json"
	"fmt"
	"mylab/cpagent/config"
	"mylab/cpagent/db"
	"net/http"

	"github.com/gorilla/websocket"
	logger "github.com/sirupsen/logrus"
)

// use default options
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // if origin not allowed it returns 403 forbidden
	},
}

func wsHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		c, err := upgrader.Upgrade(rw, req, nil)
		if err != nil {
			logger.WithField("err", err.Error()).Error("Websocket upgrader failed")
			return
		}
		defer c.Close()

		SelfTestPLC(deps)
		HeartBeatPLC(deps)

		activeWells := config.ActiveWells("activeWells")
		go func() {
			err = <-deps.ExitCh
			logger.WithField("err", err.Error()).Error("PLC failed")

		}()
		fmt.Println("websocket started")

		for {

			if Read == 1 {
				fmt.Println("Reading epxp:", experimentID)

				// retruns all targets configured for experiment
				targetDetails, err := deps.Store.ListConfTargets(req.Context(), experimentID)
				if err != nil {
					logger.WithField("err", err.Error()).Error("Error fetching target data")
					rw.WriteHeader(http.StatusInternalServerError)
					return
				}
				DBResult, err := deps.Store.GetResult(req.Context(), experimentID)
				if err != nil {
					logger.WithField("err", err.Error()).Error("Error inserting result data")
					rw.WriteHeader(http.StatusInternalServerError)
					return
				}

				// analyseResult returns data required for ploting graph
				Finalresult := analyseResult(activeWells, targetDetails, DBResult, plcStage.CycleCount)

				var Result db.FinalResultGraph
				Result.Type = "Graph"
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
					rw.WriteHeader(http.StatusInternalServerError)
					return
				}
				//get well data
				wells, err := deps.Store.ListWells(req.Context(), experimentID)
				if err != nil {
					logger.WithField("err", err.Error()).Error("Error fetching data")
					rw.WriteHeader(http.StatusInternalServerError)
					return
				}
				if len(wells) > 0 {

					welltargets, err := deps.Store.ListWellTargets(req.Context(), experimentID)
					if err != nil {
						logger.WithField("err", err.Error()).Error("Error fetching data")
						rw.WriteHeader(http.StatusInternalServerError)
						return
					}

					fmt.Println("len wells", len(wells))
					fmt.Println("len wt", len(welltargets))

					for i, w := range wells {
						for _, t := range welltargets {
							if w.Position == t.WellPosition {
								wells[i].Targets = append(wells[i].Targets, t)
							}
						}
					}

					var Result db.FinalResultWells
					Result.Type = "Wells"
					Result.Data = append(Result.Data, wells...)
					respBytes, err := json.Marshal(Result)
					if err != nil {
						logger.WithField("err", err.Error()).Error("Error marshaling Wells data")
						rw.WriteHeader(http.StatusInternalServerError)
						return
					}
					err = c.WriteMessage(1, respBytes)
					if err != nil {
						logger.WithField("err", err.Error()).Error("Websocket failed to write")
						rw.WriteHeader(http.StatusInternalServerError)
						return
					}
				}

				Read = 0

			}

		}

	})
}

func monitorExperiment(deps Dependencies) {
	// retruns all targets configured for experiment
	targetDetails, err := deps.Store.ListConfTargets(context.Background(), experimentID)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error fetching target data")
		return
	}

	activeWells := config.ActiveWells("activeWells")

	var cycle uint16
	var previousCycle uint16

	cycle = 0
	fmt.Println("Monitor Invoked")

	// ExperimentRunning is set when experiment started & if stopped then set to false
	for ExperimentRunning {
		scan, err := deps.Plc.Monitor(cycle)
		if err != nil {
			logger.WithField("err", err.Error()).Error("Error in plc monitor")
			// send error
			return
		}
		// scan.CycleComplete returns value for same cycle even when read ones, so using previousCycle to not collect already read cycle data
		if scan.CycleComplete && scan.Cycle != previousCycle {

			// write to db
			// makeResult returns data in DB result format
			result := makeResult(activeWells, scan, targetDetails, experimentID)

			// insert current cycle result into Database
			DBResult, err := deps.Store.InsertResult(context.Background(), result)
			if err != nil {
				logger.WithField("err", err.Error()).Error("Error inserting result data")
				// send error
				return
			}

			// getLastCycleResult
			var LastCycleResult []db.Result
			for _, r := range DBResult {
				if r.Cycle == scan.Cycle {
					LastCycleResult = append(LastCycleResult, r)
				}
			}
			// color analysis
			wells, err := deps.Store.ListWells(context.Background(), experimentID)
			if err != nil {
				logger.WithField("err", err.Error()).Error("Error fetching data")
				// send error
				return
			}

			welltargets, err := deps.Store.ListWellTargets(context.Background(), experimentID)
			if err != nil {
				logger.WithField("err", err.Error()).Error("Error fetching data")
				// send error
				return
			}

			// send data to color analysis

			fmt.Printf("len(LastCycleResult): %v\n len(welltargets): %v\n len(wells): %v \n", len(LastCycleResult), len(welltargets), len(wells))
			targets, well := WellColorAnalysis(LastCycleResult, welltargets, wells, scan.Cycle, plcStage.CycleCount)

			// update color in well
			if len(well) > 0 {
				for _, w := range well {
					err = deps.Store.UpdateColorWell(context.Background(), w.ColorCode, w.ID)
					if err != nil {
						// send error
						logger.WithField("err", err.Error()).Error("Error upsert wells")
						return
					}
				}
			}

			// update ct value in DB
			fmt.Println("Updated len(targets)", len(targets))
			_, err = deps.Store.UpsertWellTargets(context.Background(), targets, experimentID)
			if err != nil {
				// send error
				logger.WithField("err", err.Error()).Error("Error upsert wells")
				return
			}

			if scan.Cycle == plcStage.CycleCount {
				err = deps.Store.UpdateStopTimeExperiments(context.Background(), time.Now(), experimentID)
				if err != nil {
					logger.WithField("err", err.Error()).Error("Error fetching data")
					rw.WriteHeader(http.StatusInternalServerError)
					return
				}
				// last cycle socket closed
				ExperimentRunning = false
				break
			}
			Read = 1
			fmt.Println("cycle:", scan.Cycle)
			cycle++
			previousCycle++
		}
	}

}

func SelfTestPLC(deps Dependencies) {
	deps.Plc.SelfTest()
}

func HeartBeatPLC(deps Dependencies) {
	deps.Plc.HeartBeat()
}
