package plc

import (
	"context"
	"fmt"
	"mylab/cpagent/config"
	"mylab/cpagent/db"
	"mylab/cpagent/responses"

	logger "github.com/sirupsen/logrus"
)

// ALGORITHM
// 1. Start Heater
// 2. Reset Heater in defer
// 3. Start PID for deck
// 4. Reset PID in defer
// 5. Delay and check if PID Calib was stopped

func (d *Compact32Deck) PIDCalibration(ctx context.Context) (err error) {
	// TODO: Logging this PLC Operation
	defer func() {
		if err != nil {
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
		return err
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
		logger.Errorln("Switch on PID Calibration Error", err)
		return
	}
	// Reset PID in defer
	defer d.switchOffPIDCalibration()
	logger.Infoln(responses.PIDCalibrationStarted)

	// Sleep for given minutes
	_, err = d.AddDelay(db.Delay{DelayTime: config.GetPIDMinutes() * 60}, false)
	if err != nil {
		logger.WithField("err", "PID CALIBRATION").Errorln(err)
		return
	}

	logger.Infoln(responses.PIDCalibrationSuccess)
	return
}
