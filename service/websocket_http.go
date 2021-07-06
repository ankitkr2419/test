package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"mylab/cpagent/db"
	"mylab/cpagent/plc"
	"net/http"
	"strings"
	"time"

	"strconv"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"github.com/google/uuid"
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

		for {

			select {
			case msg := <-deps.WsMsgCh:

				msgs := msgDivision(msg)

				if msg == "read" {

					sendGraph(deps, rw, c)
					sendWells(deps, rw, c)

				} else if msg == "stop" {

					sendOnSuccess(deps, rw, c)

				} else if msg == "read_temp" {

					sendTemperature(deps, rw, c)

				} else if strings.EqualFold(msgs[0], "progress") {

					monitorOperation(deps, rw, c, msgs)

				} else if strings.EqualFold(msgs[0], "success") {

					successOperation(deps, rw, c, msgs)
				}

			case err = <-deps.ExitCh:
				var errortype, msg string
				if err.Error() == "PCR Aborted" {

					// on pre-emptive stop
					experimentRunning = false
					errortype = "ErrorPCRAborted"
					msg = "Experiment aborted by user"

				} else if err.Error() == "PCR Stopped" {
					errortype = "ErrorPCRStopped"
					msg = "PCR completed experiment"

				} else if err.Error() == "PCR Dead" {
					errortype = "ErrorPCRDead"
					msg = "Unable to connect to Hardware"

				}

				logger.WithField("err", err.Error()).Error("PLC Driver has requested exit")

				//log in Db notifications
				go LogNotification(deps, fmt.Sprintf("ExperimentId: %v, %v", experimentValues.experimentID, msg))

				sendOnFail(msg, errortype, rw, c)

			case err = <-deps.WsErrCh:

				errs := msgDivision(err.Error())
				if errs[0] == "ErrorExtractionMonitor" {

					errorData := plc.WSError{
						Message: errs[2],
						Deck:    errs[1],
					}
					wsError, err := json.Marshal(errorData)
					if err != nil {
						logger.Errorf("error in marshalling web socket data %v", err.Error())
						return
					}
					sendOnFail(string(wsError), errs[0], rw, c)
				} else {
					logger.WithField("err", err.Error()).Error("Monitor has requested exit")
					var errortype = "ErrorPCRMonitor"

					go LogNotification(deps, fmt.Sprintf("ExperimentId: %v, %v", experimentValues.experimentID, err.Error()))

					sendOnFail(err.Error(), errortype, rw, c)
				}

			}

		}

	})
}

func sendGraph(deps Dependencies, rw http.ResponseWriter, c *websocket.Conn) {

	graphResult, err := getGraph(deps, experimentValues.experimentID, experimentValues.activeWells, experimentValues.targets, experimentValues.plcStage.CycleCount)
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

	respBytes, err := getTemperatureDetails(deps, experimentValues.experimentID)
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

func sendOnFail(msg, errortype string, rw http.ResponseWriter, c *websocket.Conn) {

	r := resultOnFail{
		Type: errortype,
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

	logger.WithField("data", "Fail").Infoln("Websocket send Data")

}

func getGraph(deps Dependencies, experimentID uuid.UUID, wells []int32, targets []db.TargetDetails, t_cycles uint16) (respBytes []byte, err error) {

	DBResult, err := deps.Store.GetResult(context.Background(), experimentID)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error fetching result data")
		return
	}

	Finalresult := make([]graph, 0)

	if len(DBResult) > 0 {
		// analyseResult returns data required for ploting graph
		Finalresult = analyseResult(DBResult, wells, targets, t_cycles)
	}

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

					// show scaled value for graph
					if t.CT != "" && t.CT != undetermine {
						ct, _ := strconv.ParseFloat(t.CT, 32)
						// %.1f takes decimal value upto 1
						t.CT = fmt.Sprintf("%.1f", scaleThreshold(float32(ct)))
					}

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

func getTemperatureDetails(deps Dependencies, experimentID uuid.UUID) (respBytes []byte, err error) {
	Temp, err := deps.Store.ListExperimentTemperature(context.Background(), experimentID)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error get experiment")
		return
	}

	result := experimentTemperature{
		Type: "Temperature",
		Data: Temp,
	}

	respBytes, err = json.Marshal(result)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error marshaling result temp data")
		return
	}

	return
}

