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

func AppName() string {
	if appName == "" {
		appName = ReadEnvString("APP_NAME")
	}
	return appName
}

func AppPort() int {
	if appPort == 0 {
		appPort = ReadEnvInt("APP_PORT")
	}
	return appPort
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

func GetSecretKey() string {
	key := "SECRET_KEY"
	checkIfSet(key)
	return viper.GetString(key)
}

func GetICPosition() int {
	return ReadEnvInt("ic_position")
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
