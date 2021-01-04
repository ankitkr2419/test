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
type LabwareHandlerTestSuite struct {
	suite.Suite

	dbMock *db.DBMockStore
}

func (suite *LabwareHandlerTestSuite) SetupTest() {
	suite.dbMock = &db.DBMockStore{}
}

func TestLabwareTestSuite(t *testing.T) {
	suite.Run(t, new(LabwareHandlerTestSuite))
}

func (suite *LabwareHandlerTestSuite) TestCreateLabwareSuccess() {
	suite.dbMock.On("InsertLabware", mock.Anything, mock.Anything).Return(db.Labware{
		ID: 1, Name: "Covid", Description: "Covid Labware",
	}, nil)

	body := fmt.Sprintf(`{"id":1,"name":"Covid","description":"Covid Labware"}`)

	recorder := makeHTTPCall(http.MethodPost,
		"/labware",
		"/labware",
		body,
		createLabwareHandler(Dependencies{Store: suite.dbMock}),
	)

	output := fmt.Sprintf(`{"id":1,"name":"Covid","description":"Covid Labware"}`)

	assert.Equal(suite.T(), http.StatusCreated, recorder.Code)
	assert.Equal(suite.T(), output, recorder.Body.String())

	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *LabwareHandlerTestSuite) TestCreateLabwareFailure() {
	suite.dbMock.On("InsertLabware", mock.Anything, mock.Anything).Return(db.Labware{
		ID: 1, Name: "Covid", Description: "Covid Labware",
	}, nil)

	body := fmt.Sprintf(`{"id":1,"name":"Dengui","description":"Covid Labware"}`)

	recorder := makeHTTPCall(http.MethodPost,
		"/labware",
		"/labware",
		body,
		createLabwareHandler(Dependencies{Store: suite.dbMock}),
	)

	output := fmt.Sprintf(`{"id":1,"name":"Covid","description":"Covid Labware"}`)

	assert.Equal(suite.T(), http.StatusCreated, recorder.Code)
	assert.NotEqual(suite.T(), output, recorder.Body.String())

	suite.dbMock.AssertExpectations(suite.T())
}
