package compact32

import (
	"encoding/binary"
	"errors"
	"fmt"
	"mylab/cpagent/config"
	"mylab/cpagent/plc"

	"time"

	logger "github.com/sirupsen/logrus"
)

const (
	maxHomingTries             = 2
	homingSuccessValue         = 101
	noOfDyes                   = 4
	fValueRegisterStartAddress = 800
)

var homingCount int

// Interface Implementation Methods
func (d *Compact32) HeartBeat() {
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

		// logger.Infoln("Heartbeat output: ", beat)

		// 3 attempts to check for heartbeat of PLC and write ours!
		for i := 0; i < 3; i++ {
			if beat == 1 { // If beat is 1, PLC is alive, so write 2
				_, err = d.Driver.WriteSingleRegister(plc.MODBUS["D"][100], uint16(2))
				if err != nil {
					logger.WithField("beat", beat).Error("WriteSingleRegister:D100 : Read PLC heartbeat")
					// exit!!
					break LOOP
				}
				// logger.Infoln("Read val:", val)
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
	return
}

// dataBlock creates a sequence of uint16 data. (ref: modbus/client.go)
func dataBlock(value ...uint16) []byte {
	data := make([]byte, 2*len(value))
	for i, v := range value {
		binary.BigEndian.PutUint16(data[i*2:], v)
	}
	return data
}

func (d *Compact32) ConfigureRun(stage plc.Stage) (err error) {
	err = d.writeStageData(HOLDING_STAGE, stage)
	if err != nil {
		// propagate error immediately
		return
	}

	err = d.writeStageData(CYCLE_STAGE, stage)
	if err != nil {
		// propagate error immediately
		return
	}

	// Cycle count
	_, err = d.Driver.WriteSingleRegister(plc.MODBUS["D"][131], stage.CycleCount)
	if err != nil {
		logger.WithField("cycle_count", stage.CycleCount).Error("WriteSingleRegister:D131 : Number of Cycles")
	}
	return
}

func (d *Compact32) writeStageData(name string, stage plc.Stage) (err error) {
	// default settings for Holding stage
	steps := 4
	arr := stage.Holding
	quantity := uint16(12)
	address := plc.MODBUS["D"][101]

	if name == CYCLE_STAGE {
		steps = 6
		arr = stage.Cycle
		quantity = uint16(18)
		address = plc.MODBUS["D"][113]

	}
	targetTemp := make([]uint16, steps)
	rampUpTemp := make([]uint16, steps)
	holdTime := make([]uint16, steps)

	// Holding stage (at most 4 steps)
	for i, step := range arr {
		targetTemp[i] = (uint16)(step.TargetTemp * 10) // float32 with 1 decimal value to uint16
		rampUpTemp[i] = (uint16)(step.RampUpTemp * 10) // float32 with 1 decimal value to uint16
		holdTime[i] = (uint16)(step.HoldTime)
	}

	// Get the data as byte array!
	temp := []uint16{}
	temp = append(temp, targetTemp...)
	temp = append(temp, rampUpTemp...)
	temp = append(temp, holdTime...)
	data := dataBlock(temp...)

	logger.WithField("data", data).Debug("Writing Cycle/Holding configurations to PLC")

	// write registers data
	_, err = d.Driver.WriteMultipleRegisters(address, quantity, data)
	if err != nil {
		logger.WithFields(logger.Fields{
			"stage":    name,
			"address":  address,
			"quantity": quantity,
			"data":     data,
		}).Error("ConfigureRun: Writing Stage data")
		return
	}
	return
}

func (d *Compact32) HomingRTPCR() (err error) {

	defer func() {
		homingCount = 0
	}()

	if homingCount == maxHomingTries {
		err = fmt.Errorf("homing failed even after %v tries", maxHomingTries)
		logger.WithField("HOMING", err.Error()).Errorln("homing failed")
		d.ExitCh <- errors.New("PCR Aborted")
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

	err = d.Driver.WriteSingleCoil(plc.MODBUS["M"][100], plc.OFF)
	if err != nil {
		logger.Error("WriteSingleCoil:M100 : Start Cycle")
		return
	}
	logger.WithField("HOMING", "homing started").Infoln("HOMING STARTED")
	time.Sleep(time.Second * time.Duration(config.GetHomingTime()))
	result, err := d.Driver.ReadCoils(plc.MODBUS["M"][100], uint16(1))
	if err != nil {
		logger.Error("WriteSingleCoil:M100 ", err)
		return
	}
	logger.Infoln("homing result", result)
	if result[0] == homingSuccessValue {
		logger.WithField("HOMING", "Completed").Infoln("homing completed")
	} else {
		homingCount++
		// Try Homing Again
		err = d.HomingRTPCR()
		if err != nil {
			return
		}
	}
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
	plc.ExperimentRunning = false
	// err = d.Driver.WriteSingleCoil(plc.MODBUS["M"][102], plc.OFF)
	// if err != nil {
	// 	logger.Error("WriteSingleCoil:M102 : Stop Cycle")
	// }
	d.ExitCh <- errors.New("PCR Aborted")
	return nil
}

func (d *Compact32) Cycle() (err error) {

	if !plc.ExperimentRunning {
		return errors.New("experiment is not running or maybe aborted")
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
	err = plc.HoldSleep(int32(config.GetCycleTime()))
	if err != nil {
		logger.Errorln("Error while running cycle: ", err)
		return
	}

	plc.DataCapture = true
	// for {
	// 	cycleCompletion, err := d.Driver.ReadCoils(plc.MODBUS["M"][27], uint16(1))
	// 	if err != nil {
	// 		logger.Error("ReadSingleCoil:M27: Current PV cycle")
	// 		return err
	// 	}
	// 	fmt.Println("cycle completion ---------", cycleCompletion)
	// 	if cycleCompletion[0] == 1 {
	// 		plc.HeatingCycleComplete = true
	// 		err := d.Driver.WriteSingleCoil(plc.MODBUS["M"][27], uint16(0))
	// 		if err != nil {
	// 			logger.Error("ReadSingleCoil:M27: Current PV cycle")
	// 			return err
	// 		}
	// 		return nil
	// 	}
	// 	time.Sleep(time.Millisecond * 500)
	// }

	// for the rotation button, rotation button is required in manual move

	// err = d.Driver.WriteSingleCoil(plc.MODBUS["M"][14], plc.ON)
	// if err != nil {
	// 	logger.Error("WriteSingleCoil:M20 : Start Cycle")
	// 	return
	// }
	// err = d.Driver.WriteSingleCoil(plc.MODBUS["M"][15], plc.ON)
	// if err != nil {
	// 	logger.Error("WriteSingleCoil:M21 : Start Cycle")
	// 	return
	// }

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

		// // Read lid temperature
		// tmp, err = d.Driver.ReadSingleRegister(plc.MODBUS["D"][135])
		// if err != nil {
		// 	logger.Error("ReadSingleRegister:D135: Lid temperature")
		// 	return
		// }
		// scan.LidTemp = float32(tmp) / 10

		// // Read current cycle status
		// tmp, err = d.Driver.ReadSingleCoil(plc.MODBUS["M"][107])
		// if err != nil {
		// 	logger.Error("ReadSingleCoil:M107: Current PV cycle")
		// 	return
		// }
		// if !plc.CycleComplete { // 0x0000 means cycle is not complete
		// 	// Values would not have changed.
		// 	scan.CycleComplete = false
		// 	return
		// }
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
			start := fValueRegisterStartAddress + i*16
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
	err = d.Driver.WriteSingleCoil(plc.MODBUS["M"][197], plc.ON)
	if err != nil {
		logger.Errorln("WriteSingleCoil:M109 : Start Lid Heating")
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

func (d *Compact32) readCurrentLidTemp(expectedLidTemp uint16) (currentLidTemp uint16, err error) {
	currentLidTemp, err = d.Driver.ReadSingleRegister(plc.MODBUS["D"][135])
	if err != nil {
		logger.WithField("lid_temperature", expectedLidTemp).Errorln("readCurrentLidTemp :", err)
		return
	}
	logger.Infoln("Current Lid Temperature:", currentLidTemp)

	plc.CurrentLidTemp = float32(currentLidTemp) / 10
	return
}

func (d *Compact32) heatLidWithDeadline(maxSleepTimeSecs, expectedLidTemp uint16) (err error) {
	var lidProgressTime, currentLidTemp uint16

	currentLidTemp, err = d.readCurrentLidTemp(expectedLidTemp)

	for lidProgressTime < maxSleepTimeSecs {
		if !plc.ExperimentRunning {
			err = fmt.Errorf("No experiment running!")
			return
		}
		go func() {
			// ignore error till deadline
			currentLidTemp, _ = d.readCurrentLidTemp(expectedLidTemp)
		}()
		lidProgressTime++
		// 3 degree play
		if expectedLidTemp < (currentLidTemp + 30) {
			logger.Infoln("Lid Temperature of", currentLidTemp, " reached.")
			break
		}
		time.Sleep(time.Second)
	}
	return
}

func (d *Compact32) monitorLidTemp(expectedLidTemp uint16) {
	var currentLidTemp uint16
	var err error

	for {
		if !plc.ExperimentRunning {
			return
		}
		time.Sleep(2 * time.Second)
		//  Read lid temperature

		currentLidTemp, err = d.readCurrentLidTemp(expectedLidTemp)
		if err != nil {
			return
		}

		// Play is of +- 5 degrees
		if (currentLidTemp > (expectedLidTemp + 50)) || (currentLidTemp < (expectedLidTemp - 50)) {
			err = fmt.Errorf("lid temperature has exceeded the limits")
			logger.WithField("err:", err.Error()).Errorln("Current Lid Temp has exceeded the limits: ", currentLidTemp)
			d.ExitCh <- errors.New("PCR Aborted")
			return
		}
	}
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
