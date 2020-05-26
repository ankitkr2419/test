package compact32

import (
	"encoding/binary"
	"errors"
	"mylab/cpagent/plc"
	"time"

	logger "github.com/sirupsen/logrus"
)

// Interface Implementation Methods
func (d *Compact32) HeartBeat() {
	var err error

	logger.Info("Starting HeartBeat...")
LOOP:
	for {
		time.Sleep(200 * time.Millisecond) // sleep it off for a bit

		// Read heartbeat status to check if PLC is alive!
		var beat uint16
		beat, err = compact32.ReadSingleRegister(MODBUS["D"][100])
		if err != nil {
			break
		}

		// 3 attempts to check for heartbeat of PLC and write ours!
		for i := 0; i < 3; i++ {
			if beat == 1 { // If beat is 1, PLC is alive, so write 2
				_, err = compact32.WriteSingleRegister(MODBUS["D"][100], uint16(2))
				if err != nil {
					// exit!!
					break LOOP
				}
				continue LOOP
			}

			logger.Warn("Attempt: %d failed. PLC heartbeat value is still %d. Retrying...", i+1, beat)
			time.Sleep(200 * time.Millisecond) // sleep it off for a bit
		}

		// If we have reached here, PLC has not updated heartbeat for 3 tries, it's dead! Abort!
		logger.Warn("PLC heartbeat value is still 1 after 3 attepts. Abort!")
		err = errors.New("PLC is not responding and maybe dead. Abort!!")
		break
	}

	// something went wrong. Signal parent process
	logger.WithField("err", err.Error()).Error("Heartbeat Error. Abort!")
	d.ExitCh <- err
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

func (d *Compact32) PreRun(stage plc.Stage) (err error) {
	err = writeStageData(HOLDING_STAGE, stage)
	if err != nil {
		// propagate error immediately
		return
	}

	err = writeStageData(CYCLE_STAGE, stage)
	return
}

func writeStageData(name string, stage plc.Stage) (err error) {
	// default settings for Holding stage
	steps := 4
	arr := stage.Holding
	quantity := uint16(12)
	address := MODBUS["D"][101]

	if name == CYCLE_STAGE {
		steps = 6
		arr = stage.Cycle
		quantity = uint16(18)
		address = MODBUS["D"][113]

	}
	targetTemp := make([]uint16, steps)
	rampUpTemp := make([]uint16, steps)
	holdTime := make([]uint16, steps)

	// Holding stage (at most 4 steps)
	for i, step := range arr {
		targetTemp[i] = (uint16)(step.TargetTemp * 10)
		rampUpTemp[i] = (uint16)(step.RampUpTemp * 10)
		holdTime[i] = (uint16)(step.HoldTime * 10)
	}

	// Get the data as byte array!
	temp := []uint16{}
	temp = append(temp, targetTemp...)
	temp = append(temp, rampUpTemp...)
	temp = append(temp, holdTime...)
	data := dataBlock(temp...)

	// write registers data
	_, err = compact32.Client.WriteMultipleRegisters(address, quantity, data)
	if err != nil {
		return
	}

	// Cycle count
	_, err = compact32.Client.WriteSingleRegister(MODBUS["D"][131], stage.CycleCount)

	return
}

// Monitor periodically. If CycleComplete == true, Scan will be populated
func (d *Compact32) Monitor() (scan plc.Scan, err error) {

	// Read current cycle
	scan.Cycle, err = compact32.ReadSingleRegister(MODBUS["D"][133])
	if err != nil {
		return
	}

	// Read cycle temperature.. PLC returns 653 for 65.3 degrees
	var tmp uint16
	tmp, err = compact32.ReadSingleRegister(MODBUS["D"][132])
	if err != nil {
		return
	}
	scan.Temp = float32(tmp) / 10

	// Read lid temperature
	tmp, err = compact32.ReadSingleRegister(MODBUS["D"][135])
	if err != nil {
		return
	}
	scan.LidTemp = float32(tmp) / 10

	// Read current cycle status
	tmp, err = compact32.ReadSingleCoil(MODBUS["M"][107])
	if err != nil {
		return
	}
	if tmp == 0 { // 0x0000 means cycle is not complete
		// Values would not have changed.
		scan.CycleComplete = false
		return
	}

	scan.CycleComplete = true

	// Scan all the data from the Wells (96 x 6). Since max read is 123 registers, we shall read 96 at a time.
	offset := 2000

	for i := 0; i < 6; i++ {
		var data []byte
		data, err = compact32.ReadHoldingRegisters(MODBUS["D"][offset+(i*96)], uint16(96))
		if err != nil {
			return
		}

		// populate the wells with 6 emissions each
		offset = 0 // offset of data. increment every 2 bytes!
		for j := 0; j < 96; j++ {
			emission := plc.Emissions{}
			for k := 0; k < 6; k++ {
				emission[k] = binary.BigEndian.Uint16(data[offset : offset+2])
				j += 1
				offset += 2
			}

			scan.Wells[(i*16)+(j%16)] = emission
		}

	}

	return
}
