package service

import (
	"errors"
	"fmt"
	"mylab/cpagent/config"
	"mylab/cpagent/db"
	"mylab/cpagent/plc"
	"mylab/cpagent/plc/simulator"
	"net/http"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

// Define the suite, and absorb the built-in basic suite
// functionality from testify - including assertion methods.

type ExperimentHandlerTestSuite struct {
	suite.Suite

	dbMock  *db.DBMockStore
	plc     plc.Driver
	ExitCh  <-chan error
	WsErrCh chan error
	WsMsgCh chan string
}

func (suite *ExperimentHandlerTestSuite) SetupTest() {
	suite.dbMock = &db.DBMockStore{}
}

func TestExperimentTestSuite(t *testing.T) {
	suite.Run(t, new(ExperimentHandlerTestSuite))
}

func (suite *ExperimentHandlerTestSuite) TestListExperimentsSuccess() {
	testUUID := uuid.New()
	tempUUID := uuid.New()
	suite.dbMock.On("ListExperiments", mock.Anything).Return(
		[]db.Experiment{
			db.Experiment{ID: testUUID, Description: "blah blah", TemplateID: tempUUID, TemplateName: "test", OperatorName: "ABC"},
		},
		nil,
	)

	recorder := makeHTTPCall(
		http.MethodGet,
		"/experiments",
		"/experiments",
		"",
		listExperimentHandler(Dependencies{Store: suite.dbMock}),
	)
	output := fmt.Sprintf(`[{"id":"%s","description":"blah blah","template_id":"%s","template_name":"test","operator_name":"ABC","start_time":"0001-01-01T00:00:00Z","end_time":"0001-01-01T00:00:00Z","well_count":0}]`, testUUID, tempUUID)
	assert.Equal(suite.T(), http.StatusOK, recorder.Code)
	assert.Equal(suite.T(), output, recorder.Body.String())
	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *ExperimentHandlerTestSuite) TestListExperimentsFail() {
	suite.dbMock.On("ListExperiments", mock.Anything).Return(
		[]db.Experiment{},
		errors.New("error fetching experiments"),
	)

	recorder := makeHTTPCall(
		http.MethodGet,
		"/experiments",
		"/experiments",
		"",
		listExperimentHandler(Dependencies{Store: suite.dbMock}),
	)

	assert.Equal(suite.T(), http.StatusInternalServerError, recorder.Code)
	assert.Equal(suite.T(), "", recorder.Body.String())
	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *ExperimentHandlerTestSuite) TestCreateExperimentSuccess() {
	testUUID := uuid.New()
	tempUUID := uuid.New()
	targetUUID := uuid.New()
	suite.dbMock.On("UpsertExpTemplateTarget", mock.Anything, mock.Anything, mock.Anything).Return(
		[]db.ExpTemplateTarget{
			db.ExpTemplateTarget{ExperimentID: testUUID, TemplateID: tempUUID, TargetID: targetUUID, Threshold: 10.5},
		},
		nil,
	)
	suite.dbMock.On("ListTemplateTargets", mock.Anything, mock.Anything).Return(
		[]db.TemplateTarget{
			db.TemplateTarget{TemplateID: tempUUID, TargetID: targetUUID, Threshold: 10.5},
		},
		nil,
	)
	suite.dbMock.On("CreateExperiment", mock.Anything, mock.Anything).Return(db.Experiment{
		ID: testUUID, Description: "blah blah", TemplateID: tempUUID, TemplateName: "test", OperatorName: "ABC",
	}, nil)

	body := fmt.Sprintf(`{"description":"blah blah","template_id":"%s","operator_name":"ABC","start_time":"0001-01-01T00:00:00Z","end_time":"0001-01-01T00:00:00Z"}`, tempUUID)

	recorder := makeHTTPCall(http.MethodPost,
		"/experiments",
		"/experiments",
		body,
		createExperimentHandler(Dependencies{Store: suite.dbMock}),
	)

	output := fmt.Sprintf(`{"id":"%s","description":"blah blah","template_id":"%s","template_name":"test","operator_name":"ABC","start_time":"0001-01-01T00:00:00Z","end_time":"0001-01-01T00:00:00Z","well_count":0}`, testUUID, tempUUID)
	assert.Equal(suite.T(), http.StatusCreated, recorder.Code)
	assert.Equal(suite.T(), output, recorder.Body.String())

	suite.dbMock.AssertExpectations(suite.T())
}
func (suite *ExperimentHandlerTestSuite) TestShowExperimentSuccess() {
	testUUID := uuid.New()
	tempUUID := uuid.New()
	suite.dbMock.On("ShowExperiment", mock.Anything, mock.Anything).Return(db.Experiment{
		ID: testUUID, Description: "blah blah", TemplateID: tempUUID, TemplateName: "test", OperatorName: "ABC",
	}, nil)

	recorder := makeHTTPCall(http.MethodGet,
		"/experiments/{id}",
		"/experiments/"+testUUID.String(),
		"",
		showExperimentHandler(Dependencies{Store: suite.dbMock}),
	)
	output := fmt.Sprintf(`{"id":"%s","description":"blah blah","template_id":"%s","template_name":"test","operator_name":"ABC","start_time":"0001-01-01T00:00:00Z","end_time":"0001-01-01T00:00:00Z","well_count":0}`, testUUID, tempUUID)
	assert.Equal(suite.T(), http.StatusOK, recorder.Code)
	assert.Equal(suite.T(), output, recorder.Body.String())

	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *ExperimentHandlerTestSuite) TestRunExperimentSuccess() {
	testUUID := uuid.New()
	tempUUID := uuid.New()
	sampleID := uuid.New()
	targetID := uuid.New()
	exit := make(chan error)
	websocketMsg := make(chan string)
	websocketErr := make(chan error)

	config.Load("simulator_test")

	config.Load("config_test")

	suite.dbMock.On("ListWells", mock.Anything, mock.Anything).Return(
		[]db.Well{
			db.Well{ID: testUUID, Position: 1, SampleID: sampleID, ExperimentID: testUUID, Task: "UNKNOWN", ColorCode: "RED", Targets: []db.WellTarget{
				{WellPosition: 1,
					ExperimentID: testUUID,
					TargetID:     targetID,
					TargetName:   "COVID",
					CT:           "45"},
			}, SampleName: ""},
		},
		nil,
	)

	suite.dbMock.On("ShowExperiment", mock.Anything, mock.Anything).Return(db.Experiment{
		ID: testUUID, Description: "blah blah", TemplateID: tempUUID, OperatorName: "ABC",
	}, nil)

	stage1 := db.Stage{ID: testUUID, Type: "cycle", RepeatCount: 3, TemplateID: tempUUID, StepCount: 0}
	stage2 := db.Stage{ID: testUUID, Type: "hold", RepeatCount: 0, TemplateID: tempUUID, StepCount: 0}

	step := db.Step{TargetTemperature: 25.5, RampRate: 5.5, HoldTime: 120, DataCapture: true, StageID: testUUID}
	ss1 := db.StageStep{
		stage1, step,
	}
	ss2 := db.StageStep{
		stage2, step,
	}
	suite.dbMock.On("ListStageSteps", mock.Anything, mock.Anything).Return([]db.StageStep{
		ss1, ss2,
	}, nil)

	suite.dbMock.On("ListConfTargets", mock.Anything, mock.Anything).Return([]db.TargetDetails{}, nil)

	suite.dbMock.On("UpdateStartTimeExperiments", mock.Anything, mock.Anything, mock.Anything).Return(
		nil, nil)

	suite.dbMock.On("InsertExperimentTemperature", mock.Anything, mock.Anything).Return(
		nil, nil)

	recorder := makeHTTPCall(http.MethodGet,
		"/experiments/{experiment_id}/run",
		"/experiments/"+testUUID.String()+"/run",
		"",
		runExperimentHandler(Dependencies{Store: suite.dbMock, Plc: simulator.NewSimulator(exit), ExitCh: exit, WsErrCh: websocketErr, WsMsgCh: websocketMsg}),
	)
	<-websocketMsg // read from chn to avoid block
	assert.Equal(suite.T(), http.StatusAccepted, recorder.Code)
	assert.Equal(suite.T(), `{"code":"Warning","message":"Absence of NC,PC or NTC"}`, recorder.Body.String())

	suite.dbMock.AssertExpectations(suite.T())
}

// func (suite *ExperimentHandlerTestSuite) TestStopExperimentFail() {
// 	expUUID := uuid.New()

// 	exit := make(chan error)

// 	config.Load("simulator_test")

// 	suite.dbMock.On("UpdateStopTimeExperiments", mock.Anything, mock.Anything, mock.Anything).Return(
// 		nil, nil)
// 	recorder := makeHTTPCall(http.MethodGet,
// 		"/experiments/{experiment_id}/stop",
// 		"/experiments/"+expUUID.String()+"/stop",
// 		"",
// 		stopExperimentHandler(Dependencies{Store: suite.dbMock, Plc: simulator.NewSimulator(exit)}),
// 	)
// 	fmt.Println(recorder.Code, recorder.Body.String())
// 	fmt.Println(recorder.Body.String())

// 	assert.Equal(suite.T(), http.StatusInternalServerError, recorder.Code)
// 	assert.Equal(suite.T(), "", recorder.Body.String())

// 	suite.dbMock.AssertExpectations(suite.T())
// }
