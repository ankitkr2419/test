package service

import (
	"context"
	"encoding/json"
	"mylab/cpagent/db"
	"mylab/cpagent/plc"
	"mylab/cpagent/responses"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	logger "github.com/sirupsen/logrus"
)

func createTipOperationHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.InitialisedState, db.CreateOperation, "", responses.TipOperationInitialisedState)

		vars := mux.Vars(req)

		recipeID, err := parseUUID(vars["recipe_id"])
		// for logging error if there is any otherwise logging success
		defer func() {
			if err != nil {
				go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.ErrorState, db.CreateOperation, "", err.Error())
			} else {
				go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.CompletedState, db.CreateOperation, "", responses.TipOperationCompletedState)
			}

		}()

		if err != nil {
			// This error is already logged
			responseCodeAndMsg(rw, http.StatusBadRequest, ErrObj{Err: responses.RecipeIDInvalidError.Error()})
			return
		}

		var tipOpr db.TipOperation
		err = json.NewDecoder(req.Body).Decode(&tipOpr)
		if err != nil {
			logger.WithField("err", err.Error()).Errorln(responses.TipOperationDecodeError)
			responseCodeAndMsg(rw, http.StatusBadRequest, ErrObj{Err: responses.TipOperationDecodeError.Error()})
			return
		}

		valid, respBytes := validate(tipOpr)
		if !valid {
			logger.WithField("err", "Validation Error").Errorln(responses.TipOperationValidationError)
			responseBadRequest(rw, respBytes)
			return
		}
		if tipOpr.Type == db.PickupTip {
			err = ValidateTipPickupObject(req.Context(), deps, tipOpr, recipeID)
			if err != nil {
				logger.WithField("err", err.Error()).Error(err.Error())
				responseCodeAndMsg(rw, http.StatusInternalServerError, ErrObj{Err: err.Error()})
				return
			}
		}
		err = plc.CheckIfRecipeOrProcessSafeForCUDs(&recipeID, nil)
		if err != nil {
			responseCodeAndMsg(rw, http.StatusConflict, ErrObj{Err: err.Error()})
			logger.WithField("err", err.Error()).Error(responses.DefineCUDNotAllowedError(processC, createC))
			return
		}

		var tipOperation db.TipOperation
		tipOperation, err = deps.Store.CreateTipOperation(req.Context(), tipOpr, recipeID)
		if err != nil {
			logger.WithField("err", err.Error()).Errorln(responses.TipOperationCreateError)
			responseCodeAndMsg(rw, http.StatusInternalServerError, ErrObj{Err: responses.TipOperationCreateError.Error()})
			return
		}
		logger.Infoln(responses.TipOperationCreateSuccess)
		responseCodeAndMsg(rw, http.StatusCreated, tipOperation)
	})
}

func showTipOperationHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		//logging when the api is initialised
		go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.InitialisedState, db.ShowOperation, "", responses.TipOperationInitialisedState)

		vars := mux.Vars(req)

		id, err := parseUUID(vars["id"])
		// for logging error if there is any otherwise logging success
		defer func() {
			if err != nil {
				go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.ErrorState, db.ShowOperation, "", err.Error())

			} else {
				go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.CompletedState, db.ShowOperation, "", responses.TipOperationCompletedState)

			}

		}()

		if err != nil {
			responseCodeAndMsg(rw, http.StatusBadRequest, ErrObj{Err: responses.UUIDParseError.Error()})
			return
		}

		var tipOperation db.TipOperation

		tipOperation, err = deps.Store.ShowTipOperation(req.Context(), id)
		if err != nil {
			responseCodeAndMsg(rw, http.StatusInternalServerError, ErrObj{Err: responses.TipOperationFetchError.Error()})
			logger.WithField("err", err.Error()).Error(responses.TipOperationFetchError)
			return
		}

		logger.Infoln(responses.TipOperationFetchSuccess)
		responseCodeAndMsg(rw, http.StatusOK, tipOperation)
	})
}

func updateTipOperationHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		//logging when the api is initialised
		go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.InitialisedState, db.UpdateOperation, "", responses.TipOperationInitialisedState)

		vars := mux.Vars(req)
		id, err := parseUUID(vars["id"])

		// for logging error if there is any otherwise logging success
		defer func() {
			if err != nil {
				go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.ErrorState, db.UpdateOperation, "", err.Error())

			} else {
				go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.CompletedState, db.UpdateOperation, "", responses.TipOperationCompletedState)

			}

		}()

		if err != nil {
			responseCodeAndMsg(rw, http.StatusBadRequest, ErrObj{Err: responses.UUIDParseError.Error()})
			return
		}

		var tipOpr db.TipOperation

		err = json.NewDecoder(req.Body).Decode(&tipOpr)
		if err != nil {
			logger.WithField("err", err.Error()).Errorln(responses.TipOperationDecodeError)
			responseCodeAndMsg(rw, http.StatusBadRequest, ErrObj{Err: responses.TipOperationDecodeError.Error()})
			return
		}

		valid, respBytes := validate(tipOpr)
		if !valid {
			logger.WithField("err", "Validation Error").Errorln(responses.TipOperationValidationError)
			responseBadRequest(rw, respBytes)
			return
		}

		process, err := deps.Store.ShowProcess(req.Context(), id)
		if err != nil {
			logger.WithField("err", err.Error()).Error(responses.ProcessFetchError)
			responseCodeAndMsg(rw, http.StatusInternalServerError, ErrObj{Err: responses.ProcessFetchError.Error()})
			return
		}
		if tipOpr.Type == db.PickupTip {
			err = ValidateTipPickupObject(req.Context(), deps, tipOpr, process.RecipeID)
			if err != nil {
				logger.WithField("err", err.Error()).Error(err.Error())
				responseCodeAndMsg(rw, http.StatusInternalServerError, ErrObj{Err: err.Error()})
				return
			}
		}
		err = plc.CheckIfRecipeOrProcessSafeForCUDs(nil, &id)
		if err != nil {
			responseCodeAndMsg(rw, http.StatusConflict, ErrObj{Err: err.Error()})
			logger.WithField("err", err.Error()).Error(responses.DefineCUDNotAllowedError(processC, updateC))
			return
		}

		tipOpr.ProcessID = id
		err = deps.Store.UpdateTipOperation(req.Context(), tipOpr)
		if err != nil {
			logger.WithField("err", err.Error()).Error(responses.TipOperationUpdateError)
			responseCodeAndMsg(rw, http.StatusInternalServerError, ErrObj{Err: responses.TipOperationUpdateError.Error()})
			return
		}

		logger.Infoln(responses.TipOperationUpdateSuccess)
		responseCodeAndMsg(rw, http.StatusOK, MsgObj{Msg: responses.TipOperationUpdateSuccess})
	})
}

func ValidateTipPickupObject(ctx context.Context, deps Dependencies, to db.TipOperation, recipeID uuid.UUID) (err error) {
	//check Tip ID from recipe
	recipe, err := deps.Store.ShowRecipe(ctx, recipeID)
	if err != nil {
		logger.WithField("err", err.Error()).Error(responses.RecipeFetchError)
		return responses.RecipeFetchError
	}

	var tipID int64
	switch to.Position {
	case 1:
		if recipe.Position1 == nil {
			return responses.TipMissingError
		}
		tipID = *recipe.Position1
	case 2:
		if recipe.Position2 == nil {
			return responses.TipMissingError
		}
		tipID = *recipe.Position2
	case 3:
		if recipe.Position3 == nil {
			return responses.TipMissingError
		}
		tipID = *recipe.Position3
	case 4:
		if recipe.Position4 == nil {
			return responses.TipMissingError
		}
		tipID = *recipe.Position4
	case 5:
		if recipe.Position5 == nil {
			return responses.TipMissingError
		}
		tipID = *recipe.Position5
	}

	if !plc.DoesTipExist(tipID) {
		return responses.TipDoesNotExistError
	}

	return
}
