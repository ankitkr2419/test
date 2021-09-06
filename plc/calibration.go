package plc

import (
	"context"
	"fmt"
	"mylab/cpagent/config"
	"mylab/cpagent/responses"
	"time"

	logger "github.com/sirupsen/logrus"
)

var shaker1PIDDone, shaker2PIDDone bool

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
			return
		}
		d.WsMsgCh <- "SUCCESS_ShakerPIDTuning_ShakerPIDTuningSuccess"
	}()

	var pidTuningDone bool

	d.setShakerPIDCalibrationInProgress()
	defer d.resetShakerPIDCalibrationInProgress()

	d.WsMsgCh <- "PROGRESS_ShakerPIDTuning_ShakerPIDTuningStarted"

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
	_, err = d.switchOnShakerPIDCalibration()
	if err != nil {
		logger.Errorln("Switch on PID Shaker Calibration Error", err)
		return
	}
	// Reset PID in defer
	defer d.switchOffShakerPIDCalibration()
	logger.Infoln(responses.ShakerPIDCalibrationStarted)

	// Check until Done
	// D 504 & D 534 change from K3 to K4

	// Check if pid tuning is Done
	// 4.
	for !pidTuningDone {
		pidTuningDone, err = d.readShakerPIDCompletion()
		if err != nil {
			logger.Errorln(err)
			return
		}
		time.Sleep(10 * time.Second)
	}

	logger.Infoln(responses.ShakerPIDCalibrationSuccess)
	return
}

func (d *Compact32Deck) readShakerPIDCompletion() (pidTuningDone bool, err error) {

	var shaker1, shaker2 uint16

	if !d.isShakerPIDTuningInProgress() {
		return false, responses.ShakerPidCalibrationError
	}

	if shaker1PIDDone {
		goto skipReadingShaker1PIDStatus
	}

	shaker1, err = d.DeckDriver.ReadSingleRegister(MODBUS_EXTRACTION[d.name]["D"][504])
	if err != nil {
		logger.Errorln("Reading D 504 Error :", err)
		return
	}

	logger.Infoln("Value from PID D 504-> ", shaker1)

	if shaker1 == 4 {
		logger.Infoln("Shaker 1 PID Tuning Done for Deck : ", d.name)
		shaker1PIDDone = true
	}

skipReadingShaker1PIDStatus:

	if shaker2PIDDone {
		goto skipReadingShaker2PIDStatus
	}
	shaker2, err = d.DeckDriver.ReadSingleRegister(MODBUS_EXTRACTION[d.name]["D"][534])
	if err != nil {
		logger.Errorln("Reading D 534 Error :", err)
		return
	}

	logger.Infoln("Value from PID D 534 -> ", shaker2)

	if shaker2 == 4 {
		logger.Infoln("Shaker 2 PID Tuning Done for Deck : ", d.name)
		shaker2PIDDone = true
	}

skipReadingShaker2PIDStatus:

	return shaker1PIDDone && shaker2PIDDone, nil
}
