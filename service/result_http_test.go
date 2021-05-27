package service

/*
import (
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

type ResultHandlerTestSuite struct {
	suite.Suite
	dbMock *db.DBMockStore
}

func (suite *ResultHandlerTestSuite) SetupTest() {
	suite.dbMock = &db.DBMockStore{}
}

func TestResultTestSuite(t *testing.T) {
	suite.Run(t, new(ResultHandlerTestSuite))
}

func (suite *ResultHandlerTestSuite) TestListResultsSuccess() {
	testUUID := uuid.New()
	tempID := uuid.New()
	sampleID := uuid.New()
	experimentID := uuid.New()
	targetID := uuid.New()
	suite.dbMock.On("ListWells", mock.Anything, mock.Anything).Return(
		[]db.Well{
			db.Well{ID: testUUID, Position: 1, SampleID: sampleID, ExperimentID: experimentID, Task: "UNKNOWN", ColorCode: "RED", Targets: []db.WellTarget{
				{WellPosition: 1,
					ExperimentID: experimentID,
					TargetID:     targetID,
					TargetName:   "COVID",
					CT:           "45"},
			}, SampleName: ""},
		},
		nil,
	)
	suite.dbMock.On("ListConfTargets", mock.Anything, mock.Anything).Return([]db.TargetDetails{
		{TargetID: targetID},
	}, nil)
	suite.dbMock.On("GetResult", mock.Anything, mock.Anything, mock.Anything).Return([]db.Result{
		{ExperimentID: experimentID, TemplateID: tempID, WellPosition: 1, TargetID: targetID, Cycle: 1, FValue: 1234},
	}, nil)

	recorder := makeHTTPCall(
		http.MethodGet,
		"/experiments/{id}/emission",
		"/experiments/"+experimentID.String()+"/emission",
		"",
		getResultHandler(Dependencies{Store: suite.dbMock}),
	)
	output := fmt.Sprintf(`{"type":"Graph","data":[{"well_position":1,"target_id":"%s","experiment_id":"%s","total_cycles":3,"cycle":[1],"f_value":[0.385625],"threshold":0}]}`, targetID, experimentID)
	assert.Equal(suite.T(), http.StatusOK, recorder.Code)
	assert.Equal(suite.T(), output, recorder.Body.String())
	suite.dbMock.AssertExpectations(suite.T())
}
func (suite *ResultHandlerTestSuite) TestListResultsFail() {
	testUUID := uuid.New()
	// tempUUID := uuid.New()
	sampleID := uuid.New()
	experimentID := uuid.New()
	targetID := uuid.New()
	suite.dbMock.On("ListWells", mock.Anything, mock.Anything).Return(
		[]db.Well{
			db.Well{ID: testUUID, Position: 1, SampleID: sampleID, ExperimentID: experimentID, Task: "UNKNOWN", ColorCode: "RED", Targets: []db.WellTarget{
				{WellPosition: 1,
					ExperimentID: experimentID,
					TargetID:     targetID,
					TargetName:   "COVID",
					CT:           "45"},
			}, SampleName: ""},
		},
		nil,
	)
	suite.dbMock.On("ListConfTargets", mock.Anything, mock.Anything).Return([]db.TargetDetails{}, nil)
	suite.dbMock.On("GetResult", mock.Anything, mock.Anything, mock.Anything).Return([]db.Result{}, nil)

	recorder := makeHTTPCall(
		http.MethodGet,
		"/experiments/{id}/emission",
		"/experiments/"+experimentID.String()+"/emission",
		"",
		getResultHandler(Dependencies{Store: suite.dbMock}),
	)
	output := fmt.Sprintf(`{"type":"Graph","data":null}`)
	assert.Equal(suite.T(), http.StatusOK, recorder.Code)
	assert.Equal(suite.T(), output, recorder.Body.String())
	suite.dbMock.AssertExpectations(suite.T())
}
*/
