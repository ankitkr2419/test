package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type TEC struct {
	CurrentLimitation       float64 `json:"current_limitation"`
	VoltageLimitation       float64 `json:"voltage_limitation"`
	CurrentErrorThreshold   float64 `json:"current_error_threshold"`
	VoltageErrorThreshold   float64 `json:"voltage_error_threshold"`
	PeltierMaxCurrent       float64 `json:"peltier_max_current"`
	PeltierDeltaTemperature float64 `json:"peltier_delta_temperature"`
}

func GetCurrentLimitation() float64 {
	return ReadEnvFloat("CurrentLimitation")
}

func GetVoltageLimitation() float64 {
	return ReadEnvFloat("VoltageLimitation")
}

func GetCurrentErrorThreshold() float64 {
	return ReadEnvFloat("CurrentErrorThreshold")
}

func GetVoltageErrorThreshold() float64 {
	return ReadEnvFloat("VoltageErrorThreshold")
}

func GetPeltierMaxCurrent() float64 {
	return ReadEnvFloat("PeltierMaxCurrent")
}

func GetPeltierDeltaTemperature() float64 {
	return ReadEnvFloat("PeltierDeltaTemperature")
}

func SetCurrentLimitation(value float64) {
	viper.Set("CurrentLimitation", value)
}

func SetVoltageLimitation(value float64) {
	viper.Set("VoltageLimitation", value)
}

func SetCurrentErrorThreshold(value float64) {
	viper.Set("CurrentErrorThreshold", value)
}

func SetVoltageErrorThreshold(value float64) {
	viper.Set("VoltageErrorThreshold", value)
}

func SetPeltierMaxCurrent(value float64) {
	viper.Set("PeltierMaxCurrent", value)
}

func SetPeltierDeltaTemperature(value float64) {
	viper.Set("PeltierDeltaTemperature", value)
}

func SetTECConfigValues(tec TEC) (err error) {

	oldString, newString = []string{}, []string{}
	oldString = append(oldString,
		fmt.Sprintf("CurrentLimitation : %.2f", GetCurrentLimitation()),
		fmt.Sprintf("VoltageLimitation : %.2f", GetVoltageLimitation()),
		fmt.Sprintf("CurrentErrorThreshold: %.2f", GetCurrentErrorThreshold()),
		fmt.Sprintf("VoltageErrorThreshold: %.2f", GetVoltageErrorThreshold()),
		fmt.Sprintf("PeltierMaxCurrent: %.2f", GetPeltierMaxCurrent()),
		fmt.Sprintf("PeltierDeltaTemperature: %.2f", GetPeltierDeltaTemperature()),
	)
	newString = append(newString,
		fmt.Sprintf("CurrentLimitation : %.2f", tec.CurrentLimitation),
		fmt.Sprintf("VoltageLimitation : %.2f", tec.VoltageLimitation),
		fmt.Sprintf("CurrentErrorThreshold: %.2f", tec.CurrentErrorThreshold),
		fmt.Sprintf("VoltageErrorThreshold: %.2f", tec.VoltageErrorThreshold),
		fmt.Sprintf("PeltierMaxCurrent: %.2f", tec.PeltierMaxCurrent),
		fmt.Sprintf("PeltierDeltaTemperature: %.2f", tec.PeltierDeltaTemperature),
	)

	err = UpdateConfig(configPath)
	if err != nil {
		return
	}

	SetCurrentLimitation(tec.CurrentLimitation)
	SetVoltageLimitation(tec.VoltageLimitation)
	SetCurrentErrorThreshold(tec.CurrentErrorThreshold)
	SetVoltageErrorThreshold(tec.VoltageErrorThreshold)
	SetPeltierMaxCurrent(tec.PeltierMaxCurrent)
	SetPeltierDeltaTemperature(tec.PeltierDeltaTemperature)
	return
}

func GetTECConfigValues() TEC {
	return TEC{
		CurrentLimitation: GetCurrentLimitation(),
		VoltageLimitation:  GetVoltageLimitation(),
		CurrentErrorThreshold: GetCurrentErrorThreshold(),
		VoltageErrorThreshold: GetVoltageErrorThreshold(),
		PeltierMaxCurrent: GetPeltierMaxCurrent(),
		PeltierDeltaTemperature: GetPeltierDeltaTemperature(),
	}
}