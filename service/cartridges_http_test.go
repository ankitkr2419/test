package service

import (
	"encoding/json"
	"errors"
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
type CartridgesHandlerTestSuite struct {
	suite.Suite

	dbMock *db.DBMockStore
}

func (suite *CartridgesHandlerTestSuite) SetupTest() {
	suite.dbMock = &db.DBMockStore{}
}

func TestCartridgesTestSuite(t *testing.T) {
	suite.Run(t, new(CartridgesHandlerTestSuite))
}

var testCartridgeObj = db.Cartridge{
	ID:          1,
	Type:        "test_car1",
	Description: "test",
}

func (suite *CartridgesHandlerTestSuite) TestListCartridgesSuccess() {
	suite.dbMock.On("ListCartridges", mock.Anything, mock.Anything).Return([]db.Cartridge{testCartridgeObj}, nil)
	suite.dbMock.On("AddAuditLog", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil).Maybe()

	recorder := makeHTTPCall(http.MethodPost,
		"/cartridges/",
		"/cartridges/",
		"",
		listCartridgesHandler(Dependencies{Store: suite.dbMock}),
	)
	body, _ := json.Marshal([]db.Cartridge{testCartridgeObj})
	assert.Equal(suite.T(), http.StatusOK, recorder.Code)
	assert.Equal(suite.T(), string(body), recorder.Body.String())

	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *CartridgesHandlerTestSuite) TestListCartridgesFailure() {
	suite.dbMock.On("ListCartridges", mock.Anything, mock.Anything).Return([]db.Cartridge{}, errors.New("failed to fetch cartridge"))
	suite.dbMock.On("AddAuditLog", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil).Maybe()

	recorder := makeHTTPCall(http.MethodPost,
		"/cartridges/",
		"/cartridges/",
		"",
		listCartridgesHandler(Dependencies{Store: suite.dbMock}),
	)
	output := fmt.Sprintf(`{"err":"error fetching cartridge record"}`)

	assert.Equal(suite.T(), http.StatusInternalServerError, recorder.Code)
	assert.Equal(suite.T(), string(output), recorder.Body.String())

	suite.dbMock.AssertExpectations(suite.T())
}
