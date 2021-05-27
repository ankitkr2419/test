package service

import (
	"encoding/json"
	"mylab/cpagent/db"
	"mylab/cpagent/plc"
	"mylab/cpagent/responses"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// Define the suite, and absorb the built-in basic suite
// functionality from testify - including assertion methods.
type UpgradeHandlerTestSuite struct {
	suite.Suite
	dbMock *db.DBMockStore
	plcDeck map[string]plc.Extraction
}

func (suite *UpgradeHandlerTestSuite) SetupTest() {
	suite.dbMock = &db.DBMockStore{}
	driverA := &plc.PLCMockStore{}
	driverB := &plc.PLCMockStore{}
	suite.plcDeck = map[string]plc.Extraction{
		"A":driverA,
		"B":driverB,
	}
}

func TestUpgradeTestSuite(t *testing.T) {
	suite.Run(t, new(UpgradeHandlerTestSuite))
}
