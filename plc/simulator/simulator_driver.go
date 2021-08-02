package simulator

import (
	"encoding/binary"
	"mylab/cpagent/responses"
	"sync"
	"time"
)

type SimulatorDriver struct {
	// DeckName will be used for differentiating simulator registors
	DeckName string
}

// Below delay should be 50 for almost real simulation,
// all other delay's are its proportinate
var delay time.Duration = 50

// simulating masterLock like in compact32
var masterLock sync.Mutex

func (d *SimulatorDriver) WriteMultipleRegisters(address, quantity uint16, value []byte) (results []byte, err error) {
	masterLock.Lock()
	defer masterLock.Unlock()
	time.Sleep(delay * time.Millisecond)
	results, err = d.simulateWriteMultipleRegisters(address, quantity, value)
	return
}

func (d *SimulatorDriver) WriteSingleRegister(address, value uint16) (results []byte, err error) {
	masterLock.Lock()
	defer masterLock.Unlock()
	time.Sleep(delay * time.Millisecond)
	results, err = d.simulateWriteSingleRegister(address, value)
	return
}

func (d *SimulatorDriver) ReadHoldingRegisters(address, quantity uint16) (results []byte, err error) {
	masterLock.Lock()
	defer masterLock.Unlock()
	time.Sleep(delay * time.Millisecond)
	results, err = d.simulateReadHoldingRegisters(address, quantity)
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
	time.Sleep(delay * time.Millisecond)
	results, err = d.simulateReadCoils(address, quantity)
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
	time.Sleep(delay * time.Millisecond)
	err = d.simulateWriteSingleCoil(address, value)
	return
}

func UpdateDelay(d int) error {
	if d > 0 && d <= 100 {
		delay = time.Duration(d)
		return nil
	}
	return responses.DelayRangeInvalid
}
