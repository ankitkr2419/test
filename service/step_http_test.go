package service

import (
	"encoding/json"
	"errors"
	"mylab/cpagent/config"
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
type StepHandlerTestSuite struct {
	suite.Suite

	dbMock *db.DBMockStore
}

func (suite *StepHandlerTestSuite) SetupTest() {
	suite.dbMock = &db.DBMockStore{}
	config.SetRoomTemp(27)
	config.SetHomingTime(16)
	suite.dbMock.On("ShowStage", mock.Anything, mock.Anything).Return(testStageObj, nil).Maybe()

	suite.dbMock.On("ShowTemplate", mock.Anything, mock.Anything).Return(db.Template{
		ID:          testTemplateID,
		Name:        "test template",
		Description: "blah blah",
		Publish:     false,
	}, nil).Maybe()
	suite.dbMock.On("ListStages", mock.Anything, mock.Anything).Return(
		[]db.Stage{testStageObj},
		nil,
	).Maybe()

	suite.dbMock.On("UpdateEstimatedTime", mock.Anything, mock.Anything, mock.Anything).Return(nil).Maybe()

}

func TestStepTestSuite(t *testing.T) {
	suite.Run(t, new(StepHandlerTestSuite))
}

var stepUUID = uuid.New()
var testStepObj = db.Step{
	ID:                stepUUID,
	TargetTemperature: 25.5,
	RampRate:          5.5,
	HoldTime:          120,
	DataCapture:       true,
	StageID:           testUUID}

func (suite *StepHandlerTestSuite) TestListStepsSuccess() {

	suite.dbMock.On("ListSteps", mock.Anything, mock.Anything).Return(
		[]db.Step{testStepObj}, nil)

	recorder := makeHTTPCall(
		http.MethodGet,
		"/stages/{stage_id}/steps",
		"/stages/"+testUUID.String()+"/steps",
		"",
		listStepsHandler(Dependencies{Store: suite.dbMock}),
	)
	output, _ := json.Marshal([]db.Step{testStepObj})
	assert.Equal(suite.T(), http.StatusOK, recorder.Code)
	assert.Equal(suite.T(), string(output), recorder.Body.String())
	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *StepHandlerTestSuite) TestListStepsFail() {

	suite.dbMock.On("ListSteps", mock.Anything, mock.Anything).Return([]db.Step{}, errors.New("error fetching steps"))

	recorder := makeHTTPCall(
		http.MethodGet,
		"/stages/{stage_id}/steps",
		"/stages/"+testUUID.String()+"/steps",
		"",
		listStepsHandler(Dependencies{Store: suite.dbMock}),
	)

	assert.Equal(suite.T(), http.StatusInternalServerError, recorder.Code)
	assert.Equal(suite.T(), "", recorder.Body.String())
	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *StepHandlerTestSuite) TestCreateStepSuccess() {
	suite.dbMock.On("UpdateStepCount", mock.Anything).Return(nil, nil)

	suite.dbMock.On("CreateStep", mock.Anything, mock.Anything).Return(testStepObj, nil)
	suite.dbMock.On("ListSteps", mock.Anything, mock.Anything).Return(
		[]db.Step{testStepObj}, nil).Maybe()
	body, _ := json.Marshal(testStepObj)
	recorder := makeHTTPCall(http.MethodPost,
		"/steps",
		"/steps",
		string(body),
		createStepHandler(Dependencies{Store: suite.dbMock}),
	)
	assert.Equal(suite.T(), http.StatusCreated, recorder.Code)
	assert.Equal(suite.T(), string(body), recorder.Body.String())

	suite.dbMock.AssertExpectations(suite.T())
}
func (suite *StepHandlerTestSuite) TestUpdateStepSuccess() {

	suite.dbMock.On("UpdateStep", mock.Anything, mock.Anything).Return(nil)
	suite.dbMock.On("ListSteps", mock.Anything, mock.Anything).Return(
		[]db.Step{testStepObj}, nil).Maybe()
	body, _ := json.Marshal(testStepObj)

	recorder := makeHTTPCall(http.MethodPut,
		"/steps/{id}",
		"/steps/"+testUUID.String(),
		string(body),
		updateStepHandler(Dependencies{Store: suite.dbMock}),
	)

	assert.Equal(suite.T(), http.StatusOK, recorder.Code)
	assert.Equal(suite.T(), `{"msg":"step updated successfully"}`, recorder.Body.String())

	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *StepHandlerTestSuite) TestDeleteStepSuccess() {

	suite.dbMock.On("UpdateStepCount", mock.Anything).Return(nil, nil)
	suite.dbMock.On("DeleteStep", mock.Anything, testUUID).Return(nil)
	suite.dbMock.On("ShowStep", mock.Anything, mock.Anything).Return(testStepObj, nil)
	suite.dbMock.On("ListSteps", mock.Anything, mock.Anything).Return(
		[]db.Step{testStepObj}, nil).Maybe()
	recorder := makeHTTPCall(http.MethodDelete,
		"/steps/{id}",
		"/steps/"+testUUID.String(),
		"",
		deleteStepHandler(Dependencies{Store: suite.dbMock}),
	)
	assert.Equal(suite.T(), http.StatusOK, recorder.Code)
	assert.Equal(suite.T(), `{"msg":"step deleted successfully"}`, recorder.Body.String())

	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *StepHandlerTestSuite) TestShowStepSuccess() {

	suite.dbMock.On("ShowStep", mock.Anything, mock.Anything).Return(testStepObj, nil)

	recorder := makeHTTPCall(http.MethodGet,
		"/steps/{id}",
		"/steps/"+stepUUID.String(),
		"",
		showStepHandler(Dependencies{Store: suite.dbMock}),
	)
	output, _ := json.Marshal(testStepObj)
	assert.Equal(suite.T(), http.StatusOK, recorder.Code)
	assert.Equal(suite.T(), string(output), recorder.Body.String())

	suite.dbMock.AssertExpectations(suite.T())
}
