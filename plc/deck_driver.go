package plc

import (
	"encoding/binary"
	"fmt"
	"mylab/cpagent/responses"
	"time"

	logger "github.com/sirupsen/logrus"
)

// NOTE:
// switchOff and just off have different meanings
// for reset of operation use switchOff for only pausing use off

func (d *Compact32Deck) setupMotor(speed, pulse, ramp, direction, motorNum uint16) (response string, err error) {

	logger.Infoln("SetUpMotor Deck ", d.name, " Speed: ", speed, " Pulse: ", pulse, " Ramp: ", ramp, " Direction: ", direction, " MotorNum: ", motorNum)
	var results []byte
	var temp uint16

	if d.isMachineInAbortedState() {
		err = fmt.Errorf("Machine in ABORTED STATE for deck: %v. Please home the machine first.", d.name)
		return "", err
	}

	d.resetMotorOperationCompleted()
	if pulse < minimumPulsesThreshold {
		logger.Infoln("Current pulse: ", pulse, " is less than minimumPulsesThreshold. Avoiding Motor Movements for motor:", motorNum, ", deck: ", d.name)
		return "SUCCESS", nil
	}

	err = d.sleepIfPaused()
	if err != nil {
		return
	}

	wrotePulses.Store(d.name, uint16(0))
	executedPulses.Store(d.name, uint16(0))
	deckAndNumber := DeckNumber{Deck: d.name, Number: motorNum}

	//
	//  Detach Magnet Fully if the deck is to move and magnet is in attached State
	//

	if d.getMagnetState() != detached && motorNum == K5_Deck {
		response, err = d.detach()
		if err != nil {
			logger.Errorln(err)
			return "", fmt.Errorf("There was issue Detaching Magnet before moving the deck. Error: %v", err)
		}
	}

	logger.Infoln("Moving: ", motorNum, pulse/Motors[deckAndNumber]["steps"], "mm in ", direction, "for deck:", d.name)

	//
	// move the syringe module to rest position if the Motor Num is of deck
	// or Magnet RevFwd or Magnet UpDown motor as suggested by @Sanket
	// and syringe tips are inside of deck positions.
	//

	defer func() {
		if motorNum == K9_Syringe_Module_LHRH {
			d.setSyringeState()
		}
	}()

	// if tip discard is in progress that means avoid moving module up when motor is K5
	if d.getSyringeModuleState() == InDeck && // Syringe module has to be indeck
		((motorNum == K5_Deck && !d.isTipDiscardInProgress()) || //Tip Discard Special handling
			motorNum == K7_Magnet_Rev_Fwd || motorNum == K6_Magnet_Up_Down) { //Magnet Special Handling
		response, err = d.SyringeRestPosition()
		if err != nil {
			logger.Errorln(err)
			return "", fmt.Errorf("There was issue moving syringe module before moving the deck. Error: %v", err)
		}
	}

	// Switch OFF The motor

	if temp = d.getOnReg(); temp == highestUint16 {
		err = fmt.Errorf("on/off Register  isn't loaded!")
		return
	} else if temp != OFF {
		err = d.DeckDriver.WriteSingleCoil(MODBUS_EXTRACTION[d.name]["M"][0], OFF)
	}

	if err != nil {
		logger.Errorln("error writing Switch Off : ", err, d.name)
		return
	}
	logger.Infoln("Wrote Switch Off motor for deck", d.name)
	onReg.Store(d.name, OFF)

	err = d.DeckDriver.WriteSingleCoil(MODBUS_EXTRACTION[d.name]["M"][1], OFF)
	if err != nil {
		logger.Errorln("error writing Completion Off : ", err, d.name)
		return "", err
	}

	if temp = d.getPulseReg(); temp == highestUint16 {
		err = fmt.Errorf("pulse Register isn't loaded!")
		return
	} else if temp != pulse {
		results, err = d.DeckDriver.WriteSingleRegister(MODBUS_EXTRACTION[d.name]["D"][202], pulse)
	}

	if err != nil {
		logger.Errorln("error writing pulse : ", err, d.name)
		return "", err
	}
	logger.Infoln("Wrote Pulse for deck", d.name, ". res : ", results)
	pulseReg.Store(d.name, pulse)
	wrotePulses.Store(d.name, pulse)

	if temp = d.getSpeedReg(); temp == highestUint16 {
		err = fmt.Errorf("speed Register isn't loaded!")
		return
	} else if temp != speed {
		results, err = d.DeckDriver.WriteSingleRegister(MODBUS_EXTRACTION[d.name]["D"][200], speed)
	}

	if err != nil {
		logger.Errorln("error writing speed : ", err, d.name)
		return "", err
	}
	logger.Infoln("Wrote Speed for deck", d.name, ". res : ", results)
	speedReg.Store(d.name, speed)

	if temp = d.getRampReg(); temp == highestUint16 {
		err = fmt.Errorf("ramp Register isn't loaded!")
		return
	} else if temp != ramp {
		results, err = d.DeckDriver.WriteSingleRegister(MODBUS_EXTRACTION[d.name]["D"][204], ramp)
	}

	if err != nil {
		logger.Errorln("error writing RAMP : ", err, d.name)
		return "", err
	}
	logger.Infoln("Wrote Ramp for deck", d.name, ". res : ", results)
	rampReg.Store(d.name, ramp)

	if temp = d.getDirectionReg(); temp == highestUint16 {
		err = fmt.Errorf("direction Register isn't loaded!")
		return
	} else if temp != direction {
		results, err = d.DeckDriver.WriteSingleRegister(MODBUS_EXTRACTION[d.name]["D"][206], direction)
	}

	if err != nil {
		logger.Errorln("error writing direction : ", err, d.name)
		return "", err
	}
	logger.Infoln("Wrote direction for deck ", d.name, ". res : ", results)
	directionReg.Store(d.name, direction)

	if temp = d.getMotorNumReg(); temp == highestUint16 {
		err = fmt.Errorf("motor Number Register isn't loaded!")
		return
	} else if temp != motorNum {
		results, err = d.DeckDriver.WriteSingleRegister(MODBUS_EXTRACTION[d.name]["D"][226], motorNum)
	}

	if err != nil {
		logger.Errorln("error writing motor num: ", err, d.name)
		return "", err
	}
	logger.Infoln("Wrote motorNum", d.name, ". res : ", results)
	motorNumReg.Store(d.name, motorNum)

	results, err = d.DeckDriver.ReadCoils(MODBUS_EXTRACTION[d.name]["M"][2], uint16(1))
	if err != nil {
		logger.Errorln("error reading Sensor : ", err, d.name)
		return "", err
	}

	logger.Infoln("Sensor Returned ---> ", results[0], d.name)

	// If Sensor has already cut and motor moving TowardsSensor
	if len(results) > 0 && int(results[0]) == SensorCut && direction == TowardsSensor {
		logger.Infoln("Sensor has Cut ---> ", results[0], d.name)
		return "Sensor has cut", nil
	}

	err = d.sleepIfPaused()
	if err != nil {
		logger.Errorln(err)
		return
	}

	// Switching Motor ON
	err = d.DeckDriver.WriteSingleCoil(MODBUS_EXTRACTION[d.name]["M"][0], ON)
	if err != nil {
		logger.Errorln("error Writing On/Off : ", err, d.name)
		return "", err
	}
	logger.Infoln("Wrote Switch On motor for deck", d.name)
	onReg.Store(d.name, ON)

	// Our Run is in Progress
	logger.Infoln("Movements in Progress for deck: ", d.name)

	for {

		if temp = d.getExecutedPulses(); temp == highestUint16 {
			err = fmt.Errorf("executedPulses isn't loaded!")
			logger.Errorln(err)
			return
			// Write executed pulses to Position
		} else if d.isMachineInAbortedState() {
			logger.Infoln("position before abortion: ", Positions[deckAndNumber])
			Positions[deckAndNumber] += float64(temp) / float64(Motors[deckAndNumber]["steps"])
			logger.Infoln("position after abortion: ", Positions[deckAndNumber])
			return "", responses.AbortedError
		}

		if d.isMotorOperationCompleted() {
			logger.Warnln("Nothing to resume")
			return "Nothing to resume", nil
		}

		results, err = d.DeckDriver.ReadCoils(MODBUS_EXTRACTION[d.name]["M"][1], uint16(1))
		if err != nil {
			logger.Errorln("error while reading completion  : ", err, d.name)
			// Let this reading failure be intolerant
			return
		}

		if len(results) > 0 && int(results[0]) == 1 {
			logger.Infoln("Completion for deck", d.name, "returned ---> ", results)
			response, err = d.switchOffMotor()
			if err != nil {
				logger.Errorln("err: from setUp--> ", err, d.name)
				return
			}
			distanceMoved := float64(pulse) / float64(Motors[DeckNumber{Deck: d.name, Number: motorNum}]["steps"])
			switch direction {
			// Away from Sensor
			case REV:
				Positions[deckAndNumber] += distanceMoved
			// Towards Sensor
			case FWD:
				if (Positions[deckAndNumber] - distanceMoved) < 0 {
					logger.Errorln("Motor Just moved to negative distance", Positions[deckAndNumber]-distanceMoved, "for deck: ", d.name)
					Positions[deckAndNumber] = 0
					break
				}
				Positions[deckAndNumber] -= distanceMoved
			default:
				logger.Errorln("Unknown Direction was found")
				return "", fmt.Errorf("Unknown Direction was found: %v", direction)
			}
			logger.Infoln("pos", Positions[deckAndNumber], d.name)
			return "RUN Completed", nil
		}

		if direction == REV {
			goto skipSensor
		}
		results, err = d.DeckDriver.ReadCoils(MODBUS_EXTRACTION[d.name]["M"][2], uint16(1))
		if err != nil {
			logger.Errorln("error reading Sensor : ", err, d.name)
			return "", err
		}

		logger.Infoln("Sensor returned for deck ", d.name, "---> ", results)
		if len(results) > 0 && int(results[0]) == SensorCut {
			logger.Infoln("Sensor returned ---> ", results[0], d.name)
			response, err = d.switchOffMotor()
			if err != nil {
				logger.Errorln("Sensor err : ", err, d.name)
				return "", err
			}
			Positions[deckAndNumber] = Calibs[deckAndNumber]
			logger.Infoln("pos", Positions[deckAndNumber], d.name)
			return "RUN Completed as Sensor cut", nil
		}

	skipSensor:
		switch pulse {
		case finalSensorCutPulses:
			time.Sleep(50 * time.Millisecond)
		default:
			time.Sleep(150 * time.Millisecond)
		}
	}

}

