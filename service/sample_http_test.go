package service

import (
	"encoding/json"
	"mylab/cpagent/db"
	"net/http"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

// Define the suite, and absorb the built-in basic suite
// functionality from testify - including assertion methods.
type SampleHandlerTestSuite struct {
	suite.Suite

	dbMock *db.DBMockStore
}

func (suite *SampleHandlerTestSuite) SetupTest() {
	suite.dbMock = &db.DBMockStore{}
}

var testSampleObj = db.Sample{ID: testUUID, Name: "test sample"}

func (suite *SampleHandlerTestSuite) TestFindSamplesSuccess() {

	suite.dbMock.On("FindSamples", mock.Anything, mock.Anything).Return(
		[]db.Sample{testSampleObj}, nil)

	recorder := makeHTTPCall(
		http.MethodGet,
		"/samples/{text:[a-z]+}",
		"/samples/tes",
		"",
		findSamplesHandler(Dependencies{Store: suite.dbMock}),
	)
	output, _ := json.Marshal([]db.Sample{testSampleObj})
	assert.Equal(suite.T(), http.StatusOK, recorder.Code)
	assert.Equal(suite.T(), string(output), recorder.Body.String())
	suite.dbMock.AssertExpectations(suite.T())
}
