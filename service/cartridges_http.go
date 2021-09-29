package service

import (
	"encoding/json"
	"mylab/cpagent/db"
	"mylab/cpagent/responses"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	logger "github.com/sirupsen/logrus"
)

func listCartridgesHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		//logging when the api is initialised
		go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.InitialisedState, db.ShowOperation, "", responses.CartridgeListInitialisedState)

		var cartridges []db.Cartridge
		cartridges, err := deps.Store.ListCartridges(req.Context())

		// for logging error if there is any otherwise logging success
		defer func() {
			if err != nil {
				go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.ErrorState, db.ShowOperation, "", err.Error())
			} else {
				go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.CompletedState, db.ShowOperation, "", responses.CartridgeListCompletedState)
			}
		}()
		if err != nil {
			responseCodeAndMsg(rw, http.StatusInternalServerError, ErrObj{Err: responses.CartridgeFetchError.Error()})
			logger.WithField("err", err.Error()).Error(responses.CartridgeFetchError)
			return
		}

		logger.Infoln(responses.CartridgeFetchSuccess)
		responseCodeAndMsg(rw, http.StatusOK, cartridges)
	})
}

func createCartridgeHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		var m db.CartridgeWell

		err := json.NewDecoder(req.Body).Decode(&m)
		if err != nil {
			logger.WithField("err", err.Error()).Error(responses.CartridgeDecodeError)
			responseCodeAndMsg(rw, http.StatusBadRequest, ErrObj{Err: responses.CartridgeDecodeError.Error()})
			return
		}

		valid, respBytes := Validate(m)
		if !valid {
			responseBadRequest(rw, respBytes)
			return
		}

		err = db.SetCartridgeValues(m)
		if err != nil {
			logger.WithField("err", err.Error()).Error(responses.CartridgeCreateConfigError)
			responseCodeAndMsg(rw, http.StatusInternalServerError, ErrObj{Err: responses.CartridgeCreateConfigError.Error()})
			return
		}
		err = deps.Store.InsertCartridge(req.Context(), m.Cartridge, m.CartridgeWells)
		if err != nil {
			responseCodeAndMsg(rw, http.StatusInternalServerError, ErrObj{Err: responses.CartridegInsertionError.Error()})
			logger.WithField("err", err.Error()).Error(responses.CartridegInsertionError)
			return
		}

		responseCodeAndMsg(rw, http.StatusCreated, MsgObj{Msg: responses.CartridgeCreatedSuccess})
	})
}

func deleteCartridgeHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		id, err := strconv.ParseInt(vars["id"], 10, 64)
		if err != nil {
			logger.WithField("err", err.Error()).Error(responses.CartridgeIDParseError)
			responseCodeAndMsg(rw, http.StatusBadRequest, ErrObj{Err: responses.CartridgeIDParseError.Error()})
			return
		}

		err = deps.Store.DeleteCartridge(req.Context(), id)
		if err != nil {
			responseCodeAndMsg(rw, http.StatusInternalServerError, ErrObj{Err: responses.CartridegDeletionError.Error()})
			logger.WithField("err", err.Error()).Error(responses.CartridegDeletionError)
			return
		}

		responseCodeAndMsg(rw, http.StatusOK, MsgObj{Msg: responses.CartridgeDeletedSuccess})
	})
}

func showCartridgeHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		//logging when the api is initialised
		go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.InitialisedState, db.ShowOperation, "", responses.CartridgeInitialisedState)

		vars := mux.Vars(req)

		id, err := strconv.ParseInt(vars["id"], 10, 64)
		if err != nil {
			logger.Errorln(responses.InvalidCartridgeIDError)
			responseCodeAndMsg(rw, http.StatusBadRequest, ErrObj{Err: responses.InvalidCartridgeIDError.Error()})
			return
		}
		// for logging error if there is any otherwise logging success
		defer func() {
			if err != nil {
				go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.ErrorState, db.ShowOperation, "", err.Error())

			} else {
				go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.CompletedState, db.ShowOperation, "", responses.CartridgeCompletedState)

			}
		}()

		var cartridge db.Cartridge

		cartridge, err = deps.Store.ShowCartridge(req.Context(), id)
		if err != nil {
			responseCodeAndMsg(rw, http.StatusInternalServerError, ErrObj{Err: responses.CartridgeFetchError.Error()})
			logger.WithField("err", err.Error()).Error(responses.CartridgeFetchError)
			return
		}

		logger.Infoln(responses.CartridgeFetchSuccess)
		responseCodeAndMsg(rw, http.StatusOK, cartridge)
		return
	})
}
