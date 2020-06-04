package simulator

import "mylab/cpagent/plc"

type Simulator struct {
	ExitCh chan error
}

func NewSimulator(exit chan error) plc.Driver {
	go startPCR() // starts PCR Machine

	return &Simulator{exit}
}

func (d *Simulator) HeartBeat() {
	// TBD
	return
}

func (d *Simulator) ConfigureRun(s plc.Stage) error {
	// setting PLC resistors with stage data
	plcIO.d.stage = s
	return nil
}

func (d *Simulator) Start() (err error) {
	// start cycle
	plcIO.m.startStopCycle = 1
	return
}

func (d *Simulator) Stop() (err error) {
	// start cycle
	plcIO.m.startStopCycle = 0
	return
}

// Monitor periodically. If Status=CYCLE_COMPLETE, the Scan will be populated
func (d *Simulator) Monitor(cycle uint16) (scan plc.Scan, err error) {
	// Read current cycle
	scan.Cycle = plcIO.d.currentCycle

	// Read cycle temperature.. PLC returns 653 for 65.3 degrees
	scan.Temp = float32(plcIO.d.currentTemp) / 10

	// Read lid temperature
	scan.LidTemp = float32(plcIO.d.currentLidTemp) / 10

	/* Below Code is for cycle stage (TBD) */
	// Read current cycle status
	tmp := plcIO.m.cycleCompleted

	if tmp == 0 { // 0x0000 means cycle is not complete
		// Values would not have changed.
		scan.CycleComplete = false
		return
	}
	scan.CycleComplete = true

	// If the invoker has already read this cycle data, don't send it again!
	if cycle == scan.Cycle {
		return
	}
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