func (d *Compact32Deck) sleepIfPaused() (err error) {
	// Check if User has paused the run/operation
	for {
		if d.isMachineInAbortedState() {
			return responses.AbortedError
		} else if d.isMachineInPausedState() {
			logger.Infoln("Machine in PAUSED state for deck:", d.name)
			time.Sleep(400 * time.Millisecond)
		} else {
			return
		}
	}
}

func (d *Compact32Deck) switchOffMotor() (response string, err error) {

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

	return "SUCCESS", nil
}

func (d *Compact32Deck) offHeater() (response string, err error) {

	// Switch off Heater
	err = d.DeckDriver.WriteSingleCoil(MODBUS_EXTRACTION[d.name]["M"][3], OFF)
	if err != nil {
		logger.Errorln("err pausing the heater for deck: ", d.name, err)
		return "", err
	}
	logger.Infoln("Switched off the heater--> for deck ", d.name)

	return "SUCCESS", nil
}

func (d *Compact32Deck) switchOffHeater() (response string, err error) {

	// Switch off Heater
	err = d.DeckDriver.WriteSingleCoil(MODBUS_EXTRACTION[d.name]["M"][3], OFF)
	if err != nil {
		logger.Errorln("err Switching off the heater for deck: ", d.name, err)
		return "", err
	}
	logger.Infoln("Switched off the heater--> for deck ", d.name)

	d.resetHeaterInProgress()

	return "SUCCESS", nil
}

