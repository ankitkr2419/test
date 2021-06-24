package tec

/*
int DemoFunc(double, double);
int InitiateTEC();
int checkForErrorState();
int autoTune();
int resetDevice();
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
	"encoding/csv"
	"fmt"
	"os"

	logger "github.com/sirupsen/logrus"
	"mylab/cpagent/plc"
	"mylab/cpagent/responses"
	"time"
)

type TECTempSet struct {
	TargetTemperature float64 `json:"target_temp" validate:"gte=-20,lte=100"`
	TargetRampRate    float64 `json:"ramp_rate" validate:"gte=-20,lte=100"`
}

var errorCheckStopped, tecInProgress bool
var prevTemp float32 = 27.0

func InitiateTEC() (err error) {
	C.InitiateTEC()

	go startErrorCheck()

	return nil
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

func ConnectTEC(t TECTempSet) (err error) {
	if tecInProgress {
		return fmt.Errorf("TEC is already in Progress, please wait")
	}
	tecInProgress = true
	C.DemoFunc(C.double(t.TargetTemperature), C.double(t.TargetRampRate))
	tecInProgress = false
	return nil
}

func AutoTune() (err error) {
	C.autoTune()
	return nil
}

func ResetDevice() (err error) {
	C.resetDevice()

	if errorCheckStopped {
		startErrorCheck()
	}
	return nil
}

func Run() (err error) {
	p := plc.Stage{
		Holding: []plc.Step{
			plc.Step{65.3, 10, 5},
			plc.Step{85.3, 10, 5},
			plc.Step{95, 10, 5},
		},
		Cycle: []plc.Step{
			// plc.Step{60, 10, 10},
			plc.Step{95, 10, 5},
			plc.Step{85, 10, 5},
			plc.Step{75, 10, 5},
			plc.Step{65, 10, 5},
		},
		CycleCount: 3,
	}


	tecLogsPath := "./utils/tec"
	// logging output to file and console
	if _, err := os.Stat(tecLogsPath); os.IsNotExist(err) {
		os.MkdirAll(tecLogsPath, 0755)
		// ignore error and try creating log output file
	}

	file, err := os.Create(fmt.Sprintf("%v/output_%v.csv", tecLogsPath, time.Now().Unix()))
	if err != nil {
		logger.Errorln(responses.FileCreationError)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Start line
	err = writer.Write([]string{"Description", "Time Taken", "Initial Temp", "Final Temp"})

	err = writer.Write([]string{"Holding Stage About to start"})

	// Run Holding Stage
	logger.Infoln("Holding Stage Started")
	runStage(p.Holding, writer, 0)

	// Run Cycle Stage
	err = writer.Write([]string{"Cycle Stage About to start"})

	for i := uint16(1); i <= p.CycleCount; i++ {
		logger.Infoln("Started Cycle->", i)
		runStage(p.Cycle, writer, i)
		logger.Infoln("Holding Completed ->", p.Cycle[i-1].HoldTime, " for cycle number ", i)
	}

	// Go back to Room Temp

	logger.Infoln("Going Back to Room Temp 27 ")
	t := TECTempSet{
		TargetTemperature: 27,
		TargetRampRate:    4,
	}
	ConnectTEC(t)
	logger.Infoln("Room Temp 27 Reached ")

	return nil
}

func runStage(st []plc.Step, writer *csv.Writer, cycleNum uint16) {
	t := time.Now()
	stagePrevTemp := prevTemp
	for i, h := range st {
		t0 := time.Now()
		t := TECTempSet{
			TargetTemperature: float64(h.TargetTemp),
			TargetRampRate:    float64(h.RampUpTemp),
		}
		logger.Infoln("Started ->", t)
		ConnectTEC(t)
		writer.Write([]string{fmt.Sprintf("Time taken to complete step: %v", i+1), time.Now().Sub(t0).String(), fmt.Sprintf("%f", prevTemp), fmt.Sprintf("%f", h.TargetTemp)})

		logger.Infoln("Completed ->", t, " holding started for ", h.HoldTime)
		time.Sleep(time.Duration(h.HoldTime) * time.Second)
		logger.Infoln("Holding Completed ->", h.HoldTime)
		prevTemp = h.TargetTemp

	}
	if cycleNum != 0 {
		writer.Write([]string{fmt.Sprintf("Time taken to complete Cycle Stage %v", cycleNum), time.Now().Sub(t).String(), fmt.Sprintf("%f", stagePrevTemp), fmt.Sprintf("%f", prevTemp)})
	} else {
		writer.Write([]string{"Time taken to complete Holding Stage", time.Now().Sub(t).String(), fmt.Sprintf("%f", stagePrevTemp), fmt.Sprintf("%f", prevTemp)})
	}
}
