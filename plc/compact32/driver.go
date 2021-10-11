package compact32

import (
	"encoding/binary"
	"errors"
	"fmt"
	"math"
	"mylab/cpagent/config"
	"mylab/cpagent/db"
	"mylab/cpagent/plc"
	"mylab/cpagent/responses"

	"time"

	logger "github.com/sirupsen/logrus"
)

const (
	maxHomingTries     = 3
	homingSuccessValue = 37
	noOfDyes           = 4
	K4                 = 4
)

var homingCount int

// Interface Implementation Methods
func (d *Compact32) HeartBeat() {
	var value []byte
	var err error

	logger.Info("Starting HeartBeat...")
LOOP:
	for {
		time.Sleep(200 * time.Millisecond) // sleep it off for a bit

		// Read heartbeat status to check if PLC is alive!
		var beat uint16
		beat, err = d.Driver.ReadSingleRegister(plc.MODBUS["D"][100])
		if err != nil {
			logger.WithField("beat", beat).Error("ReadSingleRegister:D100 : Read PLC heartbeat")
			break
		}

		logger.Debugln("Heartbeat output: ", beat)

		// 3 attempts to check for heartbeat of PLC and write ours!
		for i := 0; i < 3; i++ {
			if beat == 1 { // If beat is 1, PLC is alive, so write 2
				value, err = d.Driver.WriteSingleRegister(plc.MODBUS["D"][100], uint16(2))
				if err != nil {
					logger.WithField("beat", beat).Error("WriteSingleRegister:D100 : Read PLC heartbeat")
					// exit!!
					break LOOP
				}
				logger.Debugln("Read value:", value)
				continue LOOP
			}

			logger.WithFields(logger.Fields{
				"beat":    beat,
				"attempt": i + 1,
			}).Warn("Attempt failed. PLC heartbeat value has not changed. Retrying...")
			time.Sleep(200 * time.Millisecond) // sleep it off for a bit
		}

		// If we have reached here, PLC has not updated heartbeat for 3 tries, it's dead! Abort!
		logger.Warn("PLC heartbeat value is still 1 after 3 attepts. Abort!")
		err = errors.New("PLC is not responding and maybe dead. Abort!!")
		break
	}

	// something went wrong. Signal parent process
	logger.WithField("err", err.Error()).Error("Heartbeat Error. Abort!")
	d.ExitCh <- errors.New("PCR Dead")
}

// Algorithm
// 1. Retry homing for max of maxHomingTries
// 2. Switch on M1, M2 for Homing
// 3. Switch Off cycle bit M36
// 4. Sleep for Homing Duration -- TODO: Change this to be based on bit
// 5. Switch Off M45

func (d *Compact32) HomingRTPCR() (err error) {

	homingCount++
	defer func() {
		homingCount = 0
	}()

	if homingCount == maxHomingTries {
		err = fmt.Errorf("homing failed even after %v tries", maxHomingTries)
		logger.WithField("HOMING", err.Error()).Errorln("homing failed")
		d.ExitCh <- errors.New("PCR Aborted")
		return
	}

	err = d.SwitchOffCycle()
	if err != nil {
		logger.Error("error in switching off cycle")
		return
	}

	//First Home
	err = d.Driver.WriteSingleCoil(plc.MODBUS["M"][1], plc.ON)
	if err != nil {
		logger.Error("WriteSingleCoil:M1 : Start Cycle")
		return
	}
	//First Home
	err = d.Driver.WriteSingleCoil(plc.MODBUS["M"][2], plc.ON)
	if err != nil {
		logger.Error("WriteSingleCoil:M2 : Start Cycle")
		return
	}

	err = d.Driver.WriteSingleCoil(plc.MODBUS["M"][36], plc.OFF)
	if err != nil {
		logger.Error("WriteSingleCoil:M36 : Start Cycle")
		return
	}

	logger.Infoln("HOMING Bit Reset M36")

	logger.WithField("HOMING", "homing started").Infoln("HOMING STARTED")
	time.Sleep(time.Second * time.Duration(config.GetHomingTime()))

	result, err := d.Driver.ReadCoils(plc.MODBUS["M"][36], uint16(1))
	if err != nil {
		logger.Error("ReadCoil:M36 ", err)
		return
	}
	logger.Infoln("homing result", result)
	if result[0] == homingSuccessValue {
		logger.WithField("HOMING", "Completed").Infoln("homing completed")
	} else {
		// Try Homing Again
		err = d.HomingRTPCR()
		if err != nil {
			return
		}
	}

	err = d.Driver.WriteSingleCoil(plc.MODBUS["M"][36], plc.OFF)
	if err != nil {
		logger.Error("WriteSingleCoil:M36 : Start Cycle")
		return
	}

	logger.Infoln("HOMING Bit Reset M36")

	err = d.Driver.WriteSingleCoil(plc.MODBUS["M"][45], plc.OFF)
	if err != nil {
		logger.Error("WriteSingleCoil:M45 : Start Cycle")
		return
	}

	logger.Infoln("HOMING Bit Reset M45")

	// Also Reset
	return d.Reset()
}

