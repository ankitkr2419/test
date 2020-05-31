package simulator

import (
	"math/rand"
	"mylab/cpagent/plc"
	"time"

	logger "github.com/sirupsen/logrus"
)

type emissionCases struct {
	initial []uint16 // initial emission values to start from (for 6 target)
	pass    []uint16 // values should not increase beyond threshold to pass per cycle
	fail    []uint16 // values must increase beyond threshold to fail per cycle
	test    []uint16 // any other combination if you want per cycle
}

var emissionCase = emissionCases{
	[]uint16{1100, 1100, 1100, 1100, 1100, 1100}, // initial
	[]uint16{0, 0, 0, 0, 0, 0},                   // pass
	[]uint16{10, 10, 10, 10, 10, 10},             // fail
	[]uint16{10, 0, 5, 5, 0, 10},                 // random
}

const (
	roomTemp    float32 = 30.0 // assume room temp is 30.0
	wellsCount  uint16  = 96   // number of wells to simulate
	jitterValue int     = 50   // jitter to fluctuate emission value
)

var ch chan plc.Scan

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

func (d *Simulator) ConfigureRun(stg plc.Stage) error {
	ch = make(chan plc.Scan)
	go simulate(stg, ch) // TODO: Discuss about it
	return nil
}

// Monitor periodically. If Status=CYCLE_COMPLETE, the Scan will be populated
func (d *Simulator) Monitor(cycle uint16) (scan plc.Scan, err error) {
	scan = <-ch

	if !scan.CycleComplete {
		return
	}

	// If the invoker has already read this cycle data, don't send it again!
	if cycle == scan.Cycle {
		return
	}

	return
}

func simulate(stg plc.Stage, ch chan plc.Scan) {
	// simulatiing holding stage
	simulateHoldingStage(stg, ch)

	//simulate cycle stage
	simulateCycleStage(stg, ch)
}

func simulateHoldingStage(stg plc.Stage, ch chan plc.Scan) {
	scan := plc.Scan{}
	r := roomTemp

	for _, stp := range stg.Holding {
		for { // ramp up
			time.Sleep(time.Duration(stp.HoldTime) * time.Second) // spending time - HoldTime
			scan.Temp = r                                         // simulate cycle temp
			scan.LidTemp = r + 2                                  // lid temp is always a bit more than temp, ideally 2

			if r >= stp.TargetTemp { // if the target temp is below than the next multiple of ramp up temp
				scan.Temp = stp.TargetTemp
				ch <- scan
				break
			}

			ch <- scan

			r = r + stp.RampUpTemp
		}
		r = roomTemp
	}
}

func simulateCycleStage(stg plc.Stage, ch chan plc.Scan) {
	scan := plc.Scan{}
	emissions := []plc.Emissions{}
	r := roomTemp

	for i := uint16(0); i < stg.CycleCount; i++ {
		scan.CycleComplete = false

		for s, stp := range stg.Cycle {
			for {
				scan.CycleComplete = false
				time.Sleep(time.Duration(stp.HoldTime) * time.Second) // spending time - HoldTime

				scan.Temp = r        // simulate cycle temp
				scan.LidTemp = r + 2 // lid temp is always a bit more than temp, ideally 2
				scan.Cycle = i + 1   // cycle is incrementing from 1

				if r >= stp.TargetTemp { // if the target temp is below than the next multiple of ramp up temp

					scan.Temp = stp.TargetTemp
					if s == len(stg.Cycle)-1 { // if last cycle
						scan.CycleComplete = true

						emissions = fillEmission(scan.Cycle, emissions) // populate emissions

						for x, vl := range emissions {
							scan.Wells[x] = vl
						}

					}
					ch <- scan

					break
				}
				ch <- scan

				r = r + stp.RampUpTemp
			}
			r = roomTemp
		}
		scan = plc.Scan{}
	}
}

func fillEmission(cycle uint16, ems []plc.Emissions) []plc.Emissions {

	emissions := []plc.Emissions{}
	emission := plc.Emissions{}

	for i := uint16(0); i < wellsCount; i++ {
		for x := range emission {
			emission[x] = jitter(emissionCase.initial[x])
		}

		if i < 31 { // first 32 wells are set for fail case
			for x := uint16(0); x < cycle-1; x++ {
				for i, v := range emission {
					emission[i] = v + jitter(emissionCase.fail[i])
				}
			}

			emissions = append(emissions, emission)
		}

		if i > 31 && i < 63 { // next 32 wells are set for pass case
			for x := uint16(0); x < cycle-1; x++ {
				for i, v := range emission {
					emission[i] = v + jitter(emissionCase.pass[i])
				}
			}

			emissions = append(emissions, emission)
		}

		if i > 63 && i < 95 { // next 32 wells are set for user-defined testing case
			for x := uint16(0); x < cycle-1; x++ {
				for i, v := range emission {
					emission[i] = v + jitter(emissionCase.test[i])
				}
			}

			emissions = append(emissions, emission)
		}
	}

	return emissions
}

func jitter(n uint16) uint16 {
	return n + uint16(rand.Intn(jitterValue))
}
