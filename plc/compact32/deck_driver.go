package compact32

import (
	"encoding/binary"
	"fmt"
	"time"
)

func (d *Compact32Deck) setupMotor(speed, pulse, ramp, direction, motorNum uint16) (response string, err error) {


	if d.isMachineInAbortedState() {
		err = fmt.Errorf("Machine in ABORTED STATE for deck: %v. Please home the machine first.", d.name)
		return "", err
	}

	if pulse < minimumPulsesThreshold {
		fmt.Println("Current pulse: ", pulse, " is less than minimumPulsesThreshold. Avoiding Motor Movements for motor:", motorNum, ", deck: ", d.name)
		return "SUCCESS", nil
	}

	wrotePulses.Store(d.name, 0)
	executedPulses.Store(d.name, 0)
	deckAndNumber := DeckNumber{Deck: d.name, Number: motorNum}

	var results []byte

	//
	//  Detach Magnet Fully if the deck is to move and magnet is in attached State
	//

	if d.getMagnetState() != detached && motorNum == K5_Deck {
		response, err = d.fullDetach()
		if err != nil {
			fmt.Println(err)
			return "", fmt.Errorf("There was issue Detaching Magnet before moving the deck. Error: %v", err)
		}
	}

	fmt.Println("Moving: ", motorNum, pulse/motors[deckAndNumber]["steps"], "mm in ", direction, "for deck:", d.name)

	// Switch OFF The motor

	err = d.DeckDriver.WriteSingleCoil(MODBUS_EXTRACTION[d.name]["M"][0], OFF)
	if err != nil {
		fmt.Println("error writing Switch Off : ", err, d.name)
		return
	}

	err = d.DeckDriver.WriteSingleCoil(MODBUS_EXTRACTION[d.name]["M"][1], OFF)
	if err != nil {
		fmt.Println("error writing Completion Off : ", err, d.name)
		return "", err
	}

	results, err = d.DeckDriver.WriteSingleRegister(MODBUS_EXTRACTION[d.name]["D"][202], pulse)
	if err != nil {
		fmt.Println("error writing pulse : ", err, d.name)
		return "", err
	}
	fmt.Println("Wrote Pulse. res : ", results)
	wrotePulses.Store(d.name, pulse)

	results, err = d.DeckDriver.WriteSingleRegister(MODBUS_EXTRACTION[d.name]["D"][200], speed)
	if err != nil {
		fmt.Println("error writing speed : ", err, d.name)
		return "", err
	}
	fmt.Println("Wrote Speed. res : ", results)

	results, err = d.DeckDriver.WriteSingleRegister(MODBUS_EXTRACTION[d.name]["D"][204], ramp)
	if err != nil {
		fmt.Println("error writing RAMP : ", err, d.name)
		return "", err
	}
	fmt.Println("Wrote Ramp. res : ", results)

	results, err = d.DeckDriver.WriteSingleRegister(MODBUS_EXTRACTION[d.name]["D"][206], direction)
	if err != nil {
		fmt.Println("error writing direction : ", err, d.name)
		return "", err
	}
	fmt.Println("Wrote direction. res : ", results)

	results, err = d.DeckDriver.WriteSingleRegister(MODBUS_EXTRACTION[d.name]["D"][226], motorNum)
	if err != nil {
		fmt.Println("error writing motor num: ", err, d.name)
		return "", err
	}
	fmt.Println("Wrote motorNum. res : ", results)
	// Check if User has paused the run/operation
	for {
		if d.isMachineInPausedState() {
			fmt.Println("Machine in PAUSED state for deck: %v", d.name)
			time.Sleep(400 * time.Millisecond)
		} else {
			break
		}
	}

	// Switching Motor ON
	err = d.DeckDriver.WriteSingleCoil(MODBUS_EXTRACTION[d.name]["M"][0], ON)
	if err != nil {
		fmt.Println("error Writing On/Off : ", err, d.name)
		return "", err
	}

	// Our Run is in Progress
	fmt.Println("Movements in Progress for deck: ", d.name)

	for {
		
		if temp := d.getExecutedPulses(); temp == 0 {
			err = fmt.Errorf("executedPulses isn't loaded!")
			return
			// Write executed pulses to Position
		} else if d.isMachineInAbortedState() {
			positions[deckAndNumber] += float64(temp) / float64(motors[deckAndNumber]["steps"])
			fmt.Println("pos", positions[deckAndNumber])
			err = fmt.Errorf("Operation was ABORTED!")
			return "", err
		}

		results, err = d.DeckDriver.ReadCoils(MODBUS_EXTRACTION[d.name]["M"][1], uint16(1))
		if err != nil {
			fmt.Println("error while reading completion  : ", err, d.name)
			// Let this reading failure be intolerant
			return
		}

		if len(results) > 0 {
			if int(results[0]) == 1 {
				fmt.Println("Completion returned ---> ", results, d.name)
				response, err = d.switchOffMotor()
				if err != nil {
					fmt.Println("err: from setUp--> ", err, d.name)
					return
				}
				distanceMoved := float64(pulse) / float64(motors[DeckNumber{Deck: d.name, Number: motorNum}]["steps"])
				switch direction {
				// Away from Sensor
				case REV:
					positions[deckAndNumber] += distanceMoved
				// Towards Sensor
				case FWD:
					if (positions[deckAndNumber] - distanceMoved) < 0 {
						positions[deckAndNumber] = 0
						fmt.Println("Motor Just moved to negative distance for deck: ", d.name)
					}
					positions[deckAndNumber] -= distanceMoved
				default:
					fmt.Println("Unknown Direction was found")
					return "", fmt.Errorf("Unknown Direction was found: %v", direction)
				}
				fmt.Println("pos", positions[deckAndNumber], d.name)
				return "RUN Completed", nil
			}
		}

		if direction == REV {
			goto skipSensor
		}
		results, err = d.DeckDriver.ReadCoils(MODBUS_EXTRACTION[d.name]["M"][2], uint16(1))
		if err != nil {
			fmt.Println("error reading Sensor : ", err, d.name)
			return "", err
		}

		fmt.Println("Sensor returned ---> ", results, d.name)
		if len(results) > 0 {
			if int(results[0]) == sensorCut {
				fmt.Println("Sensor returned ---> ", results[0], d.name)
				response, err = d.switchOffMotor()
				if err != nil {
					fmt.Println("Sensor err : ", err, d.name)
					return "", err
				}
				positions[deckAndNumber] = calibs[deckAndNumber]
				fmt.Println("pos", positions[deckAndNumber], d.name)
				return
			}
		}

	skipSensor:
		switch pulse {
		// Avoiding initialSensorCutMagnetPulses as its duplicate
		case initialSensorCutSyringeModulePulses, initialSensorCutDeckPulses, initialSensorCutSyringePulses:
			time.Sleep(400 * time.Millisecond)
		case finalSensorCutPulses:
			time.Sleep(50 * time.Millisecond)
		default:
			time.Sleep(500 * time.Millisecond)
		}
	}

	return "RUN Completed", nil
}