func (d *Compact32) Reset() (err error) {
	//First reset the values
	err = d.Driver.WriteSingleCoil(plc.MODBUS["M"][25], plc.ON)
	if err != nil {
		logger.Error("WriteSingleCoil:M25 ON: Reset")
		return
	}
	time.Sleep(time.Second * 1)
	err = d.Driver.WriteSingleCoil(plc.MODBUS["M"][25], plc.OFF)
	if err != nil {
		logger.Error("WriteSingleCoil:M25 OFF: Reset")
		return
	}
	return
}

func (d *Compact32) Start() (err error) {

	cycle := uint16(0)
	_, err = d.Monitor(cycle)
	if err != nil {
		logger.WithField("error", err).Error("Error in Monitoring")

	}
	return
}

func (d *Compact32) Stop() (err error) {
	if plc.LidPidTuningInProgress {
		plc.LidPidTuningInProgress = false
		plc.ExperimentRunning = false
		d.ExitCh <- errors.New("PID Error")
		return nil
	}
	plc.ExperimentRunning = false
	d.ExitCh <- errors.New("PCR Aborted")
	return nil
}

func (d *Compact32) SwitchOffCycle() (err error) {

	//For the cycle button
	err = d.Driver.WriteSingleCoil(plc.MODBUS["M"][20], plc.OFF)
	if err != nil {
		logger.Error("WriteSingleCoil:M20 : Start Cycle")
		return
	}
	err = d.Driver.WriteSingleCoil(plc.MODBUS["M"][21], plc.OFF)
	if err != nil {
		logger.Error("WriteSingleCoil:M21 : Start Cycle")
		return
	}
	return
}

func (d *Compact32) Cycle() (err error) {

	if !plc.ExperimentRunning {
		return errors.New("experiment is not running or maybe aborted")
	}

	// get blocked if homing is in progress
	for homingCount != 0 {
		time.Sleep(2 * time.Second)
		logger.Warnln("Homing is still in Progress, cycle will start once that is done.")
	}

	//For the cycle button
	err = d.Driver.WriteSingleCoil(plc.MODBUS["M"][20], plc.ON)
	if err != nil {
		logger.Error("WriteSingleCoil:M20 : Start Cycle")
		return
	}
	err = d.Driver.WriteSingleCoil(plc.MODBUS["M"][21], plc.ON)
	if err != nil {
		logger.Error("WriteSingleCoil:M21 : Start Cycle")
		return
	}
	logger.WithField("CYCLE RTPCR", "LED SWITCHED ON").Infoln("cycle started")

	err = d.checkCycleCompletion()
	if err != nil {
		return
	}

	plc.DataCapture = true

	go d.HomingRTPCR()
	return
}

