package service

import (
	"encoding/json"
	"mylab/cpagent/db"
	"mylab/cpagent/plc"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type HomingHandlerTestSuite struct {
	suite.Suite
	dbMock  *db.DBMockStore
	plcDeck map[string]plc.Extraction
}

func (suite *HomingHandlerTestSuite) SetupTest() {
	suite.dbMock = &db.DBMockStore{}
	driverA := &plc.PLCMockStore{}
	driverB := &plc.PLCMockStore{}
	suite.plcDeck = map[string]plc.Extraction{
		"A": driverA,
		"B": driverB,
	}
	suite.dbMock.On("AddAuditLog", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil).Maybe()

}

func TestHomingTestSuite(t *testing.T) {
	loadUtils()
	suite.Run(t, new(HomingHandlerTestSuite))
}

func (suite *HomingHandlerTestSuite) TestHomingSuccess() {

	deck := deckB

	suite.plcDeck[deck].(*plc.PLCMockStore).On("Homing").Return("homing success", nil).Maybe()

	recorder := makeHTTPCall(http.MethodGet,
		"/homing/{deck:[A-B]?}",
		"/homing/"+deck,
		"",
		homingHandler(Dependencies{Store: suite.dbMock, PlcDeck: suite.plcDeck}),
	)

	msg := MsgObj{Msg: `homing in progress for single deck`, Deck: deck}

	output, _ := json.Marshal(msg)

	assert.Equal(suite.T(), http.StatusOK, recorder.Code)
	assert.Equal(suite.T(), string(output), recorder.Body.String())

	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *HomingHandlerTestSuite) TestHomingFailure() {

	deck := deckB

	suite.plcDeck[deck].(*plc.PLCMockStore).On("Homing").Return("homing success", nil).Maybe()

	recorder := makeHTTPCall(http.MethodGet,
		"/homing/{deck:[A-B]?}",
		"/homing/"+invalidDeck,
		"",
		homingHandler(Dependencies{Store: suite.dbMock, PlcDeck: suite.plcDeck}),
	)

	msg := "404 page not found\n"

	assert.Equal(suite.T(), http.StatusNotFound, recorder.Code)
	assert.Equal(suite.T(), msg, recorder.Body.String())

	suite.dbMock.AssertExpectations(suite.T())
}
