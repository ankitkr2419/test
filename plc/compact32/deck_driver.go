package compact32

import (
	"encoding/binary"
	"fmt"
	"time"
)

func (d *Compact32Deck) SetupMotor(speed, pulse, ramp, direction, motorNum uint16) (response string, err error) {

	wrotePulses[d.name] = 0
	executedPulses[d.name] = 0
	deckAndNumber := DeckNumber{Deck: d.name, Number: motorNum}

	var results []byte

	if aborted[d.name] {
		err := fmt.Errorf("Machine in ABORTED STATE")
		return "", err
	}

	// Switch OFF The motor
	err = d.DeckDriver.WriteSingleCoil(MODBUS_EXTRACTION[d.name]["M"][0], OFF)
	if err != nil {
		fmt.Println("err : ", err)
		return
	}

	err = d.DeckDriver.WriteSingleCoil(MODBUS_EXTRACTION[d.name]["M"][1], OFF)
	if err != nil {
		fmt.Println("err : ", err)
		return "", err
	}

	results, err = d.DeckDriver.WriteSingleRegister(MODBUS_EXTRACTION[d.name]["D"][202], pulse)
	if err != nil {
		fmt.Println("err : ", err)
		return "", err
	}
	fmt.Println("Wrote Pulse. res : ", results)
	wrotePulses[d.name] = pulse

	results, err = d.DeckDriver.WriteSingleRegister(MODBUS_EXTRACTION[d.name]["D"][200], speed)
	if err != nil {
		fmt.Println("err : ", err)
		return "", err
	}
	fmt.Println("Wrote Speed. res : ", results)

	results, err = d.DeckDriver.WriteSingleRegister(MODBUS_EXTRACTION[d.name]["D"][204], ramp)
	if err != nil {
		fmt.Println("err : ", err)
		return "", err
	}
	fmt.Println("Wrote Ramp. res : ", results)

	results, err = d.DeckDriver.WriteSingleRegister(MODBUS_EXTRACTION[d.name]["D"][206], direction)
	if err != nil {
		fmt.Println("err : ", err)
		return "", err
	}
	fmt.Println("Wrote direction. res : ", results)

	results, err = d.DeckDriver.WriteSingleRegister(MODBUS_EXTRACTION[d.name]["D"][226], motorNum)
	if err != nil {
		fmt.Println("err : ", err)
		return "", err
	}
	fmt.Println("Wrote motorNum. res : ", results)

	for {
		if paused[d.name] {
			fmt.Println("Machine in PAUSED state")
		} else {
			break
		}
		time.Sleep(400 * time.Millisecond)
	}

	err = d.DeckDriver.WriteSingleCoil(MODBUS_EXTRACTION[d.name]["M"][0], ON)
	if err != nil {
		fmt.Println("err : ", err)
		return "", err
	}

	results, err = d.DeckDriver.ReadCoils(MODBUS_EXTRACTION[d.name]["M"][0], uint16(1))
	if err != nil {
		fmt.Println("err : ", err)
		return "", err
	}
	fmt.Printf("Read On/Off Coil. res : %+v \n", results)

	fmt.Println("Blocked")

	for {
		if aborted[d.name] {
			err = fmt.Errorf("Operation was ABORTED!")
			return "", err
		}
		results, err = d.DeckDriver.ReadCoils(MODBUS_EXTRACTION[d.name]["M"][1], uint16(1))
		if err != nil {
			fmt.Println("err : ", err)
			return "", err
		}
		if len(results) > 0 {
			if int(results[0]) == 1 {
				fmt.Println("Completion returned ---> ", results)
				d.SwitchOffMotor()
				distanceMoved := float64(pulse) / float64(motors[DeckNumber{Deck: d.name, Number: motorNum}]["steps"])
				switch direction {
				// Away from Sensor
				case 0:
					positions[deckAndNumber] += distanceMoved
				// Towards Sensor
				case 1:
					if (positions[deckAndNumber] - distanceMoved) < 0 {
						return "", fmt.Errorf("Motor Just moved to negative distnace!")
					}
					positions[deckAndNumber] -= distanceMoved
				default:
					return "", fmt.Errorf("Unknown Direction was found: %v", direction)
				}
				//statusChannel <- 1
				return "RUN Completed", nil
			}
		}

		if direction == REV && pulse != moveOppositeSensorPulses {
			goto skipSensor
		}
		results, err = d.DeckDriver.ReadCoils(MODBUS_EXTRACTION[d.name]["M"][2], uint16(1))
		if err != nil {
			fmt.Println("err : ", err)
			return "", err
		}
		fmt.Println("Sensor returned ---> ", results)
		if len(results) > 0 {
			if int(results[0]) == sensorCut && pulse != moveOppositeSensorPulses {
				fmt.Println("Sensor returned ---> ", results[0])
				response, err = d.SwitchOffMotor()
				//statusChannel <- 3
				sensorHasCut[d.name] = true
				positions[deckAndNumber] = 0
				// TODO caliberation
				return
			} else if int(results[0]) == sensorUncut && pulse == moveOppositeSensorPulses {
				fmt.Println("Sensor returned ---> ", results[0])
				response, err = d.SwitchOffMotor()
				sensorHasCut[d.name] = false
				time.Sleep(100 * time.Millisecond)
				response, err = d.SetupMotor(motors[deckAndNumber]["fast"], reverseAfterNonCutPulses, motors[deckAndNumber]["ramp"], REV, motorNum)
				//statusChannel <- 4
				return
			}
		}

	skipSensor:
		switch pulse {
		// Avoiding initialSensorCutMagnetPulses as its duplicate
		case initialSensorCutSyringeModulePulses, initialSensorCutDeckPulses, initialSensorCutSyringePulses:
			time.Sleep(100 * time.Millisecond)
		case finalSensorCutPulses:
			time.Sleep(20 * time.Millisecond)
		default:
			time.Sleep(500 * time.Millisecond)
		}
	}

	return "RUN Completed", nil
}

