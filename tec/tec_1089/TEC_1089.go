package tec_1089

/*
int DemoFunc(double, double);
int InitiateTEC();
int checkForErrorState();
int autoTune();
int resetDevice();
int getAllTEC();
double getObjectTemp();
#include <stdlib.h>
#include <time.h>
#include <fcntl.h>
#include <termios.h>
#include <unistd.h>
#include <errno.h>
#include <pthread.h>
#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>
#include <string.h>
#cgo CFLAGS  : -std=gnu99 -Wall -g -O3
#cgo LDFLAGS : -pthread -lrt
*/
import "C"
import (
	"fmt"
	"math"

	"mylab/cpagent/plc"
	"mylab/cpagent/tec"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
	logger "github.com/sirupsen/logrus"
)

type TEC1089 struct {
	ExitCh  chan error
	WsMsgCh chan string
	wsErrch chan error
}

func NewTEC1089Driver(wsMsgCh chan string, wsErrch chan error, exit chan error, test bool) tec.Driver {

	tec1089 := TEC1089{}
	tec1089.ExitCh = exit
	tec1089.WsMsgCh = wsMsgCh
	tec1089.wsErrch = wsErrch
	go tec1089.InitiateTEC()

	if test {
		tec1089.TestRun()
	}

	go startMonitor()

	return &tec1089 // tec Driver
}

var errorCheckStopped, tecInProgress bool
var prevTemp float32 = 27.0

func (t *TEC1089) InitiateTEC() (err error) {
	C.InitiateTEC()

	go startErrorCheck()

	return t.ReachRoomTemp()
}

func startMonitor() {
	go func() {
		for {
			if tec.TempMonStarted {
				target := C.getObjectTemp()
				logger.Infoln("Current Temp: ", target)
				plc.CurrentCycleTemperature = float32(target)
			}
			time.Sleep(1 * time.Second)
		}
	}()
}

func startErrorCheck() {
	go func() {
		time.Sleep(5 * time.Second)
		for {
			time.Sleep(2 * time.Second)
			var errNum C.int
			errNum = C.checkForErrorState()
			if errNum != 0 {
				errorCheckStopped = true
				logger.Errorln("Error Code for TEC: ", errNum)
				return
			}
		}
	}()
}

func (t *TEC1089) ConnectTEC(ts tec.TECTempSet) (err error) {
	if tecInProgress {
		return fmt.Errorf("TEC is already in Progress, please wait")
	}
	tec.TempMonStarted = true
	tecInProgress = true
	C.DemoFunc(C.double(ts.TargetTemperature), C.double(ts.TargetRampRate))
	tecInProgress = false
	return nil
}

func (t *TEC1089) AutoTune() (err error) {
	C.autoTune()
	err = t.InitiateTEC()
	return err
}

func (t *TEC1089) ResetDevice() (err error) {
	C.resetDevice()

	if errorCheckStopped {
		startErrorCheck()
	}
	return nil
}

func (t *TEC1089) TestRun() (err error) {
	p := plc.Stage{
		Holding: []plc.Step{
			plc.Step{65.3, 10, 5, false},
			plc.Step{85.3, 10, 5, false},
			plc.Step{95, 10, 5, false},
		},
		Cycle: []plc.Step{
			// plc.Step{60, 10, 10},
			plc.Step{95, 10, 5, false},
			plc.Step{85, 10, 5, false},
			plc.Step{75, 10, 5, false},
			plc.Step{65, 10, 5, false},
		},
		CycleCount: 3,
	}

	file := plc.GetExcelFile(tec.LogsPath, "output")

	// Start line
	headings := []interface{}{"Description", "Time Taken", "Expected Time", "Initial Temp", "Final Temp", "Ramp"}
	plc.AddRowToExcel(file, plc.TECSheet, headings)

	row := []interface{}{"Holding Stage About to start"}
	plc.AddRowToExcel(file, plc.TECSheet, row)

	// Go back to Room Temp at the end
	defer t.ReachRoomTemp()

	logger.Infoln("Room Temp 27 Reached ")
	// Run Holding Stage
	logger.Infoln("Holding Stage Started")
	t.RunStage(p.Holding, file, 0)

	// Run Cycle Stage
	row = []interface{}{"Cycle Stage About to start"}
	plc.AddRowToExcel(file, plc.TECSheet, row)

	for i := uint16(1); i <= p.CycleCount; i++ {
		logger.Infoln("Started Cycle->", i)
		t.RunStage(p.Cycle, file, i)
		logger.Infoln("Holding Completed ->", p.Cycle[len(p.Cycle)-1].HoldTime, " for cycle number ", i)
	}

	return nil
}

