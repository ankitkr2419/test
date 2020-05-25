package simulator

import "mylab/cpagent/plc"

type Simulator struct {
	ExitCh chan error
}

func NewSimulator(exit chan error) plc.Driver {
	return &Simulator{exit}
}

func (d *Simulator) HeartBeat() {
	return
}

func (d *Simulator) PreRun(plc.Stage) error {
	return nil
}

// Monitor periodically. If Status=CYCLE_COMPLETE, the Scan will be populated
func (d *Simulator) Monitor() (scan plc.Scan, status plc.Status) {
	return
}
