package tec

import (
	"mylab/cpagent/plc"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

const LogsPath = "$HOME/cpagent/utils/output"

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
