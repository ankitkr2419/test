package plc

import (
	"mylab/cpagent/db"

	"github.com/stretchr/testify/mock"
)

type PLCMockStore struct {
	mock.Mock
}

type MockCompact32Driver struct {
	mock.Mock
	ExitCh chan error

	LastAddress  uint16
	LastQuantity uint16
	LastValue    []byte
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

func (p *PLCMockStore) SetCurrentProcessNumber(step int64) {
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

func (m *MockCompact32Driver) WriteSingleRegister(address, value uint16) (results []byte, err error) {
	args := m.Called(address, value)
	return args.Get(0).([]byte), args.Error(1)
}

func (m *MockCompact32Driver) WriteMultipleRegisters(address, quantity uint16, value []byte) (results []byte, err error) {
	m.LastAddress = address
	m.LastQuantity = quantity
	m.LastValue = value

	args := m.Called(address, quantity, value)
	return args.Get(0).([]byte), args.Error(1)
}

func (m *MockCompact32Driver) ReadCoils(address, quantity uint16) (results []byte, err error) {
	args := m.Called(address, quantity)
	return args.Get(0).([]byte), args.Error(1)
}

func (m *MockCompact32Driver) ReadSingleCoil(address uint16) (value uint16, err error) {
	args := m.Called(address)
	return args.Get(0).(uint16), args.Error(1)
}

func (m *MockCompact32Driver) ReadHoldingRegisters(address, quantity uint16) (results []byte, err error) {
	args := m.Called(address, quantity)
	return args.Get(0).([]byte), args.Error(1)
}

func (m *MockCompact32Driver) ReadSingleRegister(address uint16) (value uint16, err error) {
	args := m.Called(address)
	return args.Get(0).(uint16), args.Error(1)
}
func (m *MockCompact32Driver) WriteSingleCoil(address, value uint16) (err error) {
	args := m.Called(address, value)
	return args.Error(0)
}
