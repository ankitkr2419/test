package plc

import (
	"encoding/json"
	"fmt"
	"mylab/cpagent/db"
	"mylab/cpagent/config"
	"mylab/cpagent/responses"
	"context"

	logger "github.com/sirupsen/logrus"
)

// ALGORITHM
// 1. Start Heater
// 2. Reset Heater in defer
// 3. Start PID for deck
// 4. Reset PID in defer
// 5. Sleep for 15 minutes

func (d *Compact32Deck) PIDCalibration(ctx context.Context) (err error) {
	// TODO: Logging this PLC Operation
	defer func(){
		if err != nil{
			logger.Errorln(err)
			d.WsErrCh <- fmt.Errorf("%v_%v_%v", ErrorExtractionMonitor, d.name, err.Error())
		}
	}()

	d.setPIDCalibrationInProgress()
	defer d.resetPIDCalibrationInProgress()


	// Stop Heater
	_, err = d.switchOffHeater()
	if err != nil {
		return
	}

	// Set Temperature
	//Set Temperature for heater
	result, err := d.DeckDriver.WriteSingleRegister(MODBUS_EXTRACTION[d.name]["D"][208], uint16(config.GetPIDTemp()*10))
	if err != nil {
		logger.Errorln("Error failed to write temperature: ", err)
		return  err
	}
	logger.Infoln("result from temperature set ", result, config.GetPIDTemp())

	// Start Heater
	_, err = d.switchOnHeater()
	if err != nil {
		return
	}
	// Reset Heater in defer
	defer d.switchOffHeater()
	logger.Infoln(responses.PIDCalibrationHeaterStarted)

	// Start PID for deck
	_, err = d.switchOnPIDCalibration()
	if err != nil {
		return
	}
	// Reset PID in defer
	defer d.switchOffPIDCalibration()
	logger.Infoln(responses.PIDCalibrationStarted)

	// Sleep for given minutes
	_, err = d.AddDelay(db.Delay{DelayTime: config.GetPIDMinutes() * 60}, false)
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

	return
}
