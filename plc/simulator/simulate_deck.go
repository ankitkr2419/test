package simulator

import (
	// "encoding/binary"
	"fmt"
	"mylab/cpagent/plc"
	logger "github.com/sirupsen/logrus"
)

func (d *SimulatorDriver) simulateWriteMultipleRegisters(address, quantity uint16, value []byte) (results []byte, err error) {
	// TODO: Implement this only when related method arrive
	return
}

func (d *SimulatorDriver) simulateWriteSingleRegister(address, value uint16) (results []byte, err error) {
	err = d.checkForValidAddress("D", address)
	if err != nil {
		return
	}
	REGISTERS_EXTRACTION[d.DeckName]["D"][address] = value

	results = []byte{uint8(value>>8), uint8(value&0xff) }
	
	logger.Infoln("Inside simulateWriteSingleRegister result: ", results)
	return
}

func (d *SimulatorDriver) simualateReadHoldingRegisters(address, quantity uint16) (results []byte, err error) {
	err = d.checkForValidAddress("D", address)
	if err != nil {
		return
	}
	if quantity != 1 {
		err = fmt.Errorf("invalid register reading quantity")
		return
	}

	value := REGISTERS_EXTRACTION[d.DeckName]["D"][address]
	results = []byte{uint8(value>>8), uint8(value&0xff) }

	logger.Infoln("Inside simualateReadHoldingRegisters result: ", results)
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

	value := REGISTERS_EXTRACTION[d.DeckName]["M"][address]
	results = []byte{uint8(value&0xff) }

	logger.Infoln("Inside simulateReadCoils result: ", results)
	return
}

func (d *SimulatorDriver) simulateWriteSingleCoil(address, value uint16) (err error){
	err = d.checkForValidAddress("M", address)
	if err != nil {
		return
	}
	REGISTERS_EXTRACTION[d.DeckName]["D"][address] = value

	results := []byte{uint8(value&0xff) }
	
	logger.Infoln("Inside simulateWriteSingleCoil result: ", results)
	return
}


func (d *SimulatorDriver) checkForValidAddress(registerType string, address uint16) (err error) {
	switch registerType{
	case "M":
		// valid range 0-8
		lowestMAddress := plc.MODBUS_EXTRACTION[d.DeckName][registerType][0]
		highestMAddress := plc.MODBUS_EXTRACTION[d.DeckName][registerType][8]

		if address >= lowestMAddress && address <= highestMAddress {
			return
		}
		
	case "D":
		// valid range 200-226
		lowestDAddress := plc.MODBUS_EXTRACTION[d.DeckName][registerType][200]
		highestDAddress := plc.MODBUS_EXTRACTION[d.DeckName][registerType][226]

		// check for divisibility by 2 as well
		if address >= lowestDAddress && address <= highestDAddress && address % 2 != 1 {
			return
		}

	default:
		err = fmt.Errorf("Invalid register Type")
	}

	err = fmt.Errorf("Invalid register address")
	return 
}

