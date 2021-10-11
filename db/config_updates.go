package db

import (
	"io/ioutil"
	"os"

	conf "mylab/cpagent/config"

	"github.com/lib/pq"
	"gopkg.in/yaml.v2"

	logger "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// Updates are in this package instead of configs cause of circular imports

const (
	cartridgeConfFile     = conf.ConfigStartPath + "cartridges_config.yml"
	cartridgeWellConfFile = conf.ConfigStartPath + "cartridge_wells_config.yml"
	tipTubeConfFile       = conf.ConfigStartPath + "tips_tubes_config.yml"
	motorConfFile         = conf.ConfigStartPath + "motor_config.yml"
	dyeConfFile           = conf.ConfigStartPath + "dyes.yml"
)

var TipTubeConfFile = "./conf/tips_tubes_config.yml"

// Config is used to get data from config file
type Config struct {
	Dyes []struct {
		Name      string
		Position  int
		Targets   []string `yaml:"targets,flow"`
		Tolerance float64
	} `yaml:"dyes"`
}

type MotorConfig struct {
	Motor []struct {
		ID     int
		Deck   string
		Number int
		Name   string
		Ramp   int
		Steps  int
		Slow   int
		Fast   int
	}
}

type ConsumableConfig struct {
	ConsumableDistance []struct {
		ID          int
		Name        string
		Distance    float64
		Description string
	}
}

type TipsTubesConfig struct {
	TipsTubes []struct {
		ID               int64
		Name             string
		Type             string
		AllowedPositions pq.Int64Array `yaml:"allowedPositions,flow"`
		Volume           float64
		Height           float64
		TtBase           float64
	}
}

type CartridgesConfig struct {
	Cartridges []struct {
		ID          int64         `yaml:"id"`
		Type        CartridgeType `yaml:"type"`
		Description string        `yaml:"description"`
	}
}

type CartridgeWellsConfig struct {
	CartridgeWells []struct {
		ID       int64   `yaml:"id"`
		WellNum  int64   `yaml:"wellNum"`
		Distance float64 `yaml:"distance"`
		Height   float64 `yaml:"height"`
		Volume   float64 `yaml:"volume"`
	}
}

func SetCartridgeValues(c CartridgeWell) (err error) {

	var cartridgesConfig CartridgesConfig
	var wellsConfig CartridgeWellsConfig
	err = viper.Unmarshal(&cartridgesConfig)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Unable to unmarshal cartridgesConfig")
		return
	}
	err = viper.Unmarshal(&wellsConfig)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Unable to unmarshal cartridgesConfig")
		return
	}
	carConf := CartridgesConfig{}
	carConf.Cartridges = make([]struct {
		ID          int64         `yaml:"id"`
		Type        CartridgeType `yaml:"type"`
		Description string        `yaml:"description"`
	}, len(c.Cartridge))

	carWellConf := CartridgeWellsConfig{}
	carWellConf.CartridgeWells = make([]struct {
		ID       int64   `yaml:"id"`
		WellNum  int64   `yaml:"wellNum"`
		Distance float64 `yaml:"distance"`
		Height   float64 `yaml:"height"`
		Volume   float64 `yaml:"volume"`
	}, len(c.CartridgeWells))

	for i, v := range c.Cartridge {
		carConf.Cartridges[i].ID = v.ID
		carConf.Cartridges[i].Type = v.Type
		carConf.Cartridges[i].Description = v.Description
		cartridgesConfig.Cartridges = append(cartridgesConfig.Cartridges, carConf.Cartridges[i])
	}

	for i, v := range c.CartridgeWells {
		carWellConf.CartridgeWells[i].ID = v.ID
		carWellConf.CartridgeWells[i].WellNum = v.WellNum
		carWellConf.CartridgeWells[i].Distance = v.Distance
		carWellConf.CartridgeWells[i].Height = v.Height
		carWellConf.CartridgeWells[i].Volume = v.Volume

		wellsConfig.CartridgeWells = append(wellsConfig.CartridgeWells, carWellConf.CartridgeWells[i])
	}

	res, err := yaml.Marshal(cartridgesConfig)
	if err != nil {
		logger.Errorln("error in marshalling", err)
		return
	}

	err = ioutil.WriteFile(os.ExpandEnv(cartridgeConfFile), res, 0666)
	if err != nil {
		logger.Errorln("error in writing to file", err)
		return
	}

	res, err = yaml.Marshal(wellsConfig)
	if err != nil {
		logger.Errorln("error in marshalling", err)
		return
	}

	err = ioutil.WriteFile(os.ExpandEnv(cartridgeWellConfFile), res, 0666)
	if err != nil {
		logger.Errorln("error in writing to file", err)
		return
	}

	return

}

func SetTipsTubesValues(tt []TipsTubes) (err error) {

	var config TipsTubesConfig
	err = viper.Unmarshal(&config)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Unable to unmarshal config")
		return
	}

	tipTubeConf := TipsTubesConfig{}
	tipTubeConf.TipsTubes = make([]struct {
		ID               int64
		Name             string
		Type             string
		AllowedPositions pq.Int64Array `yaml:"allowedPositions,flow"`
		Volume           float64
		Height           float64
		TtBase           float64
	}, 1)
	for i, v := range tt {
		tipTubeConf.TipsTubes[i].ID = v.ID
		tipTubeConf.TipsTubes[i].Name = v.Name
		tipTubeConf.TipsTubes[i].Type = v.Type
		tipTubeConf.TipsTubes[i].AllowedPositions = v.AllowedPositions
		tipTubeConf.TipsTubes[i].Volume = v.Volume
		tipTubeConf.TipsTubes[i].Height = v.Height
		tipTubeConf.TipsTubes[i].TtBase = v.TtBase

		config.TipsTubes = append(config.TipsTubes, tipTubeConf.TipsTubes[i])
	}
	res, err := yaml.Marshal(config)
	if err != nil {
		logger.Errorln("error in marshalling", err)
		return
	}

	err = ioutil.WriteFile(os.ExpandEnv(tipTubeConfFile), res, 0666)
	if err != nil {
		logger.Errorln("error in writing to file", err)
		return
	}
	return

}

