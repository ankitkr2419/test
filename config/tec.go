package config

import (
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