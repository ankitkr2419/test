package plc

import (
	logger "github.com/sirupsen/logrus"
	"mylab/cpagent/db"
	"context"
	"github.com/google/uuid"
"time"
	"mylab/cpagent/responses"

)

func (d *Compact32Deck) StartRecipeTimeCounter() {
	d.resetRecipeWasPaused()
	d.setRecipeStartTime()
	return
}

func (d *Compact32Deck) UpdateRecipeEstimatedTime(ctx context.Context, recipeID uuid.UUID, store db.Storer,  err *error) error{
	if *err != nil { 
		logger.Errorln("Couldn't calculate estimated recipe time as error occured during run : ", *err)
		return responses.RecipeRunError
	}

	if d.wasRecipePaused() {
		logger.Errorln("Couldn't calculate estimated recipe time as recipe was paused in between")
		return responses.RecipeWasPausedError
	}

	ti := d.getRecipeStartTime()

	// subtract time.Now(). and update recipe Time
	return store.UpdateEstimatedTimeForRecipe(ctx, recipeID, int64(time.Now().Sub(ti) / time.Second))
}