package plc

import (
	"encoding/binary"
	"fmt"
	"time"

	logger "github.com/sirupsen/logrus"
)

func (d *Compact32Deck) setupMotor(speed, pulse, ramp, direction, motorNum uint16) (response string, err error) {

	if d.isMachineInAbortedState() {
		err = fmt.Errorf("Machine in ABORTED STATE for deck: %v. Please home the machine first.", d.Name)
		return "", err
	}

	if pulse < minimumPulsesThreshold {
		logger.Infoln("Current pulse: ", pulse, " is less than minimumPulsesThreshold. Avoiding Motor Movements for motor:", motorNum, ", deck: ", d.Name)
		return "SUCCESS", nil
	}

	wrotePulses.Store(d.Name, uint16(0))
	executedPulses.Store(d.Name, uint16(0))
	deckAndNumber := DeckNumber{Deck: d.Name, Number: motorNum}

	var results []byte

	//
	//  Detach Magnet Fully if the deck is to move and magnet is in attached State
	//

	if d.getMagnetState() != detached && motorNum == K5_Deck {
		response, err = d.fullDetach()
		if err != nil {
			logger.Errorln(err)
			return "", fmt.Errorf("There was issue Detaching Magnet before moving the deck. Error: %v", err)
		}
	}

	logger.Infoln("Moving: ", motorNum, pulse/motors[deckAndNumber]["steps"], "mm in ", direction, "for deck:", d.Name)

	//
	// move the syringe module to rest position if the Motor Num is of deck
	// or Magnet Rev Fwd motor as suggested by @Sanket
	// and syringe tips are inside of deck positions.
	//

	if d.getSyringeModuleState() == InDeck && (motorNum == K5_Deck || motorNum == K7_Magnet_Rev_Fwd) {
		response, err = d.SyringeRestPosition()
		if err != nil {
			logger.Errorln(err)
			return "", fmt.Errorf("There was issue moving syringe module before moving the deck. Error: %v", err)
		}

	}

	// Switch OFF The motor

	if temp := d.getOnReg(); temp == highestUint16 {
		err = fmt.Errorf("on/off Register  isn't loaded!")
		return
	} else if temp != OFF {
		err = d.DeckDriver.WriteSingleCoil(MODBUS_EXTRACTION[d.Name]["M"][0], OFF)
	}

	if err != nil {
		logger.Errorln("error writing Switch Off : ", err, d.Name)
		return
	}
	logger.Infoln("Wrote Switch Off motor")
	onReg.Store(d.Name, OFF)

	err = d.DeckDriver.WriteSingleCoil(MODBUS_EXTRACTION[d.Name]["M"][1], OFF)
	if err != nil {
		logger.Errorln("error writing Completion Off : ", err, d.Name)
		return "", err
	}

	if temp := d.getPulseReg(); temp == highestUint16 {
		err = fmt.Errorf("pulse Register isn't loaded!")
		return
	} else if temp != pulse {
		results, err = d.DeckDriver.WriteSingleRegister(MODBUS_EXTRACTION[d.Name]["D"][202], pulse)
	}

	if err != nil {
		logger.Errorln("error writing pulse : ", err, d.Name)
		return "", err
	}
	logger.Infoln("Wrote Pulse. res : ", results)
	pulseReg.Store(d.Name, pulse)
	wrotePulses.Store(d.Name, pulse)

	if temp := d.getSpeedReg(); temp == highestUint16 {
		err = fmt.Errorf("speed Register isn't loaded!")
		return
	} else if temp != speed {
		results, err = d.DeckDriver.WriteSingleRegister(MODBUS_EXTRACTION[d.Name]["D"][200], speed)
	}

	if err != nil {
		logger.Errorln("error writing speed : ", err, d.Name)
		return "", err
	}
	logger.Infoln("Wrote Speed. res : ", results)
	speedReg.Store(d.Name, speed)

	if temp := d.getRampReg(); temp == highestUint16 {
		err = fmt.Errorf("ramp Register isn't loaded!")
		return
	} else if temp != ramp {
		results, err = d.DeckDriver.WriteSingleRegister(MODBUS_EXTRACTION[d.Name]["D"][204], ramp)
	}

	if err != nil {
		logger.Errorln("error writing RAMP : ", err, d.Name)
		return "", err
	}
	logger.Infoln("Wrote Ramp. res : ", results)
	rampReg.Store(d.Name, ramp)

	if temp := d.getDirectionReg(); temp == highestUint16 {
		err = fmt.Errorf("direction Register isn't loaded!")
		return
	} else if temp != direction {
		results, err = d.DeckDriver.WriteSingleRegister(MODBUS_EXTRACTION[d.Name]["D"][206], direction)
	}

	if err != nil {
		logger.Errorln("error writing direction : ", err, d.Name)
		return "", err
	}
	logger.Infoln("Wrote direction. res : ", results)
	directionReg.Store(d.Name, direction)

	if temp := d.getMotorNumReg(); temp == highestUint16 {
		err = fmt.Errorf("motor Number Register isn't loaded!")
		return
	} else if temp != motorNum {
		results, err = d.DeckDriver.WriteSingleRegister(MODBUS_EXTRACTION[d.Name]["D"][226], motorNum)
	}

	if err != nil {
		logger.Errorln("error writing motor num: ", err, d.Name)
		return "", err
	}
	logger.Infoln("Wrote motorNum. res : ", results)
	motorNumReg.Store(d.Name, motorNum)

	// Check if User has paused the run/operation
	for {
		if d.isMachineInPausedState() {
			logger.Infoln("Machine in PAUSED state for deck: %v", d.Name)
			time.Sleep(400 * time.Millisecond)
		} else {
			break
		}
	}

	// Switching Motor ON
	err = d.DeckDriver.WriteSingleCoil(MODBUS_EXTRACTION[d.Name]["M"][0], ON)
	if err != nil {
		logger.Errorln("error Writing On/Off : ", err, d.Name)
		return "", err
	}
	logger.Infoln("Wrote Switch On motor")
	onReg.Store(d.Name, ON)

	// Our Run is in Progress
	logger.Infoln("Movements in Progress for deck: ", d.Name)

	for {

		if temp := d.getExecutedPulses(); temp == highestUint16 {
			err = fmt.Errorf("executedPulses isn't loaded!")
			return
			// Write executed pulses to Position
		} else if d.isMachineInAbortedState() {
			logger.Infoln("position before abortion: ", positions[deckAndNumber])
			positions[deckAndNumber] += float64(temp) / float64(motors[deckAndNumber]["steps"])
			logger.Infoln("position after abortion: ", positions[deckAndNumber])
			err = fmt.Errorf("Operation was ABORTED!")
			return "", err
		}

		results, err = d.DeckDriver.ReadCoils(MODBUS_EXTRACTION[d.Name]["M"][1], uint16(1))
		if err != nil {
			logger.Errorln("error while reading completion  : ", err, d.Name)
			// Let this reading failure be intolerant
			return
		}

		if len(results) > 0 {
			if int(results[0]) == 1 {
				logger.Infoln("Completion returned ---> ", results, d.Name)
				response, err = d.switchOffMotor()
				if err != nil {
					logger.Errorln("err: from setUp--> ", err, d.Name)
					return
				}
				distanceMoved := float64(pulse) / float64(motors[DeckNumber{Deck: d.Name, Number: motorNum}]["steps"])
				switch direction {
				// Away from Sensor
				case REV:
					positions[deckAndNumber] += distanceMoved
				// Towards Sensor
				case FWD:
					if (positions[deckAndNumber] - distanceMoved) < 0 {
						positions[deckAndNumber] = 0
						logger.Errorln("Motor Just moved to negative distance for deck: ", d.Name)
					}
					positions[deckAndNumber] -= distanceMoved
				default:
					logger.Errorln("Unknown Direction was found")
					return "", fmt.Errorf("Unknown Direction was found: %v", direction)
				}
				logger.Infoln("pos", positions[deckAndNumber], d.Name)
				return "RUN Completed", nil
			}
		}

		if direction == REV {
			goto skipSensor
		}
		results, err = d.DeckDriver.ReadCoils(MODBUS_EXTRACTION[d.Name]["M"][2], uint16(1))
		if err != nil {
			logger.Errorln("error reading Sensor : ", err, d.Name)
			return "", err
		}

		logger.Infoln("Sensor returned ---> ", results, d.Name)
		if len(results) > 0 {
			if int(results[0]) == sensorCut {
				logger.Infoln("Sensor returned ---> ", results[0], d.Name)
				response, err = d.switchOffMotor()
				if err != nil {
					logger.Errorln("Sensor err : ", err, d.Name)
					return "", err
				}
				positions[deckAndNumber] = calibs[deckAndNumber]
				logger.Infoln("pos", positions[deckAndNumber], d.Name)
				return
			}
		}

	skipSensor:
		switch pulse {
		case finalSensorCutPulses:
			time.Sleep(50 * time.Millisecond)
		default:
			time.Sleep(150 * time.Millisecond)
		}
	}

	return "RUN Completed", nil
}