func (d *Compact32Deck) SwitchOffMotor() (response string, err error) {

	err = d.DeckDriver.WriteSingleCoil(MODBUS_EXTRACTION[d.name]["M"][0], OFF)
	if err != nil {
		fmt.Println("err : ", err)
		return "", err
	}

	return "SUCCESS", nil
}

func (d *Compact32Deck) ReadExecutedPulses() (response string, err error) {

	results, err := d.DeckDriver.ReadHoldingRegisters(MODBUS_EXTRACTION[d.name]["D"][212], uint16(1))
	if err != nil {
		fmt.Println("err : ", err)
		return "", err
	}

	fmt.Printf("Read D212AddressBytesUint16. res : %+v \n", results)
	if len(results) > 0 {
		executedPulses[d.name] = binary.BigEndian.Uint16(results)
	} else {
		err = fmt.Errorf("couldn't read D212")
		return "", err
	}
	fmt.Println("Read D212 Pulses -> ", executedPulses[d.name])

	return "D212 Reading SUCESS", nil

}

func (d *Compact32Deck) Homing() (response string, err error) {

	aborted[d.name] = false

	if runInProgress[d.name] {
		err = fmt.Errorf("previous run is already in progress... wait or abort it")
		return
	}

	runInProgress[d.name] = true
	defer d.ResetRunInProgress()

	fmt.Println("Moving Syringe DOWN till sensor cuts it")
	response, err = d.SyringeHoming()
	if err != nil {
		return
	}

	fmt.Println("Moving Syringe Module UP till sensor cuts it")
	response, err = d.SyringeModuleHoming()
	if err != nil {
		return
	}

	fmt.Println("Moving deck forward till sensor cuts it")
	response, err = d.DeckHoming()
	if err != nil {
		return
	}

	fmt.Println("Homing Magnet")
	response, err = d.MagnetHoming()
	if err != nil {
		return
	}

	fmt.Println("Homing Completed Successfully")

	return "HOMING SUCCESS", nil
}

// ***NOTE***
// * In Syringe Sensor is DOWN and not UP.
// * This is exactly opposite of Syringe Module and Magnet Up/Down
// * Thus we need syringeUP and syringeDOWN

