package config

import (
	"github.com/spf13/viper"
)

func GetHomingTime() int {
	return ReadEnvInt("homing_time")
}

func GetNumHomingCycles() int {
	return ReadEnvInt("num_homing_cycles")
}

func GetRoomTemp() float64 {
	return ReadEnvFloat("room_temp")
}

func GetCycleTime() int {
	return ReadEnvInt("cycle_time")
}

func GetPIDTemp() int64 {
	return int64(ReadEnvInt("pid_temp"))
}

func GetPIDMinutes() int64 {
	return int64(ReadEnvInt("pid_time"))
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

func SetPIDTemp(pT int64) {
	viper.Set("pid_temp", pT)
}

func SetPIDMinutes(pT int64) {
	viper.Set("pid_time", pT)
}