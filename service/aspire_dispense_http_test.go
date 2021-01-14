package service

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
type AspireDispenseHandlerTestSuite struct {
	suite.Suite

	dbMock *db.DBMockStore
}

func (suite *AspireDispenseHandlerTestSuite) SetupTest() {
	suite.dbMock = &db.DBMockStore{}
}

func TestAspireDispenseTestSuite(t *testing.T) {
	suite.Run(t, new(AspireDispenseHandlerTestSuite))
}

func (suite *AspireDispenseHandlerTestSuite) TestCreateAspireDispenseSuccess() {
	testUUID := uuid.New()
	suite.dbMock.On("CreateAspireDispense", mock.Anything, mock.Anything).Return(db.AspireDispense{
		ID:                 testUUID,
		Category:           "well_to_well",
		WellNoSource:       1,
		AspireHeight:       2,
		AspireMixingVolume: 3,
		AspireNoOfCycles:   4,
		AspireVolume:       5,
		DispenseHeight:     6,
		DispenseMixingVol:  7,
		DispenseNoOfCycles: 8,
		DispenseVol:        9,
		DispenseBlow:       10,
		WellToDestination:  11,
	}, nil)

	body := fmt.Sprintf(`{"id":"%s","category":"well_to_well","well_no_source":1,"aspire_height":2,"aspire_mixing_volume":3,"aspire_no_of_cycles":4,"aspire_volume":5,"dispense_height":6,"dispense_mixing_volume":7,"dispense_no_of_cycles":8,"dispense_vol":9,"dispense_blow":10,"well_to_destination":11,"created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z"}`, testUUID)
	recorder := makeHTTPCall(http.MethodPost,
		"/aspireDispense",
		"/aspireDispense",
		body,
		createAspireDispenseHandler(Dependencies{Store: suite.dbMock}),
	)
	output := fmt.Sprintf(`{"id":"%s","category":"well_to_well","well_no_source":1,"aspire_height":2,"aspire_mixing_volume":3,"aspire_no_of_cycles":4,"aspire_volume":5,"dispense_height":6,"dispense_mixing_volume":7,"dispense_no_of_cycles":8,"dispense_vol":9,"dispense_blow":10,"well_to_destination":11,"created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z"}`, testUUID)

	assert.Equal(suite.T(), http.StatusCreated, recorder.Code)
	assert.Equal(suite.T(), output, recorder.Body.String())

	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *AspireDispenseHandlerTestSuite) TestCreateAspireDispenseFailure() {
	testUUID := uuid.New()
	suite.dbMock.On("CreateAspireDispense", mock.Anything, mock.Anything).Return(db.AspireDispense{}, fmt.Errorf("Error creating aspire dispense"))

	body := fmt.Sprintf(`{"id":"%s","category":"well_to_well","well_no_source":1,"aspire_height":2,"aspire_mixing_volume":3,"aspire_no_of_cycles":4,"aspire_volume":5,"dispense_height":6,"dispense_mixing_volume":7,"dispense_no_of_cycles":8,"dispense_vol":9,"dispense_blow":10,"well_to_destination":11,"created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z"}`, testUUID)

	recorder := makeHTTPCall(http.MethodPost,
		"/aspireDispense",
		"/aspireDispense",
		body,
		createAspireDispenseHandler(Dependencies{Store: suite.dbMock}),
	)
	output := fmt.Errorf("Error creating aspire dispense")

	assert.Equal(suite.T(), http.StatusInternalServerError, recorder.Code)
	assert.NotEqual(suite.T(), output, recorder.Body.String())

	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *AspireDispenseHandlerTestSuite) TestListAspireDispenseSuccess() {
	testUUID := uuid.New()
	suite.dbMock.On("ListAspireDispense", mock.Anything, mock.Anything).Return(
		[]db.AspireDispense{
			db.AspireDispense{
				ID:                 testUUID,
				Category:           "well_to_well",
				WellNoSource:       1,
				AspireHeight:       2,
				AspireMixingVolume: 3,
				AspireNoOfCycles:   4,
				AspireVolume:       5,
				DispenseHeight:     6,
				DispenseMixingVol:  7,
				DispenseNoOfCycles: 8,
				DispenseVol:        9,
				DispenseBlow:       10,
				WellToDestination:  11,
			},
		}, nil)

	recorder := makeHTTPCall(
		http.MethodGet,
		"/aspireDispense",
		"/aspireDispense",
		"",
		listAspireDispenseHandler(Dependencies{Store: suite.dbMock}),
	)
	output := fmt.Sprintf(`[{"id":"%s","category":"well_to_well","well_no_source":1,"aspire_height":2,"aspire_mixing_volume":3,"aspire_no_of_cycles":4,"aspire_volume":5,"dispense_height":6,"dispense_mixing_volume":7,"dispense_no_of_cycles":8,"dispense_vol":9,"dispense_blow":10,"well_to_destination":11,"created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z"}]`, testUUID)
	assert.Equal(suite.T(), http.StatusOK, recorder.Code)
	assert.Equal(suite.T(), output, recorder.Body.String())
	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *AspireDispenseHandlerTestSuite) TestListAspireDispenseFailure() {
	suite.dbMock.On("ListAspireDispense", mock.Anything, mock.Anything).Return(
		[]db.AspireDispense{}, fmt.Errorf("Error fetching aspire dispense"))

	recorder := makeHTTPCall(
		http.MethodGet,
		"/aspireDispense",
		"/aspireDispense",
		"",
		listAspireDispenseHandler(Dependencies{Store: suite.dbMock}),
	)
	output := ""
	assert.Equal(suite.T(), http.StatusInternalServerError, recorder.Code)
	assert.Equal(suite.T(), output, recorder.Body.String())
	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *AspireDispenseHandlerTestSuite) TestShowAspireDispenseSuccess() {
	testUUID := uuid.New()
	suite.dbMock.On("ShowAspireDispense", mock.Anything, mock.Anything).Return(db.AspireDispense{
		ID:                 testUUID,
		Category:           "well_to_well",
		WellNoSource:       1,
		AspireHeight:       2,
		AspireMixingVolume: 3,
		AspireNoOfCycles:   4,
		AspireVolume:       5,
		DispenseHeight:     6,
		DispenseMixingVol:  7,
		DispenseNoOfCycles: 8,
		DispenseVol:        9,
		DispenseBlow:       10,
		WellToDestination:  11,
	}, nil)

	recorder := makeHTTPCall(http.MethodGet,
		"/aspireDispense/{id}",
		"/aspireDispense/"+testUUID.String(),
		"",
		showAspireDispenseHandler(Dependencies{Store: suite.dbMock}),
	)
	output := fmt.Sprintf(`{"id":"%s","category":"well_to_well","well_no_source":1,"aspire_height":2,"aspire_mixing_volume":3,"aspire_no_of_cycles":4,"aspire_volume":5,"dispense_height":6,"dispense_mixing_volume":7,"dispense_no_of_cycles":8,"dispense_vol":9,"dispense_blow":10,"well_to_destination":11,"created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z"}`, testUUID)
	assert.Equal(suite.T(), http.StatusOK, recorder.Code)
	assert.Equal(suite.T(), output, recorder.Body.String())

	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *AspireDispenseHandlerTestSuite) TestUpdateAspireDispenseSuccess() {
	testUUID := uuid.New()
	suite.dbMock.On("UpdateAspireDispense", mock.Anything, mock.Anything).Return(db.AspireDispense{
		ID:                 testUUID,
		Category:           "well_to_well",
		WellNoSource:       1,
		AspireHeight:       2,
		AspireMixingVolume: 3,
		AspireNoOfCycles:   4,
		AspireVolume:       5,
		DispenseHeight:     6,
		DispenseMixingVol:  7,
		DispenseNoOfCycles: 8,
		DispenseVol:        9,
		DispenseBlow:       10,
		WellToDestination:  11,
	}, nil)

	body := fmt.Sprintf(`{"id":"%s","category":"well_to_well","well_no_source":1,"aspire_height":2,"aspire_mixing_volume":3,"aspire_no_of_cycles":4,"aspire_volume":5,"dispense_height":6,"dispense_mixing_volume":7,"dispense_no_of_cycles":8,"dispense_vol":9,"dispense_blow":10,"well_to_destination":11,"created_at":"0001-01-01T00:00:00Z","updated_at":"0001-01-01T00:00:00Z"}`, testUUID)

	recorder := makeHTTPCall(http.MethodPut,
		"/aspireDispense/{id}",
		"/aspireDispense/"+testUUID.String(),
		body,
		updateAspireDispenseHandler(Dependencies{Store: suite.dbMock}),
	)

	assert.Equal(suite.T(), http.StatusOK, recorder.Code)
	assert.Equal(suite.T(), `{"msg":"aspire dispense updated successfully"}`, recorder.Body.String())

	suite.dbMock.AssertExpectations(suite.T())
}

func (suite *AspireDispenseHandlerTestSuite) TestDeleteAspireDispenseSuccess() {
	testUUID := uuid.New()
	suite.dbMock.On("DeleteAspireDispense", mock.Anything, mock.Anything).Return(
		testUUID,
		nil)

	recorder := makeHTTPCall(http.MethodDelete,
		"/aspireDispense/{id}",
		"/aspireDispense/"+testUUID.String(),
		"",
		deleteAspireDispenseHandler(Dependencies{Store: suite.dbMock}),
	)
	assert.Equal(suite.T(), http.StatusOK, recorder.Code)
	assert.Equal(suite.T(), `{"msg":"aspire dispense deleted successfully"}`, recorder.Body.String())

	suite.dbMock.AssertExpectations(suite.T())
}
