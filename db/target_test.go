package db

import (
	"context"
	"mylab/cpagent/config"
	"testing"

	"github.com/google/uuid"
	logger "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// Define the suite, and absorb the built-in basic suite
// functionality from testify - including assertion methods.
type TargetHandlerTestSuite struct {
	suite.Suite
	dbStore Storer
}

func TestTargetTestSuite(t *testing.T) {
	suite.Run(t, new(TargetHandlerTestSuite))
}

func (suite *TargetHandlerTestSuite) SetupTargetSuite() {
	config.Load("application_test")

	err := RunMigrations()
	if err != nil {
		logger.WithField("err", err.Error()).Error("Database init failed")
	}

	store, err := Init()
	if err != nil {
		logger.WithField("err", err.Error()).Error("Database init failed")
		return
	}
	suite.dbStore = store
}

func (suite *TargetHandlerTestSuite) TestTargetSuccess() {
	testUUID := uuid.New()
	// test create badge
	expectedT := Target{
		Name: "test targat",
		ID:   testUUID,
	}
	var err error
	createdT, err := suite.dbStore.CreateTarget(context.Background(), expectedT)

	assert.Nil(suite.T(), err)

	expectedT.ID = createdT.ID

	assert.Equal(suite.T(), expectedT, createdT)

	// test list badge
	var targetList []Target
	targetList, err = suite.dbStore.ListTarget(context.Background())

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), targetList, []Target{createdT})

}
