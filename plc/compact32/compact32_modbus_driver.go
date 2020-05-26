package compact32

import (
	"encoding/binary"
	"github.com/goburrow/modbus"
	"sync"
)

type Compact32ModbusDriver struct {
	sync.RWMutex
	Client modbus.Client
}

func (d *Compact32ModbusDriver) WriteMultipleRegisters(address, quantity uint16, value []byte) (results []byte, err error) {
	d.Lock()
	defer d.Unlock()

	results, err = d.Client.WriteMultipleRegisters(address, quantity, value)
	return
}

func (d *Compact32ModbusDriver) WriteSingleRegister(address, value uint16) (results []byte, err error) {
	d.Lock()
	defer d.Unlock()

	results, err = d.Client.WriteSingleRegister(address, value)
	return
}

func (d *Compact32ModbusDriver) ReadHoldingRegisters(address, quantity uint16) (results []byte, err error) {
	d.Lock()
	defer d.Unlock()

	results, err = d.Client.ReadHoldingRegisters(address, quantity)
	return
}

func (d *Compact32ModbusDriver) ReadSingleRegister(address uint16) (value uint16, err error) {
	// Don't take lock as ReadHoldingRegisters does take a lock! Otherwise, deadlock
	var data []byte
	data, err = C32.Driver.ReadHoldingRegisters(address, uint16(1))
	if err != nil {
		return
	}
	value = binary.BigEndian.Uint16(data)
	return
}

func (d *Compact32ModbusDriver) ReadCoils(address, quantity uint16) (results []byte, err error) {
	d.Lock()
	defer d.Unlock()

	results, err = d.Client.ReadCoils(address, quantity)
	return
}

func (d *Compact32ModbusDriver) ReadSingleCoil(address uint16) (value uint16, err error) {
	// Don't take lock as ReadCoils does take a lock! Otherwise, deadlock
	var data []byte
	data, err = C32.Driver.ReadCoils(address, uint16(1))
	if err != nil {
		return
	}
	value = binary.BigEndian.Uint16(data)
	return
}
