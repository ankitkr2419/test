package simulator

import (
	"errors"
	"math/rand"
	"mylab/cpagent/plc"
	"sync"
	"time"
)

type emissionCases struct {
	initial  []uint16 // initial emission values to start from (for 6 target)
	negative []uint16 // for negative samples, value should increase minimally or be same value for consecutive cycles
	positive []uint16 // for negative samples, increase over consecutive cycles
	testing  []uint16 // any other combination if you want per cycle like positive case with low load
}

// settings for different emission values according to targets
var emissionCase = emissionCases{
	[]uint16{1100, 1100, 1100, 1100, 1100, 1100}, // initial
	[]uint16{0, 0, 0, 0, 0, 0},                   // negative
	[]uint16{10, 10, 10, 10, 10, 10},             // positive with high load
	[]uint16{10, 0, 5, 5, 0, 10},                 // positive with low load
}

const (
	roomTemp   float32 = 30.0 // assume room temp is 30.0
	wellsCount uint16  = 96   // number of wells to simulate
)

// pcrHeartBeat sets D100 to 1 frequently
func (d *Simulator) pcrHeartBeat() {
	go func() {
		for {
			time.Sleep(200 * time.Microsecond)

			d.plcIO.d.heartbeat = 1
		}
	}()
}

func (d *Simulator) holdingStage(wg *sync.WaitGroup) {
	defer wg.Done()

	rt := roomTemp
	d.plcIO.d.currentTemp = uint16(rt * 10)

	for _, stp := range d.config.Holding {
		// ramping up temp
		for {
			// taking some time to increase the temperature
			time.Sleep(250 * time.Millisecond)

			// simulate currentLidTemp
			d.plcIO.d.currentLidTemp = jitter(uint16(d.config.IdealLidTemp*10), 0, 105)

			// simulate currentTemp
			d.plcIO.d.currentTemp = d.plcIO.d.currentTemp + uint16(stp.RampUpTemp*10)

			// if the target temp is below than the next multiple of ramp up temp
			if d.plcIO.d.currentTemp >= uint16(stp.TargetTemp*10) {
				d.plcIO.d.currentTemp = uint16(stp.TargetTemp * 10)

				// spending time - HoldTime
				time.Sleep(time.Duration(stp.HoldTime) * time.Second)
				break
			}
		}
	}
}

func (d *Simulator) cycleStage() {
	initialTemp := 45

	d.plcIO.d.currentCycle = 0
	d.plcIO.d.currentTemp = uint16(initialTemp * 10)

	for i := uint16(0); i < d.config.CycleCount; i++ { //for each cycle
		d.plcIO.m.cycleCompleted = 0
		d.plcIO.d.currentCycle++
		for _, stp := range d.config.Cycle { //for each steps
			// ramping up temp
			for {
				if d.plcIO.m.startStopCycle == 0 {
					d.ExitCh <- errors.New("recieved stop signal")
					return
				}

				// taking some time to increase the temperature
				time.Sleep(250 * time.Millisecond)

				// simulate currentLidTemp
				d.plcIO.d.currentLidTemp = jitter(uint16(d.config.IdealLidTemp*10), 0, 105)

				// simulate currentTemp
				d.plcIO.d.currentTemp = d.plcIO.d.currentTemp + uint16(stp.RampUpTemp*10)

				// if the target temp is below than the next multiple of ramp up temp
				if d.plcIO.d.currentTemp >= uint16(stp.TargetTemp*10) {
					d.plcIO.d.currentTemp = uint16(stp.TargetTemp * 10)

					// holding at target temp for some specific time (holdtime)
					time.Sleep(time.Duration(stp.HoldTime) * time.Second)
					break
				}
			}
		}

		if d.plcIO.m.emissionFlag == 1 { // Means PC did not set it to 0
			d.ExitCh <- errors.New("client not reading the emission data, stopping PCR")
			d.exitCh <- "stop" // stop cycle as client is not reading the data
		}

		// populate emmission data 96X6
		d.emit()

		d.plcIO.m.cycleCompleted = 1 // cycle completed
		d.plcIO.m.emissionFlag = 1   // PLC writing done

		// takes 1 to 3 seconds for cooling down
		time.Sleep(time.Duration(jitter(1, 1, 3)) * time.Second)

	}
	d.exitCh <- "stop"
}

func (d *Simulator) emit() {
	emission := plc.Emissions{}
	emissions := []plc.Emissions{}

	// Setting initial values
	for i := uint16(0); i < wellsCount; i++ {
		for x := range emission {
			emission[x] = jitter(emissionCase.initial[x], 1, 2)
		}

		// first 32 wells are set for positive case
		if i < 31 {
			for x := uint16(0); x < d.plcIO.d.currentCycle; x++ {
				for i, v := range emission {
					emission[i] = v + jitter(emissionCase.positive[i], 1, 10)
				}
			}
			emissions = append(emissions, emission)
		}

		// next 32 wells are set for negative case
		if i > 31 && i < 63 {
			for x := uint16(0); x < d.plcIO.d.currentCycle-1; x++ {
				for i, v := range emission {
					emission[i] = v + jitter(emissionCase.negative[i], 1, 10)
				}
			}

			emissions = append(emissions, emission)
		}

		// last 32 wells are set for user-defined testing case
		if i > 63 && i < 95 {
			for x := uint16(0); x < d.plcIO.d.currentCycle-1; x++ {
				for i, v := range emission {
					emission[i] = v + jitter(emissionCase.testing[i], 1, 10)
				}
			}

			emissions = append(emissions, emission)
		}
	}
	d.emissions = emissions
}

func jitter(n uint16, min, max int) uint16 {
	return n + uint16(rand.Intn((max-min))+min)
}
