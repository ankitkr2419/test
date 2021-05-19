package plc

import (
	
	"github.com/stretchr/testify/mock"
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