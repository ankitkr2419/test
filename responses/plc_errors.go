package responses

import (
	"fmt"
)

var (
	ProcessesAbsentError = fmt.Errorf("no process present in the recipe")
	
	InvalidOperationWebsocket = fmt.Errorf("invalid operation selected for websocket")
	
)