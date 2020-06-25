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
type WellHandlerTestSuite struct {
	suite.Suite

	dbMock *db.DBMockStore
}

func (suite *WellHandlerTestSuite) SetupTest() {
	suite.dbMock = &db.DBMockStore{}
}

func TestWellTestSuite(t *testing.T) {
	suite.Run(t, new(WellHandlerTestSuite))
}

func (suite *WellHandlerTestSuite) TestListWellsSuccess() {
	testUUID := uuid.New()
	sampleID := uuid.New()
	experimentID := uuid.New()
	targetID := uuid.New()
	suite.dbMock.On("ListWells", mock.Anything, mock.Anything).Return(
		[]db.Well{
			db.Well{ID: testUUID, Position: 1, SampleID: sampleID, ExperimentID: experimentID, Task: "UNKNOWN", ColorCode: "RED", Targets: []db.WellTarget{
				{WellID: testUUID,
					TargetID: targetID,
					CT:       "45"},
			}, SampleName: ""},
		},
		nil,
	)

	recorder := makeHTTPCall(
		http.MethodGet,
		"/experiments/{experiment_id}/wells",
		"/experiments/"+experimentID.String()+"/wells",
		"",
		listWellsHandler(Dependencies{Store: suite.dbMock}),
	)
	output := fmt.Sprintf(`[{"id":"%s","position":1,"experiment_id":"%s","sample_id":"%s","task":"UNKNOWN","color_code":"RED","targets":[{"well_id":"%s","target_id":"%s","ct":"45"}],"sample_name":""}]`, testUUID, experimentID, sampleID, testUUID, targetID)
	assert.Equal(suite.T(), http.StatusOK, recorder.Code)
	assert.Equal(suite.T(), output, recorder.Body.String())
	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *WellHandlerTestSuite) TestListWellsFail() {
	experimentID := uuid.New()
	suite.dbMock.On("ListWells", mock.Anything, mock.Anything).Return(
		[]db.Well{},
		errors.New("error fetching Wells"),
	)

	recorder := makeHTTPCall(
		http.MethodGet,
		"/experiments/{experiment_id}/wells",
		"/experiments/"+experimentID.String()+"/wells",
		"",
		listWellsHandler(Dependencies{Store: suite.dbMock}),
	)

	assert.Equal(suite.T(), http.StatusInternalServerError, recorder.Code)
	assert.Equal(suite.T(), "", recorder.Body.String())
	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *WellHandlerTestSuite) TestUpsertWellSuccess() {
	testUUID := uuid.New()
	sampleID := uuid.New()
	experimentID := uuid.New()
	targetID := uuid.New()

	suite.dbMock.On("UpsertWellTargets", mock.Anything, mock.Anything).Return(
		[]db.WellTarget{
			{
				WellID:   testUUID,
				TargetID: targetID,
				CT:       "45",
			},
		},
		nil,
	)
	suite.dbMock.On("UpsertWells", mock.Anything, mock.Anything, mock.Anything).Return(
		[]db.Well{
			db.Well{ID: testUUID, Position: 1, SampleID: sampleID, ExperimentID: experimentID, Task: "UNKNOWN", ColorCode: "RED", Targets: []db.WellTarget{
				{WellID: testUUID,
					TargetID: targetID,
					CT:       "45"},
			}, SampleName: ""},
		},
		nil,
	)

	body := fmt.Sprintf(`{"position":[1],"sample":{"id":"%s","name":"sush"},"task":"UNKNOWN","targets":["%s"]}`, sampleID, targetID)
	fmt.Println(body)
	recorder := makeHTTPCall(http.MethodPost,
		"/experiments/{experiment_id}/wells",
		"/experiments/"+experimentID.String()+"/wells",
		body,
		upsertWellHandler(Dependencies{Store: suite.dbMock}),
	)
	output := fmt.Sprintf(`[{"id":"%s","position":1,"experiment_id":"%s","sample_id":"%s","task":"UNKNOWN","color_code":"RED","targets":[{"well_id":"%s","target_id":"%s","ct":"45"},{"well_id":"%s","target_id":"%s","ct":""}],"sample_name":""}]`, testUUID, experimentID, sampleID, testUUID, targetID, testUUID, targetID)
	assert.Equal(suite.T(), http.StatusCreated, recorder.Code)
	assert.Equal(suite.T(), output, recorder.Body.String())

	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *WellHandlerTestSuite) TestDeleteWellSuccess() {
	testUUID := uuid.New()
	suite.dbMock.On("DeleteWell", mock.Anything, mock.Anything).Return(
		testUUID,
		nil)

	recorder := makeHTTPCall(http.MethodDelete,
		"/Wells/{id}",
		"/Wells/"+testUUID.String(),
		"",
		deleteWellHandler(Dependencies{Store: suite.dbMock}),
	)
	assert.Equal(suite.T(), http.StatusOK, recorder.Code)
	assert.Equal(suite.T(), `{"msg":"Well deleted successfully"}`, recorder.Body.String())

	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *WellHandlerTestSuite) TestShowWellSuccess() {
	testUUID := uuid.New()
	sampleID := uuid.New()
	experimentID := uuid.New()
	targetID := uuid.New()
	suite.dbMock.On("ListWellTargets", mock.Anything, mock.Anything).Return(
		[]db.WellTarget{
			{
				WellID:   testUUID,
				TargetID: targetID,
				CT:       "45",
			},
		},
		nil,
	)
	suite.dbMock.On("ShowWell", mock.Anything, mock.Anything).Return(
		db.Well{ID: testUUID, Position: 1, SampleID: sampleID, ExperimentID: experimentID, Task: "UNKNOWN", ColorCode: "RED", Targets: []db.WellTarget{
			{WellID: testUUID,
				TargetID: targetID,
				CT:       "45"},
		}, SampleName: ""}, nil)

	recorder := makeHTTPCall(http.MethodGet,
		"/Wells/{id}",
		"/Wells/"+testUUID.String(),
		"",
		showWellHandler(Dependencies{Store: suite.dbMock}),
	)
	output := fmt.Sprintf(`{"id":"%s","position":1,"experiment_id":"%s","sample_id":"%s","task":"UNKNOWN","color_code":"RED","targets":[{"well_id":"%s","target_id":"%s","ct":"45"}],"sample_name":""}`, testUUID, experimentID, sampleID, testUUID, targetID)
	assert.Equal(suite.T(), http.StatusOK, recorder.Code)
	assert.Equal(suite.T(), output, recorder.Body.String())

	suite.dbMock.AssertExpectations(suite.T())
}
