package simulator

import (
	logger "github.com/sirupsen/logrus"
	"sync"
)

var simulatorLock sync.Mutex

func (d *SimulatorDriver) setMotorInProgress() {
	motorInProgress.Store(d.DeckName, true)
}

func (d *SimulatorDriver) resetMotorInProgress() {
	motorInProgress.Store(d.DeckName, false)
}

func (d *SimulatorDriver) setHeaterInProgress() {
	heaterInProgress.Store(d.DeckName, true)
}

func (d *SimulatorDriver) resetHeaterInProgress() {
	heaterInProgress.Store(d.DeckName, false)
}

func (d *SimulatorDriver) setMotorDone() {
	motorDone.Store(d.DeckName, true)
}

func (d *SimulatorDriver) resetMotorDone() {
	motorDone.Store(d.DeckName, false)
}

func (d *SimulatorDriver) setSensorDone() {
	sensorDone.Store(d.DeckName, true)
}

func (d *SimulatorDriver) resetSensorDone() {
	sensorDone.Store(d.DeckName, false)
}

func (d *SimulatorDriver) isMotorInProgress() bool {
	if temp, ok := motorInProgress.Load(d.DeckName); !ok {
		logger.Errorln("motorInProgress isn't loaded!")
	} else {
		return temp.(bool)
	}
	return false
}

func (d *SimulatorDriver) isHeaterInProgress() bool {
	if temp, ok := heaterInProgress.Load(d.DeckName); !ok {
		logger.Errorln("heaterInProgress isn't loaded!")
	} else {
		return temp.(bool)
	}
	return false
}

func (d *SimulatorDriver) isMotorDone() bool {
	if temp, ok := motorDone.Load(d.DeckName); !ok {
		logger.Errorln("motorDone isn't loaded!")
	} else {
		return temp.(bool)
	}
	return false
}

func (d *SimulatorDriver) isSensorDone() bool {
	if temp, ok := sensorDone.Load(d.DeckName); !ok {
		logger.Errorln("sensorDone isn't loaded!")
	} else {
		return temp.(bool)
	}
	return false
}

// Always take Lock on REGSITERS_EXTRACTION while writing/reading
func (d *SimulatorDriver) setRegister(regType string, address, value uint16) {
	simulatorLock.Lock()
	defer simulatorLock.Unlock()
	REGISTERS_EXTRACTION[d.DeckName][regType][address] = value
}

func (d *SimulatorDriver) readRegister(regType string, address uint16) (value uint16) {
	simulatorLock.Lock()
	defer simulatorLock.Unlock()
	return REGISTERS_EXTRACTION[d.DeckName][regType][address]
}
