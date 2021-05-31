package responses

import (
	"fmt"
)

var (
	ProcessesAbsentError = fmt.Errorf("no process present in the recipe")

	InvalidOperationWebsocket = fmt.Errorf("invalid operation selected for websocket")

	RecipeUnsafeForCUDError  = fmt.Errorf("recipe is unsafe for CUDs as its run is in progress")
	ProcessUnsafeForCUDError = fmt.Errorf("process is unsafe for CUDs as its run is in progress")
	InvalidPLCRunRecipeData   = fmt.Errorf("invalid data stored for run recipe")
)
