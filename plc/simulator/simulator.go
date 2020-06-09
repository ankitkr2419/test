package simulator

import (
	"errors"
	"mylab/cpagent/plc"
	"sync"
	"time"

	logger "github.com/sirupsen/logrus"
)

type Simulator struct {
	sync.RWMutex
	plcIO     plcRegistors
	config    plc.Stage
	emissions []plc.Emissions
	exitCh    chan string
	ExitCh    chan error
}

func NewSimulator(exit chan error) plc.Driver {
	ex := make(chan string)

	s := Simulator{}
	s.exitCh = ex
	s.ExitCh = exit
	s.pcrHeartBeat()
	return &s
}

func (d *Simulator) HeartBeat() {
	go func() {
		logger.Info("Starting Simulator HeartBeat...")

		var err error

	LOOP:
		for {
			time.Sleep(500 * time.Millisecond) // sleep it off for a bit

			// 3 attempts to check for heartbeat of PLC and write ours!
			for i := 0; i < 3; i++ {
				if d.plcIO.d.heartbeat == 1 { // If beat is 1, PLC is alive, so write 2
					d.plcIO.d.heartbeat = 2
					continue LOOP
				}

				logger.WithFields(logger.Fields{
					"beat":    d.plcIO.d.heartbeat,
					"attempt": i + 1,
				}).Warn("Attempt failed. PLC heartbeat value has not changed. Retrying...")
				time.Sleep(200 * time.Millisecond) // sleep it off for a bit
			}

			// If we have reached here, PLC has not updated heartbeat for 3 tries, it's dead! Abort!
			logger.Warn("PLC heartbeat value is still 1 after 3 attepts. Abort!")
			err = errors.New("PLC is not responding and maybe dead. Abort")
			break
		}

		// something went wrong. Signal parent process
		logger.WithField("err", err.Error()).Error("Heartbeat Error. Abort!")
		d.ExitCh <- err
		return
	}()
}

func (d *Simulator) ConfigureRun(s plc.Stage) error {
	if d.config.CycleCount != 0 {
		return errors.New("PLC is already configured")
	}

	// setting config with stage data
	d.config = s

	return nil
}

func (d *Simulator) Start() (err error) {
	if d.config.CycleCount == 0 {
		err = errors.New("PLC is not configured yet")
		return
	}
	if d.plcIO.m.startStopCycle == 1 {
		err = errors.New("Cannot start again, already started")
		return
	}
	d.plcIO.m.startStopCycle = 1

	go d.simulate()

	return
}

func (d *Simulator) Stop() (err error) {
	go func() {
		if d.plcIO.m.startStopCycle == 0 {
			err = errors.New("Cannot stop, not yet started")
			return
		}

		d.plcIO.m.startStopCycle = 0
		d.exitCh <- "stop"
	}()

	return
}

func (d *Simulator) simulate() {
	var wg sync.WaitGroup
	wg.Add(1)

	//start holding stage
	d.holdingStage(&wg)

	//wait until holding stage is over
	wg.Wait()

	for {
		select {
		case msg := <-d.exitCh:
			if msg == "stop" {
				d.ExitCh <- errors.New("PCR Stopped")
				return
			}
			if msg == "abort" {
				//TBD
			}
			if msg == "pause" {
				//TBD
			}
		default:
			d.cycleStage()
		}
	}

}

// Monitor periodically. If Status=CYCLE_COMPLETE, the Scan will be populated
func (d *Simulator) Monitor(cycle uint16) (scan plc.Scan, err error) {
	d.Lock()
	defer d.Unlock()

	// Read current cycle
	scan.Cycle = d.plcIO.d.currentCycle

	// Read cycle temperature.. PLC returns 653 for 65.3 degrees
	scan.Temp = float32(d.plcIO.d.currentTemp) / 10

	// Read lid temperature
	scan.LidTemp = float32(d.plcIO.d.currentLidTemp) / 10

	// Read current cycle status
	if d.plcIO.m.cycleCompleted == 0 { // 0x0000 means cycle is not complete
		// Values would not have changed.
		scan.CycleComplete = false
		return
	}
	scan.CycleComplete = true

	// If the invoker has already read this cycle data, don't send it again!
	if cycle == scan.Cycle {
		return
	}

	// Scan all the data from the Wells (96 x 6)
	for i, data := range d.emissions {
		scan.Wells[i] = data
	}

	// PC reading done
	d.plcIO.m.emissionFlag = 0

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
