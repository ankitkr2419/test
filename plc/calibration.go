package plc

import (
	"context"
	"fmt"
	"mylab/cpagent/config"
	"mylab/cpagent/db"
	"mylab/cpagent/responses"

	logger "github.com/sirupsen/logrus"
)

var shaker1PIDDone, shaker2PIDDone bool

// ALGORITHM
// 1. Start PID for deck
// 2. Reset PID in defer
// 3. Start Heater
// 4. Reset Heater in defer
// 5. Delay and check if PID Calib was stopped

func (d *Compact32Deck) PIDCalibration(ctx context.Context) (err error) {
	// TODO: Logging this PLC Operation
	defer func() {
		if err != nil {
			logger.Errorln(err)
			if err == responses.AbortedError {
				d.WsErrCh <- fmt.Errorf("%v_%v_%v", ErrorOperationAborted, d.name, err.Error())
				return
			}
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

	logger.Infoln(responses.PIDCalibrationHeaterStarted)

	// Start PID for deck
	_, err = d.switchOnShakerPIDCalibration()
	if err != nil {
		logger.Errorln("Switch on PID Shaker Calibration Error", err)
		return
	}
	// Reset PID in defer
	defer d.switchOffShakerPIDCalibration()

	// Set Temperature
	//Set Temperature for heater
	_, err = d.switchOnHeater(uint16(config.GetPIDTemp() * 10))
	if err != nil {
		return
	}
	// Reset Heater in defer
	defer d.switchOffHeater()

	logger.Infoln(responses.ShakerPIDCalibrationStarted)

	// Check until Done
	// D 504 & D 534 change from K3 to K4

	// Check if pid tuning is Done
	for !pidTuningDone {
		if d.isMachineInAbortedState() {
			err = responses.AbortedError
			return
		}
		pidTuningDone, err = d.readShakerPIDCompletion()
		if err != nil {
			logger.Errorln(err)
			return
		}
		_, err = d.AddDelay(db.Delay{DelayTime: 10}, false)
		if err != nil {
			logger.Errorln(err)
			return
		}
	}

	// Wait for 10 sec for PID Tuning Completion
	logger.Infoln(responses.ShakerPIDCalibrationWait)
	_, err = d.AddDelay(db.Delay{DelayTime: 120}, false)
	if err != nil {
		logger.Errorln(err)
		return
	}

	logger.Infoln(responses.ShakerPIDCalibrationSuccess)
	return
}

func (d *Compact32Deck) readShakerPIDCompletion() (pidTuningDone bool, err error) {

	var shaker1, shaker2 uint16

	if !d.isShakerPIDTuningInProgress() {
		return false, responses.AbortedError
	}

	if d.isMachineInAbortedState() {
		return false, responses.AbortedError
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
	if d.isMachineInAbortedState() {
		return false, responses.AbortedError
	}

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
