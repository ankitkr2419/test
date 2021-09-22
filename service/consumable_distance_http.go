package service

import (
	"encoding/json"
	"mylab/cpagent/db"
	"mylab/cpagent/responses"
	"net/http"

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
