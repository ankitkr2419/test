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

	addDefaultUser(s)

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

// addDefaultUser to DB

func addDefaultUser(s Storer) {

	u := User{
		Username: "admin",
		Password: "admin",
	}

	err := s.InsertUser(context.Background(), u)
	if err != nil {
		return
	}

	logger.Info("Default user added")

}
