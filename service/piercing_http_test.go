package service

import (
	"encoding/json"
	"mylab/cpagent/db"
	"mylab/cpagent/responses"
	"net/http"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

// Define the suite, and absorb the built-in basic suite
// functionality from testify - including assertion methods.
type PiercingHandlerTestSuite struct {
	suite.Suite

	dbMock *db.DBMockStore
}

func (suite *PiercingHandlerTestSuite) SetupTest() {
	suite.dbMock = &db.DBMockStore{}
}

func TestPiercingTestSuite(t *testing.T) {
	suite.Run(t, new(PiercingHandlerTestSuite))
}

var testProcessUUID = uuid.New()

var testPiercingRecord = db.Piercing{
	ID:             testUUID,
	Type:           db.Cartridge1,
	CartridgeWells: []int64{1, 2},
	Discard:        "at_pickup_passing",
	ProcessID:      testProcessUUID,
}

func (suite *PiercingHandlerTestSuite) TestCreatePiercingSuccess() {

	suite.dbMock.On("CreatePiercing", mock.Anything, testPiercingRecord).Return(testPiercingRecord, nil)

	body, _ := json.Marshal(testPiercingRecord)
	recorder := makeHTTPCall(http.MethodPost,
		"/piercing",
		"/piercing",
		string(body),
		createPiercingHandler(Dependencies{Store: suite.dbMock}),
	)

	assert.Equal(suite.T(), http.StatusCreated, recorder.Code)
	assert.Equal(suite.T(), string(body), recorder.Body.String())

	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *PiercingHandlerTestSuite) TestCreatePiercingFailure() {

	suite.dbMock.On("CreatePiercing", mock.Anything, testPiercingRecord).Return(db.Piercing{}, responses.PiercingCreateError)

	body, _ := json.Marshal(testPiercingRecord)
	recorder := makeHTTPCall(http.MethodPost,
		"/piercing",
		"/piercing",
		string(body),
		createPiercingHandler(Dependencies{Store: suite.dbMock}),
	)
	output := ErrObj{Err: responses.PiercingCreateError.Error()}
	outputBytes, _ := json.Marshal(output)

	assert.Equal(suite.T(), http.StatusInternalServerError, recorder.Code)
	assert.Equal(suite.T(), outputBytes, recorder.Body.Bytes())

	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *PiercingHandlerTestSuite) TestShowPiercingSuccess() {

	suite.dbMock.On("ShowPiercing", mock.Anything, testProcessUUID).Return(testPiercingRecord, nil)

	recorder := makeHTTPCall(http.MethodGet,
		"/piercing/{id}",
		"/piercing/"+testProcessUUID.String(),
		"",
		showPiercingHandler(Dependencies{Store: suite.dbMock}),
	)

	body, _ := json.Marshal(testPiercingRecord)

	assert.Equal(suite.T(), http.StatusOK, recorder.Code)
	assert.Equal(suite.T(), string(body), recorder.Body.String())

	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *PiercingHandlerTestSuite) TestShowPiercingFailure() {

	suite.dbMock.On("ShowPiercing", mock.Anything, testProcessUUID).Return(db.Piercing{}, responses.PiercingFetchError)

	recorder := makeHTTPCall(http.MethodGet,
		"/piercing/{id}",
		"/piercing/"+testProcessUUID.String(),
		"",
		showPiercingHandler(Dependencies{Store: suite.dbMock}),
	)
	output := ErrObj{Err: responses.PiercingFetchError.Error()}
	outputBytes, _ := json.Marshal(output)

	assert.Equal(suite.T(), http.StatusInternalServerError, recorder.Code)
	assert.Equal(suite.T(), outputBytes, recorder.Body.Bytes())

	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *PiercingHandlerTestSuite) TestUpdatePiercingSuccess() {

	suite.dbMock.On("UpdatePiercing", mock.Anything, testPiercingRecord).Return(testPiercingRecord, nil)

	body, _ := json.Marshal(testPiercingRecord)

	recorder := makeHTTPCall(http.MethodPut,
		"/piercing/{id}",
		"/piercing/"+testProcessUUID.String(),
		string(body),
		updatePiercingHandler(Dependencies{Store: suite.dbMock}),
	)
	output := MsgObj{Msg: responses.PiercingUpdateSuccess}
	outputBytes, _ := json.Marshal(output)

	assert.Equal(suite.T(), http.StatusOK, recorder.Code)
	assert.Equal(suite.T(), outputBytes, recorder.Body.Bytes())

	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *PiercingHandlerTestSuite) TestUpdatePiercingFailure() {

	suite.dbMock.On("UpdatePiercing", mock.Anything, testPiercingRecord).Return(db.Piercing{}, responses.PiercingUpdateError)

	body, _ := json.Marshal(testPiercingRecord)

	recorder := makeHTTPCall(http.MethodPut,
		"/piercing/{id}",
		"/piercing/"+testProcessUUID.String(),
		string(body),
		updatePiercingHandler(Dependencies{Store: suite.dbMock}),
	)

	output := ErrObj{Err: responses.PiercingUpdateError.Error()}
	outputBytes, _ := json.Marshal(output)

	assert.Equal(suite.T(), http.StatusInternalServerError, recorder.Code)
	assert.Equal(suite.T(), outputBytes, recorder.Body.Bytes())

	suite.dbMock.AssertExpectations(suite.T())
}
