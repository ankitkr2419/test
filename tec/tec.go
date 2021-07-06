package tec

import (
	"encoding/csv"
	"mylab/cpagent/plc"
)

const LogsPath = "./utils/tec"

type TECTempSet struct {
	TargetTemperature float64 `json:"target_temp" validate:"gte=-20,lte=100"`
	TargetRampRate    float64 `json:"ramp_rate" validate:"gte=-20,lte=100"`
}

type TempProfile struct {
	Profile []plc.Step `json:"profile"`
	Cycles int64 `json:"cycles" validate:"gte=1,lte=100"`
}

type Driver interface{
	TestRun() error
	ReachRoomTemp() error
	InitiateTEC() error
	ConnectTEC(TECTempSet) error
	AutoTune() error
	ResetDevice() error
	RunStage([]plc.Step, *csv.Writer, uint16) error
	GetAllTEC() error
	RunProfile(TempProfile) error
}