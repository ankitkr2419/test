package responses

import (
	"fmt"
)

var (
	PreviousRunInProgressError = fmt.Errorf("previous run already in progress... wait or abort it")

	PIDCalibrationError  = fmt.Errorf("error doing pid calibration")
	ShakingError         = fmt.Errorf("error doing shaking")
	HeatingError         = fmt.Errorf("error doing heating")
	ProcessesAbsentError = fmt.Errorf("no process present in the recipe")
	FetchHeaterTempError = fmt.Errorf("error fetching heater temperature")

	InvalidOperationWebsocket = fmt.Errorf("invalid operation selected for websocket")

	RecipeUnsafeForCUDError  = fmt.Errorf("recipe is unsafe for CUDs as its run is in progress")
	ProcessUnsafeForCUDError = fmt.Errorf("process is unsafe for CUDs as its run is in progress")
	InvalidPLCRunRecipeData  = fmt.Errorf("invalid data stored for run recipe")
	ExcelSheetRowError       = fmt.Errorf("error in fetching excel sheet rows")
	ExcelSheetAddRowError    = fmt.Errorf("error in adding excel sheet rows")

	InvalidCurrentStep = fmt.Errorf("invalid current step, maybe process aborted")

	LidPIDTuningError        = fmt.Errorf("LID PID Tuining already in progress")
	LidPidTuningOffError     = fmt.Errorf("LID PID Tuining was stopped")
	LidPidTuningPresentError = fmt.Errorf("LID PID Tuning wasn't started!")
)
