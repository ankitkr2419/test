package simulator

import (
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
