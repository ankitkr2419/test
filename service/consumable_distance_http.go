package service

import (
	"context"
	"encoding/json"
	"mylab/cpagent/db"
	"mylab/cpagent/plc"
	"mylab/cpagent/responses"
	"net/http"

	"github.com/gorilla/mux"
	logger "github.com/sirupsen/logrus"
)

func createConsumableDistanceHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		var cd db.ConsumableDistance
		err := json.NewDecoder(req.Body).Decode(&cd)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			logger.WithField("err", err.Error()).Error("Error while decoding Consumable Distance data")
			return
		}

		valid, respBytes := Validate(cd)
		if !valid {
			responseBadRequest(rw, respBytes)
			return
		}

		err = deps.Store.InsertConsumableDistance(req.Context(), []db.ConsumableDistance{cd})
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			logger.WithField("err", err.Error()).Error("Error while inserting comsumable distance")
			return
		}

		respBytes, err = json.Marshal(cd)
		if err != nil {
			logger.WithField("err", err.Error()).Error("Error marshaling Consumable Distance data")
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}
		rw.Header().Add("Content-Type", "application/json")
		rw.WriteHeader(http.StatusCreated)
		rw.Write(respBytes)
	})
}

func listConsumableDistanceHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		//logging when the api is initialised
		go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.InitialisedState, db.ShowOperation, "", responses.ConsumableDistanceInitialisedState)

		var ConsumableDistance []db.ConsumableDistance
		ConsumableDistance, err := deps.Store.ListConsDistances()

		// for logging error if there is any otherwise logging success
		defer func() {
			if err != nil {
				go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.ErrorState, db.ShowOperation, "", err.Error())
			} else {
				go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.CompletedState, db.ShowOperation, "", responses.ConsumableDistanceCompletedState)
			}
		}()
		if err != nil {
			responseCodeAndMsg(rw, http.StatusInternalServerError, ErrObj{Err: responses.ConsumableDistanceFetchError.Error()})
			logger.WithField("err", err.Error()).Error(responses.ConsumableDistanceFetchError)
			return
		}

		logger.Infoln(responses.ConsumableDistanceFetchSuccess)
		responseCodeAndMsg(rw, http.StatusOK, ConsumableDistance)
	})
}

func listCalibrationsHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		deck := req.URL.Query().Get("deck")

		//logging when the api is initialised
		go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.InitialisedState, db.ShowOperation, "", responses.ConsumableDistanceInitialisedState)

		var min, max int64

		switch deck {
		case plc.DeckA:
			min = deckAMinCalibID
			max = deckAMaxCalibID
		case plc.DeckB:
			min = deckBMinCalibID
			max = deckBMaxCalibID
		default:
			responseCodeAndMsg(rw, http.StatusBadRequest, ErrObj{Err: responses.DeckNameInvalid.Error()})
			logger.WithField("err", "Invalid Deck").Error(responses.DeckNameInvalid)
			return
		}

		calibrations, err := deps.Store.ListConsDistancesDeck(req.Context(), min, max)
		// for logging error if there is any otherwise logging success
		defer func() {
			if err != nil {
				go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.ErrorState, db.ShowOperation, "", err.Error())
			} else {
				go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.CompletedState, db.ShowOperation, "", responses.CalibrationCompletedState)
			}
		}()
		if err != nil {
			responseCodeAndMsg(rw, http.StatusInternalServerError, ErrObj{Err: responses.CalibrationsFetchError.Error()})
			logger.WithField("err", err.Error()).Error(responses.CalibrationsFetchError)
			return
		}

		logger.Infoln(responses.CalibrationsFetchSuccess)
		responseCodeAndMsg(rw, http.StatusOK, calibrations)
	})
}

func updateConsumableDistanceHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		//logging when the api is initialised
		go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.InitialisedState, db.UpdateOperation, "", responses.ConsumableDistanceInitialisedState)
		var adobj db.ConsumableDistance

		err := json.NewDecoder(req.Body).Decode(&adobj)
		// for logging error if there is any otherwise logging success
		defer func() {
			if err != nil {
				go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.ErrorState, db.UpdateOperation, "", err.Error())
			} else {
				go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.CompletedState, db.UpdateOperation, "", responses.ConsumableDistanceCompletedState)
			}
		}()

		if err != nil {
			logger.WithField("err", err.Error()).Errorln(responses.ConsumableDistanceDecodeError)
			responseCodeAndMsg(rw, http.StatusBadRequest, ErrObj{Err: responses.ConsumableDistanceDecodeError.Error()})
			return
		}
		err = db.UpdateConsumableDistancesValues([]db.ConsumableDistance{adobj})
		if err != nil {
			logger.WithField("err", err.Error()).Error(responses.ConsumableDistanceUpdateConfigError)
			responseCodeAndMsg(rw, http.StatusInternalServerError, ErrObj{Err: responses.ConsumableDistanceUpdateConfigError.Error()})
			return
		}

		err = deps.Store.UpdateConsumableDistance(req.Context(), adobj)
		if err != nil {
			logger.WithField("err", err.Error()).Error(responses.ConsumableDistanceUpdateError)
			responseCodeAndMsg(rw, http.StatusInternalServerError, ErrObj{Err: responses.ConsumableDistanceUpdateError.Error()})
			return
		}

		logger.Infoln(responses.ConsumableDistanceUpdateSuccess)
		responseCodeAndMsg(rw, http.StatusOK, MsgObj{Msg: responses.ConsumableDistanceUpdateSuccess})
	})
}

// Sense and Set Calibration
func updateCalibrationsHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		//logging when the api is initialised
		go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.InitialisedState, db.UpdateOperation, "", responses.CalibrationInitialisedState)

		var m Manual
		vars := mux.Vars(req)
		deck := vars["deck"]

		err := json.NewDecoder(req.Body).Decode(&m)
		// for logging error if there is any otherwise logging success
		defer func() {
			if err != nil {
				go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.ErrorState, db.UpdateOperation, "", err.Error())
			} else {
				go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.CompletedState, db.UpdateOperation, "", responses.CalibrationCompletedState)
			}
		}()

		if err != nil {
			logger.WithField("err", err.Error()).Errorln(responses.CalibrationDecodeError)
			responseCodeAndMsg(rw, http.StatusBadRequest, ErrObj{Err: responses.CalibrationDecodeError.Error()})
			return
		}

		dN := plc.DeckNumber{Deck: deck, Number: uint16(m.MotorNum)}

		// Logic is different per calib position

		calibrations, err := calculateConsIDAndPosition(req.Context(), deps.Store, dN)
		if err != nil {
			logger.WithField("err", err.Error()).Error(responses.CalibrationsCalculateError)
			responseCodeAndMsg(rw, http.StatusInternalServerError, ErrObj{Err: responses.CalibrationsCalculateError.Error()})
			return
		}

		err = db.UpdateConsumableDistancesValues(calibrations)
		if err != nil {
			logger.WithField("err", err.Error()).Error(responses.CalibrationUpdateConfigError)
			responseCodeAndMsg(rw, http.StatusInternalServerError, ErrObj{Err: responses.CalibrationUpdateConfigError.Error()})
			return
		}

		err = deps.Store.UpdateConsumableDistance(req.Context(), calibrations[0])
		if err != nil {
			logger.WithField("err", err.Error()).Error(responses.CalibrationUpdateError)
			responseCodeAndMsg(rw, http.StatusInternalServerError, ErrObj{Err: responses.CalibrationUpdateError.Error()})
			return
		}

		logger.Infoln(responses.CalibrationUpdateSuccess)
		responseCodeAndMsg(rw, http.StatusOK, MsgObj{Msg: responses.CalibrationUpdateSuccess})
	})
}

func calculateConsIDAndPosition(ctx context.Context, store db.Storer, dN plc.DeckNumber) (calibrations []db.ConsumableDistance, err error) {
	logger.Infoln("Inside calculateConsIDAndPosition for ", dN)
	var consID uint16
	switch dN.Deck {
	case plc.DeckA:
		consID = deckAMinCalibID - 1 + dN.Number
	case plc.DeckB:
		consID = deckBMinCalibID - 1 + dN.Number
	}

	// Resuing ListConsDistancesDeck for single ID
	calibrations, err = store.ListConsDistancesDeck(ctx, int64(consID), int64(consID))
	if len(calibrations) == 0 {
		err = responses.CalibrationVariableMissingError
		return
	}

	// These will be updated calibrations
	calibrations, err = plc.CalculatePosition(ctx, calibrations, dN)
	if err != nil {
		err = responses.CalibrationsPositionCalculateError
		return
	}

	return
}
