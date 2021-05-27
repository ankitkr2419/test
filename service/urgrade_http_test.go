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

func (suite *UpgradeHandlerTestSuite) TestSafeToUpgradeSuccess() {
	// Deck A
	suite.plcDeck[deckA].(*plc.PLCMockStore).On("IsRunInProgress").Return(false).Once()
	suite.plcDeck[deckA].(*plc.PLCMockStore).On("SetRunInProgress").Return().Maybe()
	suite.plcDeck[deckA].(*plc.PLCMockStore).On("ResetRunInProgress").Return().Maybe()
	// Deck B
	suite.plcDeck[deckB].(*plc.PLCMockStore).On("IsRunInProgress").Return(false).Once()
	suite.plcDeck[deckB].(*plc.PLCMockStore).On("SetRunInProgress").Return().Maybe()
	suite.plcDeck[deckB].(*plc.PLCMockStore).On("ResetRunInProgress").Return().Maybe()

	recorder := makeHTTPCall(http.MethodGet,
		"/safe-to-upgrade",
		"/safe-to-upgrade",
		"",
		safeToUpgradeHandler(Dependencies{Store: suite.dbMock, PlcDeck: suite.plcDeck}),
	)

	output, _ := json.Marshal(MsgObj{Msg: responses.SafeToUpgrade})

	assert.Equal(suite.T(), http.StatusOK, recorder.Code)
	assert.Equal(suite.T(), string(output), recorder.Body.String())

	suite.dbMock.AssertExpectations(suite.T())
}