func (d *Compact32Deck) SyringeHoming() (response string, err error) {

	deckAndNumber := DeckNumber{Deck: d.name, Number: K10_Syringe_LHRH}

	fmt.Println("Syringe is moving down until sensor not cut")

	// response, err = d.SetupMotor(uint16(2000), uint16(26666), uint16(100), UP, uint16(10))
	response, err = d.SetupMotor(motors[deckAndNumber]["fast"], uint16(initialSensorCutSyringePulses), motors[deckAndNumber]["ramp"], syringeDOWN, deckAndNumber.Number)
	if err != nil {
		return
	}

	time.Sleep(100 * time.Millisecond)

	fmt.Println("Aspiring and getting cut then aspiring 2000")
	//response, err = d.SetupMotor(uint16(2000), uint16(19999), uint16(100), DOWN, uint16(10))
	response, err = d.SetupMotor(motors[deckAndNumber]["fast"], moveOppositeSensorPulses, motors[deckAndNumber]["ramp"], syringeUP, deckAndNumber.Number)
	if err != nil {
		return
	}
	time.Sleep(100 * time.Millisecond)

	fmt.Println("Syringe dispencing again")
	// response, err = d.SetupMotor(uint16(500), uint16(2999), uint16(100), UP, uint16(10))
	response, err = d.SetupMotor(motors[deckAndNumber]["slow"], finalSensorCutPulses, motors[deckAndNumber]["ramp"], syringeDOWN, deckAndNumber.Number)
	if err != nil {
		return
	}

	fmt.Println("Syringe homing is completed")

	return "SYRINGE HOMING COMPLETED", nil
}

func (d *Compact32Deck) SyringeModuleHoming() (response string, err error) {

	deckAndNumber := DeckNumber{Deck: d.name, Number: K9_Syringe_Module_LHRH}

	fmt.Println("Syringe Module moving Up")
	// response, err = d.SetupMotor(uint16(2000), initialSensorCutSyringeModulePulses, uint16(100), UP, uint16(9))
	response, err = d.SetupMotor(motors[deckAndNumber]["fast"], initialSensorCutSyringeModulePulses, motors[deckAndNumber]["ramp"], UP, deckAndNumber.Number)
	if err != nil {
		return
	}

	fmt.Println("After First Fast Moving Up and getting Cut")

	time.Sleep(100 * time.Millisecond)

	fmt.Println("Syringe Module moving Down 20 mm or More.")
	// response, err = d.SetupMotor(uint16(2000), moveOppositeSensorPulses, uint16(100), DOWN, uint16(9))
	response, err = d.SetupMotor(motors[deckAndNumber]["fast"], moveOppositeSensorPulses, motors[deckAndNumber]["ramp"], DOWN, deckAndNumber.Number)
	if err != nil {
		return
	}

	time.Sleep(100 * time.Millisecond)

	fmt.Println("Syringe Module moving Up")
	// response, err = d.SetupMotor(uint16(500), finalSensorCutPulses, uint16(100), UP, uint16(9))
	response, err = d.SetupMotor(motors[deckAndNumber]["slow"], finalSensorCutPulses, motors[deckAndNumber]["ramp"], UP, deckAndNumber.Number)
	if err != nil {
		return
	}

	fmt.Println("After Final Slow Moving Up and getting Cut")

	return "SYRINGE HOMING SUCCESS", nil
}

func (d *Compact32Deck) DeckHoming() (response string, err error) {

	deckAndNumber := DeckNumber{Deck: d.name, Number: K5_Deck}

	fmt.Println("Deck is moving forward")
	// response, err = d.SetupMotor(uint16(2000), initialSensorCutDeckPulses, uint16(100), FWD, uint16(5))
	response, err = d.SetupMotor(motors[deckAndNumber]["fast"], initialSensorCutDeckPulses, motors[deckAndNumber]["ramp"], FWD, deckAndNumber.Number)
	if err != nil {
		return
	}

	//	sensorHasCut[d.name] = false
	time.Sleep(100 * time.Millisecond)
	fmt.Println("Deck is moving back by and after not cut -> 2000")
	// response, err = d.SetupMotor(uint16(2000), moveOppositeSensorPulses, uint16(100), REV, uint16(5))
	response, err = d.SetupMotor(motors[deckAndNumber]["fast"], moveOppositeSensorPulses, motors[deckAndNumber]["ramp"], REV, deckAndNumber.Number)
	if err != nil {
		return
	}

	time.Sleep(100 * time.Millisecond)

	fmt.Println("Deck is moving forward again by 2999")
	// response, err = d.SetupMotor(uint16(500), uint16(2999), uint16(100), FWD, uint16(5))
	response, err = d.SetupMotor(motors[deckAndNumber]["slow"], finalSensorCutPulses, motors[deckAndNumber]["ramp"], FWD, deckAndNumber.Number)
	if err != nil {
		return
	}

	fmt.Println("Deck homing is completed.")

	return "DECK HOMING SUCCESS", nil
}

