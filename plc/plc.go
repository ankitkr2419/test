package plc

import (
	"mylab/cpagent/db"
)

type Status int32

const (
	OK      Status = iota // all good
	ERROR                 // something is wrong
	RESTART               // cycle was interrupted, can restart
)

type Step struct {
	TargetTemp float32 // holding temperature for step
	RampUpTemp float32 // ramp-up temperature for step
	HoldTime   int32   // hold time for step
}

// We can have at most 4 Holding steps and 6 Cycling steps.
type Stage struct {
	Holding      []Step // 4 steps
	Cycle        []Step // 6 steps
	CycleCount   uint16 // number of cycles to run
	IdealLidTemp uint16 // ideal lid temp
}

type Emissions [6]uint16

type Scan struct {
	Cycle         uint16 // current running cycle
	Wells         [96]Emissions
	Temp          float32
	LidTemp       float32
	CycleComplete bool
}

type Driver interface {
	SelfTest() Status             // Check if Homing or any other errors during bootup of PLC {ERROR | RESTART | OK}
	HeartBeat()                   // Attempt to write heartbeat 3 times else fail
	ConfigureRun(Stage) error     // Configure the various holding and cycling stages
	Start() error                 // trigger the start of the cycling process
	Stop() error                  // Stop the cycle, Status: ABORT (if pre-emptive) OK: All Cycles have completed
	Monitor(uint16) (Scan, error) // Monitor periodically. If Status=CYCLE_COMPLETE, the Scan will be populated
	Calibrate() error             // TBD
}

type DeckDriver interface {
	NameOfDeck() string
	Homing() (string, error)
	DeckHoming() (string, error)
	SyringeHoming() (string, error)
	SyringeModuleHoming() (string, error)
	MagnetHoming() (string, error)
	MagnetUpDownHoming() (string, error)
	MagnetFwdRevHoming() (string, error)
	SwitchOffMotor() (string, error)
	ReadExecutedPulses() (string, error)
	SetupMotor(uint16, uint16, uint16, uint16, uint16) (string, error)
	ManualMovement(uint16, uint16, uint16) (string, error)
	ResetRunInProgress()
	Pause() (string, error)
	Resume() (string, error)
	Abort() (string, error)
	ResumeMotorWithPulses(uint16) (string, error)
	// RunProcesses() (string, error)
	// AspireDispense(category, cartridgeType string, labwareID, source, destination, aspire_cycles, dispense_cycles int64, asp_height, asp_mix_vol, asp_vol, dis_height, dis_mix_vol, dis_vol, dis_blow float64) (response string, err error)
	AspireDispenseBeta(aspireDispense db.AspireDispense, cartridgeID int64, tipType string) (response string, err error)
}
