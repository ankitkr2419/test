package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Extraction struct {
	PIDTemperature  int64  `json:"pid_temperature" validate:"required,lte=75,gte=50"`
	PIDMinutes      int64  `json:"pid_minutes" validate:"required,lte=40,gte=20"`
}

func SetExtractionConfigValues(ex Extraction) (err error) {

	oldString, newString = []string{}, []string{}
	oldString = append(oldString,
		fmt.Sprintf("pid_temp: %d", int64(GetPIDTemp())),
		fmt.Sprintf("pid_time: %d", int64(GetPIDMinutes())),
	)
	newString = append(newString,
		fmt.Sprintf("pid_temp: %d", ex.PIDTemperature),
		fmt.Sprintf("pid_time: %d", ex.PIDMinutes),
	)

	err = UpdateConfig(configPath)
	if err != nil {
		return
	}

	SetPIDTemp(ex.PIDTemperature)
	SetPIDMinutes(ex.PIDMinutes)
	return
}

func GetPIDTemp() int64 {
	return int64(ReadEnvInt("pid_temp"))
}

func GetPIDMinutes() int64 {
	return int64(ReadEnvInt("pid_time"))
}

func SetPIDTemp(pT int64) {
	viper.Set("pid_temp", pT)
}

func SetPIDMinutes(pT int64) {
	viper.Set("pid_time", pT)
}


func GetExtractionConfigValues() Extraction {
	return Extraction{
		PIDTemperature: GetPIDTemp(),
		PIDMinutes:  GetPIDMinutes(),
	}
}