func (d *Compact32Deck) switchOnShaker(rpmPulses uint16) (response string, err error) {

	d.switchOffShaker()

	// 4 set shaker register with rpm 1
	// NOTE: Calculation of RPM involves multiplying it with 13.3
	_, err = d.DeckDriver.WriteSingleRegister(MODBUS_EXTRACTION[d.name]["D"][218], rpmPulses)
	if err != nil {
		logger.Errorln("error in setting rpm 1 value : ", err)
		return "", err
	}

	//select shaker
	_, err = d.DeckDriver.WriteSingleRegister(MODBUS_EXTRACTION[d.name]["D"][222], shaker)
	if err != nil {
		logger.Errorln("Error failed to write temperature: ", err)
		return "", err
	}

	_, err = d.startShaker()
	if err != nil {
		return
	}

	return "SUCCESS", nil
}

func (d *Compact32Deck) offShaker() (response string, err error) {

	var results []byte
	var motorNum = K8_Shaker

	if temp := d.getMotorNumReg(); temp == highestUint16 {
		err = fmt.Errorf("motor Number Register isn't loaded!")
		return
	} else if temp != motorNum {
		results, err = d.DeckDriver.WriteSingleRegister(MODBUS_EXTRACTION[d.name]["D"][226], motorNum)
	}

	if err != nil {
		logger.Errorln("error writing motor num: ", err, d.name)
		return "", err
	}
	logger.Infoln("Wrote motorNum. res : ", results)
	motorNumReg.Store(d.name, motorNum)

	// Switch off Motor
	err = d.DeckDriver.WriteSingleCoil(MODBUS_EXTRACTION[d.name]["M"][0], OFF)
	if err != nil {
		logger.Errorln("error offing motor: ", err)
		return "", err
	}
	logger.Infoln("Switched off the shaker motor--> for deck ", d.name)

	// Switch off shaker
	err = d.DeckDriver.WriteSingleCoil(MODBUS_EXTRACTION[d.name]["M"][5], OFF)
	if err != nil {
		logger.Errorln("error Switching off the shaker: ", err)
		return "", err
	}
	logger.Infoln("Switched off the shaker--> for deck ", d.name)

	// reset completion
	err = d.DeckDriver.WriteSingleCoil(MODBUS_EXTRACTION[d.name]["M"][1], OFF)
	if err != nil {
		logger.Errorln("err resetting completion bit: ", err)
		return "", err
	}

	return "SUCCESS", nil
}

