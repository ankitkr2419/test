package simulator

import (
	"fmt"
	"math"

	"mylab/cpagent/config"
	"mylab/cpagent/db"
	"mylab/cpagent/plc"
	"mylab/cpagent/tec"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
	logger "github.com/sirupsen/logrus"
)

type Simulator struct {
	ExitCh  chan error
	WsMsgCh chan string
	wsErrch chan error
}

func NewSimulatorDriver(wsMsgCh chan string, wsErrch chan error, exit chan error, test bool) tec.Driver {

	tec := Simulator{}
	tec.ExitCh = exit
	tec.WsMsgCh = wsMsgCh
	tec.wsErrch = wsErrch

	return &tec // tec Driver
}

var errorCheckStopped, tecInProgress bool
var prevTemp float32 = 27.0

// TODO: Implement Simulator
func (t *Simulator) InitiateTEC() (err error) {
	logger.Infoln("Simulating TEC Initiation...")
	logger.Infoln("TEC Initiation Successful")
	return nil
}

func (t *Simulator) SetTempAndRamp(ts tec.TECTempSet) (err error) {
	currentTemp := prevTemp
	// Reach the target temperature with given ramp rate
	timeRequiredInSecs := math.Abs(ts.TargetTemperature-float64(currentTemp)) / ts.TargetRampRate

	timeStarted := time.Now()
	var tempReached bool
	for !tempReached {
		time.Sleep(1 * time.Second)
		if float32(ts.TargetTemperature) > currentTemp {
			currentTemp += float32(ts.TargetRampRate)
		} else {
			currentTemp -= float32(ts.TargetRampRate)
		}

		logger.Infoln("New Temperature reached: ", currentTemp)
		plc.CurrentCycleTemperature = currentTemp
		if time.Now().Sub(timeStarted).Seconds() > timeRequiredInSecs {
			currentTemp = float32(ts.TargetTemperature)
			tempReached = true
		}
	}
	logger.Infoln("Target Temp reached: ", currentTemp)

	return nil
}

func (t *Simulator) AutoTune() (err error) {
	logger.Infoln("Simulating TEC Auto Tuning...")
	for i := 0; i < 101; i = i + 5 {
		time.Sleep(100 * time.Millisecond)
		logger.Infoln("Auto Tuning Percent: ", i)
	}
	return nil
}

func (t *Simulator) ResetDevice() (err error) {
	logger.Infoln("Simulating TEC Reset Device...")
	logger.Infoln("TEC Reset Device Successful")
	return nil
}

func (t *Simulator) ReachRoomTemp() error {
	logger.Infoln("Reaching Room Temp")
	ts := tec.TECTempSet{
		TargetTemperature: config.GetRoomTemp(),
		TargetRampRate:    tec.RoomTempRamp,
	}
	t.SetTempAndRamp(ts)
	logger.Infoln("Room Temp Reached")
	return nil
}

func (t *Simulator) GetAllTEC() (err error) {
	logger.Infoln("Simulating Get All TEC Data...")
	logger.Infoln("It's simulator so nothing to get")
	return nil
}

func (t *Simulator) TestRun(plcDeps plc.Driver) (err error) {
	p := plc.Stage{
		Holding: []plc.Step{
			plc.Step{TargetTemp: 65.3, RampUpTemp: 10, HoldTime: 5, DataCapture: false},
			plc.Step{TargetTemp: 85.3, RampUpTemp: 10, HoldTime: 5, DataCapture: false},
			plc.Step{TargetTemp: 95, RampUpTemp: 10, HoldTime: 5, DataCapture: false},
		},
		Cycle: []plc.Step{
			// plc.Step{60, 10, 10}
			plc.Step{TargetTemp: 95, RampUpTemp: 10, HoldTime: 5, DataCapture: false},
			plc.Step{TargetTemp: 85, RampUpTemp: 10, HoldTime: 5, DataCapture: false},
			plc.Step{TargetTemp: 75, RampUpTemp: 10, HoldTime: 5, DataCapture: false},
			plc.Step{TargetTemp: 65, RampUpTemp: 10, HoldTime: 5, DataCapture: true},
		},
		CycleCount: 3,
	}

	file := db.GetExcelFile(tec.LogsPath, "output")
	// Start line
	headings := []interface{}{"Description", "Time Taken", "Expected Time", "Initial Temp", "Final Temp", "Ramp"}
	db.AddRowToExcel(file, db.TECSheet, headings)

	row := []interface{}{"Holding Stage About to start"}
	db.AddRowToExcel(file, db.TECSheet, row)

	// Go back to Room Temp at the end
	defer t.ReachRoomTemp()

	logger.Infoln("Room Temp 27 Reached ")
	// Run Holding Stage
	logger.Infoln("Holding Stage Started")
	t.RunStage(p.Holding, plcDeps, file, 0)

	// Run Cycle Stage
	row = []interface{}{"Cycle Stage About to start"}
	db.AddRowToExcel(file, db.TECSheet, row)

	for i := uint16(1); i <= p.CycleCount; i++ {
		logger.Infoln("Started Cycle->", i)
		t.RunStage(p.Cycle, plcDeps, file, i)
		logger.Infoln("Holding Completed ->", p.Cycle[len(p.Cycle)-1].HoldTime, " for cycle number ", i)
	}

	return nil
}

