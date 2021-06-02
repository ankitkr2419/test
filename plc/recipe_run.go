package plc

import (
	"github.com/google/uuid"
	"mylab/cpagent/db"
	"mylab/cpagent/responses"

	logger "github.com/sirupsen/logrus"
)

func CheckIfRecipeOrProcessSafeForCUDs(recipeID *uuid.UUID, processID *uuid.UUID) (err error) {
	if recipeID != nil {
		// for deck A or deck B
		if (deckRecipe[deckA] != db.Recipe{} && *recipeID == deckRecipe[deckA].ID) ||
			(deckRecipe[deckB] != db.Recipe{} && *recipeID == deckRecipe[deckB].ID) {
			logger.Errorln(responses.RecipeUnsafeForCUDError)
			return responses.RecipeUnsafeForCUDError
		}
		logger.Infoln(responses.RecipeSafeForCUD)
	}

	if processID != nil {
		// for deck A
		for _, pr := range deckProcesses[deckA] {
			if pr.ID == *processID {
				logger.Errorln(responses.RecipeUnsafeForCUDError)
				return responses.ProcessUnsafeForCUDError
			}
		}
		// for deck B
		for _, pr := range deckProcesses[deckB] {
			if pr.ID == *processID {
				logger.Errorln(responses.RecipeUnsafeForCUDError)
				return responses.ProcessUnsafeForCUDError
			}
		}
		logger.Infoln(responses.ProcessSafeForCUD)
	}

	return
}

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
	d.SetCurrentProcessNumber(-1)
	logger.Infoln(responses.ResetRunRecipeDataSuccess)
}
