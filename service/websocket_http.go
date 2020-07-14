package service

import (
	"context"
	"encoding/json"
	"errors"
	"mylab/cpagent/db"
	"mylab/cpagent/plc"
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

		logger.Info("WebSocket Invoked")

		c, err := upgrader.Upgrade(rw, req, nil)
		if err != nil {
			logger.WithField("err", err.Error()).Error("Websocket upgrader failed")
			return
		}
		defer c.Close()

		deps.Plc.SelfTest()
		deps.Plc.HeartBeat()

		for {

			select {
			case msg := <-deps.WsMsgCh:

				if msg == "read" {

					sendGraph(deps, rw, c)
					sendWells(deps, rw, c)

				} else if msg == "stop" {

					sendOnSuccess(deps, rw, c)

				} else if msg == "read_temp" {

					sendTemperature(deps, rw, c)

				}

			case err = <-deps.ExitCh:

				if err.Error() == "PCR Aborted" {

					// on pre-emptive stop
					experimentRunning = false

				}

				logger.WithField("err", err.Error()).Error("PLC Driver has requested exit")

				sendOnFail(err.Error(), rw, c)

			case err = <-deps.WsErrCh:

				logger.WithField("err", err.Error()).Error("Monitor has requested exit")

				sendOnFail(err.Error(), rw, c)

			}

		}

	})
}

func sendGraph(deps Dependencies, rw http.ResponseWriter, c *websocket.Conn) {

	graphResult, err := getGraph(deps)
	if err != nil {
		logger.WithField("err", err.Error()).Error("error in fetching data")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = c.WriteMessage(1, graphResult)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Websocket failed to write")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	logger.WithField("data", "Graph").Info("Websocket send Data")
}

func sendWells(deps Dependencies, rw http.ResponseWriter, c *websocket.Conn) {

	WellResult, err := getColorCodedWells(deps)
	if err != nil {
		if err.Error() == "Wells not configured" {
			logger.Info("Wells not configured")
		} else {
			logger.WithField("err", err.Error()).Error("error in fetching data")
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}
	} else {
		err = c.WriteMessage(1, WellResult)
		if err != nil {
			logger.WithField("err", err.Error()).Error("Websocket failed to write")
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}
		logger.WithField("data", "Well").Info("Websocket send Data")
	}
}

func sendOnSuccess(deps Dependencies, rw http.ResponseWriter, c *websocket.Conn) {

	respBytes, err := getExperimentDetails(deps)
	if err != nil {
		logger.WithField("err", err.Error()).Error("error in fetching data")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = c.WriteMessage(1, respBytes)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Websocket failed to write")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	logger.WithField("data", "Success").Info("Websocket send Data")

}

func sendTemperature(deps Dependencies, rw http.ResponseWriter, c *websocket.Conn) {

	respBytes, err := getTemperatureDetails(deps)
	if err != nil {
		logger.WithField("err", err.Error()).Error("error in fetching data")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	err = c.WriteMessage(1, respBytes)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Websocket failed to write")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	logger.WithField("data", "Temperature").Info("Websocket send Data")

}

func sendOnFail(msg string, rw http.ResponseWriter, c *websocket.Conn) {

	r := resultOnFail{
		Type: "Fail",
		Data: msg,
	}

	respBytes, err := json.Marshal(r)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error marshaling result data")
		return
	}
	err = c.WriteMessage(1, respBytes)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Websocket failed to write")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	logger.WithField("data", "Fail").Info("Websocket send Data")

}

func getGraph(deps Dependencies) (respBytes []byte, err error) {

	DBResult, err := deps.Store.GetResult(context.Background(), experimentValues.experimentID)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error fetching result data")
		return
	}

	// analyseResult returns data required for ploting graph
	Finalresult := analyseResult(DBResult)

	Result := resultGraph{
		Type: "Graph",
		Data: Finalresult,
	}

	respBytes, err = json.Marshal(Result)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error marshaling graph data")
		return
	}

	return

}

func getColorCodedWells(deps Dependencies) (respBytes []byte, err error) {

	// list wells from DB
	wells, err := deps.Store.ListWells(context.Background(), experimentValues.experimentID)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error fetching data")
		return
	}
	if len(wells) > 0 {
		var welltargets []db.WellTarget

		welltargets, err = deps.Store.ListWellTargets(context.Background(), experimentValues.experimentID)
		if err != nil {
			logger.WithField("err", err.Error()).Error("Error fetching data")
			return
		}

		for i, w := range wells {
			for _, t := range welltargets {
				if w.Position == t.WellPosition {
					wells[i].Targets = append(wells[i].Targets, t)
				}
			}
		}

		Result := resultWells{
			Type: "Wells",
			Data: wells,
		}

		respBytes, err = json.Marshal(Result)
		if err != nil {
			logger.WithField("err", err.Error()).Error("Error marshaling Wells data")
			return
		}
		return
	}
	err = errors.New("Wells not configured")
	return
}

