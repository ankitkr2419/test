package service

import (
	"errors"
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
type SampleHandlerTestSuite struct {
	suite.Suite

	dbMock *db.DBMockStore
}

func (suite *SampleHandlerTestSuite) SetupTest() {
	suite.dbMock = &db.DBMockStore{}
}

func TestSampleTestSuite(t *testing.T) {
	suite.Run(t, new(SampleHandlerTestSuite))
}

func (suite *SampleHandlerTestSuite) TestListSamplesSuccess() {
	testUUID := uuid.New()
	suite.dbMock.On("ListSamples", mock.Anything).Return(
		[]db.Sample{
			db.Sample{ID: testUUID, Name: "test sample"},
		},
		nil,
	)

	recorder := makeHTTPCall(
		http.MethodGet,
		"/samples",
		"/samples",
		"",
		listSamplesHandler(Dependencies{Store: suite.dbMock}),
	)
	output := fmt.Sprintf(`[{"id":"%s","name":"test sample"}]`, testUUID)
	assert.Equal(suite.T(), http.StatusOK, recorder.Code)
	assert.Equal(suite.T(), output, recorder.Body.String())
	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *SampleHandlerTestSuite) TestListSamplesFail() {
	suite.dbMock.On("ListSamples", mock.Anything).Return(
		[]db.Sample{},
		errors.New("error fetching samples"),
	)

	recorder := makeHTTPCall(
		http.MethodGet,
		"/samples",
		"/samples",
		"",
		listSamplesHandler(Dependencies{Store: suite.dbMock}),
	)

	assert.Equal(suite.T(), http.StatusInternalServerError, recorder.Code)
	assert.Equal(suite.T(), "", recorder.Body.String())
	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *SampleHandlerTestSuite) TestCreateSampleSuccess() {
	testUUID := uuid.New()
	suite.dbMock.On("CreateSample", mock.Anything, mock.Anything).Return(db.Sample{
		ID: testUUID, Name: "test sample",
	}, nil)

	body := fmt.Sprintf(`{"name":"test sample"}`)
	recorder := makeHTTPCall(http.MethodPost,
		"/samples",
		"/samples",
		body,
		createSampleHandler(Dependencies{Store: suite.dbMock}),
	)
	output := fmt.Sprintf(`{"id":"%s","name":"test sample"}`, testUUID)
	assert.Equal(suite.T(), http.StatusCreated, recorder.Code)
	assert.Equal(suite.T(), output, recorder.Body.String())

	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *SampleHandlerTestSuite) TestUpdateSampleSuccess() {
	testUUID := uuid.New()
	suite.dbMock.On("UpdateSample", mock.Anything, mock.Anything).Return(db.Sample{
		ID: testUUID, Name: "test sample",
	}, nil)

	body := fmt.Sprintf(`{"name":"test sample"}`)

	recorder := makeHTTPCall(http.MethodPut,
		"/samples/{id}",
		"/samples/"+testUUID.String(),
		body,
		updateSampleHandler(Dependencies{Store: suite.dbMock}),
	)

	assert.Equal(suite.T(), http.StatusOK, recorder.Code)
	assert.Equal(suite.T(), `{"msg":"Sample updated successfully"}`, recorder.Body.String())

	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *SampleHandlerTestSuite) TestDeleteSampleSuccess() {
	testUUID := uuid.New()
	suite.dbMock.On("DeleteSample", mock.Anything, mock.Anything).Return(
		testUUID,
		nil)

	recorder := makeHTTPCall(http.MethodDelete,
		"/samples/{id}",
		"/samples/"+testUUID.String(),
		"",
		deleteSampleHandler(Dependencies{Store: suite.dbMock}),
	)
	assert.Equal(suite.T(), http.StatusOK, recorder.Code)
	assert.Equal(suite.T(), `{"msg":"Sample deleted successfully"}`, recorder.Body.String())

	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *SampleHandlerTestSuite) TestShowSampleSuccess() {
	testUUID := uuid.New()
	suite.dbMock.On("ShowSample", mock.Anything, mock.Anything).Return(db.Sample{
		ID: testUUID, Name: "test sample",
	}, nil)

	recorder := makeHTTPCall(http.MethodGet,
		"/samples/{id}",
		"/samples/"+testUUID.String(),
		"",
		showSampleHandler(Dependencies{Store: suite.dbMock}),
	)
	output := fmt.Sprintf(`{"id":"%s","name":"test sample"}`, testUUID)
	assert.Equal(suite.T(), http.StatusOK, recorder.Code)
	assert.Equal(suite.T(), output, recorder.Body.String())

	suite.dbMock.AssertExpectations(suite.T())
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
