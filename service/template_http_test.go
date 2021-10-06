package service

import (
	"encoding/json"
	"errors"
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
	suite.dbMock.On("AddAuditLog", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil).Maybe()

}

func TestTemplateTestSuite(t *testing.T) {
	suite.Run(t, new(TemplateHandlerTestSuite))
}

// func (suite *TemplateHandlerTestSuite) TestListTemplatesSuccess() {
// 	testUUID := uuid.New()
// 	suite.dbMock.On("ListTemplates", mock.Anything).Return(
// 		[]db.Template{
// 			db.Template{Name: "test-template", ID: testUUID, Description: "blah blah", Publish: false},
// 		},
// 		nil,
// 	)

// 	recorder := makeHTTPCall(
// 		http.MethodGet,
// 		"/templates",
// 		"/templates",
// 		"",
// 		listTemplateHandler(Dependencies{Store: suite.dbMock}),
// 	)
// 	output := fmt.Sprintf(`[{"id":"%s","name":"test-template","description":"blah blah","publish":false}]`, testUUID)
// 	assert.Equal(suite.T(), http.StatusOK, recorder.Code)
// 	assert.Equal(suite.T(), output, recorder.Body.String())
// 	suite.dbMock.AssertExpectations(suite.T())
// }

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

	suite.dbMock.On("CreateTemplate", mock.Anything, mock.Anything).Return(db.Template{
		ID:          testUUID,
		Name:        "test template",
		Description: "blah blah",
		Publish:     false,
		LidTemp:     100,
		Volume:      40,
	}, nil)

	suite.dbMock.On("CreateStages", mock.Anything, mock.Anything).Return([]db.Stage{testStageObj}, nil)

	body, _ := json.Marshal(db.Template{
		ID:          testUUID,
		Name:        "test template",
		Description: "blah blah",
		Publish:     false,
		LidTemp:     100,
		Volume:      40,
	})

	recorder := makeHTTPCall(http.MethodPost,
		"/templates",
		"/templates",
		string(body),
		createTemplateHandler(Dependencies{Store: suite.dbMock}),
	)
	testTemplate := db.Template{
		ID:          testUUID,
		Name:        "test template",
		Description: "blah blah",
		Publish:     false,
		LidTemp:     100,
		Volume:      40,
		Stages:      []db.Stage{testStageObj},
	}

	output, _ := json.Marshal(testTemplate)
	assert.Equal(suite.T(), http.StatusCreated, recorder.Code)
	assert.Equal(suite.T(), string(output), recorder.Body.String())

	suite.dbMock.AssertExpectations(suite.T())
}

// func (suite *TemplateHandlerTestSuite) TestUpdateTemplateSuccess() {
// 	testUUID := uuid.New()
// 	suite.dbMock.On("UpdateTemplate", mock.Anything, mock.Anything).Return(db.Template{
// 		ID:          testUUID,
// 		Name:        "test template",
// 		Description: "blah blah",
// 	}, nil)

// 	body := `{"name":"test template","description":"blah blah"}`

// 	recorder := makeHTTPCall(http.MethodPut,
// 		"/templates/{id}",
// 		"/templates/"+testUUID.String(),
// 		body,
// 		updateTemplateHandler(Dependencies{Store: suite.dbMock}),
// 	)

// 	assert.Equal(suite.T(), http.StatusOK, recorder.Code)
// 	assert.Equal(suite.T(), `{"msg":"template updated successfully"}`, recorder.Body.String())

// 	suite.dbMock.AssertExpectations(suite.T())
// }

// func (suite *TemplateHandlerTestSuite) TestDeleteTemplateSuccess() {
// 	testUUID := uuid.New()
// 	suite.dbMock.On("DeleteTemplate", mock.Anything, mock.Anything).Return(
// 		testUUID,
// 		nil)

// 	recorder := makeHTTPCall(http.MethodDelete,
// 		"/templates/{id}",
// 		"/templates/"+testUUID.String(),
// 		"",
// 		deleteTemplateHandler(Dependencies{Store: suite.dbMock}),
// 	)
// 	assert.Equal(suite.T(), http.StatusOK, recorder.Code)
// 	assert.Equal(suite.T(), `{"msg":"template deleted successfully"}`, recorder.Body.String())

// 	suite.dbMock.AssertExpectations(suite.T())
// }

