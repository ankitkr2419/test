package service

import (
	"mylab/cpagent/db"
	"mylab/cpagent/responses"
	"net/http"

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
