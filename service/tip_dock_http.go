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

func createTipDockHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.InitialisedState, db.CreateOperation, "", responses.TipDockingInitialisedState)

		vars := mux.Vars(req)

		recipeID, err := parseUUID(vars["recipe_id"])
		// for logging error if there is any otherwise logging success
		defer func() {
			if err != nil {
				go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.ErrorState, db.CreateOperation, "", err.Error())
			} else {
				go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.CompletedState, db.CreateOperation, "", responses.TipDockingCompletedState)
			}

		}()
		if err != nil {
			// This error is already logged
			responseCodeAndMsg(rw, http.StatusBadRequest, ErrObj{Err: responses.RecipeIDInvalidError.Error()})
			return
		}

		var tdObj db.TipDock
		err = json.NewDecoder(req.Body).Decode(&tdObj)

		if err != nil {
			logger.WithField("err", err.Error()).Errorln(responses.TipDockingDecodeError)
			responseCodeAndMsg(rw, http.StatusBadRequest, ErrObj{Err: responses.TipDockingDecodeError.Error()})
			return
		}

		valid, respBytes := validate(tdObj)
		if !valid {
			logger.WithField("err", "Validation Error").Errorln(responses.TipDockingValidationError)
			responseBadRequest(rw, respBytes)
			return
		}

		err = plc.CheckIfRecipeOrProcessSafeForCUDs(&recipeID, nil)
		if err != nil {
			responseCodeAndMsg(rw, http.StatusConflict, ErrObj{Err: err.Error()})
			logger.WithField("err", err.Error()).Error(responses.DefineCUDNotAllowedError(processC, createC))
			return
		}

		var tipDock db.TipDock
		tipDock, err = deps.Store.CreateTipDocking(req.Context(), tdObj, recipeID)
		if err != nil {
			logger.WithField("err", err.Error()).Errorln(responses.TipDockingCreateError)
			responseCodeAndMsg(rw, http.StatusInternalServerError, ErrObj{Err: responses.TipDockingCreateError.Error()})
			return
		}
		logger.Infoln(responses.TipDockingCreateSuccess)
		responseCodeAndMsg(rw, http.StatusCreated, tipDock)
	})
}

func showTipDockHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		//logging when the api is initialised

		go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.InitialisedState, db.ShowOperation, "", responses.TipDockingInitialisedState)

		vars := mux.Vars(req)
		id, err := parseUUID(vars["id"])

		// for logging error if there is any otherwise logging success
		defer func() {
			if err != nil {
				go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.ErrorState, db.ShowOperation, "", err.Error())

			} else {
				go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.CompletedState, db.ShowOperation, "", responses.TipDockingCompletedState)

			}

		}()

		if err != nil {
			responseCodeAndMsg(rw, http.StatusBadRequest, ErrObj{Err: responses.UUIDParseError.Error()})
			return
		}

		var tipDock db.TipDock
		tipDock, err = deps.Store.ShowTipDocking(req.Context(), id)
		if err != nil {
			responseCodeAndMsg(rw, http.StatusInternalServerError, ErrObj{Err: responses.TipDockingFetchError.Error()})
			logger.WithField("err", err.Error()).Error(responses.TipDockingFetchError)
			return
		}

		logger.Infoln(responses.TipDockingFetchSuccess)
		responseCodeAndMsg(rw, http.StatusOK, tipDock)
	})
}

func updateTipDockHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		//logging when the api is initialised

		go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.InitialisedState, db.UpdateOperation, "", responses.TipDockingInitialisedState)

		vars := mux.Vars(req)
		id, err := parseUUID(vars["id"])
		// for logging error if there is any otherwise logging success
		defer func() {
			if err != nil {
				go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.ErrorState, db.UpdateOperation, "", err.Error())

			} else {
				go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.CompletedState, db.UpdateOperation, "", responses.TipDockingCompletedState)

			}

		}()
		if err != nil {
			responseCodeAndMsg(rw, http.StatusBadRequest, ErrObj{Err: responses.UUIDParseError.Error()})
			return
		}
		var tdObj db.TipDock
		err = json.NewDecoder(req.Body).Decode(&tdObj)
		if err != nil {
			logger.WithField("err", err.Error()).Errorln(responses.TipDockingDecodeError)
			responseCodeAndMsg(rw, http.StatusBadRequest, ErrObj{Err: responses.TipDockingDecodeError.Error()})
			return
		}
		valid, respBytes := validate(tdObj)
		if !valid {
			logger.WithField("err", "Validation Error").Errorln(responses.TipDockingValidationError)
			responseBadRequest(rw, respBytes)
			return
		}

		err = plc.CheckIfRecipeOrProcessSafeForCUDs(nil, &id)
		if err != nil {
			responseCodeAndMsg(rw, http.StatusConflict, ErrObj{Err: err.Error()})
			logger.WithField("err", err.Error()).Error(responses.DefineCUDNotAllowedError(processC, updateC))
			return
		}

		tdObj.ProcessID = id
		err = deps.Store.UpdateTipDock(req.Context(), tdObj)
		if err != nil {
			logger.WithField("err", err.Error()).Error(responses.TipDockingUpdateError)
			responseCodeAndMsg(rw, http.StatusInternalServerError, ErrObj{Err: responses.TipDockingUpdateError.Error()})
			return
		}

		logger.Infoln(responses.TipDockingUpdateSuccess)
		responseCodeAndMsg(rw, http.StatusOK, MsgObj{Msg: responses.TipDockingUpdateSuccess})
	})
}

func ValidateTipDockObject(ctx context.Context, deps Dependencies, td db.TipDock, recipeID uuid.UUID) (err error) {

	recipe, err := deps.Store.ShowRecipe(ctx, recipeID)
	if err != nil {
		logger.WithField("err", err.Error()).Error(responses.RecipeFetchError)
		return responses.RecipeFetchError
	}
	// check if cartridage type
	switch td.Type {
	case "cartridge_1":
		if recipe.Cartridge1Position == nil {
			return responses.RecipeCartridge1Missing
		}
		cartridgeWell := createCartridgeWell(*recipe.Cartridge1Position, db.CartridgeType(td.Type), td.Position)
		if err != nil {
			return err
		}
		if !plc.IsCartridgeWellHeightSafe(cartridgeWell, td.Height) {
			return responses.InvalidTipDockWell
		}
	case "cartridge_2":
		if recipe.Cartridge2Position == nil {
			return responses.RecipeCartridge2Missing
		}

		cartridgeWell := createCartridgeWell(*recipe.Cartridge2Position, db.CartridgeType(td.Type), td.Position)
		if err != nil {
			return err
		}
		if !plc.IsCartridgeWellHeightSafe(cartridgeWell, td.Height) {
			return responses.InvalidTipDockWell
		}
	case "deck":
		if isDeckPositionInvalid(td.Position) {
			return responses.InvalidDeckPosition
		}
	}
	return
}
