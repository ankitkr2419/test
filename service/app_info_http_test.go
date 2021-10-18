package service

import (
	"encoding/json"
	"mylab/cpagent/db"
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
	dbMock *db.DBMockStore
}

func (suite *AppInfoHandlerTestSuite) SetupTest() {
	suite.dbMock = &db.DBMockStore{}
	suite.dbMock.On("AddAuditLog", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil).Maybe()
}

func TestAppInfoTestSuite(t *testing.T) {
	suite.Run(t, new(AppInfoHandlerTestSuite))
}

func (suite *AppInfoHandlerTestSuite) TestAppInfoHandler() {

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
	t := suite.T()
	t.Run("when appinfo is fetched successfully", func(t *testing.T) {

		recorder := makeHTTPCall(http.MethodPost,
			"/app-info",
			"/app-info",
			"",
			appInfoHandler(Dependencies{Store: suite.dbMock}),
		)
		body, _ := json.Marshal(appInfo)
		// TODO: Unmarshal rcorder.Body into map[string]string and compare for msg and Role
		assert.Equal(suite.T(), http.StatusOK, recorder.Code)
		assert.Equal(suite.T(), string(body), recorder.Body.String())

		suite.dbMock.AssertExpectations(suite.T())
	})

}
