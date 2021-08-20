package config

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/spf13/viper"
)

var (
	appName string
	appPort int
)

func Load(configFile string) {
	viper.SetDefault("APP_NAME", "app")
	viper.SetDefault("APP_PORT", "8002")

	viper.SetConfigName(configFile)
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./conf")
	viper.AddConfigPath("./")
	viper.AddConfigPath("./..")
	viper.AddConfigPath("./../..")
	viper.MergeInConfig()
	viper.AutomaticEnv()
}

func ReadEnvInt(key string) int {
	checkIfSet(key)
	v, err := strconv.Atoi(viper.GetString(key))
	if err != nil {
		panic(fmt.Sprintf("key %s is not a valid integer", key))
	}
	return v
}

func ReadEnvString(key string) string {
	checkIfSet(key)
	return viper.GetString(key)
}

func ReadEnvBool(key string) bool {
	checkIfSet(key)
	return viper.GetBool(key)
}

func ReadEnvFloat(key string) float64 {
	checkIfSet(key)
	return viper.GetFloat64(key)
}

func checkIfSet(key string) {
	if !viper.IsSet(key) {
		err := errors.New(fmt.Sprintf("Key %s is not set", key))
		panic(err)
	}
}

func LoadAllConfs() {
	Load("application")

	// config file to configure dye & targets
	Load("config")

	// simulator config file to configure controls & wells in simulator
	Load("simulator")

	// config file to configure motors
	Load("motor_config")

	// config file to configure consumable distance
	Load("consumable_config_v1_4")

	// config file to configure labware
	Load("labware_config")

	// config file to configure labware
	Load("tips_tubes_config")

	// config file to configure cartridge
	Load("cartridges_config")

	// config file to configure cartridge wells
	Load("cartridge_wells_config")
}
