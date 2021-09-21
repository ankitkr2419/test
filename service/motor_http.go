package service

import (
	"encoding/json"
	"mylab/cpagent/db"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	logger "github.com/sirupsen/logrus"
)

func createMotorHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		var m db.Motor
		err := json.NewDecoder(req.Body).Decode(&m)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			logger.WithField("err", err.Error()).Error("Error while decoding motor data")
			return
		}

		valid, respBytes := Validate(m)
		if !valid {
			responseBadRequest(rw, respBytes)
			return
		}

		err = db.SetMotorsValues([]db.Motor{m})
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			logger.WithField("err", err.Error()).Error("Error while inserting motor config")
			return
		}
		err = deps.Store.InsertMotor(req.Context(), []db.Motor{m})
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			logger.WithField("err", err.Error()).Error("Error while inserting motor")
			return
		}

		respBytes, err = json.Marshal(m)
		if err != nil {
			logger.WithField("err", err.Error()).Error("Error marshaling Motor data")
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}
		responseCodeAndMsg(rw, http.StatusCreated, respBytes)

	})
}

func updateMotorHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		var m db.Motor
		err := json.NewDecoder(req.Body).Decode(&m)
		if err != nil {
			responseCodeAndMsg(rw, http.StatusBadRequest, ErrObj{Err: "Error while decoding motor data"})
			logger.WithField("err", err.Error()).Error("Error while decoding motor data")
			return
		}

		valid, respBytes := Validate(m)
		if !valid {
			responseBadRequest(rw, respBytes)
			return
		}

		err = db.UpdateMotorsValues([]db.Motor{m})
		if err != nil {
			responseCodeAndMsg(rw, http.StatusInternalServerError, ErrObj{Err: "Error while while updating motor config"})
			logger.WithField("err", err.Error()).Error("Error while updating motor config")
			return
		}

		err = deps.Store.UpdateMotor(req.Context(), m)
		if err != nil {
			responseCodeAndMsg(rw, http.StatusInternalServerError, ErrObj{Err: "Error while while updating motor"})
			logger.WithField("err", err.Error()).Error("Error while updating motor")
			return
		}

		responseCodeAndMsg(rw, http.StatusOK, m)

	})
}
func deleteMotorHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		vars := mux.Vars(req)
		id, err := strconv.ParseInt(vars["id"], 10, 64)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			return
		}

		err = deps.Store.DeleteMotor(req.Context(), int(id))
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			logger.WithField("err", err.Error()).Error("Error while inserting motor")
			return
		}

		response := map[string]string{
			"msg": "motor deleted successfully",
		}

		responseCodeAndMsg(rw, http.StatusOK, response)

	})
}
func listMotorsHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		m, err := deps.Store.ListMotors(req.Context())
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			logger.WithField("err", err.Error()).Error("Error while inserting motor")
			return
		}

		responseCodeAndMsg(rw, http.StatusOK, m)

	})
}