func (d *Compact32Deck) MagnetHoming() (response string, err error) {

	deckAndNumber := DeckNumber{Deck: d.name, Number: K7_Magnet_Rev_Fwd}

	response, err = d.MagnetUpDownHoming()
	if err != nil {
		return
	}
	response, err = d.MagnetFwdRevHoming()
	if err != nil {
		return
	}

	time.Sleep(100 * time.Millisecond)
	fmt.Println("Moving Magnet Back by 50mm")

	// response, err = d.SetupMotor(uint16(2000), uint16(10000), uint16(100), REV, uint16(7))
	// TODO: Use consumable distance var instead of moveMagnetAfterFinalCutPulses
	response, err = d.SetupMotor(motors[deckAndNumber]["fast"], moveMagnetAfterFinalCutPulses, motors[deckAndNumber]["ramp"], REV, deckAndNumber.Number)
	if err != nil {
		return
	}

	return "MAGNET HOMING SUCCESS", nil
}

func (d *Compact32Deck) MagnetUpDownHoming() (response string, err error) {

	deckAndNumber := DeckNumber{Deck: d.name, Number: K6_Magnet_Up_Down}

	fmt.Println("Magnet is moving up")
	// response, err = d.SetupMotor(uint16(2000), initialSensorCutMagnetPulses, uint16(100), UP, uint16(6))
	response, err = d.SetupMotor(motors[deckAndNumber]["fast"], initialSensorCutMagnetPulses, motors[deckAndNumber]["ramp"], UP, deckAndNumber.Number)
	if err != nil {
		return
	}

	time.Sleep(100 * time.Millisecond)
	fmt.Println("Magnet is moving down by and after not cut -> 2000")
	// response, err = d.SetupMotor(uint16(2000), moveOppositeSensorPulses, uint16(100), DOWN, uint16(6))
	response, err = d.SetupMotor(motors[deckAndNumber]["fast"], moveOppositeSensorPulses, motors[deckAndNumber]["ramp"], DOWN, deckAndNumber.Number)
	if err != nil {
		return
	}

	time.Sleep(100 * time.Millisecond)

	fmt.Println("Magnet is moving up again by 2999 till sensor cuts")
	// response, err = d.SetupMotor(uint16(500), uint16(2999), uint16(100), UP, uint16(6))
	response, err = d.SetupMotor(motors[deckAndNumber]["slow"], finalSensorCutPulses, motors[deckAndNumber]["ramp"], UP, deckAndNumber.Number)

	fmt.Println("Magnet Up/Down homing is completed.")

	return "MAGNET UP/DOWN HOMING SUCCESS", nil
}

func (d *Compact32Deck) MagnetFwdRevHoming() (response string, err error) {

	deckAndNumber := DeckNumber{Deck: d.name, Number: K7_Magnet_Rev_Fwd}

	fmt.Println("Magnet is moving forward")
	// response, err = d.SetupMotor(uint16(2000), initialSensorCutMagnetPulses, uint16(100), FWD, uint16(7))
	response, err = d.SetupMotor(motors[deckAndNumber]["fast"], initialSensorCutMagnetPulses, motors[deckAndNumber]["ramp"], FWD, deckAndNumber.Number)
	if err != nil {
		return
	}

	time.Sleep(100 * time.Millisecond)
	fmt.Println("Magnet is moving back by and after not cut -> 2000")
	// response, err = d.SetupMotor(uint16(2000), moveOppositeSensorPulses, uint16(100), REV, uint16(7))
	response, err = d.SetupMotor(motors[deckAndNumber]["fast"], moveOppositeSensorPulses, motors[deckAndNumber]["ramp"], REV, deckAndNumber.Number)
	if err != nil {
		return
	}

	time.Sleep(100 * time.Millisecond)

	fmt.Println("Magnet is moving forward again by 2999")
	// response, err = d.SetupMotor(uint16(500), uint16(2999), uint16(100), FWD, uint16(7))
	response, err = d.SetupMotor(motors[deckAndNumber]["slow"], finalSensorCutPulses, motors[deckAndNumber]["ramp"], FWD, deckAndNumber.Number)

	fmt.Println("Magnet Up/Down homing is completed.")

	return "MAGNET FWD/REV HOMING SUCCESS", nil
}
