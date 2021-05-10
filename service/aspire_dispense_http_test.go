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
type AspireDispenseHandlerTestSuite struct {
	suite.Suite

	dbMock *db.DBMockStore
}

func (suite *AspireDispenseHandlerTestSuite) SetupTest() {
	suite.dbMock = &db.DBMockStore{}
}

func TestAspireDispenseTestSuite(t *testing.T) {
	suite.Run(t, new(AspireDispenseHandlerTestSuite))
}

var testAspireDispenseRecord = db.AspireDispense{
	ID:                   testUUID,
	Category:             db.WW,
	CartridgeType:        db.Cartridge1,
	SourcePosition:       1,
	AspireHeight:         2,
	AspireMixingVolume:   3,
	AspireNoOfCycles:     4,
	AspireVolume:         5,
	AspireAirVolume:      6,
	DispenseHeight:       7,
	DispenseMixingVolume: 8,
	DispenseNoOfCycles:   9,
	DestinationPosition:  10,
	ProcessID:            testProcessUUID,
}

func (suite *AspireDispenseHandlerTestSuite) TestCreateAspireDispenseSuccess() {

	suite.dbMock.On("CreateAspireDispense", mock.Anything, testAspireDispenseRecord).Return(testAspireDispenseRecord, nil)

	body, _ := json.Marshal(testAspireDispenseRecord)
	recorder := makeHTTPCall(http.MethodPost,
		"/aspire-dispense",
		"/aspire-dispense",
		string(body),
		createAspireDispenseHandler(Dependencies{Store: suite.dbMock}),
	)

	assert.Equal(suite.T(), http.StatusCreated, recorder.Code)
	assert.Equal(suite.T(), string(body), recorder.Body.String())

	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *AspireDispenseHandlerTestSuite) TestCreateAspireDispenseFailure() {

	suite.dbMock.On("CreateAspireDispense", mock.Anything, testAspireDispenseRecord).Return(db.AspireDispense{}, responses.AspireDispenseCreateError)

	body, _ := json.Marshal(testAspireDispenseRecord)
	recorder := makeHTTPCall(http.MethodPost,
		"/aspire-dispense",
		"/aspire-dispense",
		string(body),
		createAspireDispenseHandler(Dependencies{Store: suite.dbMock}),
	)
	output := ErrObj{Err: responses.AspireDispenseCreateError.Error()}
	outputBytes, _ := json.Marshal(output)

	assert.Equal(suite.T(), http.StatusInternalServerError, recorder.Code)
	assert.Equal(suite.T(), outputBytes, recorder.Body.Bytes())

	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *AspireDispenseHandlerTestSuite) TestShowAspireDispenseSuccess() {

	suite.dbMock.On("ShowAspireDispense", mock.Anything, testProcessUUID).Return(testAspireDispenseRecord, nil)

	recorder := makeHTTPCall(http.MethodGet,
		"/aspire-dispense/{id}",
		"/aspire-dispense/"+testProcessUUID.String(),
		"",
		showAspireDispenseHandler(Dependencies{Store: suite.dbMock}),
	)

	body, _ := json.Marshal(testAspireDispenseRecord)

	assert.Equal(suite.T(), http.StatusOK, recorder.Code)
	assert.Equal(suite.T(), string(body), recorder.Body.String())

	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *AspireDispenseHandlerTestSuite) TestShowAspireDispenseFailure() {

	suite.dbMock.On("ShowAspireDispense", mock.Anything, testProcessUUID).Return(db.AspireDispense{}, responses.AspireDispenseFetchError)

	recorder := makeHTTPCall(http.MethodGet,
		"/aspire-dispense/{id}",
		"/aspire-dispense/"+testProcessUUID.String(),
		"",
		showAspireDispenseHandler(Dependencies{Store: suite.dbMock}),
	)
	output := ErrObj{Err: responses.AspireDispenseFetchError.Error()}
	outputBytes, _ := json.Marshal(output)

	assert.Equal(suite.T(), http.StatusInternalServerError, recorder.Code)
	assert.Equal(suite.T(), outputBytes, recorder.Body.Bytes())

	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *AspireDispenseHandlerTestSuite) TestUpdateAspireDispenseSuccess() {

	suite.dbMock.On("UpdateAspireDispense", mock.Anything, testAspireDispenseRecord).Return(nil)

	body, _ := json.Marshal(testAspireDispenseRecord)

	recorder := makeHTTPCall(http.MethodPut,
		"/aspire-dispense/{id}",
		"/aspire-dispense/"+testProcessUUID.String(),
		string(body),
		updateAspireDispenseHandler(Dependencies{Store: suite.dbMock}),
	)
	output := MsgObj{Msg: responses.AspireDispenseUpdateSuccess}
	outputBytes, _ := json.Marshal(output)

	assert.Equal(suite.T(), http.StatusOK, recorder.Code)
	assert.Equal(suite.T(), outputBytes, recorder.Body.Bytes())

	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *AspireDispenseHandlerTestSuite) TestUpdateAspireDispenseFailure() {

	suite.dbMock.On("UpdateAspireDispense", mock.Anything, testAspireDispenseRecord).Return(responses.AspireDispenseUpdateError)

	body, _ := json.Marshal(testAspireDispenseRecord)

	recorder := makeHTTPCall(http.MethodPut,
		"/aspire-dispense/{id}",
		"/aspire-dispense/"+testProcessUUID.String(),
		string(body),
		updateAspireDispenseHandler(Dependencies{Store: suite.dbMock}),
	)

	output := ErrObj{Err: responses.AspireDispenseUpdateError.Error()}
	outputBytes, _ := json.Marshal(output)

	assert.Equal(suite.T(), http.StatusInternalServerError, recorder.Code)
	assert.Equal(suite.T(), outputBytes, recorder.Body.Bytes())

	suite.dbMock.AssertExpectations(suite.T())
}
