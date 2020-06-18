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
type ExperimentHandlerTestSuite struct {
	suite.Suite

	dbMock *db.DBMockStore
}

func (suite *ExperimentHandlerTestSuite) SetupTest() {
	suite.dbMock = &db.DBMockStore{}
}

func TestExperimentTestSuite(t *testing.T) {
	suite.Run(t, new(ExperimentHandlerTestSuite))
}

func (suite *ExperimentHandlerTestSuite) TestListExperimentsSuccess() {
	testUUID := uuid.New()
	tempUUID := uuid.New()
	suite.dbMock.On("ListExperiments", mock.Anything).Return(
		[]db.Experiment{
			db.Experiment{ID: testUUID, Description: "blah blah", TemplateID: tempUUID, OperatorName: "ABC"},
		},
		nil,
	)

	recorder := makeHTTPCall(
		http.MethodGet,
		"/experiments",
		"/experiments",
		"",
		listExperimentHandler(Dependencies{Store: suite.dbMock}),
	)
	output := fmt.Sprintf(`[{"id":"%s","description":"blah blah","template_id":"%s","operator_name":"ABC","start_time":"0001-01-01T00:00:00Z","end_time":"0001-01-01T00:00:00Z"}]`, testUUID, tempUUID)
	assert.Equal(suite.T(), http.StatusOK, recorder.Code)
	assert.Equal(suite.T(), output, recorder.Body.String())
	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *ExperimentHandlerTestSuite) TestListExperimentsFail() {
	suite.dbMock.On("ListExperiments", mock.Anything).Return(
		[]db.Experiment{},
		errors.New("error fetching experiments"),
	)

	recorder := makeHTTPCall(
		http.MethodGet,
		"/experiments",
		"/experiments",
		"",
		listExperimentHandler(Dependencies{Store: suite.dbMock}),
	)

	assert.Equal(suite.T(), http.StatusInternalServerError, recorder.Code)
	assert.Equal(suite.T(), "", recorder.Body.String())
	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *ExperimentHandlerTestSuite) TestCreateExperimentSuccess() {
	testUUID := uuid.New()
	tempUUID := uuid.New()
	targetUUID := uuid.New()
	suite.dbMock.On("UpsertExpTemplateTarget", mock.Anything, mock.Anything, mock.Anything).Return(
		[]db.ExpTemplateTarget{
			db.ExpTemplateTarget{ExperimentID: testUUID, TemplateID: tempUUID, TargetID: targetUUID, Threshold: 10.5},
		},
		nil,
	)
	suite.dbMock.On("ListTemplateTargets", mock.Anything, mock.Anything).Return(
		[]db.TemplateTarget{
			db.TemplateTarget{TemplateID: tempUUID, TargetID: targetUUID, Threshold: 10.5},
		},
		nil,
	)
	suite.dbMock.On("CreateExperiment", mock.Anything, mock.Anything).Return(db.Experiment{
		ID: testUUID, Description: "blah blah", TemplateID: tempUUID, OperatorName: "ABC",
	}, nil)

	body := fmt.Sprintf(`{"description":"blah blah","template_id":"%s","operator_name":"ABC","start_time":"0001-01-01T00:00:00Z","end_time":"0001-01-01T00:00:00Z"}`, tempUUID)

	recorder := makeHTTPCall(http.MethodPost,
		"/experiments",
		"/experiments",
		body,
		createExperimentHandler(Dependencies{Store: suite.dbMock}),
	)

	output := fmt.Sprintf(`{"id":"%s","description":"blah blah","template_id":"%s","operator_name":"ABC","start_time":"0001-01-01T00:00:00Z","end_time":"0001-01-01T00:00:00Z"}`, testUUID, tempUUID)
	assert.Equal(suite.T(), http.StatusCreated, recorder.Code)
	assert.Equal(suite.T(), output, recorder.Body.String())

	suite.dbMock.AssertExpectations(suite.T())
}
func (suite *ExperimentHandlerTestSuite) TestShowExperimentSuccess() {
	testUUID := uuid.New()
	tempUUID := uuid.New()
	suite.dbMock.On("ShowExperiment", mock.Anything, mock.Anything).Return(db.Experiment{
		ID: testUUID, Description: "blah blah", TemplateID: tempUUID, OperatorName: "ABC",
	}, nil)

	recorder := makeHTTPCall(http.MethodGet,
		"/experiments/{id}",
		"/experiments/"+testUUID.String(),
		"",
		showExperimentHandler(Dependencies{Store: suite.dbMock}),
	)
	output := fmt.Sprintf(`{"id":"%s","description":"blah blah","template_id":"%s","operator_name":"ABC","start_time":"0001-01-01T00:00:00Z","end_time":"0001-01-01T00:00:00Z"}`, testUUID, tempUUID)
	assert.Equal(suite.T(), http.StatusOK, recorder.Code)
	assert.Equal(suite.T(), output, recorder.Body.String())

	suite.dbMock.AssertExpectations(suite.T())
}