func (d *Compact32Deck) switchOffMotor() (response string, err error) {

	if temp := d.getOnReg(); temp == highestUint16 {
		err = fmt.Errorf("on/off Register  isn't loaded!")
		return
	} else if temp != OFF {
		err = d.DeckDriver.WriteSingleCoil(MODBUS_EXTRACTION[d.Name]["M"][0], OFF)
	}
	if err != nil {
		logger.Errorln("err Switching motor off: ", err)
		return "", err
	}
	logger.Infoln("Wrote Switch OFF motor")
	onReg.Store(d.Name, OFF)

	return "SUCCESS", nil
}

func (d *Compact32Deck) switchOffHeater() (response string, err error) {

	// Switch off Heater
	err = d.DeckDriver.WriteSingleCoil(MODBUS_EXTRACTION[d.Name]["M"][3], OFF)
	if err != nil {
		logger.Errorln("err Switching off the heater: ", err)
		return "", err
	}
	logger.Infoln("Switched off the heater--> for deck ", d.Name)

	return "SUCCESS", nil
}

func (d *Compact32Deck) switchOnShaker() (response string, err error) {

	// Switch on Shaker
	err = d.DeckDriver.WriteSingleCoil(MODBUS_EXTRACTION[d.Name]["M"][5], ON)
	if err != nil {
		fmt.Println("err starting shaker: ", err)
		return "", err
	}
	logger.Infoln("Switched on the shaker--> for deck ", d.Name)

	return "SUCCESS", nil
}

