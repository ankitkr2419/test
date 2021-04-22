package service

import (
	"encoding/json"
	"mylab/cpagent/db"
	"net/http"

	"github.com/gorilla/mux"

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

		vars := mux.Vars(req)
		deck := vars["deck"]
		logger.Infoln(deck)
		value, ok := userLogin.Load(deck)
		if !ok {
			rw.WriteHeader(http.StatusBadRequest)
			rw.Write([]byte(`{"error:"invalid deck name"}`))
			return
		}
		if value.(bool) == true {
			rw.WriteHeader(http.StatusForbidden)
			rw.Write([]byte(`{"error:"not allowed to login"}`))
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

		token, err := EncodeToken(u.Username, u.Role, "A", map[string]string{})

		response, err := json.Marshal(map[string]string{
			"msg":   "user logged in successfully",
			"token": token,
		})

		if err != nil {
			rw.Write([]byte(`{"msg":"user login failed"}`))
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}
		rw.Header().Add("Content-Type", "application/json")
		rw.WriteHeader(http.StatusOK)
		rw.Write(response)
	})
}
