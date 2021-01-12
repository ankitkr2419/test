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
type MotorHandlerTestSuite struct {
	suite.Suite

	dbMock *db.DBMockStore
}

func (suite *MotorHandlerTestSuite) SetupTest() {
	suite.dbMock = &db.DBMockStore{}
}

func TestMotorTestSuite(t *testing.T) {
	suite.Run(t, new(MotorHandlerTestSuite))
}

func (suite *MotorHandlerTestSuite) TestCreateMotorSuccess() {
	suite.dbMock.On("InsertMotor", mock.Anything, mock.Anything).Return(db.Motor{
		Number: 1, Name: "TestMotor", Ramp: 400, Steps: 500, Slow: 500, Fast: 2000,
	}, nil)

	body := fmt.Sprintf(`{"number":1, "name": "TestMotor", "ramp": 400, "steps": 500, "slow": 500, "fast": 2000, "created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z"}`)
	recorder := makeHTTPCall(http.MethodPost,
		"/motor",
		"/motor",
		body,
		createMotorHandler(Dependencies{Store: suite.dbMock}),
	)
	output := fmt.Sprintf(`{"number":1,"name":"TestMotor","ramp":400,"steps":500,"slow":500,"fast":2000,"created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z"}`)

	assert.Equal(suite.T(), http.StatusCreated, recorder.Code)
	assert.Equal(suite.T(), output, recorder.Body.String())

	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *MotorHandlerTestSuite) TestCreateMotorFailure() {
	suite.dbMock.On("InsertMotor", mock.Anything, mock.Anything).Return(db.Motor{
		Number: 1, Name: "TestMotor", Ramp: 400, Steps: 500, Slow: 500, Fast: 2000,
	}, nil)

	body := fmt.Sprintf(`{"number":1, "name": "TestMotor", "ramp": 500, "steps": 500, "slow": 500, "fast": 2000, "created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z"}`)
	recorder := makeHTTPCall(http.MethodPost,
		"/motor",
		"/motor",
		body,
		createMotorHandler(Dependencies{Store: suite.dbMock}),
	)
	output := fmt.Sprintf(`{"number":1,"name":"TestMotor","ramp":400,"steps":500,"slow":500,"fast":2000,"created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z"}`)

	assert.Equal(suite.T(), http.StatusCreated, recorder.Code)
	assert.NotEqual(suite.T(), output, recorder.Body.String())

	suite.dbMock.AssertExpectations(suite.T())
}
