package service

import (
	"fmt"
	"mylab/cpagent/db"
	"net/http"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

// Define the suite, and absorb the built-in basic suite
// functionality from testify - including assertion methods.

type ExperimentTemperatureHandlerTestSuite struct {
	suite.Suite
	dbMock *db.DBMockStore
}

func (suite *ExperimentTemperatureHandlerTestSuite) SetupTest() {
	suite.dbMock = &db.DBMockStore{}
}

func TestExperimentTemperatureTestSuite(t *testing.T) {
	suite.Run(t, new(ExperimentTemperatureHandlerTestSuite))
}

func (suite *ExperimentTemperatureHandlerTestSuite) TestListExperimentTemperaturesSuccess() {
	experimentID := uuid.New()
	suite.dbMock.On("ListExperimentTemperature", mock.Anything, mock.Anything).Return([]db.ExperimentTemperature{
		{ExperimentID: experimentID, Temp: 100, LidTemp: 30, Cycle: 1},
	}, nil)

	recorder := makeHTTPCall(
		http.MethodGet,
		"/experiments/{id}/temperature",
		"/experiments/"+experimentID.String()+"/temperature",
		"",
		getTemperatureHandler(Dependencies{Store: suite.dbMock}),
	)
	output := fmt.Sprintf(`{"type":"Temperature","data":[{"experiment_id":"%s","temp":100,"lid_temp":30,"cycle":1,"created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z"}]}`, experimentID)
	assert.Equal(suite.T(), http.StatusOK, recorder.Code)
	assert.Equal(suite.T(), output, recorder.Body.String())
	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *ExperimentTemperatureHandlerTestSuite) TestListExperimentTemperaturesFail() {
	experimentID := uuid.New()
	suite.dbMock.On("ListExperimentTemperature", mock.Anything, mock.Anything).Return([]db.ExperimentTemperature{}, nil)

	recorder := makeHTTPCall(
		http.MethodGet,
		"/experiments/{id}/temperature",
		"/experiments/"+experimentID.String()+"/temperature",
		"",
		getTemperatureHandler(Dependencies{Store: suite.dbMock}),
	)
	output := fmt.Sprintf(`{"type":"Temperature","data":[]}`)
	assert.Equal(suite.T(), http.StatusOK, recorder.Code)
	assert.Equal(suite.T(), output, recorder.Body.String())
	suite.dbMock.AssertExpectations(suite.T())
}
