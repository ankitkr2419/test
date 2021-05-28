package plc

import (
	"github.com/google/uuid"
	"mylab/cpagent/db"
	"mylab/cpagent/responses"

	logger "github.com/sirupsen/logrus"
)

func CheckIfRecipeOrProcessSafeForUDs(recipeID *uuid.UUID, processID *uuid.UUID) (err error) {
	if recipeID != nil {
		// for deck A or deck B
		if (deckRecipe[deckA] != db.Recipe{} && *recipeID == deckRecipe[deckA].ID) ||
			(deckRecipe[deckB] != db.Recipe{} && *recipeID == deckRecipe[deckB].ID) {
			logger.Errorln(responses.RecipeUnsafeForCRUDError)
			return responses.RecipeUnsafeForCRUDError
		}
		logger.Infoln(responses.RecipeSafeForCRUD)
	}

	if processID != nil {

		deckAcurrentProcess := getCurrentProcessNumber(deckA)
		deckBcurrentProcess := getCurrentProcessNumber(deckB)

		switch {
		// no process running on both deck
		case deckAcurrentProcess == -1 && deckBcurrentProcess == -1:
			return

		// process running on deck B only
		case deckAcurrentProcess == -1 && len(deckProcesses[deckB]) != 0 && *processID == deckProcesses[deckB][deckBcurrentProcess].ID:
			logger.Errorln(responses.ProcessUnsafeForCRUDError)
			return responses.ProcessUnsafeForCRUDError

		// process running on deck A only
		case deckAcurrentProcess == -1 && len(deckProcesses[deckA]) != 0 && *processID == deckProcesses[deckA][deckAcurrentProcess].ID:
			logger.Errorln(responses.ProcessUnsafeForCRUDError)
			return responses.ProcessUnsafeForCRUDError

			// process running on both decks
		case (len(deckProcesses[deckA]) != 0 && *processID == deckProcesses[deckA][deckAcurrentProcess].ID) ||
			(len(deckProcesses[deckB]) != 0 && *processID == deckProcesses[deckB][deckBcurrentProcess].ID):
			logger.Errorln(responses.ProcessUnsafeForCRUDError)
			return responses.ProcessUnsafeForCRUDError

		default:
			return responses.InvalidPLCRunRecipeData
		}

		logger.Infoln(responses.ProcessSafeForCRUD)
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
