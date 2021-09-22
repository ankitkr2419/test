package service

import (
	"encoding/json"
	"mylab/cpagent/config"
	"mylab/cpagent/responses"
	"net/http"
	"regexp"

	logger "github.com/sirupsen/logrus"
)

const (
	emailRegex = "^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$"
)

func getTECConfigHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		responseCodeAndMsg(rw, http.StatusOK, config.GetTECConfigValues())
	})
}

func updateTECConfigHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		var tc config.TEC
		err := json.NewDecoder(req.Body).Decode(&tc)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			logger.WithField("err", err.Error()).Errorln(responses.ConfigDataDecodeError)
			return
		}

		valid, respBytes := Validate(tc)
		if !valid {
			responseBadRequest(rw, respBytes)
			return
		}

		err = config.SetTECConfigValues(tc)
		if err != nil {
			logger.WithField("err", err.Error()).Errorln(responses.ConfigDataUpdateError)
			responseCodeAndMsg(rw, http.StatusInternalServerError, ErrObj{Err: responses.ConfigDataUpdateError.Error()})
			return
		}

		responseCodeAndMsg(rw, http.StatusOK, MsgObj{Msg: responses.UpdateConfigSuccess})
	})
}

func getRTPCRConfigHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		responseCodeAndMsg(rw, http.StatusOK, config.GetRTPCRConfigValues())
	})
}

func updateRTPCRConfigHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		var rt config.RTPCR
		err := json.NewDecoder(req.Body).Decode(&rt)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			logger.WithField("err", err.Error()).Errorln(responses.ConfigDataDecodeError)
			return
		}

		valid, respBytes := Validate(rt)
		if !valid {
			responseBadRequest(rw, respBytes)
			return
		}

		err = config.SetRTPCRConfigValues(rt)
		if err != nil {
			logger.WithField("err", err.Error()).Errorln(responses.ConfigDataUpdateError)
			responseCodeAndMsg(rw, http.StatusInternalServerError, ErrObj{Err: responses.ConfigDataUpdateError.Error()})
			return
		}

		err = deps.Plc.SetScanSpeedAndScanTime()
		if err != nil {
			logger.WithField("err", err.Error()).Errorln(responses.PLCDataUpdateError)
			responseCodeAndMsg(rw, http.StatusInternalServerError, ErrObj{Err: responses.PLCDataUpdateError.Error()})
			return
		}

		responseCodeAndMsg(rw, http.StatusOK, MsgObj{Msg: responses.UpdateConfigSuccess})
	})
}

func getExtractionConfigHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		responseCodeAndMsg(rw, http.StatusOK, config.GetExtractionConfigValues())
	})
}

func updateExtractionConfigHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		var ex config.Extraction
		err := json.NewDecoder(req.Body).Decode(&ex)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			logger.WithField("err", err.Error()).Errorln(responses.ConfigDataDecodeError)
			return
		}

		valid, respBytes := Validate(ex)
		if !valid {
			responseBadRequest(rw, respBytes)
			return
		}

		err = config.SetExtractionConfigValues(ex)
		if err != nil {
			logger.WithField("err", err.Error()).Errorln(responses.ConfigDataUpdateError)
			responseCodeAndMsg(rw, http.StatusInternalServerError, ErrObj{Err: responses.ConfigDataUpdateError.Error()})
			return
		}

		responseCodeAndMsg(rw, http.StatusOK, MsgObj{Msg: responses.UpdateConfigSuccess})
	})
}


func getCommonConfigHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		responseCodeAndMsg(rw, http.StatusOK, config.GetCommonConfigValues())
	})
}

func updateCommonConfigHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		var co config.Common
		err := json.NewDecoder(req.Body).Decode(&co)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			logger.WithField("err", err.Error()).Errorln(responses.ConfigDataDecodeError)
			return
		}

		valid, respBytes := Validate(co)
		if !valid {
			responseBadRequest(rw, respBytes)
			return
		}

		if !isEmailValid(co.ReceiverEmail) {
			logger.Errorln(responses.InvalidEmailIDError)
			responseCodeAndMsg(rw, http.StatusBadRequest, ErrObj{Err: responses.InvalidEmailIDError.Error()})
			return
		}

		err = config.SetCommonConfigValues(co)
		if err != nil {
			logger.WithField("err", err.Error()).Errorln(responses.ConfigDataUpdateError)
			responseCodeAndMsg(rw, http.StatusInternalServerError, ErrObj{Err: responses.ConfigDataUpdateError.Error()})
			return
		}

		responseCodeAndMsg(rw, http.StatusOK, MsgObj{Msg: responses.UpdateConfigSuccess})
	})
}

// isEmailValid checks if the email provided is valid by regex.
func isEmailValid(e string) bool {
	emailRegex := regexp.MustCompile(emailRegex)
	return emailRegex.MatchString(e)
}
