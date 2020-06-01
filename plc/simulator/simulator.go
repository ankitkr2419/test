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

func (d *Simulator) ConfigureRun(plc.Stage) error {
	return nil
}

func (d *Simulator) Start() (err error) {
	// TBD
	return
}

func (d *Simulator) Stop() (err error) {
	// TBD
	return
}

// Monitor periodically. If Status=CYCLE_COMPLETE, the Scan will be populated
func (d *Simulator) Monitor(cycle uint16) (scan plc.Scan, err error) {
	return
}

func (d *Simulator) SelfTest() (status plc.Status) {
	// TBD
	return
}

func (d *Simulator) Calibrate() (err error) {
	// TBD
	return
}
