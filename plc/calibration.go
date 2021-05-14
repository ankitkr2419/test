package plc

import (
	"encoding/json"
	"fmt"
	"math"
	"mylab/cpagent/db"
	"mylab/cpagent/responses"
	"strconv"

	"strings"
	"time"

	logger "github.com/sirupsen/logrus"
)

func (d *Compact32Deck) PIDCalibration(ctx context.Context) (err error) {
	// TODO: Logging this PLC Operation

	var response string

	if d.IsRunInProgress() {
		err = responses.PreviousRunInProgressError
		return
	}

	// Start Heater
	response, err = d.switchOnHeater()
	if err != nil {
		return
	}
	// Reset Heater in defer
	defer d.switchOffHeater()
	logger.Infoln(responses.HeaterStartedForPIDCalibration)

	// Start PID for deck
	response, err = d.switchOnPIDCalibration()
	if err != nil {
		return
	}
	// Reset PID in defer
	defer d.switchOffPIDCalibration()
	logger.Infoln(responses.PIDCalibrationStarted)

	// Sleep for 15 minutes
	response, err = d.AddDelay(db.Delay{DelayTime: 15 * 60})
	if err != nil {
		return
	}

	logger.Infoln(responses.PIDCalibrationSuccess)

	// send success ws data
	successWsData := WSData{
		Progress: 100,
		Deck:     d.name,
		Status:   "SUCCESS_PIDCALIBRATION",
		OperationDetails: OperationDetails{
			Message: fmt.Sprintf("successfully completed PID calibration for deck %v", d.name),
		},
	}
	wsData, err := json.Marshal(successWsData)
	if err != nil {
		logger.Errorf("error in marshalling web socket data %v", err.Error())
		return
	}
	d.WsMsgCh <- fmt.Sprintf("success_pidCalibration_%v", string(wsData))

}
