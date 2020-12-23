package db

import (
	"context"

	logger "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// Config is used to get data from config file
type Config struct {
	Dyes []struct {
		Name     string
		Position int
		Targets  []string
	}
}

type MotorConfig struct {
	Motor []struct {
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

type LabwareConfig struct {
	Labware []struct {
		ID          int
		Name        string
		Description string
	}
}

type TipsTubesConfig struct {
	TipsTubes []struct {
		LabwareID            int
		ConsumabledistanceID int
		Name                 string
		Volume               float64
		Height               float64
	}
}

type CartraidgeConfig struct {
	Cartraidge []struct {
		LabwareID   int
		Type        string
		Description string
		WellNum     int
		Distance    float64
		Height      float64
		Volume      float64
	}
}

// DBSetup initializes dye & targets in DB
func Setup(s Storer) (err error) {
	var config Config
	err = viper.Unmarshal(&config)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Unable to unmarshal config")
		return
	}

	//create dye list
	DyeList := makeDyeList(config)

	// add dye to DB
	InsertedDyes, err := s.InsertDyes(context.Background(), DyeList)
	if err != nil {
		return
	}

	logger.Info("Dyes Added in Database")

	//create target list with dye Id
	newTargets := makeTargetList(InsertedDyes, config)

	//add target to DB
	err = s.InsertTargets(context.Background(), newTargets)
	if err != nil {
		return
	}

	logger.Info("Targets Added in Database")

	return
}

func makeTargetList(dyes []Dye, config Config) (newTargets []Target) {
	for _, c := range config.Dyes {
		for _, d := range dyes {
			if c.Name == d.Name && c.Position == d.Position {
				for _, name := range c.Targets {
					t := Target{}
					t.DyeID = d.ID
					t.Name = name
					newTargets = append(newTargets, t)
				}
			}
		}
	}
	return
}

func makeDyeList(configDyes Config) (Dyes []Dye) {
	dye := Dye{}
	for _, d := range configDyes.Dyes {
		dye.Name = d.Name
		dye.Position = d.Position
		Dyes = append(Dyes, dye)
	}
	return
}

// AddDefaultUser to DB

func AddDefaultUser(s Storer, u User) error {

	err := s.InsertUser(context.Background(), u)
	if err != nil {
		return err
	}

	logger.Info("Default user added")
	return nil

}

// DBSetup initializes motors in DB
func SetupMotor(s Storer) (err error) {
	var config MotorConfig
	err = viper.Unmarshal(&config)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Unable to unmarshal config")
		return
	}

	//create motor list
	MotorList := makeMotorList(config)

	// add motor to DB
	err = s.InsertMotor(context.Background(), MotorList)
	if err != nil {
		return
	}

	logger.Info("Motors Added in Database")
	return
}

// DBSetup initializes consumable distance in DB
func SetupConsumable(s Storer) (err error) {
	var config ConsumableConfig
	err = viper.Unmarshal(&config)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Unable to unmarshal config")
		return
	}

	// create consumable list
	ConsumableDistancesList := makeConsumableDistanceList(config)

	// add distances to DB
	err = s.InsertConsumableDistance(context.Background(), ConsumableDistancesList)
	if err != nil {
		return
	}

	logger.Info("Consumable_Distance Added in Database")
	return

}

// DBSetup initializes labware in DB
func SetupLabware(s Storer) (err error) {
	var config LabwareConfig
	err = viper.Unmarshal(&config)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Unable to unmarshal config")
		return
	}

	// create labware list
	LabwareList := makeLabwareList(config)

	// add distances to DB
	err = s.InsertLabware(context.Background(), LabwareList)
	if err != nil {
		return
	}

	logger.Info("Labware Added in Database")
	return
}

// DBSetup initializes tips and tubes in DB
func SetupTipsTubes(s Storer) (err error) {
	var config TipsTubesConfig
	err = viper.Unmarshal(&config)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Unable to unmarshal config")
		return
	}

	// create tipstubes list
	TipesTubesList := makeTipesTubesList(config)

	// add distances to DB
	err = s.InsertTipsTubes(context.Background(), TipesTubesList)
	if err != nil {
		return
	}

	logger.Info("Tips and Tubes Added in Database")
	return
}

// DBSetup initializes cartraidge in DB
func SetupCartraidge(s Storer) (err error) {
	var config CartraidgeConfig
	err = viper.Unmarshal(&config)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Unable to unmarshal config")
		return
	}

	// create cartraidge list
	CartraidgeList := makeCartraidgeList(config)

	// add distances to DB
	err = s.InsertCartraidge(context.Background(), CartraidgeList)
	if err != nil {
		return
	}

	logger.Info("Cartraidge Added in Database")
	return
}

func makeMotorList(configMotors MotorConfig) (Motors []Motor) {
	motor := Motor{}
	for _, d := range configMotors.Motor {
		motor.Number = d.Number
		motor.Name = d.Name
		motor.Ramp = d.Ramp
		motor.Steps = d.Steps
		motor.Slow = d.Slow
		motor.Fast = d.Fast
		Motors = append(Motors, motor)
	}
	return
}

func makeConsumableDistanceList(configConsumable ConsumableConfig) (ConsumableDistances []ConsumableDistance) {
	consumableDistance := ConsumableDistance{}
	for _, c := range configConsumable.ConsumableDistance {
		consumableDistance.ID = c.ID
		consumableDistance.Name = c.Name
		consumableDistance.Distance = c.Distance
		consumableDistance.Description = c.Description

		ConsumableDistances = append(ConsumableDistances, consumableDistance)
	}
	return
}

func makeLabwareList(configLabware LabwareConfig) (Labwares []Labware) {
	labware := Labware{}
	for _, l := range configLabware.Labware {
		labware.ID = l.ID
		labware.Name = l.Name
		labware.Description = l.Description
		Labwares = append(Labwares, labware)
	}
	return
}

func makeTipesTubesList(configTipsTubes TipsTubesConfig) (TipsTube []TipsTubes) {
	tipstubes := TipsTubes{}
	for _, t := range configTipsTubes.TipsTubes {
		tipstubes.LabwareID = t.LabwareID
		tipstubes.ConsumabledistanceID = t.ConsumabledistanceID
		tipstubes.Name = t.Name
		tipstubes.Volume = t.Volume
		tipstubes.Height = t.Height
		TipsTube = append(TipsTube, tipstubes)
	}
	return
}

func makeCartraidgeList(configCartraidge CartraidgeConfig) (Cartraidges []Cartraidge) {
	cartraidge := Cartraidge{}
	for _, c := range configCartraidge.Cartraidge {
		cartraidge.LabwareID = c.LabwareID
		cartraidge.Type = c.Type
		cartraidge.Description = c.Description
		cartraidge.WellNum = c.WellNum
		cartraidge.Distance = c.Distance
		cartraidge.Height = c.Height
		cartraidge.Volume = c.Volume
		Cartraidges = append(Cartraidges, cartraidge)
	}
	return
}
