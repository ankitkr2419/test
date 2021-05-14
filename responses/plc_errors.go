package responses

import (
	"fmt"
)

var (
	PreviousRunInProgressError = fmt.Errorf("previous run already in progress... wait or abort it")

	PIDCalibrationError = fmt.Errorf("error doing pid calibration")
)
