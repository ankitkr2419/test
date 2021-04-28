package simulator

import (
	"fmt"
	logger "github.com/sirupsen/logrus"
	"mylab/cpagent/plc"
	"time"
)

/*
    When Switch On Motor
   - Monitor Sensor Has Cut
   - Write Pulses to D212
   - Monitor D212 Pulses

*/

// INFO: Change Log Level to Info if debugging

func (d *SimulatorDriver) simulateOnMotor() (err error) {

	for {
		if !d.isMotorInProgress() {
			break
		}
		time.Sleep(200 * time.Millisecond)
	}
	d.setMotorInProgress()
	defer d.resetMotorInProgress()

	if d.readRegister("M", plc.MODBUS_EXTRACTION[d.DeckName]["M"][5]) == plc.ON {
		err = d.simulateShakerMotor()
		return
	}

	// Reset D212
	logger.Debugln("Reset D212")
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

func (d *SimulatorDriver) simulateShakerMotor() (err error) {
	logger.Infoln("Shaker motor simulation started")
	for {
		time.Sleep(200 * time.Millisecond)

		if d.readRegister("M", plc.MODBUS_EXTRACTION[d.DeckName]["M"][5]) == plc.OFF {
			logger.Infoln("Shaker motor simulation stopped")
			return
		}
		// if motor is not K8_Shaker then return
		if plc.K8_Shaker != d.readRegister("D", plc.MODBUS_EXTRACTION[d.DeckName]["D"][226]) {
			err = fmt.Errorf("Motor is different than shaker!")
			return
		}
	}
}

func (d *SimulatorDriver) updatePulses() (err error) {
	motorNum := d.readRegister("D", plc.MODBUS_EXTRACTION[d.DeckName]["D"][226])
	speed := d.readRegister("D", plc.MODBUS_EXTRACTION[d.DeckName]["D"][200])
	pulses := d.readRegister("D", plc.MODBUS_EXTRACTION[d.DeckName]["D"][202])

	currentPulses := uint16(0)

	for {
		switch {
		case d.isMotorDone():
			logger.Debugln("completion was done for deck ", d.DeckName)
			return
		case d.isSensorDone():
			logger.Debugln("sensor has cut for deck ", d.DeckName)
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
			logger.Debugln("D212 for deck", d.DeckName, " value is: ", currentPulses)

			if currentPulses > pulses {
				// D212 updated
				d.setRegister("D", plc.MODBUS_EXTRACTION[d.DeckName]["D"][212], pulses)
				// Completion Done
				d.setRegister("M", plc.MODBUS_EXTRACTION[d.DeckName]["M"][1], uint16(1))
				// Completion is monitored here itself
				logger.Debugln("Completion is Done for deck", d.DeckName)
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
			logger.Debugln("Sensor Cut is Done for deck", d.DeckName)
			d.setSensorDone()
			return
		}

		time.Sleep(200 * time.Millisecond)
	}
}
