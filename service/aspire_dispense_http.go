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

func createAspireDispenseHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		//logging when the api is initialised
		go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.InitialisedState, db.CreateOperation, "", responses.AspireDispenseInitialisedState)

		vars := mux.Vars(req)

		recipeID, err := parseUUID(vars["recipe_id"])

		// for logging error if there is any otherwise logging success
		defer func() {
			if err != nil {
				go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.ErrorState, db.CreateOperation, "", err.Error())
			} else {
				go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.CompletedState, db.CreateOperation, "", responses.AspireDispenseCompletedState)
			}

		}()

		if err != nil {
			// This error is already logged
			responseCodeAndMsg(rw, http.StatusBadRequest, ErrObj{Err: responses.RecipeIDInvalidError.Error()})
			return
		}

		var adobj db.AspireDispense
		err = json.NewDecoder(req.Body).Decode(&adobj)
		if err != nil {
			logger.WithField("err", err.Error()).Errorln(responses.AspireDispenseDecodeError)
			responseCodeAndMsg(rw, http.StatusBadRequest, ErrObj{Err: responses.AspireDispenseDecodeError.Error()})
			return
		}

		valid, respBytes := validate(adobj)
		if !valid {
			logger.WithField("err", "Validation Error").Errorln(responses.AspireDispenseValidationError)
			responseBadRequest(rw, respBytes)
			return
		}

		err = ValidateAspireDispenceObject(req.Context(), deps, adobj, recipeID)
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

		var createdTemp db.AspireDispense
		createdTemp, err = deps.Store.CreateAspireDispense(req.Context(), adobj, recipeID)
		if err != nil {
			logger.WithField("err", err.Error()).Errorln(responses.AspireDispenseCreateError)
			responseCodeAndMsg(rw, http.StatusInternalServerError, ErrObj{Err: responses.AspireDispenseCreateError.Error()})
			return
		}
		logger.Infoln(responses.AspireDispenseCreateSuccess)
		responseCodeAndMsg(rw, http.StatusCreated, createdTemp)
	})
}

func showAspireDispenseHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		//logging when the api is initialised
		go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.InitialisedState, db.ShowOperation, "", responses.AspireDispenseInitialisedState)

		vars := mux.Vars(req)
		id, err := parseUUID(vars["id"])
		// for logging error if there is any otherwise logging success
		defer func() {
			if err != nil {
				go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.ErrorState, db.ShowOperation, "", err.Error())
			} else {
				go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.CompletedState, db.ShowOperation, "", responses.AspireDispenseCompletedState)
			}
		}()

		if err != nil {
			// This error is already logged
			responseCodeAndMsg(rw, http.StatusBadRequest, ErrObj{Err: responses.UUIDParseError.Error()})
			return
		}

		var latestT db.AspireDispense

		latestT, err = deps.Store.ShowAspireDispense(req.Context(), id)
		if err != nil {
			responseCodeAndMsg(rw, http.StatusInternalServerError, ErrObj{Err: responses.AspireDispenseFetchError.Error()})
			logger.WithField("err", err.Error()).Error(responses.AspireDispenseFetchError)
			return
		}

		logger.Infoln(responses.AspireDispenseFetchSuccess)
		responseCodeAndMsg(rw, http.StatusOK, latestT)
	})
}

func updateAspireDispenseHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		//logging when the api is initialised
		go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.InitialisedState, db.UpdateOperation, "", responses.AspireDispenseInitialisedState)

		vars := mux.Vars(req)
		id, err := parseUUID(vars["id"])

		// for logging error if there is any otherwise logging success
		defer func() {
			if err != nil {
				go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.ErrorState, db.UpdateOperation, "", err.Error())
			} else {
				go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.CompletedState, db.UpdateOperation, "", responses.AspireDispenseCompletedState)
			}
		}()

		if err != nil {
			// This error is already logged
			responseCodeAndMsg(rw, http.StatusBadRequest, ErrObj{Err: responses.UUIDParseError.Error()})
			return
		}

		var adobj db.AspireDispense

		err = json.NewDecoder(req.Body).Decode(&adobj)
		if err != nil {
			logger.WithField("err", err.Error()).Errorln(responses.AspireDispenseDecodeError)
			responseCodeAndMsg(rw, http.StatusBadRequest, ErrObj{Err: responses.AspireDispenseDecodeError.Error()})
			return
		}

		valid, respBytes := validate(adobj)
		if !valid {
			logger.WithField("err", "Validation Error").Errorln(responses.AspireDispenseValidationError)
			responseBadRequest(rw, respBytes)
			return
		}

		process, err := deps.Store.ShowProcess(req.Context(), id)
		if err != nil {
			logger.WithField("err", err.Error()).Error(responses.ProcessFetchError)
			responseCodeAndMsg(rw, http.StatusInternalServerError, ErrObj{Err: responses.ProcessFetchError.Error()})
			return
		}

		err = ValidateAspireDispenceObject(req.Context(), deps, adobj, process.RecipeID)
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

		adobj.ProcessID = id
		err = deps.Store.UpdateAspireDispense(req.Context(), adobj)
		if err != nil {
			logger.WithField("err", err.Error()).Error(responses.AspireDispenseUpdateError)
			responseCodeAndMsg(rw, http.StatusInternalServerError, ErrObj{Err: responses.AspireDispenseUpdateError.Error()})
			return
		}

		logger.Infoln(responses.AspireDispenseUpdateSuccess)
		responseCodeAndMsg(rw, http.StatusOK, MsgObj{Msg: responses.AspireDispenseUpdateSuccess})
	})
}

// ValidateAspireDispenceObject this can be called by CSV
func ValidateAspireDispenceObject(ctx context.Context, deps Dependencies, ad db.AspireDispense, recipeID uuid.UUID) (err error) {

	var aspireCartridgeWell, dispenseCartridgeWell plc.UniqueCartridge

	//check for source position validity
	if ad.SourcePosition == 0 && ad.Category != db.SW && ad.Category != db.SD {
		return responses.InvalidSourcePosition
	}

	//check for destination validity
	if ad.DestinationPosition == 0 && ad.Category != db.WS && ad.Category != db.DS {
		return responses.InvalidDestinationPosition
	}

	//check cartridge type from recipe
	recipe, err := deps.Store.ShowRecipe(ctx, recipeID)
	if err != nil {
		logger.WithField("err", err.Error()).Error(responses.RecipeFetchError)
		return responses.RecipeFetchError
	}

	switch ad.Category {
	case db.SD:
		if isDeckPositionInvalid(ad.DestinationPosition) {
			return responses.InvalidDestinationPosition
		}

	case db.DS:
		if isDeckPositionInvalid(ad.SourcePosition) {
			return responses.InvalidSourcePosition
		}

	case db.DD:
		if isDeckPositionInvalid(ad.DestinationPosition) {
			return responses.InvalidDestinationPosition
		}
		if isDeckPositionInvalid(ad.SourcePosition) {
			return responses.InvalidSourcePosition
		}

	}

	//fetch cartridge type using id
	var cartridgeID int64

	err = checkCartridgeType(recipe, ad.CartridgeType, &cartridgeID)
	if err != nil {
		return err
	}

	aspireCartridgeWell = createCartridgeWell(cartridgeID, ad.CartridgeType, ad.SourcePosition)
	dispenseCartridgeWell = createCartridgeWell(cartridgeID, ad.CartridgeType, ad.DestinationPosition)

	switch ad.Category {
	case db.WW:
		// send cartridge and both height for validation
		if !plc.IsCartridgeWellHeightSafe(aspireCartridgeWell, ad.AspireHeight) {
			return responses.InvalidAspireWell
		}
		if !plc.IsCartridgeWellHeightSafe(dispenseCartridgeWell, ad.DispenseHeight) {
			return responses.InvalidDispenseWell
		}

	case db.WD, db.WS:
		// send cartridge and aspire height for validation
		if !plc.IsCartridgeWellHeightSafe(aspireCartridgeWell, ad.AspireHeight) {
			return responses.InvalidAspireWell
		}
		return
	case db.DW, db.SW:
		// send cartridge and dispense height for validation
		if !plc.IsCartridgeWellHeightSafe(dispenseCartridgeWell, ad.DispenseHeight) {
			return responses.InvalidDispenseWell
		}
	default:
		return responses.InvalidCategoryAspireDispense
	}

	logger.Infoln(aspireCartridgeWell, dispenseCartridgeWell)

	return
}
