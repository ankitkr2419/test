package service

import (
	"fmt"
	"mylab/cpagent/db"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

// Define the suite, and absorb the built-in basic suite
// functionality from testify - including assertion methods.
type TipTubeHandlerTestSuite struct {
	suite.Suite

	dbMock *db.DBMockStore
}

func (suite *TipTubeHandlerTestSuite) SetupTest() {
	suite.dbMock = &db.DBMockStore{}
}

func TestTipTubeTestSuite(t *testing.T) {
	suite.Run(t, new(TipTubeHandlerTestSuite))
}

func (suite *TipTubeHandlerTestSuite) TestCreateTipTubeSuccess() {
	suite.dbMock.On("InsertTipsTubes", mock.Anything, mock.Anything).Return(db.TipsTubes{
		LabwareID: 1, ConsumabledistanceID: 103, Name: "piercing_tip", Volume: 700.11, Height: 93.11,
	}, nil)

	body := fmt.Sprintf(`{"labware_id":1,"consumable_distance_id":103,"name":"piercing_tip","volume":700.11,"height":93.11,"created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z"}`)

	recorder := makeHTTPCall(http.MethodPost,
		"/tiptube",
		"/tiptube",
		body,
		createTipTubeHandler(Dependencies{Store: suite.dbMock}),
	)

	output := fmt.Sprintf(`{"labware_id":1,"consumable_distance_id":103,"name":"piercing_tip","volume":700.11,"height":93.11,"created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z"}`)

	assert.Equal(suite.T(), http.StatusCreated, recorder.Code)
	assert.Equal(suite.T(), output, recorder.Body.String())

	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *TipTubeHandlerTestSuite) TestCreateTipTubeFailure() {
	suite.dbMock.On("InsertTipsTubes", mock.Anything, mock.Anything).Return(db.TipsTubes{
		LabwareID: 1, ConsumabledistanceID: 103, Name: "piercing_tip", Volume: 700.11, Height: 93.11,
	}, nil)

	body := fmt.Sprintf(`{"labware_id":1,"consumable_distance_id":103,"name":"piercing_tip","volume":700.11,"height":93.11,"created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z"}`)

	recorder := makeHTTPCall(http.MethodPost,
		"/tiptube",
		"/tiptube",
		body,
		createTipTubeHandler(Dependencies{Store: suite.dbMock}),
	)

	output := fmt.Sprintf(`{"labware_id":1,"consumable_distance_id":103,"name":"piercing_tip","volume":600.11,"height":93.11,"created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z"}`)

	assert.Equal(suite.T(), http.StatusCreated, recorder.Code)
	assert.NotEqual(suite.T(), output, recorder.Body.String())

	suite.dbMock.AssertExpectations(suite.T())
}
