package service

import (
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"mylab/cpagent/db"
	"mylab/cpagent/plc"
	"mylab/cpagent/responses"
	"net/http"

	"github.com/gorilla/mux"
	logger "github.com/sirupsen/logrus"
)

func createShakingHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.InitialisedState, db.CreateOperation, "", responses.ShakingInitialisedState)

		vars := mux.Vars(req)
		recipeID, err := parseUUID(vars["recipe_id"])
		// for logging error if there is any otherwise logging success
		defer func() {
			if err != nil {
				go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.ErrorState, db.CreateOperation, "", err.Error())
			} else {
				go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.CompletedState, db.CreateOperation, "", responses.ShakingCompletedState)
			}
		}()
		if err != nil {
			// This error is already logged
			responseCodeAndMsg(rw, http.StatusBadRequest, ErrObj{Err: responses.RecipeIDInvalidError.Error()})
			return
		}

		var shaObj db.Shaker
		err = json.NewDecoder(req.Body).Decode(&shaObj)
		if err != nil {
			logger.WithField("err", err.Error()).Errorln(responses.ShakingDecodeError)
			responseCodeAndMsg(rw, http.StatusBadRequest, ErrObj{Err: responses.ShakingDecodeError.Error()})
			return
		}

		valid, respBytes := validate(shaObj)
		if !valid {
			logger.WithField("err", "Validation Error").Errorln(responses.ShakingValidationError)
			responseBadRequest(rw, respBytes)
			return
		}

		err = ValidateShakingObject(req.Context(), deps, shaObj, recipeID)
		if err != nil {
			responseCodeAndMsg(rw, http.StatusBadRequest, ErrObj{Err: err.Error()})
			logger.WithField("err", err.Error()).Error(err.Error())
			return
		}

		err = plc.CheckIfRecipeOrProcessSafeForCUDs(&recipeID, nil)
		if err != nil {
			responseCodeAndMsg(rw, http.StatusConflict, ErrObj{Err: err.Error()})
			logger.WithField("err", err.Error()).Error(responses.DefineCUDNotAllowedError(processC, createC))
			return
		}

		var shaker db.Shaker
		shaker, err = deps.Store.CreateShaking(req.Context(), shaObj, recipeID)
		if err != nil {
			logger.WithField("err", err.Error()).Errorln(responses.ShakingCreateError)
			responseCodeAndMsg(rw, http.StatusInternalServerError, ErrObj{Err: responses.ShakingCreateError.Error()})
			return
		}
		logger.Infoln(responses.ShakingCreateSuccess)
		responseCodeAndMsg(rw, http.StatusCreated, shaker)
	})
}

func showShakingHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		//logging when the api is initialised

		go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.InitialisedState, db.ShowOperation, "", responses.ShakingInitialisedState)

		vars := mux.Vars(req)
		id, err := parseUUID(vars["id"])

		// for logging error if there is any otherwise logging success
		defer func() {
			if err != nil {
				go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.ErrorState, db.ShowOperation, "", err.Error())

			} else {
				go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.CompletedState, db.ShowOperation, "", responses.ShakingCompletedState)

			}

		}()

		if err != nil {
			responseCodeAndMsg(rw, http.StatusBadRequest, ErrObj{Err: responses.UUIDParseError.Error()})
			return
		}

		var shaking db.Shaker
		shaking, err = deps.Store.ShowShaking(req.Context(), id)
		if err != nil {
			responseCodeAndMsg(rw, http.StatusInternalServerError, ErrObj{Err: responses.ShakingFetchError.Error()})
			logger.WithField("err", err.Error()).Error(responses.ShakingFetchError)
			return
		}

		logger.Infoln(responses.ShakingFetchSuccess)
		responseCodeAndMsg(rw, http.StatusOK, shaking)
	})
}

func updateShakingHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		//logging when the api is initialised

		go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.InitialisedState, db.UpdateOperation, "", responses.ShakingInitialisedState)

		vars := mux.Vars(req)
		id, err := parseUUID(vars["id"])
		// for logging error if there is any otherwise logging success
		defer func() {
			if err != nil {
				go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.ErrorState, db.UpdateOperation, "", err.Error())

			} else {
				go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.CompletedState, db.UpdateOperation, "", responses.ShakingCompletedState)

			}

		}()

		if err != nil {
			responseCodeAndMsg(rw, http.StatusBadRequest, ErrObj{Err: responses.UUIDParseError.Error()})
			return
		}
		var shObj db.Shaker
		err = json.NewDecoder(req.Body).Decode(&shObj)
		if err != nil {
			logger.WithField("err", err.Error()).Errorln(responses.ShakingDecodeError)
			responseCodeAndMsg(rw, http.StatusBadRequest, ErrObj{Err: responses.ShakingDecodeError.Error()})
			return
		}
		valid, respBytes := validate(shObj)
		if !valid {
			logger.WithField("err", "Validation Error").Errorln(responses.ShakingValidationError)
			responseBadRequest(rw, respBytes)
			return
		}

		process, err := deps.Store.ShowProcess(req.Context(), id)
		if err != nil {
			logger.WithField("err", err.Error()).Error(responses.ProcessFetchError)
			responseCodeAndMsg(rw, http.StatusInternalServerError, ErrObj{Err: responses.ProcessFetchError.Error()})
			return
		}

		err = ValidateShakingObject(req.Context(), deps, shObj, process.RecipeID)
		if err != nil {
			logger.WithField("err", err.Error()).Error(err.Error())
			responseCodeAndMsg(rw, http.StatusBadRequest, ErrObj{Err: err.Error()})
			return
		}

		err = plc.CheckIfRecipeOrProcessSafeForCUDs(nil, &id)
		if err != nil {
			responseCodeAndMsg(rw, http.StatusConflict, ErrObj{Err: err.Error()})
			logger.WithField("err", err.Error()).Error(responses.DefineCUDNotAllowedError(processC, updateC))
			return
		}

		shObj.ProcessID = id
		err = deps.Store.UpdateShaking(req.Context(), shObj)
		if err != nil {
			logger.WithField("err", err.Error()).Error(responses.ShakingUpdateError)
			responseCodeAndMsg(rw, http.StatusInternalServerError, ErrObj{Err: responses.ShakingUpdateError.Error()})
			return
		}

		logger.Infoln(responses.ShakingUpdateSuccess)
		responseCodeAndMsg(rw, http.StatusOK, MsgObj{Msg: responses.ShakingUpdateSuccess})
	})
}

// TODO: ValidateShakingObject to be also called from CSV
func ValidateShakingObject(ctx context.Context, deps Dependencies, sh db.Shaker, recipeID uuid.UUID) (err error) {
	// Currently only validating Temperature at here
	if !sh.WithTemp {
		return
	}

	if sh.Temperature < minShakerTempAllowed || sh.Temperature > maxShakerTempAllowed {
		return responses.InvalidShakerTemp
	}
	return
}
