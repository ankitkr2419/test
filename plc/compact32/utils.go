package compact32

import (
	"encoding/binary"
	"errors"
	"fmt"
	"mylab/cpagent/plc"
	"mylab/cpagent/responses"

	"time"

	logger "github.com/sirupsen/logrus"
)

const (
	degreePlay = 50
)

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

// maxSleepTimeSecs is the Deadline
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
		if (currentLidTemp > (expectedLidTemp + degreePlay)) || (currentLidTemp < (expectedLidTemp - degreePlay)) {
			err = fmt.Errorf("lid temperature has exceeded the limits")
			logger.WithField("err:", err.Error()).Errorln("Current Lid Temp has exceeded the limits: ", currentLidTemp)
			d.ExitCh <- errors.New("PCR Aborted")
			return
		}
	}
}

func (d *Compact32) switchOnLidTemp() (err error) {
	// Switch On Lid Heating
	err = d.Driver.WriteSingleCoil(plc.MODBUS["M"][197], plc.ON)
	if err != nil {
		logger.Errorln("Switch ON Lid Heating error")
		return
	}
	logger.WithField("LID TEMP ON", "LID TEMP SWITCHED ON").Infoln("LID TEMP SWITCHED ON")
	return
}

func setLidPIDTuningInProgress() {
	plc.LidPidTuningInProgress = true
	plc.ExperimentRunning = true

}

func resetLidPIDTuningInProgress() {
	plc.LidPidTuningInProgress = false
	plc.ExperimentRunning = false
}

func (d *Compact32) switchOnLidPIDCalibration() (err error) {

	// Switch On Lid PID Tuning
	err = d.Driver.WriteSingleCoil(plc.MODBUS["M"][42], plc.ON)
	if err != nil {
		logger.Errorln("Switch ON Lid PID Tuning error")
		return
	}
	logger.WithField("LID PID TEMP ON", "LID PID TEMP SWITCHED ON").Infoln("LID PID TEMP SWITCHED ON")

	return
}

func (d *Compact32) switchOffLidPIDCalibration() (err error) {
	//Stop Lid Heating
	err = d.SwitchOffLidTemp()
	if err != nil {
		return
	}
	// Off Lid PID Tuning
	err = d.Driver.WriteSingleCoil(plc.MODBUS["M"][42], plc.OFF)
	if err != nil {
		logger.Errorln("Stop Lid PID Tuning error")
		return
	}
	logger.WithField("LID PID TEMP OFF", "LID PID TEMP SWITCHED OFF").Infoln("LID PID TEMP SWITCHED OFF")
	return
}

func (d *Compact32) readLidPIDCompletion() (pidTuningDone bool, err error) {

	if !plc.LidPidTuningInProgress {
		return false, responses.LidPidTuningOffError
	}

	result, err := d.Driver.ReadSingleRegister(plc.MODBUS["D"][504])
	if err != nil {
		logger.WithField("LID PID ERR", err).Errorln("Error Reading M43")
		return false, err
	}

	logger.Infoln("readLidPIDCompletion result: ", result)

	if result == K4 {
		return true, nil
	}

	return
}

func (d *Compact32) checkCycleCompletion() (err error) {

	var result []byte
	for {
		if !plc.ExperimentRunning {
			logger.Errorln("experiment has stoped running")
			return errors.New("experiment has stoped running")
		}

		result, err = d.Driver.ReadCoils(plc.MODBUS["M"][45], uint16(1))
		if err != nil {
			logger.Error("ReadCoil:M45 ", err)
			return
		}

		logger.Warnln("M45", result)

		if result[0] == 45 {
			logger.Warnln("Cycle completed")
			return
		}
		time.Sleep(2 * time.Second)
	}

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
