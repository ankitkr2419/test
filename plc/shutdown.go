package plc

import (
	"time"

	logger "github.com/sirupsen/logrus"
)

func (d *Compact32Deck) ShutDown() (err error) {
	err = d.DeckDriver.WriteSingleCoil(MODBUS_EXTRACTION[d.name]["M"][10], ON)
	if err != nil {
		logger.Errorln("error offing machine: ", err)
		return err
	}
	time.Sleep(time.Second * 5)
	err = d.DeckDriver.WriteSingleCoil(MODBUS_EXTRACTION[d.name]["M"][10], OFF)
	if err != nil {
		logger.Errorln("error offing machine: ", err)
		return err
	}
	return nil
}

