package service

import (
	"context"
	"mylab/cpagent/config"
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

		
		for {

			go func() {
		        	err = <- deps.ExitCh
				logger.WithField("err", err.Error()).Error("PLC failed")

			}

			if Read == 1 {
				DBResult, err := deps.Store.GetResult(req.Context(), ExperimentID)
				if err != nil {
					logger.WithField("err", err.Error()).Error("Error inserting result data")
					rw.WriteHeader(http.StatusInternalServerError)
					return
				}
				
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
				err = c.WriteMessage(1, Result)
				if err != nil {
					logger.WithField("err", err.Error()).Error("Websocket failed to write")
					rw.WriteHeader(http.StatusInternalServerError)
					return
				}

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
			return
		}
		// scan.CycleComplete returns value for same cycle even when read ones, so using previousCycle to not collect already read cycle data
		if scan.CycleComplete && scan.Cycle != previousCycle {

			// write to db
			// makeResult returns data in DB result format
			result := makeResult(activeWells, scan, targetDetails, experimentID)

			// insert current cycle result into Database
			_, err := deps.Store.InsertResult(context.Background(), result)
			if err != nil {
				logger.WithField("err", err.Error()).Error("Error inserting result data")
				return
			}

			if scan.Cycle == plcStage.CycleCount {
				// last cycle socket closed
				ExperimentRunning = false
				break
			}
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
