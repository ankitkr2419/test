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
type ConsumableDistanceHandlerTestSuite struct {
	suite.Suite

	dbMock *db.DBMockStore
}

func (suite *ConsumableDistanceHandlerTestSuite) SetupTest() {
	suite.dbMock = &db.DBMockStore{}
}

func TestConsumableDistanceTestSuite(t *testing.T) {
	suite.Run(t, new(ConsumableDistanceHandlerTestSuite))
}

func (suite *ConsumableDistanceHandlerTestSuite) TestCreateConsumableDistanceSuccess() {
	suite.dbMock.On("InsertConsumableDistance", mock.Anything, mock.Anything).Return(db.ConsumableDistance{
		ID: 1, Name: "deck_start", Distance: 1.11, Description: "deck start point",
	}, nil)

	body := fmt.Sprintf(`{"id":1,"name":"deck_start","distance":1.11,"description":"deck start point","created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z"}`)
	recorder := makeHTTPCall(http.MethodPost,
		"/consumabledistance",
		"/consumabledistance",
		body,
		createConsumableDistanceHandler(Dependencies{Store: suite.dbMock}),
	)

	output := fmt.Sprintf(`{"id":1,"name":"deck_start","distance":1.11,"description":"deck start point","created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z"}`)

	assert.Equal(suite.T(), http.StatusCreated, recorder.Code)
	assert.Equal(suite.T(), output, recorder.Body.String())

	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *ConsumableDistanceHandlerTestSuite) TestCreateConsumableDistanceFailure() {
	suite.dbMock.On("InsertConsumableDistance", mock.Anything, mock.Anything).Return(db.ConsumableDistance{
		ID: 1, Name: "deck_start", Distance: 1.11, Description: "deck start point",
	}, nil)

	body := fmt.Sprintf(`{"id":1,"name":"deck_start","distance":1.11,"description":"deck start point","created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z"}`)
	recorder := makeHTTPCall(http.MethodPost,
		"/consumabledistance",
		"/consumabledistance",
		body,
		createConsumableDistanceHandler(Dependencies{Store: suite.dbMock}),
	)

	output := fmt.Sprintf(`{"id":1,"name":"deck_start","distance":2.11,"description":"deck start point","created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z"}`)

	assert.Equal(suite.T(), http.StatusCreated, recorder.Code)
	assert.NotEqual(suite.T(), output, recorder.Body.String())

	suite.dbMock.AssertExpectations(suite.T())
}
