package tec_1089

import (

	"testing"
	"mylab/cpagent/tec"
	"mylab/cpagent/plc"

	logger "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/suite"
)

// Define the suite, and absorb the built-in basic suite
// functionality from testify - including assertion methods.
type TECTestSuite struct {
	suite.Suite
	driver tec.Driver
}

func (suite *TECTestSuite) SetupTest() {
	logger.SetLevel(logger.TraceLevel)

	exit := make(chan error)
	websocketMsg := make(chan string)
	websocketErr := make(chan error)
	test := false
	var driver plc.Driver 

	suite.driver = NewTEC1089Driver(websocketMsg, websocketErr, exit, test, driver)
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(TECTestSuite))
}

// NOTE: Currently we lack knowledge as how to write test cases for this scenario
func (suite *TECTestSuite) TestInitiateTEC() {
	
}
