package responses

import (
	"fmt"
)

var (
	ProcessesAbsentError = fmt.Errorf("no process present in the recipe")

	InvalidOperationWebsocket = fmt.Errorf("invalid operation selected for websocket")

	RecipeUnsafeForCRUDError  = fmt.Errorf("recipe is unsafe for CRUDs as its run is in Progress")
	ProcessUnsafeForCRUDError = fmt.Errorf("process is unsafe for CRUDs as its run is in Progress")
)
