package tec_1089

/*
int setTempAndRamp(double, double);
int initiateTEC();
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

	"errors"
	"mylab/cpagent/config"
	"mylab/cpagent/db"
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

func NewTEC1089Driver(wsMsgCh chan string, wsErrch chan error, exit chan error, test bool, plcDeps plc.Driver) tec.Driver {

	tec1089 := TEC1089{}
	tec1089.ExitCh = exit
	tec1089.WsMsgCh = wsMsgCh
	tec1089.wsErrch = wsErrch
	go tec1089.InitiateTEC()

	if test {
		tec1089.TestRun(plcDeps)
	}

	go startMonitor()

	return &tec1089 // tec Driver
}

var errorCheckStopped, tecInProgress bool
var prevTemp float32 = 27.0

func (t *TEC1089) InitiateTEC() (err error) {
	C.initiateTEC()

	go startErrorCheck()

	return t.ReachRoomTemp()
}

// TODO: Pass Error Chan here
func startMonitor() {
	go func() {
		for {
			if plc.ExperimentRunning {
				target := C.getObjectTemp()
				// Handle Failure, Try again 3 times in interval of 200ms
				i := 0
				for (target < 20) && (i < 3) {
					i++
					time.Sleep(200 * time.Millisecond)
					target = C.getObjectTemp()

					if (i == 3) && (target < 20) {
						logger.Errorln("Temperature couldn't be read even after 3 tries!")
					}
				}
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

func (t *TEC1089) SetTempAndRamp(ts tec.TECTempSet) (err error) {
	if tecInProgress {
		return fmt.Errorf("TEC is already in Progress, please wait")
	}
	tecInProgress = true
	tempVal := C.setTempAndRamp(C.double(ts.TargetTemperature), C.double(ts.TargetRampRate))
	tecInProgress = false
	// Handle Failure, Try again 3 times in interval of 200ms
	i := 0
	for (tempVal == -1) && (i < 3) {
		i++
		time.Sleep(200 * time.Millisecond)
		tempVal = C.setTempAndRamp(C.double(ts.TargetTemperature), C.double(ts.TargetRampRate))

		if (i == 3) && (tempVal == -1) {
			err = errors.New("Temperature couldn't be reached even after 3 tries!")
		}
	}

	return
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

func (t *TEC1089) TestRun(plcDeps plc.Driver) (err error) {
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

	file := db.GetExcelFile(tec.LogsPath, "output_test")

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

func (t *TEC1089) ReachRoomTemp() (err error) {
	logger.Infoln("Going Back to Room Temp ")
	ts := tec.TECTempSet{
		TargetTemperature: config.GetRoomTemp(),
		TargetRampRate:    tec.RoomTempRamp,
	}
	err = t.SetTempAndRamp(ts)
	if err != nil {
		logger.Errorln("Couldn't Reach Room Temp ")
		return
	}
	logger.Infoln("Room Temp Reached ")
	return nil
}

func (t *TEC1089) RunStage(st []plc.Step, plcDeps plc.Driver, file *excelize.File, cycleNum uint16) (err error) {
	ts := time.Now()
	plc.CurrentCycle = cycleNum
	stagePrevTemp := prevTemp
	for i, h := range st {
		if !plc.ExperimentRunning {
			return fmt.Errorf("Experiment is not Running!")
		}
		t0 := time.Now()
		ti := tec.TECTempSet{
			TargetTemperature: float64(h.TargetTemp),
			TargetRampRate:    float64(h.RampUpTemp),
		}
		logger.Infoln("Started ->", ti)
		t.SetTempAndRamp(ti)

		row := []interface{}{fmt.Sprintf("Time taken to complete step: %v", i+1), time.Now().Sub(t0).String(), math.Abs(float64(h.TargetTemp-prevTemp)) / float64(h.RampUpTemp), prevTemp, h.TargetTemp, h.RampUpTemp}
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
			// If this is the last step then cyceTime seconds needed for Cycle
			err = plc.HoldSleep(h.HoldTime - int32(config.GetCycleTime()))
		} else {
			err = plc.HoldSleep(h.HoldTime)
		}
		if err != nil {
			return
		}
		logger.Infoln("Holding Completed ->", h.HoldTime)

		prevTemp = h.TargetTemp

	}
	if cycleNum != 0 {
		row := []interface{}{fmt.Sprintf("Time taken to complete Cycle Stage %v", cycleNum), time.Now().Sub(ts).String(), "", stagePrevTemp, prevTemp}
		db.AddRowToExcel(file, db.TECSheet, row)
	} else {
		row := []interface{}{"Time taken to complete Holding Stage", time.Now().Sub(ts).String(), "", stagePrevTemp, prevTemp}
		db.AddRowToExcel(file, db.TECSheet, row)

	}

	plc.CycleComplete = true
	return nil
}

func (t *TEC1089) GetAllTEC() (err error) {
	C.getAllTEC()
	return nil
}

func (t *TEC1089) RunProfile(plcDeps plc.Driver, tp tec.TempProfile) (err error) {
	file := db.GetExcelFile(tec.LogsPath, "test")

	// Start line
	row := []interface{}{"Description", "Time Taken", "Expected Time", "Initial Temp", "Final Temp", "Ramp"}
	db.AddRowToExcel(file, db.TECSheet, row)

	go func() {
		for i := uint16(1); i <= uint16(tp.Cycles); i++ {
			logger.Infoln("Started Cycle->", i)
			t.RunStage(tp.Profile, plcDeps, file, i)
			logger.Infoln("Cycle Completed -> ", i)
		}
	}()

	return nil
}
