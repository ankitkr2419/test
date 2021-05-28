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

func listRecipesHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.InitialisedState, db.ShowOperation, "", responses.RecipeListInitialisedState)

		list, err := deps.Store.ListRecipes(req.Context())

		// for logging error if there is any otherwise logging success
		defer func() {
			if err != nil {
				go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.ErrorState, db.ShowOperation, "", err.Error())

			} else {
				go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.CompletedState, db.ShowOperation, "", responses.RecipeListCompletedState)

			}

		}()

		if err != nil {
			responseCodeAndMsg(rw, http.StatusInternalServerError, ErrObj{Err: responses.RecipeListFetchError.Error()})
			logger.WithField("err", err.Error()).Error(responses.RecipeListFetchError)
			return
		}

		logger.Infoln(responses.RecipeListFetchSuccess)
		responseCodeAndMsg(rw, http.StatusOK, list)
	})
}

func createRecipeHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.InitialisedState, db.CreateOperation, "", responses.RecipeInitialisedState)

		var recipe db.Recipe
		err := json.NewDecoder(req.Body).Decode(&recipe)
		// for logging error if there is any otherwise logging success
		defer func() {
			if err != nil {
				go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.ErrorState, db.CreateOperation, "", err.Error())

			} else {
				go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.CompletedState, db.CreateOperation, "", responses.RecipeCompletedState)

			}

		}()

		if err != nil {
			logger.WithField("err", err.Error()).Errorln(responses.RecipeDecodeError)
			responseCodeAndMsg(rw, http.StatusBadRequest, ErrObj{Err: responses.RecipeDecodeError.Error()})
			return
		}

		valid, respBytes := validate(recipe)
		if !valid {
			logger.WithField("err", "Validation Error").Errorln(responses.RecipeValidationError)
			responseBadRequest(rw, respBytes)

			return
		}

		var createdTemp db.Recipe
		createdTemp, err = deps.Store.CreateRecipe(req.Context(), recipe)
		if err != nil {
			logger.WithField("err", err.Error()).Errorln(responses.RecipeCreateError)
			responseCodeAndMsg(rw, http.StatusInternalServerError, ErrObj{Err: responses.RecipeCreateError.Error()})
			return
		}

		logger.Infoln(responses.RecipeCreateSuccess)
		responseCodeAndMsg(rw, http.StatusCreated, createdTemp)
	})
}

func showRecipeHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.InitialisedState, db.ShowOperation, "", responses.RecipeInitialisedState)

		vars := mux.Vars(req)

		id, err := parseUUID(vars["id"])
		// for logging error if there is any otherwise logging success
		defer func() {
			if err != nil {
				go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.ErrorState, db.ShowOperation, "", err.Error())

			} else {
				go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.CompletedState, db.ShowOperation, "", responses.RecipeCompletedState)

			}

		}()
		if err != nil {
			responseCodeAndMsg(rw, http.StatusBadRequest, ErrObj{Err: responses.UUIDParseError.Error()})
			return
		}

		var latestT db.Recipe

		latestT, err = deps.Store.ShowRecipe(req.Context(), id)
		if err != nil {
			responseCodeAndMsg(rw, http.StatusInternalServerError, ErrObj{Err: responses.RecipeFetchError.Error()})
			logger.WithField("err", err.Error()).Error(responses.RecipeFetchError)
			return
		}

		logger.Infoln(responses.RecipeFetchSuccess)
		responseCodeAndMsg(rw, http.StatusOK, latestT)
	})
}

func deleteRecipeHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.InitialisedState, db.DeleteOperation, "", responses.RecipeInitialisedState)

		vars := mux.Vars(req)
		id, err := parseUUID(vars["id"])
		// for logging error if there is any otherwise logging success
		defer func() {
			if err != nil {
				go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.ErrorState, db.DeleteOperation, "", err.Error())

			} else {
				go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.CompletedState, db.DeleteOperation, "", responses.RecipeCompletedState)

			}

		}()
		if err != nil {
			responseCodeAndMsg(rw, http.StatusBadRequest, ErrObj{Err: responses.UUIDParseError.Error()})
			return
		}

		plc.CheckIfRecipeOrProcessSafeForUDs(&id, nil)

		err = deps.Store.DeleteRecipe(req.Context(), id)
		if err != nil {
			logger.WithField("err", err.Error()).Error(responses.RecipeDeleteError)
			responseCodeAndMsg(rw, http.StatusInternalServerError, ErrObj{Err: responses.RecipeDeleteError.Error()})
			return
		}

		logger.Infoln(responses.RecipeUpdateSuccess)
		responseCodeAndMsg(rw, http.StatusOK, MsgObj{Msg: responses.RecipeDeleteSuccess})
	})
}

