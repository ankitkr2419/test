package service

import (
	"encoding/json"
	"fmt"
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

func (suite *PiercingHandlerTestSuite) TestCreatePiercingSuccess() {
	testUUID := uuid.New()
	testProcessUUID := uuid.New()
	suite.dbMock.On("CreatePiercing", mock.Anything, mock.Anything).Return(db.Piercing{
		ID: testUUID,Type: db.Cartridge1, CartridgeWells: []int64{1, 2}, Discard: "at_pickup_passing", ProcessID: testProcessUUID,
	}, nil)

	body := fmt.Sprintf(`{"type":"cartridge_1","cartridge_wells":[1,2],"discard":"at_pickup_passing","process_id":"%s"}`, testProcessUUID)
	recorder := makeHTTPCall(http.MethodPost,
		"/piercing",
		"/piercing",
		body,
		createPiercingHandler(Dependencies{Store: suite.dbMock}),
	)
	output := fmt.Sprintf(`{"id":"%s","type":"cartridge_1","cartridge_wells":[1,2],"discard":"at_pickup_passing","process_id":"%s","created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z"}`, testUUID, testProcessUUID)
	
	assert.Equal(suite.T(), http.StatusCreated, recorder.Code)
	assert.Equal(suite.T(), output, recorder.Body.String())

	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *PiercingHandlerTestSuite) TestCreatePiercingFailure() {
	testProcessUUID := uuid.New()
	suite.dbMock.On("CreatePiercing", mock.Anything, mock.Anything).Return(db.Piercing{}, responses.PiercingCreateError)

	body := fmt.Sprintf(`{"type":"cartridge_1","cartridge_wells":[1,2],"discard":"at_pickup_passing","process_id":"%s"}`, testProcessUUID)
	recorder := makeHTTPCall(http.MethodPost,
		"/piercing",
		"/piercing",
		body,
		createPiercingHandler(Dependencies{Store: suite.dbMock}),
	)
	output := ErrObj {Err: responses.PiercingCreateError.Error()}
	outputBytes,_ := json.Marshal(output)

	assert.Equal(suite.T(), http.StatusInternalServerError, recorder.Code)
	assert.Equal(suite.T(), outputBytes, recorder.Body.Bytes())

	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *PiercingHandlerTestSuite) TestShowPiercingSuccess() {
	testUUID := uuid.New()
	testProcessUUID := uuid.New()
	suite.dbMock.On("ShowPiercing", mock.Anything, mock.Anything).Return(db.Piercing{
		ID: testUUID, Type:db.Cartridge1, CartridgeWells: []int64{1, 2}, Discard: "at_pickup_passing", ProcessID: testProcessUUID,
	}, nil)

	recorder := makeHTTPCall(http.MethodGet,
		"/piercing/{id}",
		"/piercing/"+testUUID.String(),
		"",
		showPiercingHandler(Dependencies{Store: suite.dbMock}),
	)
	output := fmt.Sprintf(`{"id":"%s","type":"cartridge_1","cartridge_wells":[1,2],"discard":"at_pickup_passing","process_id":"%s","created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z"}`, testUUID, testProcessUUID)
	assert.Equal(suite.T(), http.StatusOK, recorder.Code)
	assert.Equal(suite.T(), output, recorder.Body.String())

	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *PiercingHandlerTestSuite) TestShowPiercingFailure() {
	testUUID := uuid.New()
	suite.dbMock.On("ShowPiercing", mock.Anything, mock.Anything).Return(db.Piercing{}, responses.PiercingFetchError)

	recorder := makeHTTPCall(http.MethodGet,
		"/piercing/{id}",
		"/piercing/"+testUUID.String(),
		"",
		showPiercingHandler(Dependencies{Store: suite.dbMock}),
	)
	output := ErrObj {Err: responses.PiercingFetchError.Error()}
	outputBytes,_ := json.Marshal(output)

	assert.Equal(suite.T(), http.StatusInternalServerError, recorder.Code)
	assert.Equal(suite.T(), outputBytes, recorder.Body.Bytes())

	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *PiercingHandlerTestSuite) TestUpdatePiercingSuccess() {
	testUUID := uuid.New()
	testProcessUUID := uuid.New()
	suite.dbMock.On("UpdatePiercing", mock.Anything, mock.Anything).Return(db.Piercing{
		ID: testUUID, Type:db.Cartridge1, CartridgeWells: []int64{1, 2}, Discard: "at_pickup_passing", ProcessID: testProcessUUID,
	}, nil)

	body := fmt.Sprintf(`{"id":"%s","type":"cartridge_1","cartridge_wells":[1,2],"discard":"at_pickup_passing","process_id":"%s","created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z"}`, testUUID, testProcessUUID)

	recorder := makeHTTPCall(http.MethodPut,
		"/piercing/{id}",
		"/piercing/"+testUUID.String(),
		body,
		updatePiercingHandler(Dependencies{Store: suite.dbMock}),
	)
	output := MsgObj {Msg: responses.PiercingUpdateSuccess}
	outputBytes,_ := json.Marshal(output)

	assert.Equal(suite.T(), http.StatusOK, recorder.Code)
	assert.Equal(suite.T(), outputBytes, recorder.Body.Bytes())

	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *PiercingHandlerTestSuite) TestUpdatePiercingFailure() {
	testUUID := uuid.New()
	testProcessUUID := uuid.New()
	suite.dbMock.On("UpdatePiercing", mock.Anything, mock.Anything).Return(db.Piercing{}, responses.PiercingUpdateError)

	body := fmt.Sprintf(`{"id":"%s","type":"cartridge_1","cartridge_wells":[1,2],"discard":"at_pickup_passing","process_id":"%s","created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z"}`, testUUID, testProcessUUID)

	recorder := makeHTTPCall(http.MethodPut,
		"/piercing/{id}",
		"/piercing/"+testUUID.String(),
		body,
		updatePiercingHandler(Dependencies{Store: suite.dbMock}),
	)

	output := ErrObj {Err: responses.PiercingUpdateError.Error()}
	outputBytes,_ := json.Marshal(output)

	assert.Equal(suite.T(), http.StatusInternalServerError, recorder.Code)
	assert.Equal(suite.T(), outputBytes , recorder.Body.Bytes())

	suite.dbMock.AssertExpectations(suite.T())
}