func (d *Compact32Deck) switchOffMotor() (response string, err error) {

	err = d.DeckDriver.WriteSingleCoil(MODBUS_EXTRACTION[d.name]["M"][0], OFF)
	if err != nil {
		fmt.Println("err Switching motor off: ", err)
		return "", err
	}

	return "SUCCESS", nil
}

func (d *Compact32Deck) switchOffHeater() (response string, err error) {

	// Switch off Heater
	err = d.DeckDriver.WriteSingleCoil(MODBUS_EXTRACTION[d.name]["M"][3], OFF)
	if err != nil {
		fmt.Println("err Switching off the heater: ", err)
		return "", err
	}
	fmt.Println("Switched off the heater--> for deck ", d.name)

	return "SUCCESS", nil
}

func (d *Compact32Deck) switchOnUVLight() (response string, err error) {

	err = d.DeckDriver.WriteSingleCoil(MODBUS_EXTRACTION[d.name]["M"][6], ON)
	if err != nil {
		fmt.Println("err Switching on the UV Light: ", err)
		return "", err
	}
	fmt.Println("Switched on the UV Light--> for deck ", d.name)

	return "SUCCESS", nil
}

func (d *Compact32Deck) switchOffUVLight() (response string, err error) {

	err = d.DeckDriver.WriteSingleCoil(MODBUS_EXTRACTION[d.name]["M"][6], OFF)
	if err != nil {
		fmt.Println("err Switching off the UV Light: ", err)
		return "", err
	}
	fmt.Println("Switched off the UV Light--> for deck ", d.name)

	return "SUCCESS", nil
}

func (d *Compact32Deck) readExecutedPulses() (response string, err error) {

	results, err := d.DeckDriver.ReadHoldingRegisters(MODBUS_EXTRACTION[d.name]["D"][212], uint16(1))
	if err != nil {
		fmt.Println("err : ", err, d.name)
		return "", err
	}

	fmt.Printf("Read D212AddressBytesUint16. res : %+v \n", results)
	if len(results) > 0 {
		executedPulses.Store(d.name, binary.BigEndian.Uint16(results))
	} else {
		err = fmt.Errorf("couldn't read D212")
		return "", err
	}
	fmt.Println("Read D212 Pulses -> ", binary.BigEndian.Uint16(results))

	return "D212 Reading SUCESS", nil

}

