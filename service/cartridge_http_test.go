package service

import (
	"fmt"
	"mylab/cpagent/db"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

// Define the suite, and absorb the built-in basic suite
// functionality from testify - including assertion methods.
type CartridgeHandlerTestSuite struct {
	suite.Suite

	dbMock *db.DBMockStore
}

func (suite *CartridgeHandlerTestSuite) SetupTest() {
	suite.dbMock = &db.DBMockStore{}
}

func TestCartridgeTestSuite(t *testing.T) {
	suite.Run(t, new(CartridgeHandlerTestSuite))
}

func (suite *CartridgeHandlerTestSuite) TestCreateCartridgeSuccess() {
	suite.dbMock.On("InsertCartridge", mock.Anything, mock.Anything).Return(db.Cartridge{
		ID: 1, LabwareID: 1, Type: "extraction", Description: "extraction cartridge", WellNum: 1, Distance: 17.5, Height: 2.0, Volume: 10.0,
	}, nil)

	body := fmt.Sprintf(`{"id":1,"labware_id":1,"type":"extraction","description":"extraction cartridge","wells_num":1,"distance":17.5,"height":2,"volume":10,"created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z"}`)
	recorder := makeHTTPCall(http.MethodPost,
		"/cartridge",
		"/cartridge",
		body,
		createCartridgeHandler(Dependencies{Store: suite.dbMock}),
	)

	output := fmt.Sprintf(`{"id":1,"labware_id":1,"type":"extraction","description":"extraction cartridge","wells_num":1,"distance":17.5,"height":2,"volume":10,"created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z"}`)

	assert.Equal(suite.T(), http.StatusCreated, recorder.Code)
	assert.Equal(suite.T(), output, recorder.Body.String())

	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *CartridgeHandlerTestSuite) TestCreateCartridgeFailure() {
	suite.dbMock.On("InsertCartridge", mock.Anything, mock.Anything).Return(db.Cartridge{
		ID: 1, LabwareID: 1, Type: "extraction", Description: "extraction cartridge", WellNum: 1, Distance: 17.5, Height: 2, Volume: 10,
	}, nil)

	body := fmt.Sprintf(`{"id":1,"labware_id":1,"type":"extraction","description":"extraction cartridge","wells_num":1,"distance":17.5,"height":2,"volume":10,"created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z"}`)
	recorder := makeHTTPCall(http.MethodPost,
		"/cartridge",
		"/cartridge",
		body,
		createCartridgeHandler(Dependencies{Store: suite.dbMock}),
	)

	output := fmt.Sprintf(`{"id":1,"labware_id":1,"type":"pcr","description":"extraction cartridge","wells_num":1,"distance":17.5,"height":2,"volume":10,"created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z"}`)

	assert.Equal(suite.T(), http.StatusCreated, recorder.Code)
	assert.NotEqual(suite.T(), output, recorder.Body.String())

	suite.dbMock.AssertExpectations(suite.T())
}
