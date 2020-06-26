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
type StepHandlerTestSuite struct {
	suite.Suite

	dbMock *db.DBMockStore
}

func (suite *StepHandlerTestSuite) SetupTest() {
	suite.dbMock = &db.DBMockStore{}
}

func TestStepTestSuite(t *testing.T) {
	suite.Run(t, new(StepHandlerTestSuite))
}

func (suite *StepHandlerTestSuite) TestListStepsSuccess() {
	testUUID := uuid.New()
	stgUUID := uuid.New()
	suite.dbMock.On("ListSteps", mock.Anything, mock.Anything).Return(
		[]db.Step{
			db.Step{ID: testUUID, TargetTemperature: 25.5, RampRate: 5.5, HoldTime: 120, DataCapture: true, StageID: stgUUID},
		},
		nil,
	)

	recorder := makeHTTPCall(
		http.MethodGet,
		"/stages/{stage_id}/steps",
		"/stages/"+stgUUID.String()+"/steps",
		"",
		listStepsHandler(Dependencies{Store: suite.dbMock}),
	)
	output := fmt.Sprintf(`[{"id":"%s","stage_id":"%s","ramp_rate":5.5,"target_temp":25.5,"hold_time":120,"data_capture":true}]`, testUUID, stgUUID)
	assert.Equal(suite.T(), http.StatusOK, recorder.Code)
	assert.Equal(suite.T(), output, recorder.Body.String())
	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *StepHandlerTestSuite) TestListStepsFail() {
	stgUUID := uuid.New()
	suite.dbMock.On("ListSteps", mock.Anything, mock.Anything).Return(
		[]db.Step{},
		errors.New("error fetching steps"),
	)

	recorder := makeHTTPCall(
		http.MethodGet,
		"/stages/{stage_id}/steps",
		"/stages/"+stgUUID.String()+"/steps",
		"",
		listStepsHandler(Dependencies{Store: suite.dbMock}),
	)

	assert.Equal(suite.T(), http.StatusInternalServerError, recorder.Code)
	assert.Equal(suite.T(), "", recorder.Body.String())
	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *StepHandlerTestSuite) TestCreateStepSuccess() {
	testUUID := uuid.New()
	stgUUID := uuid.New()
	suite.dbMock.On("UpdateStepCount", mock.Anything).Return(
		nil,
		nil,
	)
	suite.dbMock.On("CreateStep", mock.Anything, mock.Anything).Return(db.Step{
		ID: testUUID, StageID: stgUUID, TargetTemperature: 25.5, RampRate: 5.5, HoldTime: 120, DataCapture: true,
	}, nil)

	body := fmt.Sprintf(`{"stage_id":"%s","ramp_rate":5.5,"target_temp":25.5,"hold_time":120,"data_capture":true}`, stgUUID)
	recorder := makeHTTPCall(http.MethodPost,
		"/steps",
		"/steps",
		body,
		createStepHandler(Dependencies{Store: suite.dbMock}),
	)
	output := fmt.Sprintf(`{"id":"%s","stage_id":"%s","ramp_rate":5.5,"target_temp":25.5,"hold_time":120,"data_capture":true}`, testUUID, stgUUID)
	assert.Equal(suite.T(), http.StatusCreated, recorder.Code)
	assert.Equal(suite.T(), output, recorder.Body.String())

	suite.dbMock.AssertExpectations(suite.T())
}
func (suite *StepHandlerTestSuite) TestUpdateStepSuccess() {
	testUUID := uuid.New()
	stgUUID := uuid.New()
	suite.dbMock.On("UpdateStep", mock.Anything, mock.Anything).Return(db.Step{
		ID: testUUID, StageID: stgUUID, TargetTemperature: 25.5, RampRate: 5.5, HoldTime: 120, DataCapture: true,
	}, nil)

	body := fmt.Sprintf(`{"stage_id":"%s","ramp_rate":5.5,"target_temp":25.5,"hold_time":120,"data_capture":true}`, stgUUID)

	recorder := makeHTTPCall(http.MethodPut,
		"/steps/{id}",
		"/steps/"+testUUID.String(),
		body,
		updateStepHandler(Dependencies{Store: suite.dbMock}),
	)

	assert.Equal(suite.T(), http.StatusOK, recorder.Code)
	assert.Equal(suite.T(), `{"msg":"step updated successfully"}`, recorder.Body.String())

	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *StepHandlerTestSuite) TestDeleteStepSuccess() {
	testUUID := uuid.New()
	suite.dbMock.On("UpdateStepCount", mock.Anything).Return(
		nil,
		nil,
	)
	suite.dbMock.On("DeleteStep", mock.Anything, mock.Anything).Return(
		testUUID,
		nil)

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
	testUUID := uuid.New()
	stgUUID := uuid.New()
	suite.dbMock.On("ShowStep", mock.Anything, mock.Anything).Return(db.Step{
		ID: testUUID, StageID: stgUUID, TargetTemperature: 25.5, RampRate: 5.5, HoldTime: 120, DataCapture: true,
	}, nil)

	recorder := makeHTTPCall(http.MethodGet,
		"/steps/{id}",
		"/steps/"+testUUID.String(),
		"",
		showStepHandler(Dependencies{Store: suite.dbMock}),
	)
	output := fmt.Sprintf(`{"id":"%s","stage_id":"%s","ramp_rate":5.5,"target_temp":25.5,"hold_time":120,"data_capture":true}`, testUUID, stgUUID)
	assert.Equal(suite.T(), http.StatusOK, recorder.Code)
	assert.Equal(suite.T(), output, recorder.Body.String())

	suite.dbMock.AssertExpectations(suite.T())
}
