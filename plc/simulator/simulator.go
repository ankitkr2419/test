package simulator

import (
	"mylab/cpagent/plc"

	logger "github.com/sirupsen/logrus"
)

type Simulator struct {
	ExitCh chan error
}

func NewSimulator(exit chan error) plc.Driver {
	return &Simulator{exit}
}

func (d *Simulator) HeartBeat() {
	logger.Info("Starting HeartBeat...")
	return
}

func (d *Simulator) ConfigureRun(plc.Stage) error {
	return nil
}

// Monitor periodically. If Status=CYCLE_COMPLETE, the Scan will be populated
func (d *Simulator) Monitor(cycle uint16) (scan plc.Scan, err error) {
	scan.Cycle = cycle + 1 // for testing, can be changed as required

	scan.Temp = 150 // for testing, can be changed to 0 for setting CycleComplete to false

	scan.LidTemp = 150

	// Read current cycle status
	if scan.Temp == 0 {
		scan.CycleComplete = false
		return
	}
	scan.CycleComplete = true

	// If the invoker has already read this cycle data, don't send it again!
	if cycle == scan.Cycle {
		return
	}

	// Scan all the data from the Wells (96 x 6).
	for i := 0; i < 6; i++ {
		for j := 0; j < 16; j++ {
			emission := plc.Emissions{}

			for k := 0; k < 6; k++ {
				emission[k] = uint16(256 + k) //256 257 258 259 260 261
			}

			scan.Wells[(i*16)+j] = emission
		}
	}

	return
}
