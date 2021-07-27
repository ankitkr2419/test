package plc

import (
	"fmt"

	logger "github.com/sirupsen/logrus"
)

const (
	manualSpeed = 2000
	manualRamp  = 100
)

func (d *Compact32Deck) ManualMovement(motorNum, direction uint16, mm float32) (response string, err error) {

	if d.IsRunInProgress() {
		err = fmt.Errorf("previous run already in progress... wait or abort it")
		return "", err
	}

	d.resetAborted()
	d.SetRunInProgress()
	defer d.ResetRunInProgress()

	deckAndMotor := DeckNumber{Deck: d.name, Number: motorNum}

	response, err = d.setupMotor(manualSpeed, uint16(mm*float32(Motors[deckAndMotor]["steps"])), manualRamp, direction, motorNum)
	if err != nil {
		return "", fmt.Errorf("there was some issue doing manual movement")
	}

	return
}

func (d *Compact32Deck) Pause() (response string, err error) {

	// If machine is already PAUSED OR
	if d.isMachineInPausedState() {
		err = fmt.Errorf("Machine is already in PAUSED state")
		return "", err
	}

	// run is not in Progress
	if !d.IsRunInProgress() {
		err = fmt.Errorf("Machine is already in IDLE state")
		return "", err
	}

	if !d.isTimerInProgress() {
		response, err = d.switchOffMotor()
		if err != nil {
			return "", err
		}
	}

	if d.isUVLightInProgress() {
		response, err = d.switchOffUVLight()
		if err != nil {
			return "", err
		}
	}

	if d.isHeaterInProgress() {
		response, err = d.switchOffHeater()
		if err != nil {
			return
		}
	}

	if d.isShakerInProgress() {
		response, err = d.switchOffShaker()
		if err != nil {
			return
		}
	}

	d.SetPaused()

	return "Operation PAUSED Successfully", nil
}

func (d *Compact32Deck) Resume() (response string, err error) {

	// if paused only then resume
	if !d.isMachineInPausedState() {
		err = fmt.Errorf("System is already running, or done with the run")
		return "", err
	}

	// TODO: Check if adding && !d.isCompletionBitSet() is suited below
	// completion bit shouldn't be set or it means run was already completed

	if !d.isTimerInProgress() {
		response, err = d.readExecutedPulses()
		if err != nil {
			logger.Errorln("err : ", err)
			return "", err
		}

		if temp1 := d.getWrotePulses(); temp1 == highestUint16 {
			err = fmt.Errorf("wrotePulses isn't loaded!")
		} else if temp2 := d.getExecutedPulses(); temp2 == highestUint16 {
			err = fmt.Errorf("executedPulses isn't loaded!")
		} else if temp1 <= temp2 {
			logger.Info("executedPulses is greater than wrote Pulses that means nothing to resume for current motor.")
			wrotePulses.Store(d.name, uint16(0))
			executedPulses.Store(d.name, uint16(0))
		} else {
			// calculating wrotePulses.[d.name] - executedPulses.[d.name]
			response, err = d.resumeMotorWithPulses(temp1 - temp2)
		}
		if err != nil {
			logger.Errorln("err:", err)
			return
		}
	}

	if d.isHeaterInProgress() {
		response, err = d.switchOnHeater()
		if err != nil {
			return
		}
	}
	if d.isShakerInProgress() {
		response, err = d.switchOnShaker()
		if err != nil {
			return
		}
	}

	if d.isUVLightInProgress() {
		response, err = d.switchOnUVLight()
		if err != nil {
			return "", err
		}
	}
	d.ResetPaused()

	return "Operation RESUMED Successfully.", nil
}

func (d *Compact32Deck) Abort() (response string, err error) {

	logger.Infoln("aborting the operation....")

	if d.isUVLightInProgress() {
		//  Switch off UV Light
		response, err = d.switchOffUVLight()
		if err != nil {
			logger.Errorln("From deck ", d.name, err)
			return "", err
		}
		d.setAborted()
		return "ABORT UV LIGHT SUCCESS", nil
	}

	if d.isPIDCalibrationInProgress() {
		//  Switch off Heater
		response, err = d.switchOffHeater()
		if err != nil {
			logger.Errorln("couldn't switch OFF Heater", d.name, err)
			return "", err
		}
		//  Switch off PID Calibration
		response, err = d.switchOffPIDCalibration()
		if err != nil {
			logger.Errorln("couldn't switch OFF PID Calibration", d.name, err)
			return "", err
		}
		d.setAborted()

		return "ABORT PID CALIBRATION SUCCESS", nil
	}

	homed.Store(d.name, false)

	logger.Infoln("switching motor off....")
	response, err = d.switchOffMotor()
	if err != nil {
		logger.Errorln("From deck ", d.name, err)
		return "", err
	}

	response, err = d.switchOffHeater()
	if err != nil {
		logger.Errorln("From deck ", d.name, err)
		return "", err
	}

	// Switch off shaker
	response, err = d.switchOffShaker()
	if err != nil {
		fmt.Println("From deck ", d.name, err)
		return "", err
	}

	wrotePulses.Store(d.name, uint16(0))
	d.ResetPaused()

	// If runInProgress and no timer is in progress, that means we need to read pulses
	if d.IsRunInProgress() && !d.isTimerInProgress() {
		response, err = d.readExecutedPulses()
		if err != nil {
			logger.Errorln("err : ", err)
			return "", fmt.Errorf("Operation is ABORTED but current position was lost, please home the machine")
		}
	}

	d.setAborted()
	d.ResetRunInProgress()
	d.resetTimerInProgress()

	return "ABORT SUCCESS", nil
}

func (d *Compact32Deck) resumeMotorWithPulses(pulses uint16) (response string, err error) {

	var results []byte

	if temp := d.getOnReg(); temp == highestUint16 {
		err = fmt.Errorf("on/off Register  isn't loaded!")
		return
	} else if temp != OFF {
		err = d.DeckDriver.WriteSingleCoil(MODBUS_EXTRACTION[d.name]["M"][0], OFF)
	}
	if err != nil {
		logger.Errorln("err Switching motor off: ", err)
		return "", err
	}
	logger.Infoln("Wrote Switch OFF motor for deck", d.name)
	onReg.Store(d.name, OFF)

	if temp := d.getPulseReg(); temp == highestUint16 {
		err = fmt.Errorf("pulsesReg isn't loaded!")
		return
	} else if temp != pulses {
		results, err = d.DeckDriver.WriteSingleRegister(MODBUS_EXTRACTION[d.name]["D"][202], pulses)
	}

	if err != nil {
		logger.Errorln("err writing pulses: ", err)
		return "", err
	}
	logger.Infoln("Wrote pulses for deck", d.name, ". res : ", results)
	pulseReg.Store(d.name, pulses)
	wrotePulses.Store(d.name, pulses)

	err = d.DeckDriver.WriteSingleCoil(MODBUS_EXTRACTION[d.name]["M"][0], ON)
	if err != nil {
		logger.Errorln("err : ", err)
		return "", err
	}

	logger.Infoln("Wrote Switch ON motor")
	onReg.Store(d.name, ON)

	return "RESUMED with pulses.", nil
}
