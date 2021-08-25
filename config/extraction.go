package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Extraction struct {
	PIDTemperature    int64 `json:"pid_temperature" validate:"required,lte=75,gte=50"`
	PIDMinutes        int64 `json:"pid_minutes" validate:"required,lte=40,gte=20"`
	ShakerStepsPerRev int64 `json:"shaker_steps_per_revolution" validate:"required,lte=20000,gte=200"`
	MicroLitrePulses  int64 `json:"micro_lit_pulses" validate:"required,gte=1"`
}

func SetExtractionConfigValues(ex Extraction) (err error) {

	oldString, newString = []string{}, []string{}
	oldString = append(oldString,
		fmt.Sprintf("pid_temp: %d", int64(GetPIDTemp())),
		fmt.Sprintf("pid_time: %d", int64(GetPIDMinutes())),
		fmt.Sprintf("shaker_steps_per_revolution: %d", int64(GetShakerStepsPerRev())),
		fmt.Sprintf("micro_lit_pulses: %d", int64(GetMicroLitrePulses())),
	)
	newString = append(newString,
		fmt.Sprintf("pid_temp: %d", ex.PIDTemperature),
		fmt.Sprintf("pid_time: %d", ex.PIDMinutes),
		fmt.Sprintf("shaker_steps_per_revolution: %d", ex.ShakerStepsPerRev),
		fmt.Sprintf("micro_lit_pulses: %d", ex.MicroLitrePulses),
	)

	err = UpdateConfig(configPath)
	if err != nil {
		return
	}

	SetPIDTemp(ex.PIDTemperature)
	SetPIDMinutes(ex.PIDMinutes)
	SetShakerStepsPerRev(ex.ShakerStepsPerRev)
	SetMicroLitrePulses(ex.MicroLitrePulses)

	return
}

func GetPIDTemp() int64 {
	return int64(ReadEnvInt("pid_temp"))
}

func GetPIDMinutes() int64 {
	return int64(ReadEnvInt("pid_time"))
}

func GetShakerStepsPerRev() int64 {
	return int64(ReadEnvInt("shaker_steps_per_revolution"))
}

func GetMicroLitrePulses() int64 {
	return int64(ReadEnvInt("micro_lit_pulses"))
}

func SetPIDTemp(pT int64) {
	viper.Set("pid_temp", pT)
}

func SetPIDMinutes(pT int64) {
	viper.Set("pid_time", pT)
}

func SetShakerStepsPerRev(value int64) {
	viper.Set("shaker_steps_per_revolution", value)
}

func SetMicroLitrePulses(value int64) {
	viper.Set("micro_lit_pulses", value)
}

func GetExtractionConfigValues() Extraction {
	return Extraction{
		PIDTemperature: GetPIDTemp(),
		PIDMinutes:     GetPIDMinutes(),
		ShakerStepsPerRev:GetShakerStepsPerRev(),
		MicroLitrePulses: GetMicroLitrePulses(),
	}
}