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
type StageHandlerTestSuite struct {
	suite.Suite

	dbMock *db.DBMockStore
}

func (suite *StageHandlerTestSuite) SetupTest() {
	suite.dbMock = &db.DBMockStore{}
}

func TestStageTestSuite(t *testing.T) {
	suite.Run(t, new(StageHandlerTestSuite))
}

func (suite *StageHandlerTestSuite) TestListStagesSuccess() {
	testUUID := uuid.New()
	tempUUID := uuid.New()
	suite.dbMock.On("ListStages", mock.Anything, mock.Anything).Return(
		[]db.Stage{
			db.Stage{Name: "test-stage", ID: testUUID, Type: "Repeat", RepeatCount: 3, TemplateID: tempUUID},
		},
		nil,
	)

	recorder := makeHTTPCall(
		http.MethodGet,
		"/stages/{template_id}",
		"/stages/"+tempUUID.String(),
		"",
		listStagesHandler(Dependencies{Store: suite.dbMock}),
	)
	output := fmt.Sprintf(`[{"id":"%s","name":"test-stage","type":"Repeat","repeat_count":3,"template_id":"%s"}]`, testUUID, tempUUID)
	assert.Equal(suite.T(), http.StatusOK, recorder.Code)
	assert.Equal(suite.T(), output, recorder.Body.String())
	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *StageHandlerTestSuite) TestListStagesFail() {
	tempUUID := uuid.New()
	suite.dbMock.On("ListStages", mock.Anything, mock.Anything).Return(
		[]db.Stage{},
		errors.New("error fetching templates"),
	)

	recorder := makeHTTPCall(
		http.MethodGet,
		"/stages/{template_id}",
		"/stages/"+tempUUID.String(),
		"",
		listStagesHandler(Dependencies{Store: suite.dbMock}),
	)

	assert.Equal(suite.T(), http.StatusInternalServerError, recorder.Code)
	assert.Equal(suite.T(), "", recorder.Body.String())
	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *StageHandlerTestSuite) TestCreateStageSuccess() {
	testUUID := uuid.New()
	tempUUID := uuid.New()
	suite.dbMock.On("CreateStage", mock.Anything, mock.Anything).Return(db.Stage{
		Name: "test-stage", ID: testUUID, Type: "Repeat", RepeatCount: 3, TemplateID: tempUUID,
	}, nil)

	body := fmt.Sprintf(`{"name":"test-stage","type":"Repeat", "repeat_count":3, "template_id":"%s"}`, tempUUID)
	recorder := makeHTTPCall(http.MethodPost,
		"/stage",
		"/stage",
		body,
		createStageHandler(Dependencies{Store: suite.dbMock}),
	)
	output := fmt.Sprintf(`{"id":"%s","name":"test-stage","type":"Repeat","repeat_count":3,"template_id":"%s"}`, testUUID, tempUUID)
	assert.Equal(suite.T(), http.StatusCreated, recorder.Code)
	assert.Equal(suite.T(), output, recorder.Body.String())

	suite.dbMock.AssertExpectations(suite.T())
}
func (suite *StageHandlerTestSuite) TestUpdateStageSuccess() {
	testUUID := uuid.New()
	tempUUID := uuid.New()
	suite.dbMock.On("UpdateStage", mock.Anything, mock.Anything).Return(db.Stage{
		Name: "test-stage", ID: testUUID, Type: "Repeat", RepeatCount: 3, TemplateID: tempUUID,
	}, nil)

	body := fmt.Sprintf(`{"name":"test-stage","type":"Repeat", "repeat_count":3, "template_id":"%s"}`, tempUUID)

	recorder := makeHTTPCall(http.MethodPut,
		"/stage/{id}",
		"/stage/"+testUUID.String(),
		body,
		updateStageHandler(Dependencies{Store: suite.dbMock}),
	)

	assert.Equal(suite.T(), http.StatusOK, recorder.Code)
	assert.Equal(suite.T(), "stage updated successfully", recorder.Body.String())

	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *StageHandlerTestSuite) TestDeleteStageSuccess() {
	testUUID := uuid.New()
	suite.dbMock.On("DeleteStage", mock.Anything, mock.Anything).Return(
		testUUID,
		nil)

	recorder := makeHTTPCall(http.MethodDelete,
		"/stage/{id}",
		"/stage/"+testUUID.String(),
		"",
		deleteStageHandler(Dependencies{Store: suite.dbMock}),
	)
	assert.Equal(suite.T(), http.StatusOK, recorder.Code)
	assert.Equal(suite.T(), "", recorder.Body.String())

	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *StageHandlerTestSuite) TestShowStageSuccess() {
	testUUID := uuid.New()
	tempUUID := uuid.New()
	suite.dbMock.On("ShowStage", mock.Anything, mock.Anything).Return(db.Stage{
		Name: "test-stage", ID: testUUID, Type: "Repeat", RepeatCount: 3, TemplateID: tempUUID,
	}, nil)

	recorder := makeHTTPCall(http.MethodGet,
		"/stage/{id}",
		"/stage/"+testUUID.String(),
		"",
		showStageHandler(Dependencies{Store: suite.dbMock}),
	)
	output := fmt.Sprintf(`{"id":"%s","name":"test-stage","type":"Repeat","repeat_count":3,"template_id":"%s"}`, testUUID, tempUUID)
	assert.Equal(suite.T(), http.StatusOK, recorder.Code)
	assert.Equal(suite.T(), output, recorder.Body.String())

	suite.dbMock.AssertExpectations(suite.T())
}