func (suite *TemplateHandlerTestSuite) TestShowTemplateSuccess() {

	suite.dbMock.On("ShowTemplate", mock.Anything, mock.Anything).Return(db.Template{
		ID:          testUUID,
		Name:        "test template",
		Description: "blah blah",
		Publish:     false,
		LidTemp:     100,
		Volume:      40,
	}, nil)

	recorder := makeHTTPCall(http.MethodGet,
		"/templates/{id}",
		"/templates/"+testUUID.String(),
		"",
		showTemplateHandler(Dependencies{Store: suite.dbMock}),
	)
	testTemplate := db.Template{
		ID:          testUUID,
		Name:        "test template",
		Description: "blah blah",
		Publish:     false,
		LidTemp:     100,
		Volume:      40,
	}

	output, _ := json.Marshal(testTemplate)

	assert.Equal(suite.T(), http.StatusOK, recorder.Code)
	assert.Equal(suite.T(), string(output), recorder.Body.String())

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

	stage1 := db.Stage{ID: testUUID, Type: "cycle", RepeatCount: 5, TemplateID: tempUUID, StepCount: 0}
	stage2 := db.Stage{ID: testUUID, Type: "hold", RepeatCount: 0, TemplateID: tempUUID, StepCount: 0}

	step := db.Step{TargetTemperature: 25.5, RampRate: 5.5, HoldTime: 120, DataCapture: true, StageID: testUUID}
	ss1 := db.StageStep{
		Stage: stage1, Step: step,
	}
	ss2 := db.StageStep{
		Stage: stage2, Step: step,
	}
	suite.dbMock.On("ListStageSteps", mock.Anything, mock.Anything).Return([]db.StageStep{
		ss1, ss2,
	}, nil)

	suite.dbMock.On("PublishTemplate", mock.Anything, mock.Anything).Return(
		nil,
		nil)

	suite.dbMock.On("CheckIfICTargetAdded", mock.Anything, mock.Anything).Return(
		db.WarnResponse{},
		nil)

	recorder := makeHTTPCall(http.MethodPut,
		"/templates/{id}/publish",
		"/templates/"+tempUUID.String()+"/publish",
		"",
		publishTemplateHandler(Dependencies{Store: suite.dbMock}),
	)

	assert.Equal(suite.T(), http.StatusOK, recorder.Code)
	assert.Equal(suite.T(), `{"msg":"template published successfully"}`, recorder.Body.String())

	suite.dbMock.AssertExpectations(suite.T())
}

// func (suite *TemplateHandlerTestSuite) TestListPublishedTemplatesSuccess() {
// 	testUUID := uuid.New()
// 	suite.dbMock.On("ListPublishedTemplates", mock.Anything).Return(
// 		[]db.Template{
// 			db.Template{Name: "test-template", ID: testUUID, Description: "blah blah", Publish: true},
// 		},
// 		nil,
// 	)

// 	recorder := makeHTTPCall(
// 		http.MethodGet,
// 		"/templates/publish",
// 		"/templates/publish",
// 		"",
// 		listPublishedTemplateHandler(Dependencies{Store: suite.dbMock}),
// 	)
// 	output := fmt.Sprintf(`[{"id":"%s","name":"test-template","description":"blah blah","publish":true}]`, testUUID)
// 	assert.Equal(suite.T(), http.StatusOK, recorder.Code)
// 	assert.Equal(suite.T(), output, recorder.Body.String())
// 	suite.dbMock.AssertExpectations(suite.T())
// }

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

func (suite *TemplateHandlerTestSuite) TestPublishTemplateFail() {
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
		Stage: stage1, Step: step,
	}
	ss2 := db.StageStep{
		Stage: stage2, Step: step,
	}
	suite.dbMock.On("ListStageSteps", mock.Anything, mock.Anything).Return([]db.StageStep{
		ss1, ss2,
	}, nil)

	recorder := makeHTTPCall(http.MethodPut,
		"/templates/{id}/publish",
		"/templates/"+tempUUID.String()+"/publish",
		"",
		publishTemplateHandler(Dependencies{Store: suite.dbMock}),
	)

	assert.Equal(suite.T(), http.StatusBadRequest, recorder.Code)
	assert.Equal(suite.T(), `{"error":{"code":"invalid_data","message":"Please provide valid template data","fields":{"repeatCount":"Invalid repeat_count in cycle stage"}}}`, recorder.Body.String())

	suite.dbMock.AssertExpectations(suite.T())
}