func SetMotorsValues(m []Motor) (err error) {
	var config MotorConfig
	err = viper.Unmarshal(&config)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Unable to unmarshal config")
		return
	}

	motorConf := MotorConfig{}
	motorConf.Motor = make([]struct {
		ID     int
		Deck   string
		Number int
		Name   string
		Ramp   int
		Steps  int
		Slow   int
		Fast   int
	}, 1)
	for i, v := range m {
		motorConf.Motor[i].ID = v.ID
		motorConf.Motor[i].Deck = v.Deck
		motorConf.Motor[i].Number = v.Number
		motorConf.Motor[i].Name = v.Name
		motorConf.Motor[i].Ramp = v.Ramp
		motorConf.Motor[i].Steps = v.Steps
		motorConf.Motor[i].Slow = v.Slow
		motorConf.Motor[i].Fast = v.Fast

		config.Motor = append(config.Motor, motorConf.Motor[i])
	}
	res, err := yaml.Marshal(config)
	if err != nil {
		logger.Errorln("error in marshalling", err)
		return
	}

	err = ioutil.WriteFile(os.ExpandEnv(motorConfFile), res, 0666)
	if err != nil {
		logger.Errorln("error in writing to file", err)
		return
	}

	return
}

func UpdateMotorsValues(m []Motor) (err error) {
	var config MotorConfig
	err = viper.Unmarshal(&config)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Unable to unmarshal config")
		return
	}

	motorConf := MotorConfig{}
	motorConf.Motor = make([]struct {
		ID     int
		Deck   string
		Number int
		Name   string
		Ramp   int
		Steps  int
		Slow   int
		Fast   int
	}, 1)
	for _, confMotor := range config.Motor {
		for _, motor := range m {

			if motor.ID == confMotor.ID {
				confMotor.Deck = motor.Deck
				confMotor.Number = motor.Number
				confMotor.Name = motor.Name
				confMotor.Ramp = motor.Ramp
				confMotor.Steps = motor.Steps
				confMotor.Slow = motor.Slow
				confMotor.Fast = motor.Fast
			}

			motorConf.Motor = append(motorConf.Motor, confMotor)
		}
	}
	res, err := yaml.Marshal(config)
	if err != nil {
		logger.Errorln("error in marshalling", err)
		return
	}

	err = ioutil.WriteFile(os.ExpandEnv(motorConfFile), res, 0666)
	if err != nil {
		logger.Errorln("error in writing to file", err)
		return
	}

	return
}

func UpdateConsumableDistancesValues(m []ConsumableDistance) (err error) {
	var config ConsumableConfig
	err = viper.Unmarshal(&config)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Unable to unmarshal config")
		return
	}

	ConsumableDConf := ConsumableConfig{}
	ConsumableDConf.ConsumableDistance = make([]struct {
		ID          int
		Name        string
		Distance    float64
		Description string
	}, 0)
	for _, confConsDist := range config.ConsumableDistance {
		for _, consDist := range m {

			if consDist.ID == confConsDist.ID {
				confConsDist.ID = consDist.ID
				confConsDist.Name = consDist.Name
				logger.Debugln("At conf", consDist.Distance)
				confConsDist.Distance = consDist.Distance
				confConsDist.Description = consDist.Description

			}

			ConsumableDConf.ConsumableDistance = append(ConsumableDConf.ConsumableDistance, confConsDist)
		}
	}
	res, err := yaml.Marshal(ConsumableDConf)
	if err != nil {
		logger.Errorln("error in marshalling", err)
		return
	}

	consumableDistanceConfFile := conf.GetConsumableDistanceFilePath()

	err = ioutil.WriteFile(os.ExpandEnv(consumableDistanceConfFile), res, 0666)
	if err != nil {
		logger.Errorln("error in writing to file", err)
		return
	}

	return
}

func UpdateDyesTolerance(dyes []Dye) (err error) {

	var config Config
	err = viper.Unmarshal(&config)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Unable to unmarshal config")
		return
	}

	dyeConf := Config{}
	dyeConf.Dyes = make([]struct {
		Name      string
		Position  int
		Targets   []string `yaml:"targets,flow"`
		Tolerance float64
	}, 0)
	for _, confDyes := range config.Dyes {
		for _, v := range dyes {
			if v.Name == confDyes.Name && v.Position == confDyes.Position {
				confDyes.Tolerance = v.Tolerance
			}
		}
		dyeConf.Dyes = append(dyeConf.Dyes, confDyes)
	}
	res, err := yaml.Marshal(dyeConf)
	if err != nil {
		logger.Errorln("error in marshalling", err)
		return
	}

	err = ioutil.WriteFile(os.ExpandEnv(dyeConfFile), res, 0666)
	if err != nil {
		logger.Errorln("error in writing to file", err)
		return
	}
	return
}