func (d *Compact32Deck) switchOffShaker() (response string, err error) {

	response, err = d.offShaker()
	if err != nil {
		return
	}

	d.resetShakerInProgress()

	return "SUCCESS", nil
}

func (d *Compact32Deck) switchOnHeater(temp uint16) (response string, err error) {

	response, err = d.switchOffHeater()
	if err != nil {
		return
	}

	err = d.sleepIfPaused()
	if err != nil {
		return
	}

	//validation for shaker
	if shaker > 3 || shaker < 1 {
		err = fmt.Errorf("%v not in valid range of 1-3", shaker)
		logger.Errorln("Error shaker number not in valid range: ", err)
		return "", err
	}

	//select shaker for heating
	result, err := d.DeckDriver.WriteSingleRegister(MODBUS_EXTRACTION[d.name]["D"][222], shaker)
	if err != nil {
		logger.Errorln("Error failed to write temperature: ", err)
		return "", err
	}

	logger.Infoln("result from shaker selection", result)

	//Step 3: Set Temperature
	//Set Temperature for heater
	result, err = d.DeckDriver.WriteSingleRegister(MODBUS_EXTRACTION[d.name]["D"][208], temp)
	if err != nil {
		logger.Errorln("Error failed to write temperature: ", err)
		return "", err
	}
	logger.Infoln("result from temperature set ", result, temp)

	_, err = d.startHeater()
	if err != nil {
		return
	}

	return "SUCCESS", nil
}

