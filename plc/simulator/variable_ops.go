package simulator

import (
	logger "github.com/sirupsen/logrus"
)

func (d *SimulatorDriver) setMotorInProgress() {
	motorInProgress.Store(d.DeckName, true)
}

func (d *SimulatorDriver) resetMotorInProgress() {
	motorInProgress.Store(d.DeckName, false)
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
