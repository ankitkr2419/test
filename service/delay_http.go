package service

import (
	"encoding/json"
	"mylab/cpagent/db"
	"mylab/cpagent/plc"
	"mylab/cpagent/responses"
	"net/http"

	"github.com/gorilla/mux"
	logger "github.com/sirupsen/logrus"
)

func createDelayHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.InitialisedState, db.CreateOperation, "", responses.DelayInitialisedState)

		vars := mux.Vars(req)
		recipeID, err := parseUUID(vars["recipe_id"])
		// for logging error if there is any otherwise logging success
		defer func() {
			if err != nil {
				go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.ErrorState, db.CreateOperation, "", err.Error())

			} else {
				go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.CompletedState, db.CreateOperation, "", responses.DelayCompletedState)

			}

		}()

		if err != nil {
			// This error is already logged
			responseCodeAndMsg(rw, http.StatusBadRequest, ErrObj{Err: responses.RecipeIDInvalidError.Error()})
			return
		}

		var delay db.Delay
		err = json.NewDecoder(req.Body).Decode(&delay)
		if err != nil {
			logger.WithField("err", err.Error()).Errorln(responses.DelayDecodeError)
			responseCodeAndMsg(rw, http.StatusBadRequest, ErrObj{Err: responses.DelayDecodeError.Error()})
			return
		}

		valid, respBytes := validate(delay)
		if !valid {
			logger.WithField("err", "Validation Error").Errorln(responses.DelayValidationError)
			responseBadRequest(rw, respBytes)
			return
		}

		err = plc.CheckIfRecipeOrProcessSafeForCUDs(&recipeID, nil)
		if err != nil {
			responseCodeAndMsg(rw, http.StatusConflict, ErrObj{Err: err.Error()})
			logger.WithField("err", err.Error()).Error(responses.DefineCUDNotAllowedError(processC, createC))
			return
		}

		var createdTemp db.Delay
		createdTemp, err = deps.Store.CreateDelay(req.Context(), delay, recipeID)
		if err != nil {
			logger.WithField("err", err.Error()).Errorln(responses.DelayCreateError)
			responseCodeAndMsg(rw, http.StatusInternalServerError, ErrObj{Err: responses.DelayCreateError.Error()})
			return
		}

		logger.Infoln(responses.DelayCreateSuccess)
		responseCodeAndMsg(rw, http.StatusCreated, createdTemp)
	})
}

func showDelayHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		//logging when the api is initialised
		go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.InitialisedState, db.ShowOperation, "", responses.DelayInitialisedState)

		vars := mux.Vars(req)

		id, err := parseUUID(vars["id"])

		// for logging error if there is any otherwise logging success
		defer func() {
			if err != nil {
				go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.ErrorState, db.ShowOperation, "", err.Error())

			} else {
				go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.CompletedState, db.ShowOperation, "", responses.DelayCompletedState)

			}

		}()

		if err != nil {
			// This error is already logged
			responseCodeAndMsg(rw, http.StatusBadRequest, ErrObj{Err: responses.UUIDParseError.Error()})
			return
		}

		var delay db.Delay

		delay, err = deps.Store.ShowDelay(req.Context(), id)
		if err != nil {
			responseCodeAndMsg(rw, http.StatusInternalServerError, ErrObj{Err: responses.DelayFetchError.Error()})
			logger.WithField("err", err.Error()).Error(responses.DelayFetchError)
			return
		}

		logger.Infoln(responses.DelayFetchSuccess)
		responseCodeAndMsg(rw, http.StatusOK, delay)
	})
}

func updateDelayHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		//logging when the api is initialised
		go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.InitialisedState, db.UpdateOperation, "", responses.DelayInitialisedState)

		vars := mux.Vars(req)
		id, err := parseUUID(vars["id"])
		// for logging error if there is any otherwise logging success
		defer func() {
			if err != nil {
				go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.ErrorState, db.UpdateOperation, "", err.Error())

			} else {
				go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.CompletedState, db.UpdateOperation, "", responses.DelayCompletedState)

			}

		}()

		if err != nil {
			responseCodeAndMsg(rw, http.StatusBadRequest, ErrObj{Err: responses.UUIDParseError.Error()})
			return
		}

		var delay db.Delay

		err = json.NewDecoder(req.Body).Decode(&delay)
		if err != nil {
			logger.WithField("err", err.Error()).Errorln(responses.DelayDecodeError)
			responseCodeAndMsg(rw, http.StatusBadRequest, ErrObj{Err: responses.DelayDecodeError.Error()})
			return
		}

		valid, respBytes := validate(delay)
		if !valid {
			logger.WithField("err", "Validation Error").Errorln(responses.DelayValidationError)
			responseBadRequest(rw, respBytes)
			return
		}

		err = plc.CheckIfRecipeOrProcessSafeForCUDs(nil, &id)
		if err != nil {
			responseCodeAndMsg(rw, http.StatusConflict, ErrObj{Err: err.Error()})
			logger.WithField("err", err.Error()).Error(responses.DefineCUDNotAllowedError(processC, updateC))
			return
		}

		delay.ProcessID = id
		err = deps.Store.UpdateDelay(req.Context(), delay)
		if err != nil {
			logger.WithField("err", err.Error()).Error(responses.DelayUpdateError)
			responseCodeAndMsg(rw, http.StatusInternalServerError, ErrObj{Err: responses.DelayUpdateError.Error()})
			return
		}

		logger.Infoln(responses.DelayUpdateSuccess)
		responseCodeAndMsg(rw, http.StatusOK, MsgObj{Msg: responses.DelayUpdateSuccess})
	})
}
