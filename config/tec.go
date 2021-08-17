package config

import (
	"fmt"
	"github.com/spf13/viper"
)

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

type TEC struct {
	CurrentLimitation       float64 `json:"current_limitation"`
	VoltageLimitation       float64 `json:"voltage_limitation"`
	CurrentErrorThreshold   float64 `json:"current_error_threshold"`
	VoltageErrorThreshold   float64 `json:"voltage_error_threshold"`
	PeltierMaxCurrent       float64 `json:"peltier_max_current"`
	PeltierDeltaTemperature float64 `json:"peltier_delta_temperature"`
}

func SetTECValues(tec TEC) (err error) {

	cl := GetCurrentLimitation()
	vl := GetVoltageLimitation()
	cet := GetCurrentErrorThreshold()
	vet := GetVoltageErrorThreshold()
	pmc := GetPeltierMaxCurrent()
	pdt := GetPeltierDeltaTemperature()

	oldString, newString = []string{}, []string{}
	oldString = append(oldString,

		fmt.Sprintf("CurrentLimitation : %.2f", cl),
		fmt.Sprintf("VoltageLimitation : %.2f", vl),
		fmt.Sprintf("CurrentErrorThreshold: %.2f", cet),
		fmt.Sprintf("VoltageErrorThreshold: %.2f", vet),
		fmt.Sprintf("PeltierMaxCurrent: %.2f", pmc),
		fmt.Sprintf("PeltierDeltaTemperature: %.2f", pdt),
	)
	newString = append(newString,
		fmt.Sprintf("CurrentLimitation : %.2f", tec.CurrentLimitation),
		fmt.Sprintf("VoltageLimitation : %.2f", tec.VoltageLimitation),
		fmt.Sprintf("CurrentErrorThreshold: %.2f", tec.CurrentErrorThreshold),
		fmt.Sprintf("VoltageErrorThreshold: %.2f", tec.VoltageErrorThreshold),
		fmt.Sprintf("PeltierMaxCurrent: %.2f", tec.PeltierMaxCurrent),
		fmt.Sprintf("PeltierDeltaTemperature: %.2f", tec.PeltierDeltaTemperature),
	)

	err = UpdateConfig(configPath, oldString, newString)
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
