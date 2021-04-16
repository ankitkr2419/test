package simulator

import (
	"fmt"
	"mylab/cpagent/plc"
	"sync"
)

var motorDone, sensorDone, motorInProgress sync.Map

func LoadUtils() {
	motorDone.Store("A", false)
	motorDone.Store("B", false)
	sensorDone.Store("A", false)
	sensorDone.Store("B", false)
	motorInProgress.Store("A", false)
	motorInProgress.Store("B", false)
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
