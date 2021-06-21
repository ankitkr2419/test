package service

import (
	"fmt"
	"mylab/cpagent/db"
	"mylab/cpagent/plc"
	"mylab/cpagent/responses"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	logger "github.com/sirupsen/logrus"
)

func discardAndHomeHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.InitialisedState, db.ExecuteOperation, "", responses.DiscardTipHomeInitialisedState)

		var discard bool
		var err error
		// for logging error if there is any otherwise logging success
		defer func() {
			if err != nil {
				go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.ErrorState, db.ExecuteOperation, "", err.Error())
			} else {
				go deps.Store.AddAuditLog(req.Context(), db.ApiOperation, db.CompletedState, db.ExecuteOperation, "", responses.DiscardTipHomeCompletedState)
			}
		}()

		vars := mux.Vars(req)
		deck := vars["deck"]

		discardTip := vars["discard"]

		if discard, err = strconv.ParseBool(discardTip); err != nil {
			responseCodeAndMsg(rw, http.StatusBadRequest, ErrObj{Err: responses.DiscardBoolOptionError.Error()})
			logger.WithField("err", err.Error()).Error(responses.DiscardBoolOptionError)
			err := fmt.Errorf("Discard option Should be boolean only")
			deps.WsErrCh <- err
		}

		switch deck {
		case plc.DeckA, plc.DeckB:
			go deps.PlcDeck[deck].DiscardTipAndHome(discard)
			logger.Infoln(responses.DiscardTipHomeSuccess)
			responseCodeAndMsg(rw, http.StatusOK, MsgObj{Msg: responses.DiscardTipHomeSuccess, Deck: deck})

		default:
			err := fmt.Errorf("Check your deck name")
			deps.WsErrCh <- err
		}

		if err != nil {
			responseCodeAndMsg(rw, http.StatusBadRequest, ErrObj{Err: err.Error()})
			logger.WithField("err", err.Error()).Error(err)
		}

	})
}