// Monitor periodically. If CycleComplete == true, Scan will be populated
func (d *Compact32) Monitor(cycle uint16) (scan plc.Scan, err error) {

	logger.Infoln("---------------------------MONITOR------------------------")
	// Read current cycle

	scan.Temp = plc.CurrentCycleTemperature
	scan.LidTemp = float32(plc.CurrentLidTemp)
	logger.Infoln("	scan.Temp: ", scan.Temp, "\tscan.LidTemp: ", scan.LidTemp)
	if !plc.ExperimentRunning {
		return scan, errors.New("experiment is not running or maybe aborted")
	}

	if plc.CycleComplete {
		scan.CycleComplete = true

		// If the invoker has already read this cycle data, don't send it again!
		if cycle == scan.Cycle {
			logger.Println("cycle----------scan cycle------EQUAL", cycle, scan.Cycle)
			return
		}
	}
	if plc.DataCapture {
		var data []byte
		for i := 0; i < noOfDyes; i++ {
			start := plc.FValueRegisterStartAddress + i*16
			data, err = d.Driver.ReadHoldingRegisters(plc.MODBUS["D"][start], uint16(16))
			if err != nil {
				logger.WithField("register", plc.MODBUS["D"][start]).Error("ReadHoldingRegisters: Wells emission data")

			}
			//need to change just for testing
			// if data[1] != 0 {
			// 	logger.Println("data received-------------->", data, "\n start", start)
			// 	break LOOP
			// }
			offset := 0 // offset of data. increment every 2 bytes!
			for j := 0; j < 16; j++ {
				scan.Wells[j][i] = binary.BigEndian.Uint16(data[offset : offset+2])
				offset += 2
			}
		}
		scan.Cycle = plc.CurrentCycle
		//write values to the file
	}
	return
}

func (d *Compact32) SelfTest() (status plc.Status) {
	// TBD
	return
}

func (d *Compact32) Calibrate() (err error) {
	// TBD
	return
}

func (d *Compact32) SetLidTemp(expectedLidTemp uint16) (err error) {

	var currentLidTemp, maxSleepTimeSecs uint16

	if !plc.ExperimentRunning {
		logger.Warnln("No experiment in progress... avoiding Lid Temp Setting")
		return
	}
	// Off Lid Heating
	err = d.SwitchOffLidTemp()
	if err != nil {
		logger.Errorln("switch off lid temp", err)
		return
	}
	_, err = d.Driver.WriteSingleRegister(plc.MODBUS["D"][134], expectedLidTemp)
	if err != nil {
		logger.WithField("lid_temperature", expectedLidTemp).Errorln("WriteSingleRegister:D134 :", err)
		return
	}
	logger.WithField("LID TEMP", "LID TEMP started").Infoln("LID TEMP STARTED")
	currentLidTemp, err = d.Driver.ReadSingleRegister(plc.MODBUS["D"][135])
	if err != nil {
		logger.WithField("lid_temperature", expectedLidTemp).Errorln("ReadSingleRegister:D135 :", err)
		return
	}
	logger.Infoln("Current Lid Temperature:", currentLidTemp)
	plc.CurrentLidTemp = float32(currentLidTemp) / 10

	// Start Lid Heating
	err = d.switchOnLidTemp()
	if err != nil {
		return
	}

	if expectedLidTemp <= currentLidTemp {
		goto monitorLidTemp
	}

	// NOTE: If temperature doesn't reach in this time interval then
	// experiment should be aborted
	// give 0.1 degree per sec increment
	// expected Sleep time secs:= ((expectedLidTemp - currentLidTemp)/10) * 10
	maxSleepTimeSecs = expectedLidTemp - currentLidTemp
	logger.Infoln("Waiting for ", maxSleepTimeSecs, " secs at Max for Lid to reach the Expected Temp of: ", expectedLidTemp)

	// monitor lid temp accurately till maxSleepTimeSecs is reached
	if err = d.heatLidWithDeadline(maxSleepTimeSecs, expectedLidTemp); err != nil {
		return
	}

monitorLidTemp:
	go d.monitorLidTemp(expectedLidTemp)

	return
}

func (d *Compact32) SwitchOffLidTemp() (err error) {
	// Off Lid Heating
	err = d.Driver.WriteSingleCoil(plc.MODBUS["M"][197], plc.OFF)
	if err != nil {
		logger.Errorln("Stop Lid Heating error")
		return
	}
	logger.WithField("LID TEMP OFF", "LID TEMP SWITCHED OFF").Infoln("LID TEMP SWITCHED OFF")
	return
}

// 1. Set LID Tuning
// 2. switch on pid bit
// 3. Write PID Temp to D 134
// 4. switch on lid temp
// 4. Continuously read M 43 till PID Tuning Success