// NOTE: no need to call CheckIfRecipeOrProcessSafeForCRUDs
//  as before run itself complete recipe data is stored
func updateRecipeHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.InitialisedState, db.UpdateOperation, "", responses.RecipeInitialisedState)

		vars := mux.Vars(req)
		id, err := parseUUID(vars["id"])
		// for logging error if there is any otherwise logging success
		defer func() {
			if err != nil {
				go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.ErrorState, db.UpdateOperation, "", err.Error())

			} else {
				go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.CompletedState, db.UpdateOperation, "", responses.RecipeCompletedState)

			}

		}()
		if err != nil {
			responseCodeAndMsg(rw, http.StatusBadRequest, ErrObj{Err: responses.UUIDParseError.Error()})
			return
		}

		var recipe db.Recipe

		err = json.NewDecoder(req.Body).Decode(&recipe)
		if err != nil {
			logger.WithField("err", err.Error()).Errorln(responses.RecipeDecodeError)
			responseCodeAndMsg(rw, http.StatusBadRequest, ErrObj{Err: responses.RecipeDecodeError.Error()})
			return
		}

		valid, respBytes := validate(recipe)
		if !valid {
			logger.WithField("err", "Validation Error").Errorln(responses.RecipeValidationError)
			responseBadRequest(rw, respBytes)
			return
		}

		recipe.ID = id
		err = deps.Store.UpdateRecipe(req.Context(), recipe)
		if err != nil {
			logger.WithField("err", err.Error()).Error(responses.RecipeUpdateError)
			responseCodeAndMsg(rw, http.StatusInternalServerError, ErrObj{Err: responses.RecipeUpdateError.Error()})
			return
		}

		logger.Infoln(responses.RecipeUpdateSuccess)
		responseCodeAndMsg(rw, http.StatusOK, MsgObj{Msg: responses.RecipeUpdateSuccess})
	})
}

func publishRecipeHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.InitialisedState, db.UpdateOperation, "", responses.RecipePublishedState)

		var publishFlag bool
		var successMsg string
		vars := mux.Vars(req)
		id, err := parseUUID(vars["id"])

		// for logging error if there is any otherwise logging success
		defer func() {
			if err != nil {
				go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.ErrorState, db.UpdateOperation, "", err.Error())

			} else {
				go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.CompletedState, db.UpdateOperation, "", responses.RecipePublishedState)

			}

		}()
		if err != nil {
			responseCodeAndMsg(rw, http.StatusBadRequest, ErrObj{Err: responses.UUIDParseError.Error()})
			return
		}

		//for more clarity, we take whole keywords(publish and unpublish) from the url
		publish := vars["publish"]

		switch publish {
		case "publish":
			publishFlag = true
		case "unpublish":
			publishFlag = false
		default:
			responseCodeAndMsg(rw, http.StatusBadRequest, ErrObj{Err: responses.UrlArgumentInvalid.Error()})
			return
		}

		var recipe db.Recipe

		recipe.ID = id
		recipe, err = deps.Store.ShowRecipe(req.Context(), id)
		if err != nil {
			responseCodeAndMsg(rw, http.StatusInternalServerError, ErrObj{Err: responses.RecipeFetchError.Error()})
			logger.WithField("err", err.Error()).Error(responses.RecipeFetchError)
			return
		}

		if publishFlag == recipe.IsPublished {
			responseCodeAndMsg(rw, http.StatusInternalServerError, ErrObj{Err: responses.RecipePublishError.Error()})
			logger.WithField("err", "PUBLISH_ERROR").Error(responses.RecipePublishError)
			return

		}
		//TODO : check if the recipe is published on the cloud and if there then delete
		// it from the cloud.
		recipe.IsPublished = publishFlag

		err = deps.Store.UpdateRecipe(req.Context(), recipe)
		if err != nil {
			logger.WithField("err", err.Error()).Error(responses.RecipeUpdateError)
			responseCodeAndMsg(rw, http.StatusInternalServerError, ErrObj{Err: responses.RecipeUpdateError.Error()})
			return
		}

		if recipe.IsPublished {
			successMsg = responses.RecipePublishSuccess
		} else {
			successMsg = responses.RecipeUnPublishSuccess
		}

		logger.Infoln(responses.RecipeUpdateSuccess)
		responseCodeAndMsg(rw, http.StatusOK, MsgObj{Msg: successMsg})
	})
}
