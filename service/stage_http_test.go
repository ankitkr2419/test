package service

import (
	"encoding/json"
	"errors"
	"mylab/cpagent/db"
	"net/http"
	"testing"

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

var testStageObj = db.Stage{
	ID:          testUUID,
	Type:        "Repeat",
	RepeatCount: 3,
	TemplateID:  testTemplateID,
	StepCount:   0}

func (suite *StageHandlerTestSuite) TestListStagesSuccess() {

	suite.dbMock.On("ListStages", mock.Anything, mock.Anything).Return(
		[]db.Stage{testStageObj},
		nil,
	)

	recorder := makeHTTPCall(
		http.MethodGet,
		"/templates/{template_id}/stages",
		"/templates/"+testTemplateID.String()+"/stages",
		"",
		listStagesHandler(Dependencies{Store: suite.dbMock}),
	)
	output, _ := json.Marshal([]db.Stage{testStageObj})
	assert.Equal(suite.T(), http.StatusOK, recorder.Code)
	assert.Equal(suite.T(), string(output), recorder.Body.String())
	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *StageHandlerTestSuite) TestListStagesFail() {

	suite.dbMock.On("ListStages", mock.Anything, mock.Anything).Return(
		[]db.Stage{},
		errors.New("error fetching templates"),
	)

	recorder := makeHTTPCall(
		http.MethodGet,
		"/templates/{template_id}/stages",
		"/templates/"+testTemplateID.String()+"/stages",
		"",
		listStagesHandler(Dependencies{Store: suite.dbMock}),
	)

	assert.Equal(suite.T(), http.StatusInternalServerError, recorder.Code)
	assert.Equal(suite.T(), "", recorder.Body.String())
	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *StageHandlerTestSuite) TestShowStageSuccess() {

	suite.dbMock.On("ShowStage", mock.Anything, mock.Anything).Return(testStageObj, nil)

	recorder := makeHTTPCall(http.MethodGet,
		"/stages/{id}",
		"/stages/"+testUUID.String(),
		"",
		showStageHandler(Dependencies{Store: suite.dbMock}),
	)
	output, _ := json.Marshal(testStageObj)
	assert.Equal(suite.T(), http.StatusOK, recorder.Code)
	assert.Equal(suite.T(), string(output), recorder.Body.String())

	suite.dbMock.AssertExpectations(suite.T())
}

/*
func (suite *StageHandlerTestSuite) TestUpdateStageSuccess() {
	testUUID := uuid.New()
	tempUUID := uuid.New()
	suite.dbMock.On("UpdateStage", mock.Anything, mock.Anything).Return(db.Stage{
		ID: testUUID, Type: "Repeat", RepeatCount: 3, TemplateID: tempUUID, StepCount: 0,
	}, nil)

	body := fmt.Sprintf(`{"type":"Repeat", "repeat_count":3, "template_id":"%s","step_count":0}`, tempUUID)

	recorder := makeHTTPCall(http.MethodPut,
		"/stages/{id}",
		"/stages/"+testUUID.String(),
		body,
		updateStageHandler(Dependencies{Store: suite.dbMock}),
	)

	assert.Equal(suite.T(), http.StatusOK, recorder.Code)
	assert.Equal(suite.T(), `{"msg":"stage updated successfully"}`, recorder.Body.String())

	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *StageHandlerTestSuite) TestDeleteStageSuccess() {
	testUUID := uuid.New()
	suite.dbMock.On("DeleteStage", mock.Anything, mock.Anything).Return(
		testUUID,
		nil)

	recorder := makeHTTPCall(http.MethodDelete,
		"/stages/{id}",
		"/stages/"+testUUID.String(),
		"",
		deleteStageHandler(Dependencies{Store: suite.dbMock}),
	)
	assert.Equal(suite.T(), http.StatusOK, recorder.Code)
	assert.Equal(suite.T(), `{"msg":"stage deleted successfully"}`, recorder.Body.String())

	suite.dbMock.AssertExpectations(suite.T())
}
*/
