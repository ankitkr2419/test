package compact32

import (
	"encoding/binary"
	"mylab/cpagent/config"
	"mylab/cpagent/plc"
	"sync"
	"time"

	"github.com/goburrow/modbus"
)

const (
	HOLDING_STAGE = "hold"
	CYCLE_STAGE   = "cycle"
)

type Compact32 struct {
	sync.RWMutex
	Client modbus.Client
	ExitCh chan error
}

var compact32 Compact32 = Compact32{}

func NewCompact32Driver(exit chan error) plc.Driver {
	/* Modbus RTU/ASCII */
	handler := modbus.NewRTUClientHandler(config.ReadEnvString("MODBUS_TTY"))
	handler.BaudRate = 9600
	handler.DataBits = 8
	handler.Parity = "E"
	handler.StopBits = 1
	handler.SlaveId = 1
	handler.Timeout = 2 * time.Second

	handler.Connect()
	compact32.Client = modbus.NewClient(handler)
	compact32.ExitCh = exit

	// Start the Heartbeat
	go compact32.HeartBeat()

	return &compact32
}

// Helper Routines to ensure sync!
func (d *Compact32) WriteMultipleRegisters(address, quantity uint16, value []byte) (results []byte, err error) {
	d.Lock()
	defer d.Unlock()

	results, err = d.Client.WriteMultipleRegisters(address, quantity, value)
	return
}

func (d *Compact32) WriteSingleRegister(address, value uint16) (results []byte, err error) {
	d.Lock()
	defer d.Unlock()

	results, err = d.Client.WriteSingleRegister(address, value)
	return
}

func (d *Compact32) ReadHoldingRegisters(address, quantity uint16) (results []byte, err error) {
	d.Lock()
	defer d.Unlock()

	results, err = d.Client.ReadHoldingRegisters(address, quantity)
	return
}

func (d *Compact32) ReadSingleRegister(address uint16) (value uint16, err error) {
	// Don't take lock as ReadHoldingRegisters does take a lock! Otherwise, deadlock
	var data []byte
	data, err = compact32.ReadHoldingRegisters(address, uint16(1))
	if err != nil {
		return
	}
	value = binary.BigEndian.Uint16(data)
	return
}

func (d *Compact32) ReadCoils(address, quantity uint16) (results []byte, err error) {
	d.Lock()
	defer d.Unlock()

	results, err = d.Client.ReadCoils(address, quantity)
	return
}

func (d *Compact32) ReadSingleCoil(address uint16) (value uint16, err error) {
	// Don't take lock as ReadCoils does take a lock! Otherwise, deadlock
	var data []byte
	data, err = compact32.ReadCoils(address, uint16(1))
	if err != nil {
		return
	}
	value = binary.BigEndian.Uint16(data)
	return
}
