package simulator

import (
	"fmt"
	logger "github.com/sirupsen/logrus"
	"mylab/cpagent/plc"
	"time"
)

func (d *SimulatorDriver) simulateWriteMultipleRegisters(address, quantity uint16, value []byte) (results []byte, err error) {
	// TODO: Implement this only when related method arrive
	return
}

func (d *SimulatorDriver) simulateWriteSingleRegister(address, value uint16) (results []byte, err error) {
	err = d.checkForValidAddress("D", address)
	if err != nil {
		return
	}

	d.setRegister("D", address, value)

	results = []byte{uint8(value >> 8), uint8(value & 0xff)}

	logger.Infoln("Inside simulateWriteSingleRegister for deck ", d.DeckName, " result: ", results, ". address: ", address)
	return
}

func (d *SimulatorDriver) simulateReadHoldingRegisters(address, quantity uint16) (results []byte, err error) {
	err = d.checkForValidAddress("D", address)
	if err != nil {
		return
	}
	if quantity != 1 {
		err = fmt.Errorf("invalid register reading quantity")
		return
	}

	value := d.readRegister("D", address)

	results = []byte{uint8(value >> 8), uint8(value & 0xff)}

	logger.Infoln("Inside simulateReadHoldingRegisters for deck ", d.DeckName, " result: ", results, ". address: ", address)
	return
}

func (d *SimulatorDriver) simulateReadCoils(address, quantity uint16) (results []byte, err error) {
	err = d.checkForValidAddress("M", address)
	if err != nil {
		return
	}

	if quantity != 1 {
		err = fmt.Errorf("invalid coil reading quantity")
		return
	}

	value := d.readRegister("M", address)

	results = []byte{uint8(value & 0xff)}

	logger.Infoln("Inside simulateReadCoils for deck ", d.DeckName, " result: ", results, ". address: ", address)
	return
}

func (d *SimulatorDriver) simulateWriteSingleCoil(address, value uint16) (err error) {
	err = d.checkForValidAddress("M", address)
	if err != nil {
		return
	}

	d.setRegister("M", address, value)

	results := []byte{uint8(value & 0xff)}

	logger.Infoln("Inside simulateWriteSingleCoil for deck ", d.DeckName, " result: ", results, ". address: ", address)

	// Calling in go routine so that masterLock is free
	switch address {
	case plc.MODBUS_EXTRACTION[d.DeckName]["M"][0]:
		if value == plc.ON {
			go d.simulateOnMotor()
		}

	case plc.MODBUS_EXTRACTION[d.DeckName]["M"][3]:
		go d.simulateHeater(value)
	}

	return
}

func (d *SimulatorDriver) checkForValidAddress(registerType string, address uint16) (err error) {
	switch registerType {
	case "M":
		// valid range 0-8
		lowestMAddress := plc.MODBUS_EXTRACTION[d.DeckName][registerType][0]
		highestMAddress := plc.MODBUS_EXTRACTION[d.DeckName][registerType][8]

		if address >= lowestMAddress && address <= highestMAddress {
			return
		}

	case "D":
		// valid range 200-226
		lowestDAddress := plc.MODBUS_EXTRACTION[d.DeckName][registerType][200]
		highestDAddress := plc.MODBUS_EXTRACTION[d.DeckName][registerType][226]

		// check for divisibility by 2 as well
		if address >= lowestDAddress && address <= highestDAddress && address%2 != 1 {
			return
		}

	default:
		err = fmt.Errorf("Invalid register Type")
	}

	err = fmt.Errorf("Invalid register address")
	return
}

// TODO: Handle Motor Movements
/*
    When Switch On Motor
   - Monitor Sensor Has Cut
   - Write Pulses to D212
   - Monitor D212 Pulses

*/

func (d *SimulatorDriver) simulateOnMotor() (err error) {

	for {
		if !d.isMotorInProgress() {
			break
		}
		time.Sleep(200 * time.Millisecond)
	}
	d.setMotorInProgress()
	defer d.resetMotorInProgress()

	// Reset D212
	logger.Infoln("Reset D212")
	d.setRegister("D", plc.MODBUS_EXTRACTION[d.DeckName]["D"][212], 0)
	// Reset Sensor Cut
	d.setRegister("M", plc.MODBUS_EXTRACTION[d.DeckName]["M"][2], plc.SensorUncut)

	// Reset map vars
	d.resetSensorDone()
	d.resetMotorDone()

	// Pulses Register
	// If Pulses greater than 0 && Direction is towards Sensor
	if d.readRegister("D", plc.MODBUS_EXTRACTION[d.DeckName]["D"][202]) > 0 &&
		d.readRegister("D", plc.MODBUS_EXTRACTION[d.DeckName]["D"][206]) == plc.TowardsSensor {
		// Call MonitorSensorCut in a go routine
		// Only makes sense when we are going towards Sensor
		go d.monitorSensorCut()
	}

	// Update Pulses every 100 Millisecond
	err = d.updatePulses()

	return
}

