package compact32

import (
	"errors"
	"fmt"
	"mylab/cpagent/plc"
	"testing"
	"time"

	logger "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

// Define the suite, and absorb the built-in basic suite
// functionality from testify - including assertion methods.
type Compact32DriverTestSuite struct {
	suite.Suite
	driver *plc.MockCompact32Driver
	C32    *Compact32
}

func (suite *Compact32DriverTestSuite) SetupTest() {
	logger.SetLevel(logger.TraceLevel)
	suite.driver = &plc.MockCompact32Driver{}

	suite.C32 = &Compact32{
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

	go suite.C32.HeartBeat()
	err := <-suite.C32.ExitCh

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

	go suite.C32.HeartBeat()

	for i := 0; i < 4; i++ {
		select {
		case err := <-suite.C32.ExitCh:
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
			plc.Step{TargetTemp: 65.3, RampUpTemp: 2.1, HoldTime: 5, DataCapture: false},
			plc.Step{TargetTemp: 85.3, RampUpTemp: 2.2, HoldTime: 3, DataCapture: false},
			plc.Step{TargetTemp: 95, RampUpTemp: 2, HoldTime: 5, DataCapture: false},
		},
	}

	suite.C32.writeStageData(HOLDING_STAGE, p)

	assert.Equal(suite.T(), suite.driver.LastAddress, plc.MODBUS["D"][101])
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
			plc.Step{TargetTemp: 65.3, RampUpTemp: 2.1, HoldTime: 5, DataCapture: false},
			plc.Step{TargetTemp: 85.3, RampUpTemp: 2.2, HoldTime: 3, DataCapture: false},
			plc.Step{TargetTemp: 95, RampUpTemp: 2, HoldTime: 5, DataCapture: false},
		},
		Cycle: []plc.Step{
			plc.Step{TargetTemp: 55, RampUpTemp: 2, HoldTime: 5, DataCapture: false},
			plc.Step{TargetTemp: 65, RampUpTemp: 2, HoldTime: 5, DataCapture: false},
			plc.Step{TargetTemp: 75, RampUpTemp: 2, HoldTime: 5, DataCapture: false},
			plc.Step{TargetTemp: 85, RampUpTemp: 2, HoldTime: 5, DataCapture: false},
			plc.Step{TargetTemp: 95, RampUpTemp: 2, HoldTime: 5, DataCapture: true},
		},
		CycleCount: 45,
	}

	suite.C32.ConfigureRun(p)

	assert.Equal(suite.T(), suite.driver.LastAddress, plc.MODBUS["D"][113]) // cycling stage

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

func (suite *Compact32DriverTestSuite) TestMonitorCycleNotComplete() {
	suite.driver.On("ReadSingleRegister", plc.MODBUS["D"][133]).Return(
		uint16(32), // cycle counter
		nil,
	)

	suite.driver.On("ReadSingleRegister", plc.MODBUS["D"][132]).Return(
		uint16(655), // PV Temperature
		nil,
	)
	suite.driver.On("ReadSingleRegister", plc.MODBUS["D"][135]).Return(
		uint16(953), // Lid Temperature
		nil,
	)
	suite.driver.On("ReadSingleCoil", plc.MODBUS["M"][107]).Return(
		uint16(0x00), // Cycle complete
		nil,
	)
	suite.driver.On("WriteSingleCoil", plc.MODBUS["M"][106], plc.OFF).Return(
		nil,
	)
	suite.driver.On("ReadHoldingRegisters", mock.Anything, mock.Anything).Return(
		func() []byte {
			// Return: [0 0 0 2 0 4 0 6 0 8 0 10 0 12 0 14 0 16 0 18 0 20 0 22 0 24 0 26 0 28 0 30 0 32 0 34 0 36 0 38 0 40 0 42 0 44 0 46 0 48 0 50 0 52 0 54 0 56 0 58 0 60 0 62 0 64 0 66 0 68 0 70 0 72 0 74 0 76 0 78 0 80 0 82 0 84 0 86 0 88 0 90 0 92 0 94 0 96 0 98 0 100 0 102 0 104 0 106 0 108 0 110 0 112 0 114 0 116 0 118 0 120 0 122 0 124 0 126 0 128 0 130 0 132 0 134 0 136 0 138 0 140 0 142 0 144 0 146 0 148 0 150 0 152 0 154 0 156 0 158 0 160 0 162 0 164 0 166 0 168 0 170 0 172 0 174 0 176 0 178 0 180 0 182 0 184 0 186 0 188 0 190]
			data := make([]byte, 192)

			for i := 0; i < 192; i += 2 {
				data[i] = 0
				data[i+1] = byte(i)
			}
			return data
		}(),
		nil,
	)

	scan, _ := suite.C32.Monitor(uint16(1))

	assert.Equal(suite.T(), uint16(32), scan.Cycle)
	assert.Equal(suite.T(), float32(65.5), scan.Temp)
	assert.Equal(suite.T(), float32(95.3), scan.LidTemp)

	// Emissions should NOT be populated!
	assert.Equal(suite.T(), scan.Wells[0], plc.Emissions{0, 0, 0, 0})
}

func (suite *Compact32DriverTestSuite) TestMonitorEmissionSuccess() {
	suite.driver.On("ReadSingleRegister", plc.MODBUS["D"][133]).Return(
		uint16(32), // cycle counter
		nil,
	)

	suite.driver.On("ReadSingleRegister", plc.MODBUS["D"][132]).Return(
		uint16(655), // PV Temperature
		nil,
	)
	suite.driver.On("ReadSingleRegister", plc.MODBUS["D"][135]).Return(
		uint16(953), // Lid Temperature
		nil,
	)
	suite.driver.On("ReadSingleCoil", plc.MODBUS["M"][107]).Return(
		uint16(0xFF00), // Cycle complete
		nil,
	)
	suite.driver.On("WriteSingleCoil", plc.MODBUS["M"][106], plc.OFF).Return(
		nil,
	)
	suite.driver.On("ReadHoldingRegisters", mock.Anything, mock.Anything).Return(
		func() []byte {
			// Return: [0 0 0 2 0 4 0 6 0 8 0 10 0 12 0 14 0 16 0 18 0 20 0 22 0 24 0 26 0 28 0 30 0 32 0 34 0 36 0 38 0 40 0 42 0 44 0 46 0 48 0 50 0 52 0 54 0 56 0 58 0 60 0 62 0 64 0 66 0 68 0 70 0 72 0 74 0 76 0 78 0 80 0 82 0 84 0 86 0 88 0 90 0 92 0 94 0 96 0 98 0 100 0 102 0 104 0 106 0 108 0 110 0 112 0 114 0 116 0 118 0 120 0 122 0 124 0 126 0 128 0 130 0 132 0 134 0 136 0 138 0 140 0 142 0 144 0 146 0 148 0 150 0 152 0 154 0 156 0 158 0 160 0 162 0 164 0 166 0 168 0 170 0 172 0 174 0 176 0 178 0 180 0 182 0 184 0 186 0 188 0 190]
			data := make([]byte, 192)

			for i := 0; i < 192; i += 2 {
				data[i] = 0
				data[i+1] = byte(i)
			}
			return data
		}(),
		nil,
	)

	scan, _ := suite.C32.Monitor(uint16(1))

	assert.Equal(suite.T(), uint16(32), scan.Cycle)
	assert.Equal(suite.T(), float32(65.5), scan.Temp)
	assert.Equal(suite.T(), float32(95.3), scan.LidTemp)

	assert.Equal(suite.T(), scan.Wells[0], plc.Emissions{0, 2, 4, 6})
}
