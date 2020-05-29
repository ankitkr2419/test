package compact32

import (
	"errors"
	"fmt"
	"mylab/cpagent/plc"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

// Define the suite, and absorb the built-in basic suite
// functionality from testify - including assertion methods.
type Compact32DriverTestSuite struct {
	suite.Suite
	driver *MockCompact32Driver
}

func (suite *Compact32DriverTestSuite) SetupTest() {
	suite.driver = &MockCompact32Driver{}

	C32 = Compact32{
		Driver: suite.driver,
		ExitCh: make(chan error),
	}

}

func TestSuite(t *testing.T) {
	suite.Run(t, new(Compact32DriverTestSuite))
}

func (suite *Compact32DriverTestSuite) TestHeartBeatPLCFailure() {
	suite.driver.On("ReadSingleRegister", mock.Anything).Return(
		uint16(2), // PLC has NOT written 1
		nil,
	)

	go C32.HeartBeat()
	err := <-C32.ExitCh

	assert.Equal(suite.T(), err, errors.New("PLC is not responding and maybe dead. Abort!!"))
	suite.driver.AssertExpectations(suite.T())
}

func (suite *Compact32DriverTestSuite) TestHeartBeatSuccess() {
	suite.driver.On("ReadSingleRegister", mock.Anything).Return(
		uint16(1), // PLC has written 1.
		nil,
	)

	suite.driver.On("WriteSingleRegister", mock.Anything, mock.Anything).Return(
		[]byte{}, // Software  has written
		nil,
	)

	go C32.HeartBeat()

	for i := 0; i < 4; i++ {
		select {
		case err := <-C32.ExitCh:
			assert.FailNow(suite.T(), err.Error())
		default:
			time.Sleep(500 * time.Millisecond)

		}
	}

	// Heartbeat is supposed to hang!
	suite.driver.AssertExpectations(suite.T())
}

func (suite *Compact32DriverTestSuite) TestConfigureRunStaging() {
	suite.driver.On("WriteMultipleRegisters", mock.Anything, mock.Anything, mock.Anything).Return(
		[]byte{},
		nil,
	)

	p := plc.Stage{
		Holding: []plc.Step{
			plc.Step{65.3, 2.1, 5},
			plc.Step{85.3, 2.2, 3},
			plc.Step{95, 2, 5},
		},
	}

	//C32.ConfigureRun(p)
	writeStageData(HOLDING_STAGE, p)

	assert.Equal(suite.T(), suite.driver.LastAddress, MODBUS["D"][101])
	// Expected []Bytes: [2 141 3 85 3 182 0 0 0 21 0 22 0 20 0 0 0 5 0 3 0 5 0 0]
	//                    ----- ---- ----- --- ---- ---- ---- --- --- --- --- ---
	/*
		2 = 0x02, 141 = 0x8D, 0x028D = 653
		3 = 0x03, 85  = 0x55, 0x0355 = 853
		3 = 0x03, 182 = 0xB6, 0x03B6 = 950
	*/
	expected := fmt.Sprintf("%v", []byte{2, 141, 3, 85, 3, 182, 0, 0, 0, 21, 0, 22, 0, 20, 0, 0, 0, 5, 0, 3, 0, 5, 0, 0})
	assert.Equal(suite.T(), expected, fmt.Sprintf("%v", suite.driver.LastValue))
}

func (suite *Compact32DriverTestSuite) TestConfigureRunSuccess() {
	suite.driver.On("WriteMultipleRegisters", mock.Anything, mock.Anything, mock.Anything).Return(
		[]byte{},
		nil,
	)

	suite.driver.On("WriteSingleRegister", mock.Anything, mock.Anything).Return(
		[]byte{}, // Software  has written
		nil,
	)

	p := plc.Stage{
		Holding: []plc.Step{
			plc.Step{65.3, 2.1, 5},
			plc.Step{85.3, 2.2, 3},
			plc.Step{95, 2, 5},
		},
		Cycle: []plc.Step{
			plc.Step{55, 2, 5},
			plc.Step{65, 2, 5},
			plc.Step{75, 2, 5},
			plc.Step{85, 2, 5},
			plc.Step{95, 2, 5},
		},
		CycleCount: 45,
	}

	C32.ConfigureRun(p)

	assert.Equal(suite.T(), suite.driver.LastAddress, MODBUS["D"][113]) // cycling stage

	// Expected []Bytes: [[2 38 2 138 2 238 3 82 3 182 0 0 0 20 0 20 0 20 0 20 0 20 0 0 0 5 0 5 0 5 0 5 0 5 0 0]
	//                     ---- ----- ----- ---- ----- --- ---- ---- ---- ---- ---- --- --- --- --- --- --- ---
	/*
		2 = 0x02, 38  = 0x26, 0x0226 = 550
		2 = 0x02, 138 = 0x8A, 0x028A = 650
		2 = 0x02, 238 = 0xEE, 0x02EE = 750
		3 = 0x03, 82  = 0x52, 0x0352 = 850
	*/

	expected := fmt.Sprintf("%v", []byte{2, 38, 2, 138, 2, 238, 3, 82, 3, 182, 0, 0, 0, 20, 0, 20, 0, 20, 0, 20, 0, 20, 0, 0, 0, 5, 0, 5, 0, 5, 0, 5, 0, 5, 0, 0})
	assert.Equal(suite.T(), expected, fmt.Sprintf("%v", suite.driver.LastValue))

	suite.driver.AssertExpectations(suite.T())
}