func (d *Compact32Deck) startHeater() (response string, err error) {

	// on Heater
	err = d.DeckDriver.WriteSingleCoil(MODBUS_EXTRACTION[d.name]["M"][3], ON)
	if err != nil {
		logger.Errorln("error starting on the heater: ", err)
		return "", err
	}
	logger.Infoln("Switched on the heater--> for deck ", d.name)

	d.setHeaterInProgress()

	return "SUCCESS", nil
}

func (d *Compact32Deck) startShaker() (response string, err error) {

	// Switch on Motor
	err = d.DeckDriver.WriteSingleCoil(MODBUS_EXTRACTION[d.name]["M"][0], ON)
	if err != nil {
		logger.Errorln("error starting motor: ", err)
		return "", err
	}
	logger.Infoln("Switched on the shaker motor--> for deck ", d.name)

	// Switch on Shaker
	err = d.DeckDriver.WriteSingleCoil(MODBUS_EXTRACTION[d.name]["M"][5], ON)
	if err != nil {
		logger.Errorln("error starting shaker: ", err)
		return "", err
	}
	logger.Infoln("Switched on the shaker--> for deck ", d.name)

	d.setShakerInProgress()

	return "SUCCESS", nil
}

func (d *Compact32Deck) switchOnUVLight() (response string, err error) {

	err = d.DeckDriver.WriteSingleCoil(MODBUS_EXTRACTION[d.name]["M"][6], ON)
	if err != nil {
		logger.Errorln("err Switching on the UV Light: ", err)
		return "", err
	}
	logger.Infoln("Switched on the UV Light--> for deck ", d.name)

	d.setUVLightInProgress()

	return "SUCCESS", nil
}

func (d *Compact32Deck) switchOnLight() (response string, err error) {
	err = d.DeckDriver.WriteSingleCoil(MODBUS_EXTRACTION[d.name]["M"][7], ON)
	if err != nil {
		logger.Errorln("err Switching on the Light: ", err)
		return "", err
	}
	logger.Infoln("Switched on the Light--> for deck ", d.name)

	return "SUCCESS", nil
}

func (d *Compact32Deck) switchOffLight() (response string, err error) {
	err = d.DeckDriver.WriteSingleCoil(MODBUS_EXTRACTION[d.name]["M"][7], OFF)
	if err != nil {
		logger.Errorln("err Switching off the Light: ", err)
		return "", err
	}
	logger.Infoln("Switched off the Light--> for deck ", d.name)

	return "SUCCESS", nil
}

func (d *Compact32Deck) offUVLight() (response string, err error) {

	err = d.DeckDriver.WriteSingleCoil(MODBUS_EXTRACTION[d.name]["M"][6], OFF)
	if err != nil {
		logger.Errorln("err Pausing the UV Light: ", err)
		return "", err
	}
	logger.Infoln("Switched off the UV Light--> for deck ", d.name)

	return "SUCCESS", nil
}

func (d *Compact32Deck) switchOffUVLight() (response string, err error) {

	response, err = d.offUVLight()
	if err != nil {
		return
	}

	// reset a little late to avoid bug in uv light delay
	go func() {
		time.Sleep(1 * time.Second)
		d.resetUVLightInProgress()
	}()

	return "SUCCESS", nil
}

