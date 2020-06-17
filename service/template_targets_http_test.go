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

type TempTargetHandlerTestSuite struct {
	suite.Suite

	dbMock *db.DBMockStore
}

func (suite *TempTargetHandlerTestSuite) SetupTest() {
	suite.dbMock = &db.DBMockStore{}
}

func TestTempTargetTestSuite(t *testing.T) {
	suite.Run(t, new(TempTargetHandlerTestSuite))
}

func (suite *TempTargetHandlerTestSuite) TestListTempTargetSuccess() {
	targetUUID := uuid.New()
	tempUUID := uuid.New()
	suite.dbMock.On("ListTemplateTargets", mock.Anything, mock.Anything).Return(
		[]db.TemplateTarget{
			db.TemplateTarget{TemplateID: tempUUID, TargetID: targetUUID, Threshold: 10.5},
		},
		nil,
	)

	recorder := makeHTTPCall(
		http.MethodGet,
		"/templates/{template_id}/targets",
		"/templates/"+tempUUID.String()+"targets",
		"",
		listTempTargetsHandler(Dependencies{Store: suite.dbMock}),
	)
	output := fmt.Sprintf(`[{"template_id":"%s","target_id":"%s","threshold":10.5}]`, tempUUID, targetUUID)
	assert.Equal(suite.T(), http.StatusOK, recorder.Code)
	assert.Equal(suite.T(), output, recorder.Body.String())
	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *TempTargetHandlerTestSuite) TestListTemplatesFail() {
	tempUUID := uuid.New()
	suite.dbMock.On("ListTemplateTargets", mock.Anything, mock.Anything).Return(
		[]db.TemplateTarget{},
		errors.New("error fetching template target"),
	)

	recorder := makeHTTPCall(
		http.MethodGet,
		"/templates/{template_id}/targets",
		"/templates/"+tempUUID.String()+"targets",
		"",
		listTempTargetsHandler(Dependencies{Store: suite.dbMock}),
	)

	assert.Equal(suite.T(), http.StatusInternalServerError, recorder.Code)
	assert.Equal(suite.T(), "", recorder.Body.String())
	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *TempTargetHandlerTestSuite) TestUpsertTempTargetSuccess() {
	targetUUID := uuid.New()
	tempUUID := uuid.New()
	suite.dbMock.On("UpsertTemplateTarget", mock.Anything, mock.Anything, mock.Anything).Return(
		[]db.TemplateTarget{
			db.TemplateTarget{TemplateID: tempUUID, TargetID: targetUUID, Threshold: 10.5},
		},
		nil,
	)

	body := fmt.Sprintf(`[{"template_id":"%s","target_id":"%s","threshold":10.5}]`, tempUUID, targetUUID)

	recorder := makeHTTPCall(http.MethodPost,
		"/templates/{template_id}/targets",
		"/templates/"+tempUUID.String()+"targets",
		body,
		updateTempTargetsHandler(Dependencies{Store: suite.dbMock}),
	)

	output := fmt.Sprintf(`[{"template_id":"%s","target_id":"%s","threshold":10.5}]`, tempUUID, targetUUID)
	assert.Equal(suite.T(), http.StatusCreated, recorder.Code)
	assert.Equal(suite.T(), output, recorder.Body.String())

	suite.dbMock.AssertExpectations(suite.T())
}
