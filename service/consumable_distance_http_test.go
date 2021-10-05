package service

import (
	"encoding/json"
	"errors"
	"mylab/cpagent/config"
	"mylab/cpagent/db"
	"mylab/cpagent/responses"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

// Define the suite, and absorb the built-in basic suite
// functionality from testify - including assertion methods.
type ConsumableDistanceHandlerTestSuite struct {
	suite.Suite

	dbMock *db.DBMockStore
}

func (suite *ConsumableDistanceHandlerTestSuite) SetupTest() {
	suite.dbMock = &db.DBMockStore{}
	config.SetConsumableDistanceFilePath("../conf/test_config.yml")
}

func TestConsumableDistanceTestSuite(t *testing.T) {
	suite.Run(t, new(ConsumableDistanceHandlerTestSuite))
}

var testConsumableObj = db.ConsumableDistance{
	ID:          0,
	Name:        "test consumable",
	Distance:    3.5,
	Description: "test consumable",
}

func (suite *ConsumableDistanceHandlerTestSuite) TestListConsumableDistanceSuccess() {
	suite.dbMock.On("ListConsDistances", mock.Anything, mock.Anything).Return([]db.ConsumableDistance{testConsumableObj}, nil)
	suite.dbMock.On("AddAuditLog", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil).Maybe()

	recorder := makeHTTPCall(http.MethodPost,
		"/consumable-distance/",
		"/consumable-distance/",
		"",
		listConsumableDistanceHandler(Dependencies{Store: suite.dbMock}),
	)
	body, _ := json.Marshal([]db.ConsumableDistance{testConsumableObj})
	assert.Equal(suite.T(), http.StatusOK, recorder.Code)
	assert.Equal(suite.T(), string(body), recorder.Body.String())

	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *ConsumableDistanceHandlerTestSuite) TestListConsumableDistanceFailure() {
	suite.dbMock.On("ListConsDistances", mock.Anything, mock.Anything).Return([]db.ConsumableDistance{}, errors.New("failed to fetch distances"))
	suite.dbMock.On("AddAuditLog", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil).Maybe()

	recorder := makeHTTPCall(http.MethodPost,
		"/consumable-distance/",
		"/consumable-distance/",
		"",
		listConsumableDistanceHandler(Dependencies{Store: suite.dbMock}),
	)

	resp, _ := json.Marshal(ErrObj{Err: responses.ConsumableDistanceFetchError.Error()})
	assert.Equal(suite.T(), http.StatusInternalServerError, recorder.Code)
	assert.Equal(suite.T(), string(resp), recorder.Body.String())

	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *ConsumableDistanceHandlerTestSuite) TestCreateConsumableDistanceSuccess() {
	suite.dbMock.On("InsertConsumableDistance", mock.Anything, mock.Anything).Return(nil)
	suite.dbMock.On("AddAuditLog", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil).Maybe()

	body, _ := json.Marshal(testConsumableObj)
	recorder := makeHTTPCall(http.MethodPost,
		"/consumable-distance/",
		"/consumable-distance/",
		string(body),
		createConsumableDistanceHandler(Dependencies{Store: suite.dbMock}),
	)

	assert.Equal(suite.T(), http.StatusCreated, recorder.Code)
	assert.Equal(suite.T(), string(body), recorder.Body.String())

	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *ConsumableDistanceHandlerTestSuite) TestCreateConsumableDistanceFailure() {
	suite.dbMock.On("InsertConsumableDistance", mock.Anything, mock.Anything).Return(errors.New("failed to insert new distances"))
	suite.dbMock.On("AddAuditLog", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil).Maybe()

	body, _ := json.Marshal(testConsumableObj)
	recorder := makeHTTPCall(http.MethodPost,
		"/consumable-distance/",
		"/consumable-distance/",
		string(body),
		createConsumableDistanceHandler(Dependencies{Store: suite.dbMock}),
	)

	assert.Equal(suite.T(), http.StatusInternalServerError, recorder.Code)

	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *ConsumableDistanceHandlerTestSuite) TestUpdateConsumableDistanceSuccess() {
	suite.dbMock.On("UpdateConsumableDistance", mock.Anything, mock.Anything).Return(nil)
	suite.dbMock.On("AddAuditLog", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil).Maybe()

	body, _ := json.Marshal(testConsumableObj)
	recorder := makeHTTPCall(http.MethodPost,
		"/consumable-distance/",
		"/consumable-distance/",
		string(body),
		updateConsumableDistanceHandler(Dependencies{Store: suite.dbMock}),
	)

	output, _ := json.Marshal(MsgObj{Msg: "consumable distance record updated successfully"})
	assert.Equal(suite.T(), http.StatusOK, recorder.Code)
	assert.Equal(suite.T(), string(output), recorder.Body.String())

	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *ConsumableDistanceHandlerTestSuite) TestUpdateConsumableDistanceFailure() {
	suite.dbMock.On("UpdateConsumableDistance", mock.Anything, mock.Anything).Return(errors.New("failed to update new distances"))
	suite.dbMock.On("AddAuditLog", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil).Maybe()

	body, _ := json.Marshal(testConsumableObj)
	recorder := makeHTTPCall(http.MethodPost,
		"/consumable-distance/",
		"/consumable-distance/",
		string(body),
		updateConsumableDistanceHandler(Dependencies{Store: suite.dbMock}),
	)
	output, _ := json.Marshal(ErrObj{Err: responses.ConsumableDistanceUpdateError.Error()})
	assert.Equal(suite.T(), http.StatusInternalServerError, recorder.Code)
	assert.Equal(suite.T(), string(output), recorder.Body.String())

	suite.dbMock.AssertExpectations(suite.T())
}
