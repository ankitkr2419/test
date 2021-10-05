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

var testExpID = uuid.New()
var testTemplateID = uuid.New()
var testTargetID = uuid.New()
var testExpTemplateTargetobj = db.ExpTemplateTarget{
	ExperimentID: testExpID,
	TemplateID:   testTemplateID,
	TargetID:     testTargetID,
	Threshold:    10.5,
}

var testTargTempObj = []db.ExpTempTargeTDye{
	db.ExpTempTargeTDye{
		DyeName:           "test",
		ExpTemplateTarget: testExpTemplateTargetobj,
	},
}

var testTempTarObj = db.TemplateTarget{
	TemplateID: testTemplateID,
	TargetID:   testTargetID,
	Threshold:  0.5,
}

func (suite *ExpTempTargetHandlerTestSuite) TestListExpTempTargetSuccess() {

	suite.dbMock.On("ListExpTemplateTargets", mock.Anything, mock.Anything).Return(
		testTargTempObj,
		nil,
	)

	recorder := makeHTTPCall(
		http.MethodGet,
		"/experiments/{experiment_id}/targets",
		"/experiments/"+testExpID.String()+"/targets",
		"",
		listExpTempTargetsHandler(Dependencies{Store: suite.dbMock}),
	)
	output, _ := json.Marshal(testTargTempObj)
	assert.Equal(suite.T(), http.StatusOK, recorder.Code)
	assert.Equal(suite.T(), string(output), recorder.Body.String())
	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *ExpTempTargetHandlerTestSuite) TestListExpTemplatesFail() {
	expID := uuid.New()
	suite.dbMock.On("ListExpTemplateTargets", mock.Anything, mock.Anything).Return(
		testTargTempObj,
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

	suite.dbMock.On("UpsertExpTemplateTarget", mock.Anything, mock.Anything, mock.Anything).Return(
		testTargTempObj,
		nil,
	)

	body, _ := json.Marshal([]db.ExpTemplateTarget{testExpTemplateTargetobj})

	recorder := makeHTTPCall(http.MethodPost,
		"/experiments/{experiment_id}/targets",
		"/experiments/"+testExpID.String()+"/targets",
		string(body),
		updateExpTempTargetsHandler(Dependencies{Store: suite.dbMock}),
	)

	output, _ := json.Marshal(testTargTempObj)
	assert.Equal(suite.T(), http.StatusCreated, recorder.Code)
	assert.Equal(suite.T(), string(output), recorder.Body.String())

	suite.dbMock.AssertExpectations(suite.T())
}