func (d *SimulatorDriver) updatePulses() (err error) {
	motorNum := d.readRegister("D", plc.MODBUS_EXTRACTION[d.DeckName]["D"][226])
	speed := d.readRegister("D", plc.MODBUS_EXTRACTION[d.DeckName]["D"][200])
	pulses := d.readRegister("D", plc.MODBUS_EXTRACTION[d.DeckName]["D"][202])

	currentPulses := uint16(0)

	for {
		switch {
		case d.isMotorDone():
			logger.Infoln("completion was done for deck ", d.DeckName)
			return
		case d.isSensorDone():
			logger.Infoln("sensor has cut for deck ", d.DeckName)
			return
		default:
			time.Sleep(200 * time.Millisecond)
			// if motor is OFF then return
			if plc.OFF == d.readRegister("M", plc.MODBUS_EXTRACTION[d.DeckName]["M"][0]) {
				return
			}
			// if motor is changed then return
			if motorNum != d.readRegister("D", plc.MODBUS_EXTRACTION[d.DeckName]["D"][226]) {
				return
			}

			// We are updating D212 after every 0.1 second
			d212Val := d.readRegister("D", plc.MODBUS_EXTRACTION[d.DeckName]["D"][212])
			d.setRegister("D", plc.MODBUS_EXTRACTION[d.DeckName]["D"][212], d212Val+uint16(float64(speed)*0.1))
			currentPulses = d.readRegister("D", plc.MODBUS_EXTRACTION[d.DeckName]["D"][212])
			logger.Infoln("D212 for deck", d.DeckName, " value is: ", currentPulses)

			if currentPulses > pulses {
				// D212 updated
				d.setRegister("D", plc.MODBUS_EXTRACTION[d.DeckName]["D"][212], pulses)
				// Completion Done
				d.setRegister("M", plc.MODBUS_EXTRACTION[d.DeckName]["M"][1], uint16(1))
				// Completion is monitored here itself
				logger.Infoln("Completion is Done for deck", d.DeckName)
				d.setMotorDone()
			}
		}
	}
}

func (d *SimulatorDriver) monitorSensorCut() (err error) {

	deckAndMotor := plc.DeckNumber{Deck: d.DeckName, Number: d.readRegister("D", plc.MODBUS_EXTRACTION[d.DeckName]["D"][226])}

	// shift = D212 Pulses / Motor steps
	for {
		// putting declaration inside the loop cause what if motor change happens within 100 ms!!
		// if motor is OFF then return
		if plc.OFF == d.readRegister("M", plc.MODBUS_EXTRACTION[d.DeckName]["M"][0]) {
			return
		}
		// if motor is changed then return
		if deckAndMotor.Number != d.readRegister("D", plc.MODBUS_EXTRACTION[d.DeckName]["D"][226]) {
			return
		}
		// if direction is changed then return
		if plc.TowardsSensor != d.readRegister("D", plc.MODBUS_EXTRACTION[d.DeckName]["D"][206]) {
			return
		}

		// Check For Sensor Cut
		shift := float64(d.readRegister("D", plc.MODBUS_EXTRACTION[d.DeckName]["D"][212])) / float64(plc.Motors[deckAndMotor]["steps"])
		if plc.Positions[deckAndMotor]-shift <= plc.Calibs[deckAndMotor] {
			d.setRegister("M", plc.MODBUS_EXTRACTION[d.DeckName]["M"][2], plc.SensorCut)
			logger.Infoln("Sensor Cut is Done for deck", d.DeckName)
			d.setSensorDone()
			return
		}

		time.Sleep(200 * time.Millisecond)
	}
}

// TODO: Handle Heater M 3
/*
   When Heater On
   - Monitor Heater
   - Set Present Value
   - Calibrate Temperature
   When Heater Off
   - Close Heater Monitoring
*/
func (d *SimulatorDriver) simulateHeater(value uint16) (err error) {
	// TODO: Implement this with help from @GautamRege
	return
}
