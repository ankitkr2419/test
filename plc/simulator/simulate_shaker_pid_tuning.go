package simulator

import (
	logger "github.com/sirupsen/logrus"
	"mylab/cpagent/plc"
	"time"
)

func (d *SimulatorDriver) simulateShakerPIDTuning() (err error) {

	if d.isShakerPIDCalibrationInProgress() {
		return
	}
	logger.Infoln("Simulating shaker PID Tuning")

	d.setShakerPIDCalibrationInProgress()
	defer d.resetShakerPIDCalibrationInProgress()

	d.setRegister("D", plc.MODBUS_EXTRACTION[d.DeckName]["D"][504], pidProgressResponse)
	d.setRegister("D", plc.MODBUS_EXTRACTION[d.DeckName]["D"][534], pidProgressResponse)
	go d.simulateOnHeater()

	logger.Warnln("AtS")
	time.Sleep(time.Duration(pidTuningTime*delay) * time.Second)
	d.setRegister("M", plc.MODBUS_EXTRACTION[d.DeckName]["M"][3], plc.OFF)
	d.setRegister("D", plc.MODBUS_EXTRACTION[d.DeckName]["D"][504], pidDoneResponse)
	d.setRegister("D", plc.MODBUS_EXTRACTION[d.DeckName]["D"][534], pidDoneResponse)

	return
}
