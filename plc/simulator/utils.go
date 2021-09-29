package simulator

import (
	"fmt"
	"mylab/cpagent/plc"
	"sync"
)

const (
	pidTuningTime       = 10
	pidProgressResponse = 0x0003
	pidDoneResponse     = 0x0004
)

var motorDone, sensorDone, motorInProgress, heaterInProgress, shakerPIDCalibrationInProgress sync.Map

func loadUtils() {
	motorDone.Store(plc.DeckA, false)
	motorDone.Store(plc.DeckB, false)
	sensorDone.Store(plc.DeckA, false)
	sensorDone.Store(plc.DeckB, false)
	motorInProgress.Store(plc.DeckA, false)
	motorInProgress.Store(plc.DeckB, false)
	heaterInProgress.Store(plc.DeckA, false)
	heaterInProgress.Store(plc.DeckB, false)
	shakerPIDCalibrationInProgress.Store(plc.DeckA, false)
	shakerPIDCalibrationInProgress.Store(plc.DeckB, false)
}

func (d *SimulatorDriver) checkForValidAddress(registerType string, address uint16) (err error) {
	switch registerType {
	case "M":
		// valid range 0-8
		lowestMAddress := plc.MODBUS_EXTRACTION[d.DeckName][registerType][0]
		highestMAddress := plc.MODBUS_EXTRACTION[d.DeckName][registerType][47]

		if address >= lowestMAddress && address <= highestMAddress {
			return
		}

	case "D":
		// valid range 200-226
		lowestDAddress := plc.MODBUS_EXTRACTION[d.DeckName][registerType][200]
		highestDAddress := plc.MODBUS_EXTRACTION[d.DeckName][registerType][534]

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