func (t *Simulator) RunStage(st []plc.Step, plcDeps plc.Driver, file *excelize.File, cycleNum uint16) (err error) {
	var row []interface{}
	ts := time.Now()
	plc.CurrentCycle = cycleNum
	stagePrevTemp := prevTemp
	for i, h := range st {
		if !plc.ExperimentRunning {
			return fmt.Errorf("Experiment is not Running or aborted!")
		}
		t0 := time.Now()
		ti := tec.TECTempSet{
			TargetTemperature: float64(h.TargetTemp),
			TargetRampRate:    float64(h.RampUpTemp),
		}
		logger.Infoln("Started ->", ti)
		t.SetTempAndRamp(ti)
		row = []interface{}{fmt.Sprintf("Time taken to complete step: %v", i+1), time.Now().Sub(t0).String(), math.Abs(float64(h.TargetTemp-prevTemp)) / float64(h.RampUpTemp), prevTemp, h.TargetTemp, h.RampUpTemp}
		db.AddRowToExcel(file, db.TECSheet, row)
		logger.Infoln("Time taken to complete step: ", i+1, "\t cycle num: ", cycleNum, "\nTime Taken: ", time.Now().Sub(t0), "\nExpected Time: ", math.Abs(float64(h.TargetTemp-prevTemp))/float64(h.RampUpTemp), "\nInitial Temp:", prevTemp, "\nTarget Temp: ", h.TargetTemp, "\nRamp Rate: ", h.RampUpTemp)
		logger.Infoln("Completed ->", ti, " holding started for ", h.HoldTime)

		if h.DataCapture {
			// Cycle in Plc
			err = plcDeps.Cycle()
			if err != nil {
				logger.Errorln("Couldn't Complete PLC Cycling")
				return
			}
			logger.Infoln("PLC cycle Completed ->", h.HoldTime)
			// If this is the last step then cycleTime seconds needed for Cycle
			err = plc.HoldSleep(h.HoldTime - int32(config.GetCycleTime()))

		} else {
			err = plc.HoldSleep(h.HoldTime)

		}
		if err != nil {
			logger.Errorln("Couldn't complete hold time")
			return
		}
		plc.CurrentCycleTemperature = h.TargetTemp
		prevTemp = h.TargetTemp

	}

	if cycleNum != 0 {
		row = []interface{}{fmt.Sprintf("Time taken to complete Cycle Stage %v", cycleNum), time.Now().Sub(ts).String(), "", stagePrevTemp, prevTemp}
		db.AddRowToExcel(file, db.TECSheet, row)
	} else {
		row = []interface{}{"Time taken to complete Holding Stage", time.Now().Sub(ts).String(), "", stagePrevTemp, prevTemp}
		db.AddRowToExcel(file, db.TECSheet, row)
	}

	plc.CycleComplete = true
	return nil
}

func (t *Simulator) RunProfile(plcDeps plc.Driver, tp tec.TempProfile) (err error) {

	file := db.GetExcelFile(tec.LogsPath, "output")

	// Start line
	headings := []interface{}{"Description", "Time Taken", "Expected Time", "Initial Temp", "Final Temp", "Ramp"}
	db.AddRowToExcel(file, db.TECSheet, headings)

	for i := uint16(1); i <= uint16(tp.Cycles); i++ {
		logger.Infoln("Started Cycle->", i)
		t.RunStage(tp.Profile, plcDeps, file, i)
		logger.Infoln("Cycle Completed -> ", i)
	}

	return nil
}
