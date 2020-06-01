package compact32

import (
	"github.com/stretchr/testify/mock"
)

type MockCompact32Driver struct {
	mock.Mock
	ExitCh chan error

	LastAddress  uint16
	LastQuantity uint16
	LastValue    []byte
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
