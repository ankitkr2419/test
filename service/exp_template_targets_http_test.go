package service
/*
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

type ExpTempTargetHandlerTestSuite struct {
	suite.Suite

	dbMock *db.DBMockStore
}

func (suite *ExpTempTargetHandlerTestSuite) SetupTest() {
	suite.dbMock = &db.DBMockStore{}
}

func TestExpTempTargetTestSuite(t *testing.T) {
	suite.Run(t, new(ExpTempTargetHandlerTestSuite))
}

func (suite *ExpTempTargetHandlerTestSuite) TestListExpTempTargetSuccess() {
	expID := uuid.New()
	targetUUID := uuid.New()
	tempUUID := uuid.New()
	suite.dbMock.On("ListExpTemplateTargets", mock.Anything, mock.Anything).Return(
		[]db.ExpTemplateTarget{
			db.ExpTemplateTarget{ExperimentID: expID, TemplateID: tempUUID, TargetID: targetUUID, Threshold: 10.5},
		},
		nil,
	)

	recorder := makeHTTPCall(
		http.MethodGet,
		"/experiments/{experiment_id}/targets",
		"/experiments/"+expID.String()+"/targets",
		"",
		listExpTempTargetsHandler(Dependencies{Store: suite.dbMock}),
	)
	output := fmt.Sprintf(`[{"experiment_id":"%s","template_id":"%s","target_id":"%s","threshold":10.5,"target_name":""}]`, expID, tempUUID, targetUUID)
	assert.Equal(suite.T(), http.StatusOK, recorder.Code)
	assert.Equal(suite.T(), output, recorder.Body.String())
	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *ExpTempTargetHandlerTestSuite) TestListExpTemplatesFail() {
	expID := uuid.New()
	suite.dbMock.On("ListExpTemplateTargets", mock.Anything, mock.Anything).Return(
		[]db.ExpTemplateTarget{},
		errors.New("error fetching experiment template target"),
	)

	recorder := makeHTTPCall(
		http.MethodGet,
		"/experiments/{experiment_id}/targets",
		"/experiments/"+expID.String()+"/targets",
		"",
		listExpTempTargetsHandler(Dependencies{Store: suite.dbMock}),
	)

	assert.Equal(suite.T(), http.StatusInternalServerError, recorder.Code)
	assert.Equal(suite.T(), "", recorder.Body.String())
	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *ExpTempTargetHandlerTestSuite) TestUpsertExpTempTargetSuccess() {
	expID := uuid.New()
	targetUUID := uuid.New()
	tempUUID := uuid.New()
	suite.dbMock.On("UpsertExpTemplateTarget", mock.Anything, mock.Anything, mock.Anything).Return(
		[]db.ExpTemplateTarget{
			db.ExpTemplateTarget{ExperimentID: expID, TemplateID: tempUUID, TargetID: targetUUID, Threshold: 10.5},
		},
		nil,
	)

	body := fmt.Sprintf(`[{"experiment_id":"%s","template_id":"%s","target_id":"%s","threshold":10.5}]`, expID, tempUUID, targetUUID)

	recorder := makeHTTPCall(http.MethodPost,
		"/experiments/{experiment_id}/targets",
		"/experiments/"+expID.String()+"/targets",
		body,
		updateExpTempTargetsHandler(Dependencies{Store: suite.dbMock}),
	)

	output := fmt.Sprintf(`[{"experiment_id":"%s","template_id":"%s","target_id":"%s","threshold":10.5,"target_name":""}]`, expID, tempUUID, targetUUID)
	assert.Equal(suite.T(), http.StatusCreated, recorder.Code)
	assert.Equal(suite.T(), output, recorder.Body.String())

	suite.dbMock.AssertExpectations(suite.T())
}
*/