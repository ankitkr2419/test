package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type RTPCR struct {
	HomingTime      int64  `json:"homing_time" validate:"required,lte=30,gte=16"`
	NumHomingCycles int64  `json:"no_of_homing_cycles" validate:"required,lte=100,gte=0"`
	CycleTime       int64  `json:"cycle_time" validate:"required,lte=30,gte=2"`
}

func SetRTPCRConfigValues(rt RTPCR) (err error) {

	oldString, newString = []string{}, []string{}
	oldString = append(oldString,
		fmt.Sprintf("homing_time: %d", GetHomingTime()),
		fmt.Sprintf("num_homing_cycles: %d", GetNumHomingCycles()),
		fmt.Sprintf("cycle_time: %d", int64(GetCycleTime())),
	)
	newString = append(newString,
		fmt.Sprintf("homing_time: %d", rt.HomingTime),
		fmt.Sprintf("num_homing_cycles: %d", rt.NumHomingCycles),
		fmt.Sprintf("cycle_time: %d", rt.CycleTime),
	)

	err = UpdateConfig(configPath)
	if err != nil {
		return
	}

	SetHomingTime(rt.HomingTime)
	SetNumHomingCycles(rt.NumHomingCycles)
	SetCycleTime(rt.CycleTime)
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

func GetRoomTemp() int {
	return ReadEnvInt("room_temp")
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

func SetRoomTemp(rT int64) {
	viper.Set("room_temp", rT)
}

func SetCycleTime(cT int64) {
	viper.Set("cycle_time", cT)
}

func GetRTPCRConfigValues() RTPCR {
	return RTPCR{
		HomingTime: int64(GetHomingTime()),
		NumHomingCycles: int64(GetNumHomingCycles()),
		CycleTime: int64(GetCycleTime()),
	}
}