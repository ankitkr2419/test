package service

import (
	"encoding/json"
	"mylab/cpagent/config"
	"mylab/cpagent/responses"
	"net/http"

	logger "github.com/sirupsen/logrus"
)

func getConfigHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		c, err := getConfigDetails()
		if err != nil {
			logger.WithField("err", err.Error()).Errorln(responses.ConfigDataFetchError)
			responseCodeAndMsg(rw, http.StatusInternalServerError, ErrObj{Err: responses.ConfigDataFetchError.Error()})
			return
		}

		responseCodeAndMsg(rw, http.StatusOK, c)
	})
}

func updateConfigHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		var c config.Conf
		err := json.NewDecoder(req.Body).Decode(&c)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			logger.WithField("err", err.Error()).Errorln(responses.ConfigDataDecodeError)
			return
		}

		valid, respBytes := validate(c)
		if !valid {
			responseBadRequest(rw, respBytes)
			return
		}

		err = config.SetValues(c)
		if err != nil {
			logger.WithField("err", err.Error()).Errorln(responses.ConfigDataUpdateError)
			responseCodeAndMsg(rw, http.StatusInternalServerError, ErrObj{Err: responses.ConfigDataUpdateError.Error()})
			return
		}

		responseCodeAndMsg(rw, http.StatusOK, MsgObj{Msg: responses.UpdateConfigSuccess})
	})
}

func getConfigDetails() (c config.Conf, err error) {
	c = config.Conf{
		RoomTemperature: int64(config.GetRoomTemp()),
		HomingTime:      int64(config.GetHomingTime()),
		NumHomingCycles: int64(config.GetNumHomingCycles()),
		CycleTime:       int64(config.GetCycleTime()),
		PIDMinutes:      int64(config.GetPIDMinutes()),
		PIDTemperature:  int64(config.GetPIDTemp()),
	}

	return
}
