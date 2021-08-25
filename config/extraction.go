package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Extraction struct {
	PIDTemperature   int64 `json:"pid_temperature" validate:"required,lte=75,gte=50"`
	PIDMinutes       int64 `json:"pid_minutes" validate:"required,lte=40,gte=20"`
	ShakerRPM        int64 `json:"shaker_steps_per_revolution" validate:"lte=20000,gte=200"`
	MicroLitrePulses int64 `json:"micro_lit_pulses" validate:"gte=1"`
}

func SetExtractionConfigValues(ex Extraction) (err error) {

	oldString, newString = []string{}, []string{}
	oldString = append(oldString,
		fmt.Sprintf("pid_temp: %d", int64(GetPIDTemp())),
		fmt.Sprintf("pid_time: %d", int64(GetPIDMinutes())),
		fmt.Sprintf("shaker_rpm: %d", int64(GetShakerRPM())),
		fmt.Sprintf("micro_lit_pulses: %d", int64(GetMicroLitrePulses())),
	)
	newString = append(newString,
		fmt.Sprintf("pid_temp: %d", ex.PIDTemperature),
		fmt.Sprintf("pid_time: %d", ex.PIDMinutes),
		fmt.Sprintf("shaker_rpm: %d", ex.ShakerRPM),
		fmt.Sprintf("micro_lit_pulses: %d", ex.MicroLitrePulses),
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
		PIDMinutes:     GetPIDMinutes(),
		ShakerRPM:      GetShakerRPM(),
	}
}

func GetShakerRPM() (shRpm int64) {
	return int64(ReadEnvInt("shaker_rpm"))
}
func GetMicroLitrePulses() (micLitPulses int64) {
	return int64(ReadEnvInt("micro_lit_pulses"))
}
