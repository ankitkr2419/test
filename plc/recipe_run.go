package plc

import (
	"mylab/cpagent/db"
	"mylab/cpagent/responses"

	logger "github.com/sirupsen/logrus"
)

func (d *Compact32Deck) RunRecipeWebsocketData(recipe db.Recipe, processes []db.Process) (err error) {
	deckRecipe[d.name] = recipe
	deckProcesses[d.name] = processes
	if recipe.ProcessCount == 0 {
		return responses.ProcessesAbsentError
	}

	d.SetCurrentProcessNumber(0)
	go d.AddDelay(db.Delay{DelayTime: recipe.TotalTime}, true)
	logger.Infoln(responses.SetRunRecipeDataSuccess)
	return
}

func (d *Compact32Deck) resetRunRecipeData() {
	deckRecipe[d.name] = db.Recipe{}
	deckProcesses[d.name] = []db.Process{}
	d.SetCurrentProcessNumber(0)
	logger.Infoln(responses.ResetRunRecipeDataSuccess)
}
