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
type TemplateHandlerTestSuite struct {
	suite.Suite

	dbMock *db.DBMockStore
}

func (suite *TemplateHandlerTestSuite) SetupTest() {
	suite.dbMock = &db.DBMockStore{}
}

func TestTemplateTestSuite(t *testing.T) {
	suite.Run(t, new(TemplateHandlerTestSuite))
}

func (suite *TemplateHandlerTestSuite) TestListTemplatesSuccess() {
	testUUID := uuid.New()
	suite.dbMock.On("ListTemplates", mock.Anything).Return(
		[]db.Template{
			db.Template{Name: "test-template", ID: testUUID, Description: "blah blah", Publish: false},
		},
		nil,
	)

	recorder := makeHTTPCall(
		http.MethodGet,
		"/templates",
		"/templates",
		"",
		listTemplateHandler(Dependencies{Store: suite.dbMock}),
	)
	output := fmt.Sprintf(`[{"id":"%s","name":"test-template","description":"blah blah","Publish":false}]`, testUUID)
	assert.Equal(suite.T(), http.StatusOK, recorder.Code)
	assert.Equal(suite.T(), output, recorder.Body.String())
	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *TemplateHandlerTestSuite) TestListTemplatesFail() {
	suite.dbMock.On("ListTemplates", mock.Anything).Return(
		[]db.Template{},
		errors.New("error fetching templates"),
	)

	recorder := makeHTTPCall(
		http.MethodGet,
		"/templates",
		"/templates",
		"",
		listTemplateHandler(Dependencies{Store: suite.dbMock}),
	)

	assert.Equal(suite.T(), http.StatusInternalServerError, recorder.Code)
	assert.Equal(suite.T(), "", recorder.Body.String())
	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *TemplateHandlerTestSuite) TestCreateTemplateSuccess() {
	testUUID := uuid.New()
	suite.dbMock.On("CreateTemplate", mock.Anything, mock.Anything).Return(db.Template{
		ID:          testUUID,
		Name:        "test template",
		Description: "blah blah",
		Publish:     false,
	}, nil)

	stageUUID := uuid.New()
	suite.dbMock.On("CreateStages", mock.Anything, mock.Anything).Return([]db.Stage{
		{ID: stageUUID, Type: "Repeat", RepeatCount: 3, TemplateID: testUUID, StepCount: 0},
	}, nil)

	body := `{"name":"test template","description":"blah blah"}`

	recorder := makeHTTPCall(http.MethodPost,
		"/templates",
		"/templates",
		body,
		createTemplateHandler(Dependencies{Store: suite.dbMock}),
	)
	output := fmt.Sprintf(`{"template":{"id":"%s","name":"test template","description":"blah blah"},"stages":[{"id":"%s","type":"Repeat","repeat_count":3,"template_id":"%s","step_count":0,"created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z"}]}`, testUUID, stageUUID, testUUID)
	assert.Equal(suite.T(), http.StatusCreated, recorder.Code)
	assert.Equal(suite.T(), output, recorder.Body.String())

	suite.dbMock.AssertExpectations(suite.T())
}
func (suite *TemplateHandlerTestSuite) TestUpdateTemplateSuccess() {
	testUUID := uuid.New()
	suite.dbMock.On("UpdateTemplate", mock.Anything, mock.Anything).Return(db.Template{
		ID:          testUUID,
		Name:        "test template",
		Description: "blah blah",
	}, nil)

	body := `{"name":"test template","description":"blah blah"}`

	recorder := makeHTTPCall(http.MethodPut,
		"/templates/{id}",
		"/templates/"+testUUID.String(),
		body,
		updateTemplateHandler(Dependencies{Store: suite.dbMock}),
	)

	assert.Equal(suite.T(), http.StatusOK, recorder.Code)
	assert.Equal(suite.T(), `{"msg":"template updated successfully"}`, recorder.Body.String())

	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *TemplateHandlerTestSuite) TestDeleteTemplateSuccess() {
	testUUID := uuid.New()
	suite.dbMock.On("DeleteTemplate", mock.Anything, mock.Anything).Return(
		testUUID,
		nil)

	recorder := makeHTTPCall(http.MethodDelete,
		"/templates/{id}",
		"/templates/"+testUUID.String(),
		"",
		deleteTemplateHandler(Dependencies{Store: suite.dbMock}),
	)
	assert.Equal(suite.T(), http.StatusOK, recorder.Code)
	assert.Equal(suite.T(), `{"msg":"template deleted successfully"}`, recorder.Body.String())

	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *TemplateHandlerTestSuite) TestShowTemplateSuccess() {
	testUUID := uuid.New()
	suite.dbMock.On("ShowTemplate", mock.Anything, mock.Anything).Return(db.Template{
		ID:          testUUID,
		Name:        "test template",
		Description: "blah blah",
		Publish:     false,
	}, nil)

	recorder := makeHTTPCall(http.MethodGet,
		"/templates/{id}",
		"/templates/"+testUUID.String(),
		"",
		showTemplateHandler(Dependencies{Store: suite.dbMock}),
	)
	output := fmt.Sprintf(`{"id":"%s","name":"test template","description":"blah blah","Publish":false}`, testUUID)
	assert.Equal(suite.T(), http.StatusOK, recorder.Code)
	assert.Equal(suite.T(), output, recorder.Body.String())

	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *TemplateHandlerTestSuite) TestPublishTemplateSuccess() {
	testUUID := uuid.New()
	targetUUID := uuid.New()
	tempUUID := uuid.New()
	suite.dbMock.On("ListTemplateTargets", mock.Anything, mock.Anything).Return(
		[]db.TemplateTarget{
			db.TemplateTarget{TemplateID: tempUUID, TargetID: targetUUID, Threshold: 10.5},
		},
		nil,
	)

	stage1 := db.Stage{ID: testUUID, Type: "cycle", RepeatCount: 3, TemplateID: tempUUID, StepCount: 0}
	stage2 := db.Stage{ID: testUUID, Type: "hold", RepeatCount: 0, TemplateID: tempUUID, StepCount: 0}

	step := db.Step{TargetTemperature: 25.5, RampRate: 5.5, HoldTime: 120, DataCapture: true, StageID: testUUID}
	ss1 := db.StageStep{
		stage1, step,
	}
	ss2 := db.StageStep{
		stage2, step,
	}
	suite.dbMock.On("ListStageSteps", mock.Anything, mock.Anything).Return([]db.StageStep{
		ss1, ss2,
	}, nil)

	suite.dbMock.On("PublishTemplate", mock.Anything, mock.Anything).Return(
		nil,
		nil)

	body := ``

	recorder := makeHTTPCall(http.MethodPut,
		"/templates/{id}/publish",
		"/templates/"+tempUUID.String()+"/publish",
		body,
		publishTemplateHandler(Dependencies{Store: suite.dbMock}),
	)

	assert.Equal(suite.T(), http.StatusOK, recorder.Code)
	assert.Equal(suite.T(), `{"msg":"template published successfully"}`, recorder.Body.String())

	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *TemplateHandlerTestSuite) TestListPublishedTemplatesSuccess() {
	testUUID := uuid.New()
	suite.dbMock.On("ListPublishedTemplates", mock.Anything).Return(
		[]db.Template{
			db.Template{Name: "test-template", ID: testUUID, Description: "blah blah", Publish: false},
		},
		nil,
	)

	recorder := makeHTTPCall(
		http.MethodGet,
		"/templates/publish",
		"/templates/publish",
		"",
		listPublishedTemplateHandler(Dependencies{Store: suite.dbMock}),
	)
	output := fmt.Sprintf(`[{"id":"%s","name":"test-template","description":"blah blah","Publish":false}]`, testUUID)
	assert.Equal(suite.T(), http.StatusOK, recorder.Code)
	assert.Equal(suite.T(), output, recorder.Body.String())
	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *TemplateHandlerTestSuite) TestListPublishedTemplatesFail() {
	suite.dbMock.On("ListPublishedTemplates", mock.Anything).Return(
		[]db.Template{},
		errors.New("error fetching templates"),
	)

	recorder := makeHTTPCall(
		http.MethodGet,
		"/templates/publish",
		"/templates/publish",
		"",
		listPublishedTemplateHandler(Dependencies{Store: suite.dbMock}),
	)

	assert.Equal(suite.T(), http.StatusInternalServerError, recorder.Code)
	assert.Equal(suite.T(), "", recorder.Body.String())
	suite.dbMock.AssertExpectations(suite.T())
}
