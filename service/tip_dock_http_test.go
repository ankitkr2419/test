package service

import (
	"encoding/json"
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
type TipDockHandlerTestSuite struct {
	suite.Suite

	dbMock *db.DBMockStore
}

func (suite *TipDockHandlerTestSuite) SetupTest() {
	suite.dbMock = &db.DBMockStore{}
}

func TestTipDockTestSuite(t *testing.T) {
	suite.Run(t, new(TipDockHandlerTestSuite))
}

var testTipDockRecord = db.TipDock{
	ID:        testUUID,
	Type:      "deck",
	Position:  1,
	Height:    1.3,
	ProcessID: testProcessUUID,
}

func (suite *TipDockHandlerTestSuite) TestCreateTipDockSuccess() {

	suite.dbMock.On("CreateTipDocking", mock.Anything, testTipDockRecord).Return(testTipDockRecord, nil)

	body, _ := json.Marshal(testTipDockRecord)
	recorder := makeHTTPCall(http.MethodPost,
		"/tip-docking",
		"/tip-docking",
		string(body),
		createTipDockHandler(Dependencies{Store: suite.dbMock}),
	)

	assert.Equal(suite.T(), http.StatusCreated, recorder.Code)
	assert.Equal(suite.T(), string(body), recorder.Body.String())

	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *TipDockHandlerTestSuite) TestCreateTipDockFailure() {

	suite.dbMock.On("CreateTipDocking", mock.Anything, testTipDockRecord).Return(testTipDockRecord, responses.TipDockingCreateError)

	body, _ := json.Marshal(testTipDockRecord)
	recorder := makeHTTPCall(http.MethodPost,
		"/tip-docking",
		"/tip-docking",
		string(body),
		createTipDockHandler(Dependencies{Store: suite.dbMock}),
	)
	output := ErrObj{Err: responses.TipDockingCreateError.Error()}
	outputBytes, _ := json.Marshal(output)

	assert.Equal(suite.T(), http.StatusInternalServerError, recorder.Code)
	assert.Equal(suite.T(), outputBytes, recorder.Body.Bytes())

	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *TipDockHandlerTestSuite) TestShowTipDockSuccess() {

	suite.dbMock.On("ShowTipDocking", mock.Anything, testProcessUUID).Return(testTipDockRecord, nil)

	recorder := makeHTTPCall(http.MethodGet,
		"/tip-docking/{id}",
		"/tip-docking/"+testProcessUUID.String(),
		"",
		showTipDockHandler(Dependencies{Store: suite.dbMock}),
	)

	body, _ := json.Marshal(testTipDockRecord)

	assert.Equal(suite.T(), http.StatusOK, recorder.Code)
	assert.Equal(suite.T(), string(body), recorder.Body.String())

	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *TipDockHandlerTestSuite) TestShowTipDockFailure() {

	suite.dbMock.On("ShowTipDocking", mock.Anything, testProcessUUID).Return(testTipDockRecord, responses.TipDockingFetchError)

	recorder := makeHTTPCall(http.MethodGet,
		"/tip-docking/{id}",
		"/tip-docking/"+testProcessUUID.String(),
		"",
		showTipDockHandler(Dependencies{Store: suite.dbMock}),
	)
	output := ErrObj{Err: responses.TipDockingFetchError.Error()}
	outputBytes, _ := json.Marshal(output)

	assert.Equal(suite.T(), http.StatusInternalServerError, recorder.Code)
	assert.Equal(suite.T(), outputBytes, recorder.Body.Bytes())

	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *TipDockHandlerTestSuite) TestUpdateTipDockSuccess() {

	suite.dbMock.On("UpdateTipDock", mock.Anything, testTipDockRecord).Return(nil)

	body, _ := json.Marshal(testTipDockRecord)

	recorder := makeHTTPCall(http.MethodPut,
		"/tip-docking/{id}",
		"/tip-docking/"+testProcessUUID.String(),
		string(body),
		updateTipDockHandler(Dependencies{Store: suite.dbMock}),
	)
	output := MsgObj{Msg: responses.TipDockingUpdateSuccess}
	outputBytes, _ := json.Marshal(output)

	assert.Equal(suite.T(), http.StatusOK, recorder.Code)
	assert.Equal(suite.T(), outputBytes, recorder.Body.Bytes())

	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *TipDockHandlerTestSuite) TestUpdateTipDockFailure() {

	suite.dbMock.On("UpdateTipDock", mock.Anything, testTipDockRecord).Return(responses.TipDockingUpdateError)

	body, _ := json.Marshal(testTipDockRecord)

	recorder := makeHTTPCall(http.MethodPut,
		"/tip-docking/{id}",
		"/tip-docking/"+testProcessUUID.String(),
		string(body),
		updateTipDockHandler(Dependencies{Store: suite.dbMock}),
	)

	output := ErrObj{Err: responses.TipDockingUpdateError.Error()}
	outputBytes, _ := json.Marshal(output)

	assert.Equal(suite.T(), http.StatusInternalServerError, recorder.Code)
	assert.Equal(suite.T(), outputBytes, recorder.Body.Bytes())

	suite.dbMock.AssertExpectations(suite.T())
}
