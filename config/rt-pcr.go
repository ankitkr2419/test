package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type RTPCR struct {
	HomingTime      int64 `json:"homing_time" validate:"required,lte=30,gte=16"`
	NumHomingCycles int64 `json:"no_of_homing_cycles" validate:"required,lte=100,gte=0"`
	CycleTime       int64 `json:"cycle_time" validate:"required,lte=30,gte=2"`
	LidPIDTemp      int64 `json:"pid_lid_temp" validate:"required,lte=110,gte=50"`
	ScanSpeed       int64 `json:"scan_speed" validate:"required,lte=300,gte=1"` // unit: RPM
	ScanTime        int64 `json:"scan_time" validate:"required,lte=1000,gte=1"` // unit: ms
	StartCycle      int64 `json:"start_cycle" validate:"required,lte=100,gte=0"`
	EndCycle        int64 `json:"end_cycle" validate:"required,lte=100,gte=0"`
}

func SetRTPCRConfigValues(rt RTPCR) (err error) {

	oldString, newString = []string{}, []string{}
	oldString = append(oldString,
		fmt.Sprintf("homing_time: %d", GetHomingTime()),
		fmt.Sprintf("num_homing_cycles: %d", GetNumHomingCycles()),
		fmt.Sprintf("cycle_time: %d", int64(GetCycleTime())),
		fmt.Sprintf("pid_lid_temp: %d", int64(GetLidPIDTemp())),
		fmt.Sprintf("scan_speed: %d", int64(GetScanSpeed())),
		fmt.Sprintf("scan_time: %d", int64(GetScanTime())),
		fmt.Sprintf("start_cycle: %d", int64(GetStartCycle())),
		fmt.Sprintf("end_cycle: %d", int64(GetEndCycle())),
	)
	newString = append(newString,
		fmt.Sprintf("homing_time: %d", rt.HomingTime),
		fmt.Sprintf("num_homing_cycles: %d", rt.NumHomingCycles),
		fmt.Sprintf("cycle_time: %d", rt.CycleTime),
		fmt.Sprintf("pid_lid_temp: %d", rt.LidPIDTemp),
		fmt.Sprintf("scan_speed: %d", rt.ScanSpeed),
		fmt.Sprintf("scan_time: %d", rt.ScanTime),
		fmt.Sprintf("start_cycle: %d", rt.StartCycle),
		fmt.Sprintf("end_cycle: %d", rt.EndCycle),
	)

	err = UpdateConfig(configPath)
	if err != nil {
		return
	}

	SetHomingTime(rt.HomingTime)
	SetNumHomingCycles(rt.NumHomingCycles)
	SetCycleTime(rt.CycleTime)
	SetLidPIDTemp(rt.LidPIDTemp)
	SetScanSpeed(rt.ScanSpeed)
	SetScanTime(rt.ScanTime)
	SetStartCycle(rt.StartCycle)
	SetEndCycle(rt.EndCycle)
	return
}

func ActiveWells(key string) (activeWells []int32) {
	checkIfSet(key)
	wells := viper.GetIntSlice(key)
	for _, w := range wells {
		activeWells = append(activeWells, int32(w))
	}
	return
}

func GetColorLimits(key string) uint16 {
	checkIfSet(key)
	return uint16(viper.GetInt(key))
}

func GetICPosition() int {
	return ReadEnvInt("ic_position")
}

func GetHomingTime() int {
	return ReadEnvInt("homing_time")
}

func GetNumHomingCycles() int {
	return ReadEnvInt("num_homing_cycles")
}

func GetCycleTime() int {
	return ReadEnvInt("cycle_time")
}

func GetLidPIDTemp() int64 {
	return int64(ReadEnvInt("pid_lid_temp"))
}

func GetScanSpeed() int64 {
	return int64(ReadEnvInt("scan_speed"))
}

func GetScanTime() int64 {
	return int64(ReadEnvInt("scan_time"))
}

func GetStartCycle() int64 {
	return int64(ReadEnvInt("start_cycle"))
}

func GetEndCycle() int64 {
	return int64(ReadEnvInt("end_cycle"))
}

func SetHomingTime(hT int64) {
	viper.Set("homing_time", hT)
}

func SetNumHomingCycles(hC int64) {
	viper.Set("num_homing_cycles", hC)
}

func SetCycleTime(cT int64) {
	viper.Set("cycle_time", cT)
}

func SetLidPIDTemp(pT int64) {
	viper.Set("pid_lid_temp", pT)
}

func SetScanSpeed(value int64) {
	viper.Set("scan_speed", value)
}

func SetScanTime(value int64) {
	viper.Set("scan_time", value)
}

func SetStartCycle(value int64) {
	viper.Set("start_cycle", value)
}

func SetEndCycle(value int64) {
	viper.Set("end_cycle", value)
}

func GetCycleRange() (startCycle, endCycle uint16) {
	startCycle = uint16(viper.GetInt("start_cycle"))
	endCycle = uint16(viper.GetInt("end_cycle"))
	return
}

func GetRTPCRConfigValues() RTPCR {
	return RTPCR{
		HomingTime:      int64(GetHomingTime()),
		NumHomingCycles: int64(GetNumHomingCycles()),
		CycleTime:       int64(GetCycleTime()),
		LidPIDTemp:      int64(GetLidPIDTemp()),
		ScanSpeed:       int64(GetScanSpeed()),
		ScanTime:        int64(GetScanTime()),
		StartCycle:      int64(GetStartCycle()),
		EndCycle:        int64(GetEndCycle()),
	}
}

func GetEngineerCycleCount() int64 {
	return int64(ReadEnvInt("engineer_cycle_count"))
}
