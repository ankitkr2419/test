package simulator

import (
	"crypto/rand"
	"errors"
	"math"
	"math/big"
	"mylab/cpagent/config"
	"mylab/cpagent/db"
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
	logger.Info("Starting cycleStage: ")
	plc.HeatingCycleComplete = false
	d.plcIO.d.currentCycle = 0
	d.plcIO.m.emissionFlag = 0
	logger.Println("cycle count", d.config.CycleCount)

	for i := uint16(0); i < d.config.CycleCount; i++ { //for each cycle
		// Check for Stop signal
		logger.Println("perform cycle", i)
		if d.plcIO.m.startStopCycle == 0 {
			d.ErrCh <- errors.New("recieved stop signal")
			break
		}

		d.plcIO.m.cycleCompleted = 0
		d.plcIO.d.currentCycle++
		d.performSteps(d.config.Cycle)
		logger.Println("emission flag", d.plcIO.m.emissionFlag)

		if d.plcIO.m.emissionFlag == 1 { // Means PC did not set it to 0
			d.ErrCh <- errors.New("client not reading the emission data, stopping PCR")
			break // stop cycle as client is not reading the data
		}
		logger.Println("perform cycle done", i)

		// populate emmission data 96X6
		d.emit()

		d.plcIO.m.cycleCompleted = 1 // cycle completed
		plc.HeatingCycleComplete = true
		d.plcIO.m.emissionFlag = 1 // PLC writing done

		// takes 1 to 3 seconds for cooling down
		time.Sleep(time.Duration(jitter(1, 1, 3)) * time.Second)
	}

	d.ExitCh <- errors.New("stop")
}

func (d *Simulator) performSteps(steps []plc.Step) {
	for _, stp := range steps { //for each steps
		// ramping up temp
		for {
			if d.plcIO.m.startStopCycle == 0 {
				d.ErrCh <- errors.New("recieved stop signal")
				break
			}

			// taking some time to increase the temperature
			//time.Sleep(200 * time.Millisecond)
			time.Sleep(time.Duration(jitter(0, 1, 3)) * time.Second) // sleep for 1 to 3 seconds

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
		logger.WithField("emission", emission).Info("EMISSIONS:")
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

func (d *Simulator) Cycle() (err error) {

	if plc.ExperimentRunning {
		logger.WithField("CYCLE RTPCR", "LED SWITCHED ON").Infoln("cycle started")
		err = plc.HoldSleep(int32(config.GetCycleTime()))
		if err != nil {
			logger.Errorln("Error while running cycle: ", err)
			return
		}
		plc.DataCapture = true
	}
	return
}
func (d *Simulator) HomingRTPCR() (err error) {
	logger.WithField("HOMING", "Started").Infoln("homing started")
	time.Sleep(time.Second * time.Duration(config.GetHomingTime()))
	logger.WithField("HOMING", "Completed").Infoln("homing completed")
	return
}
func (d *Simulator) Reset() (err error) { return }

func (d *Simulator) SetLidTemp(expectedLidTemp uint16) (err error) {
	logger.WithField("LID TEMP", "LID TEMP started").Infoln("LID TEMP STARTED")

	// simulate currentLidTemp
	if plc.ExperimentRunning {
		time.Sleep(2 * time.Second)
		plc.CurrentLidTemp = float32(expectedLidTemp) / 10
		logger.Infoln("Current Lid Temp: ", plc.CurrentLidTemp)
	}
	return
}

func (d *Simulator) SwitchOffLidTemp() (err error) {
	// Off Lid Heating
	plc.CurrentLidTemp = float32(config.GetRoomTemp())
	logger.WithField("LID TEMP OFF", "LID TEMP SWITCHED OFF").Infoln("LID TEMP SWITCHED OFF")

	return
}

func (d *Simulator) SetScanSpeedAndScanTime() (err error) {
	// TBD
	return
}

func (d *Simulator) CalculateOpticalResult(dye db.Dye, kitID string, knownValue, cycleCount int64) (opticalResult []db.DyeWellTolerance, err error) {

	wellsData := make(map[int][]uint16, cycleCount)

	plc.ExperimentRunning = true
	for i := 0; i < int(cycleCount); i++ {
		d.Cycle()
		d.emit()
		logger.Println("emissions------------------>", d.emissions)
		wellsData[i] = make([]uint16, 16)
		for j := 0; j < 16; j++ {
			for k, data := range d.emissions {
				if k == j {
					logger.Infoln(data[dye.Position-1])
					wellsData[i][j] = data[dye.Position-1]
				}
			}
		}
	}

	for j := 0; j < 16; j++ {
		var finalValue uint16
		var deviatedResult db.DyeWellTolerance
		for i := 0; i < int(cycleCount); i++ {

			finalValue += wellsData[i][j]
		}
		finalAvg := int64(finalValue) / cycleCount
		deviatedValue := float64(knownValue) - float64(finalAvg)
		deviatedResult.OpticalResult = math.Abs(deviatedValue / float64(knownValue) * float64(100))
		deviatedResult.DyeID = dye.ID
		deviatedResult.KitID = kitID
		deviatedResult.WellNo = j
		if deviatedResult.OpticalResult > dye.Tolerance {
			deviatedResult.Valid = false
		}

		opticalResult = append(opticalResult, deviatedResult)
	}
	return
}
