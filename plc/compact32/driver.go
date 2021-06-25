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
		beat, err = d.Driver.ReadSingleRegister(plc.MODBUS["D"][100])
		if err != nil {
			logger.WithField("beat", beat).Error("ReadSingleRegister:D100 : Read PLC heartbeat")
			break
		}

		// logger.Infoln("Heartbeat output: ", beat)

		// 3 attempts to check for heartbeat of PLC and write ours!
		for i := 0; i < 3; i++ {
			if beat == 1 { // If beat is 1, PLC is alive, so write 2
				_, err := d.Driver.WriteSingleRegister(plc.MODBUS["D"][100], uint16(2))
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
	d.ExitCh <- errors.New("PLC is not responding and maybe dead. Abort!!")
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
	// err = d.Driver.WriteSingleCoil(plc.MODBUS["M"][25], plc.OFF)
	// if err != nil {
	// 	logger.Error("WriteSingleCoil:M102 : Start Cycle")
	// }

	// err = d.Driver.WriteSingleCoil(plc.MODBUS["M"][102], plc.ON)
	// if err != nil {
	// 	logger.Error("WriteSingleCoil:M102 : Start Cycle")
	// }
	return
}

func (d *Compact32) Stop() (err error) {
	err = d.Driver.WriteSingleCoil(plc.MODBUS["M"][102], plc.OFF)
	if err != nil {
		logger.Error("WriteSingleCoil:M102 : Stop Cycle")
	}
	return
}

func (d *Compact32) Cycle() (err error) {
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

	time.Sleep(time.Second * 1)

	// for the rotation button
	err = d.Driver.WriteSingleCoil(plc.MODBUS["M"][14], plc.ON)
	if err != nil {
		logger.Error("WriteSingleCoil:M20 : Start Cycle")
		return
	}
	err = d.Driver.WriteSingleCoil(plc.MODBUS["M"][15], plc.ON)
	if err != nil {
		logger.Error("WriteSingleCoil:M21 : Start Cycle")
		return
	}

	// Let the rotation complete
	time.Sleep(time.Second * 20)

	return
}

// Monitor periodically. If CycleComplete == true, Scan will be populated
func (d *Compact32) Monitor(cycle uint16) (scan plc.Scan, err error) {

	// Read current cycle
	// scan.Cycle, err = d.Driver.ReadSingleRegister(plc.MODBUS["D"][133])
	// if err != nil {
	// 	logger.Error("ReadSingleRegister:D133: current cycle")
	// 	return
	// }

	// // Read cycle temperature.. PLC returns 653 for 65.3 degrees
	// var tmp uint16
	// tmp, err = d.Driver.ReadSingleRegister(plc.MODBUS["D"][132])
	// if err != nil {
	// 	logger.Error("ReadSingleRegister:D132: Cycle Temperature")
	// 	return
	// }
	// scan.Temp = float32(tmp) / 10

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
	// if tmp == 0 { // 0x0000 means cycle is not complete
	// 	// Values would not have changed.
	// 	scan.CycleComplete = false
	// 	return
	// }
	// scan.CycleComplete = true

	// // If the invoker has already read this cycle data, don't send it again!
	// if cycle == scan.Cycle {
	// 	return
	// }

	// Scan all the data from the Wells (96 x 6). Since max read is 123 registers, we shall read 96 at a time.
	// start := plc.MODBUS["D"][23]

	// err = d.Driver.WriteSingleCoil(plc.MODBUS["M"][20], plc.OFF)
	// if err != nil {
	// 	logger.Error("WriteSingleCoil:M20 : Start Cycle")
	// 	return
	// }
	// err = d.Driver.WriteSingleCoil(plc.MODBUS["M"][21], plc.OFF)
	// if err != nil {
	// 	logger.Error("WriteSingleCoil:M21 : Start Cycle")
	// 	return
	// }

	// err = d.Driver.WriteSingleCoil(plc.MODBUS["M"][14], plc.OFF)
	// if err != nil {
	// 	logger.Error("WriteSingleCoil:M14 : Start Cycle")
	// 	return
	// }
	// err = d.Driver.WriteSingleCoil(plc.MODBUS["M"][15], plc.OFF)
	// if err != nil {
	// 	logger.Error("WriteSingleCoil:M15 : Start Cycle")
	// 	return
	// }

	start := 44
	var data []byte

	for i := 0; i < 2; i++ {
		start = start + (16 * i)

	LOOP:
		for {

			data, err = d.Driver.ReadHoldingRegisters(plc.MODBUS["D"][start], uint16(16))
			if err != nil {
				logger.WithField("register", plc.MODBUS["D"][start]).Error("ReadHoldingRegisters: Wells emission data")

			}

			if len(data) != 0 {
				logger.Println("data received-------------->", data, "\n start", start)
				break LOOP
			}

		}

		offset := 0 // offset of data. increment every 2 bytes!

		for j := 0; j < 4; j++ {
			// populate each wells with 2 emissions each

			scan.Wells[j][i] = binary.BigEndian.Uint16(data[offset : offset+2])
			logger.Println("emission----", scan.Wells[j][i])
			offset += 8

			logger.Println("well----", j, "value", scan.Wells[j])

		}

	}
	//write values to the file

	logger.Println(scan.Wells)

	// logger.WithField("scan", scan).Debug("Monitored data")

	// Write to inform PLC that reading is completed
	// err = d.Driver.WriteSingleCoil(plc.MODBUS["M"][106], plc.OFF)
	// if err != nil {
	// 	logger.Error("WriteSingleCoil:M106: PC reading done")
	// 	return
	// }

	d.Reset()

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