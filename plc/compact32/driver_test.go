package compact32

import (
	"errors"
	"testing"

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
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(Compact32DriverTestSuite))
}

func (suite *Compact32DriverTestSuite) TestHeartBeatSuccess() {

	suite.driver.On("ReadSingleRegister", mock.Anything).Return(
		uint16(2), // PLC has written
		nil,
	)

	suite.driver.On("WriteSingleRegister", mock.Anything, mock.Anything).Return(
		[]byte{}, // PLC has written
		nil,
	)

	C32 = Compact32{
		Driver: suite.driver,
		ExitCh: make(chan error),
	}

	go C32.HeartBeat()

	err := <-C32.ExitCh

	assert.Equal(suite.T(), err, errors.New("PLC is not responding and maybe dead. Abort!!"))

	suite.driver.AssertExpectations(suite.T())
}
