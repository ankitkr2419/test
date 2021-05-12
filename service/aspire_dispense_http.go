package service

import (
	"encoding/json"
	"mylab/cpagent/db"
	"mylab/cpagent/responses"
	"net/http"

	"github.com/gorilla/mux"
	logger "github.com/sirupsen/logrus"
)

func createAspireDispenseHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)

		recipeID, err := parseUUID(vars["recipe_id"])
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
		vars := mux.Vars(req)

		id, err := parseUUID(vars["id"])
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
		vars := mux.Vars(req)
		id, err := parseUUID(vars["id"])
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
