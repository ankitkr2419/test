package simulator

import (
	"fmt"
	logger "github.com/sirupsen/logrus"
	"mylab/cpagent/plc"
	"time"
)

// INFO: Change Log Level to Info if debugging

/*
   When Heater On
   - Monitor Heater
   - Set Present Value
   - Calibrate Temperature
   When Heater Off
   - Close Heater Monitoring
*/

type Celsius float64

const (
	// roomTemp already declared in pcr.go
	// for first 80% of heating up
	rampUpFastRate Celsius = 1 // increase temp by 1 degree celsius per second when heater is ON
	// rampUpTemp is temperature increase with heater ON
	// for last 20% of heating up
	rampUpSlowRate Celsius = 0.2 // increase temp by 0.2 degree celsius per second when heater is ON
	// rampDownTemp is temperature decrease with heater OFF
	rampDownFastRate Celsius = 0.2  // decrease temp by 0.2 degree celsius per second when heater is OFF
	rampDownSlowRate Celsius = 0.05 // decrease temp by 0.05 degree celsius per second when heater is OFF
)

// Currently we only handle if both shaker's start together and not separate

// WARN: Be careful with temperature, its multiplied by 10 for machine
func (d *SimulatorDriver) simulateOnHeater() (err error) {

	// We don't need separate handle for separate shaker part
	heaterNum := d.readRegister("D", plc.MODBUS_EXTRACTION[d.DeckName]["D"][222])
	if heaterNum != uint16(3) {
		err = fmt.Errorf("simulator support is only provided for LH, RH shaker heaters together only")
		return
	}
	for {
		if !d.isHeaterInProgress() {
			break
		}
		time.Sleep(delay * 4 * time.Millisecond)
	}
	d.setHeaterInProgress()
	defer d.resetHeaterInProgress()

	logger.Debugln("Heater has started!!")

	// Update Temperature every sec
	d.updateTemperature()

	return
}

func (d *SimulatorDriver) updateTemperature() {

	var temp uint16 = 0

	for {
		time.Sleep(delay * 20 * time.Millisecond)
		if plc.OFF == d.readRegister("M", plc.MODBUS_EXTRACTION[d.DeckName]["M"][3]) {
			go d.coolDown()
			return
		}

		targetTemp := d.readRegister("D", plc.MODBUS_EXTRACTION[d.DeckName]["D"][208])
		currentTempLH := d.readRegister("D", plc.MODBUS_EXTRACTION[d.DeckName]["D"][210])
		currentTempRH := d.readRegister("D", plc.MODBUS_EXTRACTION[d.DeckName]["D"][224])

		// 1 degree extra
		logger.Debugln("Heating Up -> targetTemp :", targetTemp, ", currentTempLH :", currentTempLH, ", currentTempRH", currentTempRH)

		if currentTempLH < targetTemp+10 {
			if (float64(currentTempLH) / float64(targetTemp)) < 0.8 {
				temp = currentTempLH + uint16(rampUpFastRate*10)
			} else {
				temp = currentTempLH + uint16(rampUpSlowRate*10)
			}

			d.setRegister("D", plc.MODBUS_EXTRACTION[d.DeckName]["D"][210], temp)

		}
		if currentTempRH < targetTemp+10 {
			if (float64(currentTempRH) / float64(targetTemp)) < 0.8 {
				temp = currentTempRH + uint16(rampUpFastRate*10)
			} else {
				temp = currentTempRH + uint16(rampUpSlowRate*10)
			}
			d.setRegister("D", plc.MODBUS_EXTRACTION[d.DeckName]["D"][224], temp)
		}
	}
}

// cool Down to room Temperature if heater isn't in progress, update every 2 sec
func (d *SimulatorDriver) coolDown() {
	heater1Cooled, heater2Cooled := false, false
	var temp uint16 = 0

	for {
		time.Sleep(delay * 40 * time.Millisecond)
		if d.isHeaterInProgress() || (heater1Cooled && heater2Cooled) {
			return
		}

		currentTempLH := d.readRegister("D", plc.MODBUS_EXTRACTION[d.DeckName]["D"][210])
		currentTempRH := d.readRegister("D", plc.MODBUS_EXTRACTION[d.DeckName]["D"][224])

		logger.Debugln("Cooling Down -> currentTempLH :", currentTempLH, ", currentTempRH", currentTempRH)
		// 1 degree extra
		if currentTempLH > uint16(roomTemp*10+10) {
			if (float64(currentTempLH) / float64(roomTemp)) < 0.8 {
				temp = currentTempLH - uint16(rampDownFastRate*20)
			} else {
				temp = currentTempLH - uint16(rampDownSlowRate*20)
			}
			d.setRegister("D", plc.MODBUS_EXTRACTION[d.DeckName]["D"][210], temp)
		} else {
			heater1Cooled = true
		}
		if currentTempRH > uint16(roomTemp*10+10) {
			if (float64(currentTempRH) / float64(roomTemp)) < 0.8 {
				temp = currentTempRH - uint16(rampDownFastRate*20)
			} else {
				temp = currentTempRH - uint16(rampDownSlowRate*20)
			}
			d.setRegister("D", plc.MODBUS_EXTRACTION[d.DeckName]["D"][224], temp)
		} else {
			heater2Cooled = true
		}
	}
}
