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

func createPiercingHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.InitialisedState, db.CreateOperation, "", responses.PiercingInitialisedState)

		vars := mux.Vars(req)

		recipeID, err := parseUUID(vars["recipe_id"])
		// for logging error if there is any otherwise logging success
		defer func() {
			if err != nil {
				go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.ErrorState, db.CreateOperation, "", err.Error())

			} else {
				go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.CompletedState, db.CreateOperation, "", responses.PiercingCompletedState)

			}

		}()

		if err != nil {
			// This error is already logged
			responseCodeAndMsg(rw, http.StatusBadRequest, ErrObj{Err: responses.RecipeIDInvalidError.Error()})
			return
		}

		var pobj db.Piercing

		err = json.NewDecoder(req.Body).Decode(&pobj)
		if err != nil {
			logger.WithField("err", err.Error()).Errorln(responses.PiercingDecodeError)
			responseCodeAndMsg(rw, http.StatusBadRequest, ErrObj{Err: responses.PiercingDecodeError.Error()})
			return
		}

		valid, respBytes := Validate(pobj)
		if !valid {
			logger.WithField("err", "Validation Error").Errorln(responses.PiercingValidationError)
			responseBadRequest(rw, respBytes)
			return
		}

		err = ValidatePiercingObject(req.Context(), deps, &pobj, recipeID)
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

		var createdTemp db.Piercing
		createdTemp, err = deps.Store.CreatePiercing(req.Context(), pobj, recipeID)
		if err != nil {
			logger.WithField("err", err.Error()).Errorln(responses.PiercingCreateError)
			responseCodeAndMsg(rw, http.StatusInternalServerError, ErrObj{Err: responses.PiercingCreateError.Error()})
			return
		}
		logger.Infoln(responses.PiercingCreateSuccess)
		responseCodeAndMsg(rw, http.StatusCreated, createdTemp)
	})
}

func showPiercingHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		//logging when the api is initialised
		go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.InitialisedState, db.ShowOperation, "", responses.PiercingInitialisedState)

		vars := mux.Vars(req)

		id, err := parseUUID(vars["id"])

		// for logging error if there is any otherwise logging success
		defer func() {
			if err != nil {
				go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.ErrorState, db.ShowOperation, "", err.Error())

			} else {
				go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.CompletedState, db.ShowOperation, "", responses.PiercingCompletedState)

			}

		}()
		if err != nil {
			// This error is already logged
			responseCodeAndMsg(rw, http.StatusBadRequest, ErrObj{Err: responses.UUIDParseError.Error()})
			return
		}

		var latestT db.Piercing

		latestT, err = deps.Store.ShowPiercing(req.Context(), id)
		if err != nil {
			responseCodeAndMsg(rw, http.StatusInternalServerError, ErrObj{Err: responses.PiercingFetchError.Error()})
			logger.WithField("err", err.Error()).Error(responses.PiercingFetchError)
			return
		}

		logger.Infoln(responses.PiercingFetchSuccess)
		responseCodeAndMsg(rw, http.StatusOK, latestT)
	})
}

func updatePiercingHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		//logging when the api is initialised
		go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.InitialisedState, db.UpdateOperation, "", responses.PiercingInitialisedState)

		vars := mux.Vars(req)
		id, err := parseUUID(vars["id"])
		// for logging error if there is any otherwise logging success
		defer func() {
			if err != nil {
				go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.ErrorState, db.UpdateOperation, "", err.Error())

			} else {
				go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.CompletedState, db.UpdateOperation, "", responses.PiercingCompletedState)

			}

		}()
		if err != nil {
			// This error is already logged
			responseCodeAndMsg(rw, http.StatusBadRequest, ErrObj{Err: responses.UUIDParseError.Error()})
			return
		}

		var pobj db.Piercing

		err = json.NewDecoder(req.Body).Decode(&pobj)
		if err != nil {
			logger.WithField("err", err.Error()).Errorln(responses.PiercingDecodeError)
			responseCodeAndMsg(rw, http.StatusBadRequest, ErrObj{Err: responses.PiercingDecodeError.Error()})
			return
		}

		valid, respBytes := Validate(pobj)
		if !valid {
			responseBadRequest(rw, respBytes)
			return
		}

		process, err := deps.Store.ShowProcess(req.Context(), id)
		if err != nil {
			logger.WithField("err", err.Error()).Error(responses.ProcessFetchError)
			responseCodeAndMsg(rw, http.StatusInternalServerError, ErrObj{Err: responses.ProcessFetchError.Error()})
			return
		}

		err = ValidatePiercingObject(req.Context(), deps, &pobj, process.RecipeID)
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

		pobj.ProcessID = id
		err = deps.Store.UpdatePiercing(req.Context(), pobj)
		if err != nil {
			logger.WithField("err", err.Error()).Error(responses.PiercingUpdateError)
			responseCodeAndMsg(rw, http.StatusInternalServerError, ErrObj{Err: responses.PiercingUpdateError.Error()})
			return
		}

		logger.Infoln(responses.PiercingUpdateSuccess)
		responseCodeAndMsg(rw, http.StatusOK, MsgObj{Msg: responses.PiercingUpdateSuccess})
	})
}

// ValidatePiercingObject this can be called by CSV
func ValidatePiercingObject(ctx context.Context, deps Dependencies, pi *db.Piercing, recipeID uuid.UUID) (err error) {

	//fetch cartridge type using id
	var cartridgeID int64

	//check cartridge type from recipe
	recipe, err := deps.Store.ShowRecipe(ctx, recipeID)
	if err != nil {
		logger.WithField("err", err.Error()).Error(responses.RecipeFetchError)
		return responses.RecipeFetchError
	}

	err = checkCartridgeType(recipe, pi.Type, &cartridgeID)
	if err != nil {
		return
	}

	// sorting logic needs to be here cause CSV will be using this
	return validateWellsLengthAndSortWells(pi, cartridgeID)
}

// check per well basis
func validateWellsLengthAndSortWells(pi *db.Piercing, cartridgeID int64) error {
	var wellsToBePierced, piercingHeights []int
	var cartridgeWell plc.UniqueCartridge

	if len(pi.CartridgeWells) == 0 {
		logger.Errorln(responses.MissingCartridgeWellsError)
		return responses.MissingCartridgeWellsError
	}

	if len(pi.CartridgeWells) != len(pi.Heights) {
		logger.Errorln(responses.CartridgeWellsHeightsMismatchError)
		return responses.CartridgeWellsMismatchWithHeightError
	}

	for i := range pi.CartridgeWells {
		wellsToBePierced = append(wellsToBePierced, int(pi.CartridgeWells[i]))
		piercingHeights = append(piercingHeights, int(pi.Heights[i]))
	}

	t := db.NewWellsSlice(wellsToBePierced, piercingHeights)

	logger.Debugln("Wells -> ", t)

	for i := range t.IntSlice {
		pi.CartridgeWells[i] = int64(t.IntSlice[i])
		pi.Heights[i] = int64(t.Heights[i])

		//  check if the individual well exists
		cartridgeWell = createCartridgeWell(cartridgeID, pi.Type, int64(t.IntSlice[i]))
		if !plc.IsCartridgeWellHeightSafe(cartridgeWell, float64(t.Heights[i])) {
			return responses.InvalidPiercingWell
		}
	}
	return nil
}
