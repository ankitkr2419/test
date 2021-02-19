package compact32

import (
	"encoding/binary"
	"sync"
	"time"

	"github.com/goburrow/modbus"
)

type Compact32ModbusDriver struct {
	Client modbus.Client
}

// We have 2 masters, only 1 should be allowed and that too with
// 200ms delay for 9600 baud rate
// 100ms delay for 19600 baud rate
var masterLock sync.Mutex

func (d *Compact32ModbusDriver) WriteMultipleRegisters(address, quantity uint16, value []byte) (results []byte, err error) {
	masterLock.Lock()
	defer masterLock.Unlock()
	time.Sleep(200 * time.Millisecond)
	results, err = d.Client.WriteMultipleRegisters(address, quantity, value)
	masterLock.Unlock()
	return
}

func (d *Compact32ModbusDriver) WriteSingleRegister(address, value uint16) (results []byte, err error) {
	masterLock.Lock()
	defer masterLock.Unlock()
	time.Sleep(200 * time.Millisecond)
	results, err = d.Client.WriteSingleRegister(address, value)
	return
}

func (d *Compact32ModbusDriver) ReadHoldingRegisters(address, quantity uint16) (results []byte, err error) {
	masterLock.Lock()
	defer masterLock.Unlock()
	time.Sleep(200 * time.Millisecond)
	results, err = d.Client.ReadHoldingRegisters(address, quantity)
	return
}

func (d *Compact32ModbusDriver) ReadSingleRegister(address uint16) (value uint16, err error) {
	// Don't take lock as ReadHoldingRegisters does take a lock! Otherwise, deadlock
	var data []byte
	data, err = d.ReadHoldingRegisters(address, uint16(1))
	if err != nil {
		return
	}
	value = binary.BigEndian.Uint16(data)
	return
}

func (d *Compact32ModbusDriver) ReadCoils(address, quantity uint16) (results []byte, err error) {
	masterLock.Lock()
	defer masterLock.Unlock()
	time.Sleep(200 * time.Millisecond)
	results, err = d.Client.ReadCoils(address, quantity)
	return
}

func (d *Compact32ModbusDriver) ReadSingleCoil(address uint16) (value uint16, err error) {
	// Don't take lock as ReadCoils does take a lock! Otherwise, deadlock
	var data []byte
	data, err = d.ReadCoils(address, uint16(1))
	if err != nil {
		return
	}
	// ReadCoil returns a single  byte!
	value = binary.BigEndian.Uint16([]byte{data[0], 0x00})
	return
}

func (d *Compact32ModbusDriver) WriteSingleCoil(address, value uint16) (err error) {
	masterLock.Lock()
	defer masterLock.Unlock()
	time.Sleep(200 * time.Millisecond)
	_, err = d.Client.WriteSingleCoil(address, value)
	return
}