func (t *TEC1089) ReachRoomTemp() (err error) {
	logger.Infoln("Going Back to Room Temp 27 ")
	ts := tec.TECTempSet{
		TargetTemperature: 27,
		TargetRampRate:    4,
	}
	err = t.ConnectTEC(ts)
	if err != nil {
		logger.Errorln("Couldn't Reach Room Temp 27")
		return
	}
	logger.Infoln("Room Temp 27 Reached ")
	tec.TempMonStarted = false
	return nil
}

func (t *TEC1089) RunStage(st []plc.Step, file *excelize.File, cycleNum uint16) (err error) {
	ts := time.Now()
	stagePrevTemp := prevTemp
	for i, h := range st {
		t0 := time.Now()
		ti := tec.TECTempSet{
			TargetTemperature: float64(h.TargetTemp),
			TargetRampRate:    float64(h.RampUpTemp),
		}
		logger.Infoln("Started ->", ti)
		t.ConnectTEC(ti)
		plc.DataCapture = h.DataCapture
		row := []interface{}{fmt.Sprintf("Time taken to complete step: %v", i+1), time.Now().Sub(t0).String(), math.Abs(float64(h.TargetTemp-prevTemp)) / float64(h.RampUpTemp), prevTemp, h.TargetTemp, h.RampUpTemp}
		plc.AddRowToExcel(file, plc.TECSheet, row)

		logger.Infoln("Time taken to complete step: ", i+1, "\t cycle num: ", cycleNum, "\nTime Taken: ", time.Now().Sub(t0), "\nExpected Time: ", math.Abs(float64(h.TargetTemp-prevTemp))/float64(h.RampUpTemp), "\nInitial Temp:", prevTemp, "\nTarget Temp: ", h.TargetTemp, "\nRamp Rate: ", h.RampUpTemp)
		logger.Infoln("Completed ->", ti, " holding started for ", h.HoldTime)
		if i == (len(st) - 1) {
			// If this is the last step then 16 seconds needed for Cycle
			time.Sleep(time.Duration(h.HoldTime-16) * time.Second)
		} else {
			time.Sleep(time.Duration(h.HoldTime) * time.Second)
		}
		logger.Infoln("Holding Completed ->", h.HoldTime)

		prevTemp = h.TargetTemp

	}
	if cycleNum != 0 {
		row := []interface{}{fmt.Sprintf("Time taken to complete Cycle Stage %v", cycleNum), time.Now().Sub(ts).String(), "", stagePrevTemp, prevTemp}
		plc.AddRowToExcel(file, plc.TECSheet, row)
	} else {
		row := []interface{}{"Time taken to complete Holding Stage", time.Now().Sub(ts).String(), "", stagePrevTemp, prevTemp}
		plc.AddRowToExcel(file, plc.TECSheet, row)

	}

	plc.CurrentCycle = cycleNum
	plc.HeatingCycleComplete = true
	return nil
}

func (t *TEC1089) GetAllTEC() (err error) {
	C.getAllTEC()
	return nil
}

func (t *TEC1089) RunProfile(tp tec.TempProfile) (err error) {
	file := plc.GetExcelFile(tec.LogsPath, "test")

	// Start line
	row := []interface{}{"Description", "Time Taken", "Expected Time", "Initial Temp", "Final Temp", "Ramp"}
	plc.AddRowToExcel(file, plc.TECSheet, row)

	go func() {
		for i := uint16(1); i <= uint16(tp.Cycles); i++ {
			logger.Infoln("Started Cycle->", i)
			t.RunStage(tp.Profile, file, i)
			logger.Infoln("Cycle Completed -> ", i)
		}
	}()

	return nil
}
