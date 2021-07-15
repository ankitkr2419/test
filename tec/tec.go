package tec

import (
	"errors"
	"mylab/cpagent/plc"
	"time"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
	logger "github.com/sirupsen/logrus"
)

const LogsPath = "./utils/output"
const RoomTempRamp = 6

var TecTempLogFile *excelize.File

type TECTempSet struct {
	TargetTemperature float64 `json:"target_temp" validate:"gte=-20,lte=100"`
	TargetRampRate    float64 `json:"ramp_rate" validate:"gte=-20,lte=100"`
}

type TempProfile struct {
	Profile []plc.Step `json:"profile"`
	Cycles  int64      `json:"cycles" validate:"gte=1,lte=100"`
}

type Driver interface {
	TestRun(plc.Driver) error
	ReachRoomTemp() error
	InitiateTEC() error
	SetTempAndRamp(TECTempSet) error
	AutoTune() error
	ResetDevice() error
	RunStage([]plc.Step, plc.Driver, *excelize.File, uint16) error
	GetAllTEC() error
	RunProfile(plc.Driver, TempProfile) error
}

func HoldSleep(sleepTime int32) (err error) {

	var elaspedTime int32
	for {
		logger.Infoln("plc.ExperimentRunning && elaspedTime < sleepTime ", plc.ExperimentRunning, elaspedTime, sleepTime)
		if plc.ExperimentRunning && elaspedTime < sleepTime {
			time.Sleep(time.Second * 1)
			logger.Infoln("sleeping in holdsleep")
		} else {
			if !plc.ExperimentRunning {

				return errors.New("experiment has stoped running")
			}
			return nil
		}
		elaspedTime = elaspedTime + 1
	}
}
