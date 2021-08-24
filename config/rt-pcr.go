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
}

func SetRTPCRConfigValues(rt RTPCR) (err error) {

	oldString, newString = []string{}, []string{}
	oldString = append(oldString,
		fmt.Sprintf("homing_time: %d", GetHomingTime()),
		fmt.Sprintf("num_homing_cycles: %d", GetNumHomingCycles()),
		fmt.Sprintf("cycle_time: %d", int64(GetCycleTime())),
		fmt.Sprintf("pid_lid_temp: %d", int64(GetLidPIDTemp())),
	)
	newString = append(newString,
		fmt.Sprintf("homing_time: %d", rt.HomingTime),
		fmt.Sprintf("num_homing_cycles: %d", rt.NumHomingCycles),
		fmt.Sprintf("cycle_time: %d", rt.CycleTime),
		fmt.Sprintf("pid_lid_temp: %d", rt.LidPIDTemp),
	)

	err = UpdateConfig(configPath)
	if err != nil {
		return
	}

	SetHomingTime(rt.HomingTime)
	SetNumHomingCycles(rt.NumHomingCycles)
	SetCycleTime(rt.CycleTime)
	SetLidPIDTemp(rt.LidPIDTemp)
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

func SetHomingTime(hT int64) {
	viper.Set("homing_time", hT)
}

func SetNumHomingCycles(hC int64) {
	viper.Set("num_homing_cycles", hC)
}

func SetCycleTime(cT int64) {
	viper.Set("cycle_time", cT)
}

func GetRTPCRConfigValues() RTPCR {
	return RTPCR{
		HomingTime:      int64(GetHomingTime()),
		NumHomingCycles: int64(GetNumHomingCycles()),
		CycleTime:       int64(GetCycleTime()),
		LidPIDTemp:      int64(GetLidPIDTemp()),
	}
}

func GetLidPIDTemp() int64 {
	return int64(ReadEnvInt("pid_lid_temp"))
}

func SetLidPIDTemp(pT int64) {
	viper.Set("pid_lid_temp", pT)
}

func GetCycleRange() (startCycle, endCycle uint16) {
	startCycle = uint16(viper.GetInt("start_cycle"))
	endCycle = uint16(viper.GetInt("end_cycle"))
	return
}