func getExperimentDetails(deps Dependencies) (respBytes []byte, err error) {
	latestE, err := deps.Store.ShowExperiment(context.Background(), experimentValues.experimentID)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error get experiment")
		return
	}

	result := resultOnSuccess{
		Type: "Success",
		Data: latestE,
	}

	respBytes, err = json.Marshal(result)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error marshaling experiment data")
		return
	}
	return
}

func getTemperatureDetails(deps Dependencies) (respBytes []byte, err error) {
	Temp, err := deps.Store.ListResultTemperature(context.Background(), experimentValues.experimentID)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error get experiment")
		return
	}

	data := makeResultTemp(Temp)

	result := resultTemperature{
		Type: "Temperature",
		Data: data,
	}

	respBytes, err = json.Marshal(result)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error marshaling result temp data")
		return
	}

	return
}

func monitorExperiment(deps Dependencies) {

	var cycle uint16
	var previousCycle uint16
	var previousTemp float32

	cycle = 0

	// experimentRunning is set when experiment started & if stopped then set to false
	for experimentRunning {

		scan, err := deps.Plc.Monitor(cycle)
		if err != nil {
			logger.WithField("err", err.Error()).Error("Error in plc monitor")
			deps.WsErrCh <- err
			return
		}
		// fmt.Printf("scan: %+v\n", scan)

		// scan.CycleComplete returns value for same cycle even when read ones, so using previousCycle to not collect already read cycle data
		if scan.CycleComplete && scan.Cycle != previousCycle {

			logger.Info("Received Emmissions from PLC for cycle: ", scan.Cycle)

			DBResult, err := WriteResult(deps, scan)
			if err != nil {
				return
			}
			WriteColorCTValues(deps, DBResult, scan)
			if err != nil {
				return
			}
			deps.WsMsgCh <- "read"
			if scan.Cycle == experimentValues.plcStage.CycleCount {
				err = deps.Store.UpdateStopTimeExperiments(context.Background(), time.Now(), experimentValues.experimentID)
				if err != nil {
					logger.WithField("err", err.Error()).Error("Error updating stop time")
					deps.WsErrCh <- err
					return
				}
				deps.WsMsgCh <- "stop"
				experimentRunning = false
				break
			}

			cycle++
			previousCycle++
		}

		// writes temp on every step against time in DB
		if scan.Temp != previousTemp {
			err = WriteResultTemperature(deps, scan)
			if err != nil {
				return
			} else {
				previousTemp = scan.Temp
				deps.WsMsgCh <- "read_temp"
			}
		}

	}
	logger.Info("Stop monitoring experiment")
}

func WriteResult(deps Dependencies, scan plc.Scan) (DBResult []db.Result, err error) {

	// makeResult returns data in DB result format
	result := makeResult(scan)

	// insert current cycle result into Database
	DBResult, err = deps.Store.InsertResult(context.Background(), result)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error inserting result data")
		// send error
		deps.WsErrCh <- err
		return
	}
	return
}

func WriteColorCTValues(deps Dependencies, DBResult []db.Result, scan plc.Scan) (err error) {

	// getLastCycleResult
	var LastCycleResult []db.Result
	for _, r := range DBResult {
		if r.Cycle == scan.Cycle {
			LastCycleResult = append(LastCycleResult, r)
		}
	}

	// color analysis
	wells, err := deps.Store.ListWells(context.Background(), experimentValues.experimentID)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error fetching data")
		// send error
		deps.WsErrCh <- err
		return
	}

	welltargets, err := deps.Store.ListWellTargets(context.Background(), experimentValues.experimentID)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error fetching data")
		deps.WsErrCh <- err
		// send error
		return
	}

	// send data to color analysis

	targets, well := wellColorAnalysis(LastCycleResult, welltargets, wells, scan.Cycle)

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
	_, err = deps.Store.UpsertWellTargets(context.Background(), targets, experimentValues.experimentID)
	if err != nil {
		// send error
		logger.WithField("err", err.Error()).Error("Error upsert wells")
		deps.WsErrCh <- err
		return
	}
	return
}

func WriteResultTemperature(deps Dependencies, scan plc.Scan) (err error) {

	// makeResultTemp returns data in DB resultTemp format

	resultTemp := db.ResultTemperature{
		ExperimentID: experimentValues.experimentID,
		Temp:         scan.Temp,
		LidTemp:      scan.LidTemp,
		Cycle:        scan.Cycle,
	}

	// insert every cycle  result temp into Database
	err = deps.Store.InsertResultTemperature(context.Background(), resultTemp)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error inserting resultTemp data")
		// send error
		deps.WsErrCh <- err
		return
	}
	return
}
