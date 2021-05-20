package service

import (
	"encoding/json"
	"mylab/cpagent/responses"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// Define the suite, and absorb the built-in basic suite
// functionality from testify - including assertion methods.
type PingHandlerTestSuite struct {
	suite.Suite
}

func TestPingTestSuite(t *testing.T) {
	suite.Run(t, new(PingHandlerTestSuite))
}


func (suite *PingHandlerTestSuite) TestPingSuccess(){

	recorder := makeHTTPCall(http.MethodGet,
		"/ping",
		"/ping",
		"",
		pingHandler,
	)
	output := MsgObj{Msg: responses.Pong}
	outputBytes, _ := json.Marshal(output)

	assert.Equal(suite.T(), http.StatusOK, recorder.Code)
	assert.Equal(suite.T(), outputBytes, recorder.Body.Bytes())

}