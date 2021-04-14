package simulator

import (
	"encoding/binary"
	"sync"
	"time"
)

type SimulatorDriver struct {
	// DeckName will be used for differentiating simulator registors 
	DeckName string
}

const delay = 50

// We have 2 masters, only 1 should be allowed and that too with
// 200ms delay for 9600 baud rate
// 100ms delay for 19600 baud rate
// 50ms delay for 37000 baud rate
// 40ms delay for 57000 baud rate
// NOTE: Only 9600 works!!!
var masterLock sync.Mutex

func (d *SimulatorDriver) WriteMultipleRegisters(address, quantity uint16, value []byte) (results []byte, err error) {
	masterLock.Lock()
	defer masterLock.Unlock()
	time.Sleep(time.Duration(delay) * time.Millisecond)
	results, err = simulateWriteMultipleRegisters(address, quantity, value)
	return
}

func (d *SimulatorDriver) WriteSingleRegister(address, value uint16) (results []byte, err error) {
	masterLock.Lock()
	defer masterLock.Unlock()
	time.Sleep(time.Duration(delay) * time.Millisecond)
	results, err = simulateWriteSingleRegister(address, value)
	return
}

func (d *SimulatorDriver) ReadHoldingRegisters(address, quantity uint16) (results []byte, err error) {
	masterLock.Lock()
	defer masterLock.Unlock()
	time.Sleep(time.Duration(delay) * time.Millisecond)
	results, err = simualateReadHoldingRegisters(address, quantity)
	return
}

func (d *SimulatorDriver) ReadSingleRegister(address uint16) (value uint16, err error) {
	// Don't take lock as ReadHoldingRegisters does take a lock! Otherwise, deadlock
	var data []byte
	data, err = d.ReadHoldingRegisters(address, uint16(1))
	if err != nil {
		return
	}
	value = binary.BigEndian.Uint16(data)
	return
}

func (d *SimulatorDriver) ReadCoils(address, quantity uint16) (results []byte, err error) {
	masterLock.Lock()
	defer masterLock.Unlock()
	time.Sleep(time.Duration(delay) * time.Millisecond)
	results, err = simulateReadCoils(address, quantity)
	return
}

func (d *SimulatorDriver) ReadSingleCoil(address uint16) (value uint16, err error) {
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

func (d *SimulatorDriver) WriteSingleCoil(address, value uint16) (err error) {
	masterLock.Lock()
	defer masterLock.Unlock()
	time.Sleep(time.Duration(delay) * time.Millisecond)
	err = simulateWriteSingleCoil(address, value)
	return
}
