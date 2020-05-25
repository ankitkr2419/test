package simulator

import "mylab/cpagent/plc"

type Simulator struct {
}

func NewSimulator() plc.Driver {
	return &Simulator{}
}

func (d *Simulator) HeartBeat() error {
	return nil
}

func (d *Simulator) PreRun(plc.Stage) error {
	return nil
}

// Monitor periodically. If Status=CYCLE_COMPLETE, the Scan will be populated
func (d *Simulator) Monitor() (scan plc.Scan, status plc.Status) {
	return
}
