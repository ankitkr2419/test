package service

import (
	"net/http"
	"mylab/cpagent/config"

	logger "github.com/sirupsen/logrus"
)

func getConfigHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		c, err := getConfigDetails()
		if err != nil {
			logger.WithField("err", err.Error()).Error("Error fetching Config data")
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}

		responseCodeAndMsg(rw, http.StatusOK, c)
	})
}


func getConfigDetails() (c Conf, err error){
	c = Conf{
		RoomTemperature: int64(config.GetRoomTemp()),
		HomingTime: int64(config.GetHomingTime()),
		NumHomingCycles: int64(config.GetNumHomingCycles()),
	}

	return
}

type Conf struct{
	RoomTemperature int64 `json:"room_temperature"`
	HomingTime int64 `json:"homing_time"`
	NumHomingCycles int64 `json:"no_of_homing_cycles"`
}
