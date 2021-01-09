package compact32

import (
	"mylab/cpagent/config"
	"mylab/cpagent/plc"
	"time"

	"github.com/goburrow/modbus"
	logger "github.com/sirupsen/logrus"
)

const (
	HOLDING_STAGE = "hold"
	CYCLE_STAGE   = "cycle"
)

// Internal Interface to ensure sync'ing and testing modbus interfaces
type Compact32Driver interface {
	WriteSingleRegister(address, value uint16) (results []byte, err error)
	WriteMultipleRegisters(address, quantity uint16, value []byte) (results []byte, err error)
	ReadCoils(address, quantity uint16) (results []byte, err error)
	ReadSingleCoil(address uint16) (value uint16, err error)
	ReadHoldingRegisters(address, quantity uint16) (results []byte, err error)
	ReadSingleRegister(address uint16) (value uint16, err error)
	WriteSingleCoil(address, value uint16) (err error)
}

type Compact32 struct {
	ExitCh chan error
	Driver Compact32Driver
}

type Compact32Deck struct {
	name       string
	ExitCh     chan error
	DeckDriver Compact32Driver
}

func NewCompact32Driver(exit chan error, test bool) plc.Driver {
	/* Modbus RTU/ASCII */
	handler := modbus.NewRTUClientHandler(config.ReadEnvString("MODBUS_TTY"))
	handler.BaudRate = 9600
	handler.DataBits = 8
	handler.Parity = "E"
	handler.StopBits = 1
	handler.SlaveId = 5 // THis is hard-coded as the PLC RS485 is configured as SlaveID-5
	handler.Timeout = 500 * time.Millisecond

	handler.Connect()
	driver := Compact32ModbusDriver{}
	driver.Client = modbus.NewClient(handler)

	C32 := Compact32{}
	C32.Driver = &driver
	C32.ExitCh = exit

	// Start the Heartbeat
	go C32.HeartBeat()

	// Specifically for testing!
	if test {
		p := plc.Stage{
			Holding: []plc.Step{
				plc.Step{65.3, 3.2, 1},
				plc.Step{85.3, 3.1, 1},
				plc.Step{95, 2.8, 1},
			},
			Cycle: []plc.Step{
				plc.Step{55, 3, 1},
				plc.Step{65, 3, 1},
				plc.Step{75, 3, 1},
				plc.Step{85, 3, 1},
				plc.Step{95, 3, 1},
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

	return &C32 // plc Driver
}

// Compact32 Driver for Deck A
func NewCompact32DeckADriver(exit chan error, test bool) plc.DeckDriver {
	/* Modbus RTU/ASCII */
	handler := modbus.NewRTUClientHandler(config.ReadEnvString("MODBUS_TTY"))
	handler.BaudRate = 9600
	handler.DataBits = 8
	handler.Parity = "E"
	handler.StopBits = 1
	handler.SlaveId = byte(1)
	handler.Timeout = 200 * time.Millisecond

	handler.Connect()
	driver := Compact32ModbusDriver{}
	driver.Client = modbus.NewClient(handler)

	C32 := Compact32Deck{}
	C32.DeckDriver = &driver
	C32.ExitCh = exit
	C32.name = "A"

	return &C32 // plc Driver
}

// Compact32 Driver for Deck B
func NewCompact32DeckBDriver(exit chan error, test bool) plc.DeckDriver {
	/* Modbus RTU/ASCII */
	handler := modbus.NewRTUClientHandler(config.ReadEnvString("MODBUS_TTY"))
	handler.BaudRate = 9600
	handler.DataBits = 8
	handler.Parity = "E"
	handler.StopBits = 1
	handler.SlaveId = byte(2)
	handler.Timeout = 200 * time.Millisecond

	handler.Connect()
	driver := Compact32ModbusDriver{}
	driver.Client = modbus.NewClient(handler)

	C32 := Compact32Deck{}
	C32.DeckDriver = &driver
	C32.ExitCh = exit
	C32.name = "B"

	return &C32 // plc Driver
}
