package simulator

import (
	logger "github.com/sirupsen/logrus"
	"mylab/cpagent/plc"
	"time"
)

func (d *SimulatorDriver) simulateShakerPIDTuning() (err error) {

	var pidDone bool

	if d.isShakerPIDCalibrationInProgress() {
		return
	}
	logger.Infoln("Simulating shaker PID Tuning")

	d.setShakerPIDCalibrationInProgress()
	defer d.resetShakerPIDCalibrationInProgress()

	d.setRegister("D", plc.MODBUS_EXTRACTION[d.DeckName]["D"][504], pidProgressResponse)
	d.setRegister("D", plc.MODBUS_EXTRACTION[d.DeckName]["D"][534], pidProgressResponse)
	go d.simulateOnHeater()

	pidDone = d.sleepOrCheckForAbort(int(delay) * pidTuningTime)
	if !pidDone{
		return
	}

	d.setRegister("M", plc.MODBUS_EXTRACTION[d.DeckName]["M"][3], plc.OFF)
	d.setRegister("D", plc.MODBUS_EXTRACTION[d.DeckName]["D"][504], pidDoneResponse)
	d.setRegister("D", plc.MODBUS_EXTRACTION[d.DeckName]["D"][534], pidDoneResponse)

	return
}


func(d *SimulatorDriver) sleepOrCheckForAbort(maxPIDTime int) bool{

	for  maxPIDTime > 0{
		maxPIDTime--
		time.Sleep(time.Second)
		if plc.OFF == d.readRegister("M", plc.MODBUS_EXTRACTION[d.DeckName]["M"][4]) {
			return false
		}
	}
	return true
}