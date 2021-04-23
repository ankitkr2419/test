package service

import (
	"encoding/json"
	"mylab/cpagent/db"
	"net/http"

	logger "github.com/sirupsen/logrus"
)

func listCartridgesHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		var cartridge []db.Cartridge
		cartridge, err := deps.Store.ListCartridges()
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			logger.WithField("err", err.Error()).Error("Error showing cartridges")
			return
		}

		respBytes, err := json.Marshal(cartridge)
		if err != nil {
			logger.WithField("err", err.Error()).Error("Error marshaling cartridges data")
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}
		rw.Header().Add("Content-Type", "application/json")
		rw.WriteHeader(http.StatusOK)
		rw.Write(respBytes)
	})
}
