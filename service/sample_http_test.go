package service

import (
	"fmt"
	"mylab/cpagent/db"
	"net/http"

	"github.com/google/uuid"
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

func (suite *SampleHandlerTestSuite) TestFindSamplesSuccess() {
	testUUID := uuid.New()
	suite.dbMock.On("FindSamples", mock.Anything, mock.Anything).Return(
		[]db.Sample{
			db.Sample{ID: testUUID, Name: "test sample"},
		},
		nil,
	)

	recorder := makeHTTPCall(
		http.MethodGet,
		"/samples/{text:[a-z]+}",
		"/samples/tes",
		"",
		findSamplesHandler(Dependencies{Store: suite.dbMock}),
	)
	output := fmt.Sprintf(`[{"id":"%s","name":"test sample"}]`, testUUID)
	assert.Equal(suite.T(), http.StatusOK, recorder.Code)
	assert.Equal(suite.T(), output, recorder.Body.String())
	suite.dbMock.AssertExpectations(suite.T())
}
