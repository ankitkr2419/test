package db

import (
	"context"

	logger "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
	Dyes    []Dye
	Targets []Target
}

func DBSetup(s Storer) (err error) {
	var config Config
	err = viper.Unmarshal(&config)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Unable to unmarshal config")
		return
	}

	// add dye to DB
	InsertedDyes, err := s.InsertDyes(context.Background(), config.Dyes)
	if err != nil {
		return
	}

	logger.Info("Dyes Added in Database")

	//create target list with dye Id
	newTargets := makeTargetList(InsertedDyes, config.Targets)

	//add target to DB
	err = s.InsertTargets(context.Background(), newTargets)
	if err != nil {
		return
	}

	logger.Info("Targets Added in Database")

	return
}

func makeTargetList(dyes []Dye, targets []Target) (newTargets []Target) {
	for _, t := range targets {
		for _, d := range dyes {
			if t.Position == d.Position {
				t.DyeID = d.ID
				newTargets = append(newTargets, t)
			}
		}
	}
	return
}