func monitorExperiment(deps Dependencies, file *excelize.File) {

	var cycle uint16
	var previousCycle uint16

	cycle = 0
	// Start line
	heading := []string{"ExperimentID", "Well Position", "Cycle", "Dye Position", "TargetID", "FValue", "Temperature"}
	plc.AddRowToExcel(file, plc.RTPCRSheet, heading)
	// experimentRunning is set when experiment started & if stopped then set to false
	for experimentRunning {
		time.Sleep(500 * time.Millisecond)

		scan, err := deps.Plc.Monitor(cycle)
		if err != nil {
			logger.WithField("err", err.Error()).Error("Error in plc monitor")
			deps.WsErrCh <- err
			return
		}
		//Add to excel
		row := []string{time.Now().String(), fmt.Sprintf("%v", scan.Temp), fmt.Sprintf("%v", scan.LidTemp)}
		plc.AddRowToExcel(file, plc.TempLogs, row)

		// writes temp on every step against time in DB
		err = WriteExperimentTemperature(deps, scan)
		if err != nil {
			fmt.Println("Write Exp Temp Error")
			return
		} else {
			deps.WsMsgCh <- "read_temp"
		}
		// scan.CycleComplete returns value for same cycle even when read ones, so using previousCycle to not collect already read cycle data
		if plc.HeatingCycleComplete && scan.CycleComplete {

			logger.Info("Received Emmissions from PLC for cycle: ", scan.Cycle, scan)

			DBResult, err := WriteResult(deps, scan, file)
			if err != nil {
				logger.WithField("err", err.Error()).Error("Error in dbresult")
				return
			}
			WriteColorCTValues(deps, DBResult, scan)
			if err != nil {
				logger.WithField("err", err.Error()).Error("Error in ct values")
				return
			}
			deps.WsMsgCh <- "read"
			if scan.Cycle == experimentValues.plcStage.CycleCount {
				err = deps.Store.UpdateStopTimeExperiments(context.Background(), time.Now(), experimentValues.experimentID, "successful")
				if err != nil {
					logger.WithField("err", err.Error()).Error("Error updating stop time")
					deps.WsErrCh <- err
					return
				}
				deps.WsMsgCh <- "stop"
				fmt.Println("exit chan 2--------------------------------")

				experimentRunning = false
				break
			}
			cycle++
			previousCycle++
			plc.CycleComplete = false
			plc.HeatingCycleComplete = false

		}
		// adding delay of 0.5s to reduce the cpu usage
	}
	logger.Info("Stop monitoring experiment")
}

func WriteResult(deps Dependencies, scan plc.Scan, file *excelize.File) (DBResult []db.Result, err error) {

	// makeResult returns data in DB result format
	result := makeResult(scan, file)

	// for cycle one , preceed default data [0,0] for cycle 0 ,needed to plot the graph
	if scan.Cycle == 1 {
		addResultForZerothCycle(deps, result)
	}

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

//addResultForZerothCycle for graph
func addResultForZerothCycle(deps Dependencies, result []db.Result) {

	// set default value [0,0]
	var zerothResult []db.Result
	for _, v := range result {
		var r db.Result

		// copy all fields
		r = v

		// set cycle & FValue to [0,0]
		r.Cycle = 0
		r.FValue = 0

		zerothResult = append(zerothResult, r)

	}

	// insert current cycle result into Database
	_, err := deps.Store.InsertResult(context.Background(), zerothResult)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error inserting result data")
		// send error
		deps.WsErrCh <- err
		return
	}
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
	_, err = deps.Store.UpsertWellTargets(context.Background(), targets, experimentValues.experimentID, false)
	if err != nil {
		// send error
		logger.WithField("err", err.Error()).Error("Error upsert wells")
		deps.WsErrCh <- err
		return
	}
	return
}

func WriteExperimentTemperature(deps Dependencies, scan plc.Scan) (err error) {

	// makeexpTemp returns data in DB expTemp format
	expTemp := db.ExperimentTemperature{
		ExperimentID: experimentValues.experimentID,
		Temp:         scan.Temp,
		LidTemp:      scan.LidTemp,
		Cycle:        scan.Cycle,
	}

	// insert every cycle  result temp into Database
	err = deps.Store.InsertExperimentTemperature(context.Background(), expTemp)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error inserting experiment_Temperatures data")
		// send error
		deps.WsErrCh <- err
		return
	}
	return
}

func monitorOperation(deps Dependencies, rw http.ResponseWriter, c *websocket.Conn, msg []string) (err error) {

	monitorOpr := oprProgress{
		Type: fmt.Sprintf("PROGRESS_%s", strings.ToUpper(msg[1])),
		Data: fmt.Sprintf("%s", msg[2]),
	}

	respBytes, err := json.Marshal(monitorOpr)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error marshaling result temp data")
		return
	}

	err = c.WriteMessage(1, respBytes)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Websocket failed to write")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	logger.WithField("data", monitorOpr).Info("Websocket send Data")
	return
}

func successOperation(deps Dependencies, rw http.ResponseWriter, c *websocket.Conn, msg []string) (err error) {

	successOpr := oprSuccess{
		Type: fmt.Sprintf("SUCCESS_%s", strings.ToUpper(msg[1])),
		Data: fmt.Sprintf("%s", msg[2]),
	}

	respBytes, err := json.Marshal(successOpr)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error marshaling result temp data")
		return
	}

	err = c.WriteMessage(1, respBytes)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Websocket failed to write")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	logger.WithField("data", successOpr).Info("Websocket send Data")
	return
}

func msgDivision(msg string) (msgs []string) {
	msgs = strings.SplitN(msg, "_", 3)
	return
}
