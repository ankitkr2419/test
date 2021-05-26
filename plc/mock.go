package plc

import (
	"github.com/stretchr/testify/mock"
	"mylab/cpagent/db"
)

type PLCMockStore struct {
	mock.Mock
}

func (p *PLCMockStore) IsMachineHomed() (homed bool) {
	args := p.Called()
	return args.Get(0).(bool)
}

func (p *PLCMockStore) IsRunInProgress() (pro bool) {
	args := p.Called()
	return args.Get(0).(bool)
}

func (p *PLCMockStore) SetRunInProgress() {
	_ = p.Called()
	return
}

func (p *PLCMockStore) ResetRunInProgress() {
	_ = p.Called()
	return
}

func (p *PLCMockStore) AspireDispense(ad db.AspireDispense, cartridgeID int64, tipType string) (response string, err error) {
	args := p.Called(ad, cartridgeID, tipType)
	return args.Get(0).(string), args.Error(1)
}

func (p *PLCMockStore) AttachDetach(ad db.AttachDetach) (response string, err error) {
	args := p.Called(ad)
	return args.Get(0).(string), args.Error(1)
}

func (p *PLCMockStore) DiscardBoxCleanup() (response string, err error) {
	args := p.Called()
	return args.Get(0).(string), args.Error(1)
}

func (p *PLCMockStore) RestoreDeck() (response string, err error) {
	args := p.Called()
	return args.Get(0).(string), args.Error(1)
}

func (p *PLCMockStore) UVLight(uvTime string) (response string, err error) {
	args := p.Called(uvTime)
	return args.Get(0).(string), args.Error(1)
}

func (p *PLCMockStore) AddDelay(delay db.Delay, runRecipe bool) (response string, err error) {
	args := p.Called(delay, runRecipe)
	return args.Get(0).(string), args.Error(1)
}

func (p *PLCMockStore) SetCurrentProcess(step int64) {
	_ = p.Called(step)
	return
}

func (p *PLCMockStore) RunRecipeWebsocketData(recipe db.Recipe, processes []db.Process) (err error){
	args := p.Called(recipe, processes)
	return args.Error(0)
}


func (p *PLCMockStore) DiscardTipAndHome(discard bool) (response string, err error) {
	args := p.Called(discard)
	return args.Get(0).(string), args.Error(1)
}

func (p *PLCMockStore) Heating(ht db.Heating) (response string, err error) {
	args := p.Called(ht)
	return args.Get(0).(string), args.Error(1)
}

func (p *PLCMockStore) Homing() (response string, err error) {
	args := p.Called()
	return args.Get(0).(string), args.Error(1)
}

func (p *PLCMockStore) ManualMovement(motorNum, direction, pulses uint16) (response string, err error) {
	args := p.Called(motorNum, direction, pulses)
	return args.Get(0).(string), args.Error(1)
}

func (p *PLCMockStore) Resume() (response string, err error) {
	args := p.Called()
	return args.Get(0).(string), args.Error(1)
}

func (p *PLCMockStore) Pause() (response string, err error) {
	args := p.Called()
	return args.Get(0).(string), args.Error(1)
}

func (p *PLCMockStore) Abort() (response string, err error) {
	args := p.Called()
	return args.Get(0).(string), args.Error(1)
}

func (p *PLCMockStore) Piercing(pi db.Piercing, cartridgeID int64) (response string, err error) {
	args := p.Called(pi, cartridgeID)
	return args.Get(0).(string), args.Error(1)
}

func (p *PLCMockStore) Shaking(shakerData db.Shaker) (response string, err error) {
	args := p.Called(shakerData)
	return args.Get(0).(string), args.Error(1)
}

func (p *PLCMockStore) TipDocking(td db.TipDock, cartridgeID int64) (response string, err error) {
	args := p.Called(td, cartridgeID)
	return args.Get(0).(string), args.Error(1)
}

func (p *PLCMockStore) TipOperation(to db.TipOperation) (response string, err error) {
	args := p.Called(to)
	return args.Get(0).(string), args.Error(1)
}
