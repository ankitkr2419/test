package simulator

import (
	"encoding/csv"
	"mylab/cpagent/plc"
	"mylab/cpagent/tec"
)


type Simulator struct {
	ExitCh  chan error
	WsMsgCh chan string
	wsErrch chan error
}


func NewSimulatorDriver(wsMsgCh chan string, wsErrch chan error, exit chan error, test bool) tec.Driver {

	tec := Simulator{}
	tec.ExitCh = exit
	tec.WsMsgCh = wsMsgCh
	tec.wsErrch = wsErrch

	return &tec // tec Driver
}

var errorCheckStopped, tecInProgress bool
var prevTemp float32 = 27.0

// TODO: Implement Simulator
func (t *Simulator) InitiateTEC() (err error) {

	return nil
}


func (t *Simulator) ConnectTEC(ts tec.TECTempSet) (err error) {

	return nil
}

func (t *Simulator) AutoTune() (err error) {

	return nil
}

func (t *Simulator) ResetDevice() (err error) {

	return nil
}

func (t *Simulator) TestRun() (err error) {
	
	return nil
}

func (t *Simulator) ReachRoomTemp() error{
	
	return nil
}

func (t *Simulator) RunStage(st []plc.Step, writer *csv.Writer, cycleNum uint16) (err error){
	
	return nil
}
