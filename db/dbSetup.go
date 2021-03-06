package db

import (
	"context"

	logger "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// DBSetup initializes dye & targets in DB
func SetupDyeAndTargets(s Storer) (err error) {
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
		dye.Tolerance = d.Tolerance
		Dyes = append(Dyes, dye)
	}
	return
}

// DBSetup initializes motors in DB
func setupMotor(s Storer) (err error) {
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
func setupConsumable(s Storer) (err error) {
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

// DBSetup initializes tips and tubes in DB
func setupTipsTubes(s Storer) (err error) {
	var config TipsTubesConfig
	err = viper.Unmarshal(&config)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Unable to unmarshal config")
		return
	}

	// create tipstubes list
	tipsTubesList := makeTipsTubesList(config)

	// add distances to DB
	err = s.InsertTipsTubes(context.Background(), tipsTubesList)
	if err != nil {
		return
	}

	logger.Info("Tips and Tubes Added in Database")
	return
}

// DBSetup initializes cartridge in DB
func setupCartridges(s Storer) (err error) {
	var cartridgesConfig CartridgesConfig
	var wellsConfig CartridgeWellsConfig
	err = viper.Unmarshal(&cartridgesConfig)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Unable to unmarshal cartridgesConfig")
		return
	}
	err = viper.Unmarshal(&wellsConfig)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Unable to unmarshal cartridgeWellsConfig")
		return
	}

	// create cartridge list
	cartridgesList := makeCartridgeList(cartridgesConfig)
	wellsList := makeCartridgeWellsList(wellsConfig)

	ctx := context.WithValue(context.Background(), ContextKeyUsername, "main")
	// add distances to DB
	err = s.InsertCartridge(ctx, cartridgesList, wellsList)
	if err != nil {
		return
	}

	logger.Info("Cartridges Added in Database")
	return
}

func makeMotorList(configMotors MotorConfig) (Motors []Motor) {
	motor := Motor{}
	for _, d := range configMotors.Motor {
		motor.ID = d.ID
		motor.Deck = d.Deck
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

func makeTipsTubesList(configTipsTubes TipsTubesConfig) (allTipsTubes []TipsTubes) {
	tipstubes := TipsTubes{}
	for _, t := range configTipsTubes.TipsTubes {
		tipstubes.ID = t.ID
		tipstubes.Type = t.Type
		tipstubes.AllowedPositions = t.AllowedPositions
		tipstubes.Name = t.Name
		tipstubes.Volume = t.Volume
		tipstubes.Height = t.Height
		tipstubes.TtBase = t.TtBase
		allTipsTubes = append(allTipsTubes, tipstubes)
	}
	return
}

func makeCartridgeList(configCartridge CartridgesConfig) (Cartridges []Cartridge) {
	cartridge := Cartridge{}
	for _, c := range configCartridge.Cartridges {
		cartridge.ID = c.ID
		cartridge.Type = c.Type
		cartridge.Description = c.Description
		Cartridges = append(Cartridges, cartridge)
	}
	return
}

func makeCartridgeWellsList(configCartridge CartridgeWellsConfig) (cartridgeWells []CartridgeWells) {
	cartridgeWell := CartridgeWells{}
	for _, c := range configCartridge.CartridgeWells {
		cartridgeWell.ID = c.ID
		cartridgeWell.WellNum = c.WellNum
		cartridgeWell.Distance = c.Distance
		cartridgeWell.Height = c.Height
		cartridgeWell.Volume = c.Volume
		cartridgeWells = append(cartridgeWells, cartridgeWell)
	}
	return
}

func LoadAllDBSetups(s Storer) (err error) {

	// setup Db with dyes & targets
	err = SetupDyeAndTargets(s)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Setup Dyes & Targets failed")
		return
	}

	// setup Db with motors
	err = setupMotor(s)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Setup Motors failed")
		return
	}

	// setup Db with consumable distance
	err = setupConsumable(s)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Setup Cosumable Distance failed")
		return
	}

	// setup Db with tipstube
	err = setupTipsTubes(s)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Setup TipsTubes failed")
		return
	}

	// setup Db with cartridge
	err = setupCartridges(s)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Setup Cartridge failed")
		return
	}

	return nil
}
