package compact32

import (
	"mylab/cpagent/config"
	"mylab/cpagent/plc"
	"time"

	"github.com/goburrow/modbus"
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
}

type Compact32 struct {
	ExitCh chan error
	Driver Compact32Driver
}

var C32 Compact32

func NewCompact32Driver(exit chan error) plc.Driver {
	/* Modbus RTU/ASCII */
	handler := modbus.NewRTUClientHandler(config.ReadEnvString("MODBUS_TTY"))
	handler.BaudRate = 9600
	handler.DataBits = 8
	handler.Parity = "E"
	handler.StopBits = 1
	handler.SlaveId = 5
	handler.Timeout = 2 * time.Second

	handler.Connect()
	driver := Compact32ModbusDriver{}
	driver.Client = modbus.NewClient(handler)

	C32 = Compact32{}
	C32.Driver = &driver
	C32.ExitCh = exit

	// Start the Heartbeat
	go C32.HeartBeat()

	return &C32 // plc Driver
}
