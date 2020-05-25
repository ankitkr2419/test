package plc

type Status int32

const (
	OK             Status = iota // all good
	ERROR                        // something is wrong
	RESTART                      // cycle was interrupted, can restart
	ABORT                        // abort a running cycle
	CYCLE_RUNNING                // cycle is running
	CYCLE_COMPLETE               // cycle is complete
)

type Step struct {
	HoldingTemp float32 // holding temperature for step
	RampUpTemp  float32 // ramp-up temperature for step
	HoldTime    int32   // hold time for step
}

// We can have at most 4 Holding steps and 6 Cycling steps.
type Stage struct {
	Holding    [4]Step
	Cycle      [6]Step
	CycleCount int32 // number of cycles to run
}

type Emissions [6]int32

type Scan struct {
	Cycle   int32 // current running cycle
	Wells   [96]Emissions
	Temp    float32
	LidTemp float32
}

type Driver interface {
	HeartBeat() error        // Attempt to write heartbeat 3 times else fail
	PreRun(Stage) error      // Configure the various holding and cycling stages
	Monitor() (Scan, Status) // Monitor periodically. If Status=CYCLE_COMPLETE, the Scan will be populated
}

type RealDriver interface {
	SelfTest() Status        // Check if Homing or any other errors during bootup of PLC
	HeartBeat() error        // Attempt to write heartbeat 3 times else fail
	PreRun(Stage) error      // Configure the various holding and cycling stages
	Start() error            // trigger the start of the cycling process
	Monitor() (Scan, Status) // Monitor periodically. If Status=CYCLE_COMPLETE, the Scan will be populated
	Stop(Status) error       // Stop the cycle, Status: ABORT (if pre-emptive) OK: All Cycles have completed
	Calibrate() error        // TBD
}
