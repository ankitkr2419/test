package service

import (
	"encoding/json"
	"mylab/cpagent/db"
	"mylab/cpagent/responses"
	"net/http"

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

		valid, respBytes := validate(pobj)
		if !valid {
			logger.WithField("err", "Validation Error").Errorln(responses.PiercingValidationError)
			responseBadRequest(rw, respBytes)
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

		valid, respBytes := validate(pobj)
		if !valid {
			responseBadRequest(rw, respBytes)
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
