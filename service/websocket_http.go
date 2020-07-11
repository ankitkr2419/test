package service

import (
	"context"
	"encoding/json"
	"mylab/cpagent/config"
	"mylab/cpagent/db"
	"net/http"
	"time"

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

		// The exit plan incase there is a feedback from the driver to abort/exit
		go func() {
			for {
				err = <-deps.ExitCh
				logger.WithField("err", err.Error()).Error("PLC Driver has requested exit")
				ExperimentRunning = false // on pre-emptive stop
			// TODO: Handle exit gracefully
			// We need to call the API on the Web to display the error and restart, abort or call service!
			}
			logger.WithField("msg", "Exit").Info("WebSocket Thread Reading PLC Err Channel Exit")

		}()

		go func() {
			// read channel from websocket err
			err = <-deps.WsErrCh
			logger.WithField("err", err.Error()).Error("Monitor has requested exit")
			logger.WithField("msg", "Exit").Info("WebSocket Thread Reading WB Err Channel Exit")

		}()

		for {

			msg := <-deps.WsMsgCh
			switch msg {
			case "read":

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
				logger.WithField("msg", "Sent").Info("Graph data")
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
			case "stop":
				// send exp stop data

				latestE, err := deps.Store.ShowExperiment(req.Context(), experimentID)
				if err != nil {
					rw.WriteHeader(http.StatusInternalServerError)
					logger.WithField("err", err.Error()).Error("Error get experiment")
					return
				}
				var result db.FinalResultOnSuccess
				result.Type = "Success"
				result.Data = latestE

				respBytes, err := json.Marshal(result)
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
				break

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

	// ExperimentRunning is set when experiment started & if stopped then set to false
	for ExperimentRunning {
		scan, err := deps.Plc.Monitor(cycle)
		if err != nil {
			logger.WithField("err", err.Error()).Error("Error in plc monitor")
			// send error
			deps.WsErrCh <- err
			return
		}
		
		// scan.CycleComplete returns value for same cycle even when read ones, so using previousCycle to not collect already read cycle data
		if scan.CycleComplete && scan.Cycle != previousCycle {
			logger.Info("Received data for cycle: ",scan.Cycle)

			// write to db
			// makeResult returns data in DB result format
			result := makeResult(activeWells, scan, targetDetails, experimentID)

			// insert current cycle result into Database
			DBResult, err := deps.Store.InsertResult(context.Background(), result)
			if err != nil {
				logger.WithField("err", err.Error()).Error("Error inserting result data")
				// send error
				deps.WsErrCh <- err
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
				deps.WsErrCh <- err
				return
			}

			welltargets, err := deps.Store.ListWellTargets(context.Background(), experimentID)
			if err != nil {
				logger.WithField("err", err.Error()).Error("Error fetching data")
				deps.WsErrCh <- err
				// send error
				return
			}

			// send data to color analysis

			targets, well := WellColorAnalysis(LastCycleResult, welltargets, wells, scan.Cycle, plcStage.CycleCount)

			// update color in well
			if len(well) > 0 {
				for _, w := range well {
					err = deps.Store.UpdateColorWell(context.Background(), w.ColorCode, w.ID)
					if err != nil {
						// send error
						logger.WithField("err", err.Error()).Error("Error upsert wells")
						deps.WsErrCh <- err
						return
					}
				}
			}

			// update ct value in DB
			_, err = deps.Store.UpsertWellTargets(context.Background(), targets, experimentID)
			if err != nil {
				// send error
				logger.WithField("err", err.Error()).Error("Error upsert wells")
				deps.WsErrCh <- err
				return
			}
			deps.WsMsgCh <- "read"
			if scan.Cycle == plcStage.CycleCount {
				logger.Info("Received data for last cycle")
				err = deps.Store.UpdateStopTimeExperiments(context.Background(), time.Now(), experimentID)
				if err != nil {
					logger.WithField("err", err.Error()).Error("Error fetching data")
					// send error
					deps.WsErrCh <- err
					return
				}
				// last cycle socket closed
				deps.WsMsgCh <- "stop"
				ExperimentRunning = false
				break
			}

			cycle++
			previousCycle++
		}
	}
	logger.Info("Stop monitoring experiment")

}

func SelfTestPLC(deps Dependencies) {
	deps.Plc.SelfTest()
}

func HeartBeatPLC(deps Dependencies) {
	deps.Plc.HeartBeat()
}
