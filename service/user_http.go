package service

import (
	"encoding/json"
	"mylab/cpagent/db"
	"net/http"

	logger "github.com/sirupsen/logrus"
)

func validateUserHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		var u db.User
		err := json.NewDecoder(req.Body).Decode(&u)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			logger.WithField("err", err.Error()).Error("Error while decoding user data")
			return
		}

		valid, respBytes := validate(u)
		if !valid {
			responseBadRequest(rw, respBytes)
			return
		}

		//hash password to validate
		u.Password = MD5Hash(u.Password)

		err = deps.Store.ValidateUser(req.Context(), u)
		if err != nil {
			if err.Error() == "Record Not Found" {
				rw.Header().Add("Content-Type", "application/json")
				rw.WriteHeader(http.StatusExpectationFailed)
				rw.Write([]byte(`{"msg":"Invalid User"}`))
				return
			}
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}
		rw.Header().Add("Content-Type", "application/json")
		rw.WriteHeader(http.StatusOK)
		rw.Write([]byte(`{"msg":"Validated User Sucessfully"}`))
	})
}

func createUserHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		// ASK and VEFRIFY: @Paramita only admin/supervisor are allowed to create roles, 
		// so this API chould be only called by these 2 roles
		var u db.User
		rw.Header().Add("Content-Type", "application/json")
		err := json.NewDecoder(req.Body).Decode(&u)
		if err != nil {
			logger.WithField("err", err.Error()).Error("Error while decoding user data: ", req.Body)
			rw.WriteHeader(http.StatusBadRequest)
			rw.Write([]byte(`{"msg":"Error while decoding user data"}`))
			return
		}

		valid, respBytes := validate(u)
		if !valid {
			responseBadRequest(rw, respBytes)
			return
		}

		//hash password to validate
		u.Password = MD5Hash(u.Password)

		err = deps.Store.InsertUser(req.Context(), u)
		if err != nil {
			logger.WithField("err", err.Error()).Error("Error while inserting user", u)	
			rw.WriteHeader(http.StatusInternalServerError)
			rw.Write([]byte(`{"msg":"Error while inserting user"}`))
			return
		}
		
		logger.Infoln(u, "user inserted successfully")
		rw.WriteHeader(http.StatusCreated)
		rw.Write([]byte(`{"msg":"Created User Sucessfully"}`))
		return
	})
}