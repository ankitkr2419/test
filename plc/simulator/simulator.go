package simulator

import (
	"errors"
	"mylab/cpagent/config"
	"mylab/cpagent/plc"
	"sync"
	"time"

	logger "github.com/sirupsen/logrus"
)

type Well struct {
	// emissions plc.Emissions // dye emmissions.
	control string    // "", positive, negative, internal or no_template
	goals   [6]string // "", "high", "low"
}

type Simulator struct {
	sync.RWMutex
	plcIO     plcRegistors
	config    plc.Stage
	emissions []plc.Emissions
	ExitCh    chan string
	ErrCh     chan error
	wells     []Well
}

func NewSimulator(exit chan error) plc.Driver {
	ex := make(chan string)

	s := Simulator{}
	s.ExitCh = ex
	s.ErrCh = exit
	s.pcrHeartBeat()

	go s.HeartBeat()

	return &s
}

func (d *Simulator) HeartBeat() {
	go func() {
		logger.Info("Starting Simulator HeartBeat...")

		var err error

	LOOP:
		for {
			time.Sleep(2000 * time.Millisecond) // sleep it off for a bit

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
		d.ExitCh <- "dead"
		return
	}()
}

func (d *Simulator) ConfigureRun(s plc.Stage) error {
	// NOTE: commented to run new exp when stop
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
	// Abort running process

	if d.plcIO.m.startStopCycle == 0 {
		err = errors.New("Cannot stop, not yet started")
		return
	}

	d.plcIO.m.startStopCycle = 0

	d.ExitCh <- "abort"
	return
}

func (d *Simulator) simulate() {
	//set wells with goals
	d.setWells()

	//start holding stage
	d.holdingStage()

	// Start Cycling stage in a different go-routine and listen for events on exitCh and errCh
	go d.cycleStage()

	for {
		// Intentionally don't have a default, so that it blocks on either one of the channels.
		select {
		case msg := <-d.ExitCh:
			logger.WithField("msg", msg).Info("simulate: ExitCh received data")
			if msg == "stop" {
				d.ErrCh <- errors.New("PCR Stopped")

				// reset to start new experiment
				d.config = plc.Stage{}
				d.plcIO = plcRegistors{}
				d.wells = []Well{}

			}
			if msg == "abort" {

				d.ErrCh <- errors.New("PCR Aborted")

				// reset to start new experiment
				d.config = plc.Stage{}
				d.plcIO = plcRegistors{}
				d.wells = []Well{}

			}
			if msg == "pause" {
				//TBD
			}
			if msg == "dead" {

				// heart beat failes, pcr is not responding
				d.ErrCh <- errors.New("PCR Dead")

			}
			/* This ErrCh will never be used between simulator and PCR
			case err := <-d.ErrCh:
				// Some error flagged
				logger.WithField("err", err.Error()).Error("simulate: errCh recevied data")
				return
			*/
		}
	}
}

func (d *Simulator) setWells() {
	// Reading simulator configurations
	//env := config.ReadEnvInt("environment") - TBD

	// wells count
	wc := config.ReadEnvInt("wells.count")

	// controls
	pc := config.ReadEnvInt("controls.positive")
	nc := config.ReadEnvInt("controls.negative")
	ic := config.GetICPosition() - 1 // ic to be added in all the wells //-1 as positions start from 1
	ntc := config.ReadEnvInt("controls.no_template")

	/* TBD, the logic needs to be optimised, too many conditions,
	also by this logic pcr wont run only with control wells - need to fix */
	for i := 1; i <= wc; {
		well := Well{}

		if i == pc {
			well.control = "positive"
			for i := 0; i < 6; i++ {
				well.goals[i] = "high"
			}

		} else if i == nc {
			well.control = "negative"
			for i := 0; i < 6; i++ {
				well.goals[i] = ""
			}

		} else if i == ic {
			well.control = "internal"
			well.goals = [6]string{"", "", "", "", "", "high"} //TODO: discuss

		} else if i == ntc {
			well.control = "no_template"
			for i := 0; i < 6; i++ {
				well.goals[i] = "0"
			}

		} else {
			well.control = "" // patient sample

			for i := 0; i < 6; i++ {
				if i != ic { // for all targets accept ic assign random goals
					switch goal := jitter(0, 1, 4); goal { // randomization of goals
					case 1:
						well.goals[i] = "" // negative
					case 2:
						well.goals[i] = "high"
					case 3:
						well.goals[i] = "low"
					}
				} else {
					well.goals[i] = "high" //	set internal control "high"
				}
			}
		}
		d.wells = append(d.wells, well)
		i++
	}
}

// Monitor periodically. If Status=CYCLE_COMPLETE, the Scan will be populated
func (d *Simulator) Monitor(cycle uint16) (scan plc.Scan, err error) {
	d.Lock()
	defer d.Unlock()

	// Read current cycle
	scan.Cycle = d.plcIO.d.currentCycle

	//do not need this as handled in the software

	// Read cycle temperature.. PLC returns 653 for 65.3 degrees
	//scan.Temp = float32(d.plcIO.d.currentTemp) / 10

	// Read lid temperature
	//scan.LidTemp = float32(d.plcIO.d.currentLidTemp) / 10

	// Read current cycle status
	// if d.plcIO.m.cycleCompleted == 0 { // 0x0000 means cycle is not complete
	// 	// Values would not have changed.
	// 	scan.CycleComplete = false
	// 	return
	// }
	// scan.CycleComplete = true

	// If the invoker has already read this cycle data, don't send it again!
	// if cycle == scan.Cycle {
	// 	return
	// }

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
