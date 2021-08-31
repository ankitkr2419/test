package compact32

import (
	"encoding/binary"
	"math"
	"mylab/cpagent/config"
	"mylab/cpagent/db"
	"mylab/cpagent/plc"
	"time"

	"github.com/goburrow/modbus"
	logger "github.com/sirupsen/logrus"
)

const (
	HOLDING_STAGE     = "hold"
	CYCLE_STAGE       = "cycle"
	baudRate          = 9600
	dataBits          = 8
	parity            = "E"
	stopBits          = 1
	station_1_SlaveID = byte(1)
	station_2_SlaveID = byte(2)
	timeoutMs         = 50
)

type Compact32 struct {
	ExitCh  chan error
	WsMsgCh chan string
	wsErrch chan error
	Driver  plc.Compact32Driver
}

func NewCompact32Driver(wsMsgCh chan string, wsErrch chan error, exit chan error, test bool) plc.Driver {
	/* Modbus RTU/ASCII */
	handler := modbus.NewRTUClientHandler(config.ReadEnvString("MODBUS_TTY"))
	handler.BaudRate = baudRate
	handler.DataBits = dataBits
	handler.Parity = parity
	handler.StopBits = stopBits
	handler.SlaveId = station_1_SlaveID // THis is hard-coded as the PLC RS485 is configured as SlaveID-5
	handler.Timeout = timeoutMs * time.Millisecond

	handler.Connect()
	driver := Compact32ModbusDriver{}
	driver.Client = modbus.NewClient(handler)

	C32 := Compact32{}
	C32.Driver = &driver
	C32.ExitCh = exit
	C32.WsMsgCh = wsMsgCh

	// Start the Heartbeat
	// TODO: Uncomment this after RT-PCR m/c is ready
	go C32.HeartBeat()

	// Specifically for testing!
	// ASK: Should testing logic be present here?
	// It doesn't make sense as now testing is dependent on TEC as well.
	if test {
		p := plc.Stage{
			Holding: []plc.Step{
				plc.Step{TargetTemp: 65.3, RampUpTemp: 2.1, HoldTime: 5, DataCapture: false},
				plc.Step{TargetTemp: 85.3, RampUpTemp: 2.2, HoldTime: 3, DataCapture: false},
				plc.Step{TargetTemp: 95, RampUpTemp: 2, HoldTime: 5, DataCapture: false},
			},
			Cycle: []plc.Step{
				plc.Step{TargetTemp: 55, RampUpTemp: 2, HoldTime: 5, DataCapture: false},
				plc.Step{TargetTemp: 65, RampUpTemp: 2, HoldTime: 5, DataCapture: false},
				plc.Step{TargetTemp: 75, RampUpTemp: 2, HoldTime: 5, DataCapture: false},
				plc.Step{TargetTemp: 85, RampUpTemp: 2, HoldTime: 5, DataCapture: false},
				plc.Step{TargetTemp: 95, RampUpTemp: 2, HoldTime: 5, DataCapture: true},
			},
			CycleCount: 3,
		}

		C32.Stop()
		C32.ConfigureRun(p)
		C32.Start()

		cycle := uint16(0)
		for {
			scan, err := C32.Monitor(cycle)
			if err != nil {
				logger.WithField("error", err).Error("Error in Monitoring")
				break
			}
			// Log everytime to scan changes!
			if scan.Cycle != cycle {
				cycle = scan.Cycle
				logger.WithField("scan", scan).Info("Monitoring..")
			}
			logger.WithField("temperature", scan.Temp).Info("Monitoring..")

			if cycle == p.CycleCount {
				// Cycling done .. issue stop!
				break
			}
			time.Sleep(200 * time.Millisecond)
		}
		C32.Stop()
	}
	err := C32.SwitchOffLidTemp()
	if err != nil {
		logger.Warnln("Failed to switch off lid temp", err)
	}
	return &C32 // plc Driver
}

// Compact32 Driver for Deck A
// TODO: Use test to Configure the Deck Operations
func NewCompact32DeckDriverA(wsMsgCh chan string, wsErrch chan error, exit chan error, test bool) (plc.Extraction, *modbus.RTUClientHandler) {
	/* Modbus RTU/ASCII */
	handler := modbus.NewRTUClientHandler(config.ReadEnvString("MODBUS_TTY"))
	handler.BaudRate = baudRate
	handler.DataBits = dataBits
	handler.Parity = parity
	handler.StopBits = stopBits
	handler.SlaveId = station_1_SlaveID
	handler.Timeout = timeoutMs * time.Millisecond

	handler.Connect()
	driver := Compact32ModbusDriver{}

	driver.Client = modbus.NewClient(handler)

	C32 := plc.Compact32Deck{}
	C32.DeckDriver = &driver
	C32.ExitCh = exit
	C32.WsMsgCh = wsMsgCh
	C32.WsErrCh = wsErrch

	plc.SetDeckName(&C32, plc.DeckA)

	return &C32, handler
}

func NewCompact32DeckDriverB(wsMsgCh chan string, exit chan error, test bool, handler *modbus.RTUClientHandler) plc.Extraction {
	/* Modbus RTU/ASCII */
	handler2 := modbus.NewRTUClientHandler(config.ReadEnvString("MODBUS_TTY"))
	handler2.BaudRate = baudRate
	handler2.DataBits = dataBits
	handler2.Parity = parity
	handler2.StopBits = stopBits
	handler2.SlaveId = station_2_SlaveID
	handler2.Timeout = timeoutMs * time.Millisecond

	handler2.Connect()
	driver := Compact32ModbusDriver{}
	driver.Client = modbus.NewClient2(handler2, handler)

	C32 := plc.Compact32Deck{}
	C32.DeckDriver = &driver
	C32.ExitCh = exit
	C32.WsMsgCh = wsMsgCh

	plc.SetDeckName(&C32, plc.DeckB)

	return &C32
}

func (d *Compact32) CalculateOpticalResult(dye db.Dye, kitID string, knownValue, cycleCount int64) (opticalResult []db.DyeWellTolerance, err error) {

	wellsData := make(map[int][]uint16, cycleCount)

	plc.ExperimentRunning = true
	start := plc.FValueRegisterStartAddress + (dye.Position-1)*16
	for i := 0; i < int(cycleCount); i++ {
		d.Cycle()
		data, err := d.Driver.ReadHoldingRegisters(plc.MODBUS["D"][start], uint16(16))
		if err != nil {
			logger.WithField("register", plc.MODBUS["D"][start]).Error("ReadHoldingRegisters: Wells emission data")
		}
		offset := 0 // offset of data. increment every 2 bytes!
		for j := 0; j < 16; j++ {
			wellsData[i][j] = binary.BigEndian.Uint16(data[offset : offset+2])
			offset += 2
		}
	}

	for j := 0; j < 16; j++ {
		var finalValue uint16
		var deviatedResult db.DyeWellTolerance
		for i := 0; i < int(cycleCount); i++ {

			finalValue += wellsData[i][j]
		}
		finalAvg := int64(finalValue) / cycleCount
		deviatedValue := knownValue - int64(finalAvg)
		deviatedResult.OpticalResult = math.Abs(float64((deviatedValue / knownValue) * 100))
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
