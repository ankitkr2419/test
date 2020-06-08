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
			db.Template{Name: "test-template", ID: testUUID, Description: "blah blah"},
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
	output := fmt.Sprintf(`[{"id":"%s","name":"test-template","description":"blah blah"}]`, testUUID)
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
	}, nil)

	body := `{"name":"test template","description":"blah blah"}`

	recorder := makeHTTPCall(http.MethodPost,
		"/template",
		"/template",
		body,
		createTemplateHandler(Dependencies{Store: suite.dbMock}),
	)
	output := fmt.Sprintf(`{"id":"%s","name":"test template","description":"blah blah"}`, testUUID)
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
		"/template/{id}",
		"/template/"+testUUID.String(),
		body,
		updateTemplateHandler(Dependencies{Store: suite.dbMock}),
	)

	assert.Equal(suite.T(), http.StatusOK, recorder.Code)
	assert.Equal(suite.T(), "template updated successfully", recorder.Body.String())

	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *TemplateHandlerTestSuite) TestDeleteTemplateSuccess() {
	testUUID := uuid.New()
	suite.dbMock.On("DeleteTemplate", mock.Anything, mock.Anything).Return(
		testUUID,
		nil)

	recorder := makeHTTPCall(http.MethodDelete,
		"/template/{id}",
		"/template/"+testUUID.String(),
		"",
		deleteTemplateHandler(Dependencies{Store: suite.dbMock}),
	)
	assert.Equal(suite.T(), http.StatusOK, recorder.Code)
	assert.Equal(suite.T(), "", recorder.Body.String())

	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *TemplateHandlerTestSuite) TestShowTemplateSuccess() {
	testUUID := uuid.New()
	suite.dbMock.On("ShowTemplate", mock.Anything, mock.Anything).Return(db.Template{
		ID:          testUUID,
		Name:        "test template",
		Description: "blah blah",
	}, nil)

	recorder := makeHTTPCall(http.MethodGet,
		"/template/{id}",
		"/template/"+testUUID.String(),
		"",
		showTemplateHandler(Dependencies{Store: suite.dbMock}),
	)
	output := fmt.Sprintf(`{"id":"%s","name":"test template","description":"blah blah"}`, testUUID)
	assert.Equal(suite.T(), http.StatusOK, recorder.Code)
	assert.Equal(suite.T(), output, recorder.Body.String())

	suite.dbMock.AssertExpectations(suite.T())
}