func (d *Compact32Deck) switchOffShaker() (response string, err error) {

	// Switch off shaker
	err = d.DeckDriver.WriteSingleCoil(MODBUS_EXTRACTION[d.Name]["M"][5], OFF)
	if err != nil {
		fmt.Println("err Switching off the shaker: ", err)
		return "", err
	}
	fmt.Println("Switched off the shaker--> for deck ", d.Name)
	return "SUCCESS", nil

}

func (d *Compact32Deck) switchOnHeater() (response string, err error) {

	// Switch off Heater
	err = d.DeckDriver.WriteSingleCoil(MODBUS_EXTRACTION[d.Name]["M"][3], ON)
	if err != nil {
		fmt.Println("err Switching on the heater: ", err)
		return "", err
	}
	fmt.Println("Switched on the heater--> for deck ", d.Name)

	return "SUCCESS", nil
}

func (d *Compact32Deck) switchOnUVLight() (response string, err error) {

	err = d.DeckDriver.WriteSingleCoil(MODBUS_EXTRACTION[d.Name]["M"][6], ON)
	if err != nil {
		logger.Errorln("err Switching on the UV Light: ", err)
		return "", err
	}
	logger.Infoln("Switched on the UV Light--> for deck ", d.Name)

	return "SUCCESS", nil
}

func (d *Compact32Deck) switchOffUVLight() (response string, err error) {

	err = d.DeckDriver.WriteSingleCoil(MODBUS_EXTRACTION[d.Name]["M"][6], OFF)
	if err != nil {
		logger.Errorln("err Switching off the UV Light: ", err)
		return "", err
	}
	logger.Infoln("Switched off the UV Light--> for deck ", d.Name)

	return "SUCCESS", nil
}

func (d *Compact32Deck) readExecutedPulses() (response string, err error) {

	results, err := d.DeckDriver.ReadHoldingRegisters(MODBUS_EXTRACTION[d.Name]["D"][212], uint16(1))
	if err != nil {
		logger.Errorln("err : ", err, d.Name)
		return "", err
	}

	fmt.Printf("Read D212AddressBytesUint16. res : %+v \n", results)
	if len(results) > 0 {
		executedPulses.Store(d.Name, binary.BigEndian.Uint16(results))
	} else {
		err = fmt.Errorf("couldn't read D212")
		return "", err
	}
	logger.Infoln("Read D212 Pulses -> ", binary.BigEndian.Uint16(results))

	return "D212 Reading SUCESS", nil

}