func (d *Compact32) LidPIDCalibration() (err error) {
	// TODO: Logging this PLC Operation

	var pidTuningDone bool

	defer func() {
		if err != nil {
			logger.Errorln(err)
			d.ExitCh <- fmt.Errorf(plc.ErrorLidPIDTuning)
		}
		d.WsMsgCh <- "SUCCESS_LidPIDTuning_LidPIDTuningSuccess"
	}()

	// 1.
	setLidPIDTuningInProgress()
	defer resetLidPIDTuningInProgress()
	// Start PID for deck
	// 2.
	err = d.switchOnLidPIDCalibration()
	if err != nil {
		return
	}

	// 3.
	result, err := d.Driver.WriteSingleRegister(plc.MODBUS["D"][134], uint16(config.GetLidPIDTemp()*10))
	if err != nil {
		logger.Errorln("Error failed to write lid pid temperature: ", err)
		return err
	}
	logger.Infoln("result from lid pid temperature set ", result, config.GetLidPIDTemp())

	d.WsMsgCh <- "PROGRESS_LidPIDTuning_LidPIDTuningStarted"

	//4. switch on temp
	err = d.switchOnLidTemp()
	if err != nil {
		logger.Errorln("Switch ON Lid Temp error")
		return
	}

	// Reset PID in defer
	defer d.switchOffLidPIDCalibration()
	logger.Infoln(responses.PIDCalibrationStarted)

	// Check if pid tuning is Done
	// 5.
	for !pidTuningDone {
		pidTuningDone, err = d.readLidPIDCompletion()
		if err != nil {
			return
		}
		plc.HoldSleep(1)
	}

	plc.HoldSleep(120)

	logger.Infoln(responses.PIDCalibrationSuccess)

	return
}

func (d *Compact32) SetScanSpeedAndScanTime() (err error) {

	result, err := d.Driver.WriteSingleRegister(plc.MODBUS["D"][462], uint16(config.GetScanSpeed()))
	if err != nil {
		logger.Errorln("Error failed to write scan speed: ", err)
		return err
	}
	logger.Infoln("result from scan speed set ", result, config.GetScanSpeed())

	result, err = d.Driver.WriteSingleRegister(plc.MODBUS["D"][464], uint16(config.GetScanTime()))
	if err != nil {
		logger.Errorln("Error failed to write scan time: ", err)
		return err
	}
	logger.Infoln("result from scan time set ", result, config.GetScanTime())
	return
}

func (d *Compact32) CalculateOpticalResult(dye db.Dye, kitID string, knownValue, cycleCount int64) (opticalResult []db.DyeWellTolerance, err error) {

	wellsData := make(map[int][]uint16, cycleCount)
	defer func() {

		if err != nil {
			logger.WithField("ERR", err.Error()).Errorln("error in calculating the optical result", err.Error())
			d.wsErrch <- err
		}
	}()
	plc.ExperimentRunning = true
	start := plc.FValueRegisterStartAddress + (dye.Position-1)*16
	for i := 0; i < int(cycleCount); i++ {
		d.Cycle()
		wellsData[i] = make([]uint16, 16)
		data, err := d.Driver.ReadHoldingRegisters(plc.MODBUS["D"][start], uint16(16))
		if err != nil {
			logger.WithField("register", plc.MODBUS["D"][start]).Error("ReadHoldingRegisters: Wells emission data")
		}
		offset := 0 // offset of data. increment every 2 bytes!
		for j := 0; j < 16; j++ {
			wellsData[i][j] = binary.BigEndian.Uint16(data[offset : offset+2])
			offset += 2
		}

		d.WsMsgCh <- "PROGRESS_OPTCALIB_" + fmt.Sprintf("%d", (i+1)*(100/int(cycleCount)))
	}

	for j := 0; j < 16; j++ {
		var finalValue uint16
		var deviatedResult db.DyeWellTolerance
		for i := 0; i < int(cycleCount); i++ {

			finalValue += wellsData[i][j]
		}
		finalAvg := float64(finalValue) / float64(cycleCount)

		deviatedValue := float64(knownValue) - finalAvg
		deviatedResult.OpticalResult = math.Abs((deviatedValue / float64(knownValue)) * 100)

		deviatedResult.DyeID = dye.ID
		deviatedResult.KitID = kitID
		deviatedResult.WellNo = j
		deviatedResult.Valid = true
		if deviatedResult.OpticalResult > dye.Tolerance {
			deviatedResult.Valid = false
		}

		opticalResult = append(opticalResult, deviatedResult)
	}
	return
}