func (d *Compact32Deck) switchOnShakerPIDCalibration() (response string, err error) {

	// PID calibration LH ON
	err = d.DeckDriver.WriteSingleCoil(MODBUS_EXTRACTION[d.name]["M"][4], ON)
	if err != nil {
		logger.Errorln("err Switching ON PID calibration LH: ", err)
		return "", err
	}
	logger.Infoln("Switched ON PID calibration LH for deck ", d.name)
	// PID calibration RH ON
	err = d.DeckDriver.WriteSingleCoil(MODBUS_EXTRACTION[d.name]["M"][8], ON)
	if err != nil {
		logger.Errorln("err Switching ON PID calibration RH: ", err)
		return "", err
	}
	logger.Infoln("Switched ON PID calibration RH for deck ", d.name)

	return "SUCCESS", nil
}

func (d *Compact32Deck) switchOffShakerPIDCalibration() (response string, err error) {

	// PID calibration LH OFF
	err = d.DeckDriver.WriteSingleCoil(MODBUS_EXTRACTION[d.name]["M"][4], OFF)
	if err != nil {
		logger.Errorln("err Switching OFF PID calibration LH: ", err)
		return "", err
	}
	logger.Infoln("Switched OFF PID calibration LH for deck ", d.name)
	// PID calibration RH OFF
	err = d.DeckDriver.WriteSingleCoil(MODBUS_EXTRACTION[d.name]["M"][8], OFF)
	if err != nil {
		logger.Errorln("err Switching OFF PID calibration RH: ", err)
		return "", err
	}
	logger.Infoln("Switched OFF PID calibration RH for deck ", d.name)

	d.resetShakerPIDCalibrationInProgress()

	return "SUCCESS", nil
}

func (d *Compact32Deck) readExecutedPulses() (response string, err error) {

	results, err := d.DeckDriver.ReadHoldingRegisters(MODBUS_EXTRACTION[d.name]["D"][212], uint16(1))
	if err != nil {
		logger.Errorln("err : ", err, d.name)
		return "", err
	}

	logger.Infof("Read D212AddressBytesUint16. res : %+v \n", results)
	if len(results) > 0 {
		executedPulses.Store(d.name, binary.BigEndian.Uint16(results))
	} else {
		err = fmt.Errorf("couldn't read D212")
		return "", err
	}
	logger.Infoln("Read D212 Pulses -> ", binary.BigEndian.Uint16(results))

	return "D212 Reading SUCESS", nil
}

func (d *Compact32Deck) SwitchOffAllCoils() (response string, err error) {
	var tempErr error
	_, tempErr = d.switchOffMotor()
	if tempErr != nil {
		err = tempErr
	}
	_, tempErr = d.switchOffShaker()
	if tempErr != nil {
		err = fmt.Errorf("%v\n%v", err, tempErr)
	}
	_, tempErr = d.switchOffHeater()
	if tempErr != nil {
		err = fmt.Errorf("%v\n%v", err, tempErr)
	}
	_, tempErr = d.switchOffUVLight()
	if tempErr != nil {
		err = fmt.Errorf("%v\n%v", err, tempErr)
	}

	// Reset M60
	response, err = d.switchCoilRegister(60, OFF, "Attach Sensor M60")
	if err != nil {
		logger.Errorln(err)
		return
	}

	// Reset M61
	response, err = d.switchCoilRegister(61, OFF, "Attach Sensor M61")
	if err != nil {
		logger.Errorln(err)
		return
	}

	// reset completion bit
	tempErr = d.DeckDriver.WriteSingleCoil(MODBUS_EXTRACTION[d.name]["M"][1], OFF)
	if tempErr != nil {
		logger.Errorln("error writing Completion Off : ", tempErr, d.name)
		err = fmt.Errorf("%v\n%v", err, tempErr)
	}

	_, tempErr = d.switchOffShakerPIDCalibration()
	if tempErr != nil {
		logger.Errorln("error switching off pid calibration bits: ", tempErr, d.name)
		err = fmt.Errorf("%v\n%v", err, tempErr)
	}
	return "SUCCESS", err
}
