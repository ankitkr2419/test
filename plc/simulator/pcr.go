package simulator

import (
	"crypto/rand"
	"errors"
	"math"
	"math/big"
	"mylab/cpagent/plc"
	"time"

	logger "github.com/sirupsen/logrus"
)

const (
	roomTemp float32 = 30.0 // assume room temp is 30.0
)

// pcrHeartBeat sets D100 to 1 frequently
func (d *Simulator) pcrHeartBeat() {
	go func() {
		for {
			time.Sleep(5000 * time.Microsecond)

			d.plcIO.d.heartbeat = 1
		}
	}()
}

func (d *Simulator) holdingStage() {
	rt := roomTemp
	d.plcIO.d.currentTemp = uint16(rt * 10)
	d.performSteps(d.config.Holding)

}

func (d *Simulator) cycleStage() {
	logger.Info("Starting cycleStage")
	d.plcIO.d.currentCycle = 0

	for i := uint16(0); i < d.config.CycleCount; i++ { //for each cycle
		// Check for Stop signal
		if d.plcIO.m.startStopCycle == 0 {
			d.ErrCh <- errors.New("recieved stop signal")
			break
		}

		d.plcIO.m.cycleCompleted = 0
		d.plcIO.d.currentCycle++
		d.performSteps(d.config.Cycle)

		if d.plcIO.m.emissionFlag == 1 { // Means PC did not set it to 0
			d.errCh <- errors.New("client not reading the emission data, stopping PCR")
			break // stop cycle as client is not reading the data
		}

		// populate emmission data 96X6
		d.emit()

		d.plcIO.m.cycleCompleted = 1 // cycle completed
		d.plcIO.m.emissionFlag = 1   // PLC writing done

		// takes 1 to 3 seconds for cooling down
		time.Sleep(time.Duration(jitter(1, 1, 3)) * time.Second)
	}
	d.ExitCh <- "stop"
}

func (d *Simulator) performSteps(steps []plc.Step) {
	for _, stp := range steps { //for each steps
		// ramping up temp
		for {
			if d.plcIO.m.startStopCycle == 0 {
				d.ErrCh <- errors.New("recieved stop signal")
				return
			}

			// taking some time to increase the temperature
			//time.Sleep(200 * time.Millisecond)
			time.Sleep(time.Duration(jitter(0, 1, 3)) * time.Second) // sleep for 1 to 3 seconds

			// simulate currentLidTemp
			d.plcIO.d.currentLidTemp = jitter(uint16(d.config.IdealLidTemp*10), 0, 5)

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
}

func (d *Simulator) emit() {
	emissions := []plc.Emissions{}

	// iterating all wells
	for _, well := range d.wells {

		// setting all 6 emissions for each well
		emission := plc.Emissions{}
		for i := range emission {
			emission[i] = calculate(d.plcIO.d.currentCycle, well.goals[i])
		}
		logger.WithField("emission", emission).Debug("EMISSIONS:")
		emissions = append(emissions, emission)
	}
	d.emissions = emissions
}

// this gives absolute value of emission for each dye, not difference (delta)
func calculate(n uint16, s string) uint16 {
	//for no template control
	if s == "0" {
		return 0
	}

	// for initial cycle, eratic values between 1000 to 2000
	if n <= 10 {
		return jitter(0, 1000, 3000)
	}

	// for high viral load after 20 cycles, pleatus
	if n >= 30 && s == "high" {
		return jitter(29000, 1500, 3000)
	}

	// for high viral load during 10-30 cycles, exponential growth approximately between 2000 to 30000
	if s == "high" {
		return jitter(uint16(2000*(math.Pow(float64(n)/float64(10), 2.5))), 1000, 2000)
	}

	// for low viral load after 20 cycles, pleatus
	if n >= 46 && s == "low" {
		return jitter(29000, 1000, 3000)
	}

	// for low viral load till 24 cycles, do not increase exponentially
	if n <= 24 && s == "low" {
		return jitter(0, 1000, 3000) // not too much eratic either...
	}

	// for low viral load during 25-45 cycles, exponential growth approximately between 2000 to 30000
	if s == "low" {
		return jitter(uint16(2000*(math.Pow(float64(n)/float64(25), 4.7))), 1000, 2000)
	}

	// for negative goal, will be contant between 2500-2550
	if s == "" {
		return jitter(2500, 0, 50)
	}

	return 0
}

func jitter(n uint16, min, max int) uint16 {
	nBig := big.NewInt(int64(0))
	for int(nBig.Int64()) < min {
		nBig, _ = rand.Int(rand.Reader, big.NewInt(int64(max)))
	}

	final := n + uint16(int(nBig.Int64()))

	logger.WithFields(logger.Fields{
		"n":     n,
		"min":   min,
		"max":   max,
		"final": final,
	}).Debug("inside jitter...")
	return final
}
