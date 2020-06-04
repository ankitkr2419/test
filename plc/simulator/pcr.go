package simulator

import (
	"math/rand"
	"time"
)

const (
	roomTemp float32 = 30.0 // assume room temp is 30.0
)

var pcrStatus = "stopped"

func startPCR() {
	for {
		// start simulating cycles if D100 is set to 1
		if plcIO.m.startStopCycle == 1 && pcrStatus == "stopped" {
			go simulate()
			pcrStatus = "started"
		}

		// stop simulating cycles if D100 is set to 0
		if plcIO.m.startStopCycle == 0 && pcrStatus == "started" {
			// TBD
			// Stop
		}
	}
}

// pcrHeartBeat checks the D100 register and read/write on it accordingly
func pcrHeartBeat() {
	// TBD
}

func simulate() {
	// simulatiing holding stage
	holdingStage()

	//simulate cycle stage
	cycleStage()
}

func holdingStage() {
	rt := roomTemp

	for _, stp := range plcIO.d.stage.Holding {
		plcIO.d.currentTemp = uint16(rt * 10)

		// ramping up temp
		for {
			// taking some time to increase the temperature
			time.Sleep(250 * time.Millisecond)

			// simulate currentLidTemp
			plcIO.d.currentLidTemp = jitter(uint16(plcIO.d.stage.IdealLidTemp*10), 0, 105)

			// simulate currentTemp
			plcIO.d.currentTemp = plcIO.d.currentTemp + uint16(stp.RampUpTemp*10)

			// if the target temp is below than the next multiple of ramp up temp
			if plcIO.d.currentTemp >= uint16(stp.TargetTemp*10) {
				plcIO.d.currentTemp = uint16(stp.TargetTemp * 10)

				// spending time - HoldTime
				time.Sleep(time.Duration(stp.HoldTime) * time.Second)
				break
			}
		}
	}
	// set currentTemp back to room temp after holding cycle ends
	plcIO.d.currentTemp = uint16(rt * 10)
	// set currentLidTemp back to IdealLidTemp after holding cycle ends
	plcIO.d.currentLidTemp = plcIO.d.stage.IdealLidTemp * 10
}

func cycleStage() {
	// TBD
}

func jitter(n uint16, min, max int) uint16 {
	return n + uint16(rand.Intn((max-min))+min)
}
