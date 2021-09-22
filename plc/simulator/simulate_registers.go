package simulator

import (
	"fmt"
	logger "github.com/sirupsen/logrus"
	"mylab/cpagent/plc"
)

// INFO: Change Log Level to Info if debugging

func (d *SimulatorDriver) simulateWriteMultipleRegisters(address, quantity uint16, value []byte) (results []byte, err error) {
	// TODO: Implement this only when related method arrive
	return
}

func (d *SimulatorDriver) simulateWriteSingleRegister(address, value uint16) (results []byte, err error) {
	err = d.checkForValidAddress("D", address)
	if err != nil {
		return
	}

	d.setRegister("D", address, value)

	results = []byte{uint8(value >> 8), uint8(value & 0xff)}

	logger.Debugln("Inside simulateWriteSingleRegister for deck ", d.DeckName, " result: ", results, ". address: ", address)
	return
}

func (d *SimulatorDriver) simulateReadHoldingRegisters(address, quantity uint16) (results []byte, err error) {
	err = d.checkForValidAddress("D", address)
	if err != nil {
		return
	}
	if quantity != 1 {
		err = fmt.Errorf("invalid register reading quantity")
		return
	}

	value := d.readRegister("D", address)

	results = []byte{uint8(value >> 8), uint8(value & 0xff)}

	logger.Debugln("Inside simulateReadHoldingRegisters for deck ", d.DeckName, " result: ", results, ". address: ", address)
	return
}

func (d *SimulatorDriver) simulateReadCoils(address, quantity uint16) (results []byte, err error) {
	err = d.checkForValidAddress("M", address)
	if err != nil {
		return
	}

	if quantity != 1 {
		err = fmt.Errorf("invalid coil reading quantity")
		return
	}

	value := d.readRegister("M", address)

	results = []byte{uint8(value & 0xff)}

	logger.Debugln("Inside simulateReadCoils for deck ", d.DeckName, " result: ", results, ". address: ", address)
	return
}

func (d *SimulatorDriver) simulateWriteSingleCoil(address, value uint16) (err error) {
	err = d.checkForValidAddress("M", address)
	if err != nil {
		return
	}

	d.setRegister("M", address, value)

	results := []byte{uint8(value & 0xff)}

	logger.Debugln("Inside simulateWriteSingleCoil for deck ", d.DeckName, " result: ", results, ". address: ", address)

	// If its not ON event then it doesn't need handling
	if value != plc.ON {
		return
	}

	// Calling in go routine so that masterLock is free

	switch address {
	case plc.MODBUS_EXTRACTION[d.DeckName]["M"][0]:
		go d.simulateOnMotor()
	case plc.MODBUS_EXTRACTION[d.DeckName]["M"][3]:
		go d.simulateOnHeater()
	case plc.MODBUS_EXTRACTION[d.DeckName]["M"][4], plc.MODBUS_EXTRACTION[d.DeckName]["M"][8]:
		go d.simulateShakerPIDTuning()
	}
	return
}
