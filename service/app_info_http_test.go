package service

import (
	"encoding/json"
	"mylab/cpagent/db"
	"mylab/cpagent/plc"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

// Define the suite, and absorb the built-in basic suite
// functionality from testify - including assertion methods.
type AppInfoHandlerTestSuite struct {
	suite.Suite
	dbMock  *db.DBMockStore
	plcMock map[string]plc.Extraction
}

func (suite *AppInfoHandlerTestSuite) SetupTest() {
	suite.dbMock = &db.DBMockStore{}
	suite.plcMock = map[string]plc.Extraction{
		plc.DeckA: &plc.PLCMockStore{},
		plc.DeckB: &plc.PLCMockStore{},
	}
	suite.dbMock.On("AddAuditLog", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil).Maybe()
}

func (suite *AppInfoHandlerTestSuite) TearDownTest() {
	suite.dbMock.AssertExpectations(suite.T())
}

func TestAppInfoTestSuite(t *testing.T) {
	suite.Run(t, new(AppInfoHandlerTestSuite))
}

func (suite *AppInfoHandlerTestSuite) TestApplication() {

	t := suite.T()

	t.Run("Application Info Fetch", func(t *testing.T) {
		// Setup initial values as these are setup at compile time
		Application = RTPCR
		Version = "1.3.1"
		User = "josh"
		BuiltOn = "Mon Jun 14 11:49:45 IST 2021"
		CommitID = "51351df30541ce089316f955b6a249eadc0f0fcf"
		Branch = "story/application_info"
		Machine = "dell-latitude-5490"

		var appInfo = struct {
			Application string `json:"app"`
			Version     string `json:"version"`
			User        string `json:"user"`
			Machine     string `json:"machine"`
			CommitID    string `json:"commit_id"`
			Branch      string `json:"branch"`
			BuiltOn     string `json:"built_on"`
		}{
			Application: Application,
			Version:     Version,
			User:        User,
			Machine:     Machine,
			CommitID:    CommitID,
			Branch:      Branch,
			BuiltOn:     BuiltOn,
		}

		recorder := makeHTTPCall(http.MethodGet,
			"/app-info",
			"/app-info",
			"",
			appInfoHandler(Dependencies{Store: suite.dbMock}),
		)

		body, _ := json.Marshal(appInfo)

		assert.Equal(suite.T(), http.StatusOK, recorder.Code)
		assert.Equal(suite.T(), string(body), recorder.Body.String())
	})

	t.Run("Print Binary Info", func(t *testing.T) {
		PrintBinaryInfo()
	})

	/*
		t.Run("Shut Down Extraction Gracefully success", func(t *testing.T) {

			suite.plcMock[plc.DeckA].(*plc.PLCMockStore).On("SwitchOffAllCoils").Return("SUCCESS", nil).Once()
			suite.plcMock[plc.DeckB].(*plc.PLCMockStore).On("SwitchOffAllCoils").Return("SUCCESS", errors.New("Some")).Once()

			Application = Extraction
			err := ShutDownGracefully(Dependencies{Store: suite.dbMock, PlcDeck: suite.plcMock})
			assert.NotNil(suite.T(), err)
		})

		t.Run("Shut Down RTPCR Gracefully success", func(t *testing.T) {

		})

		t.Run("Reach Room Temp RTPCR failure", func(t *testing.T) {

		})

		t.Run("Lid Temp Switch Off RTPCR failure", func(t *testing.T) {

		})
	*/

}
