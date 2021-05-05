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
		var token string
		err := json.NewDecoder(req.Body).Decode(&u)
		if err != nil {
			rw.WriteHeader(http.StatusBadRequest)
			logger.WithField("err", err.Error()).Error("Error while decoding user data")
			return
		}

		vars := mux.Vars(req)
		deck := vars["deck"]

		if deck != "" {
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

		//create a new user_auth record
		authID, err := deps.Store.InsertUserAuths(req.Context(), u.Username)
		if err != nil {
			rw.Write([]byte(`{"msg":"user login failed"}`))
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}

		if deck != "" {
			token, err = EncodeToken(u.Username, authID, u.Role, deck, map[string]string{})
			userLogin.Store(deck, true)
		} else {
			token, err = EncodeToken(u.Username, authID, u.Role, "", map[string]string{})
		}

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

func createUserHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

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

func logoutUserHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		token := extractToken(req.Header.Get("Authorization"))
		vars := mux.Vars(req)
		deck := vars["deck"]
		validRoles := []string{admin, engineer, supervisor, operator}

		userAuth, err := getUserAuth(token, deck, deps, validRoles...)
		if err != nil {
			logger.WithField("err", err.Error()).Error("Error while fetching user authentication data", userAuth)
			rw.WriteHeader(http.StatusForbidden)
			// TODO check what message to give here.
			rw.Write([]byte(`{"error":"invalid user authentication data"}`))
			return
		}

		err = deps.Store.DeleteUserAuth(req.Context(), userAuth)
		if err != nil {
			logger.WithField("err", err.Error()).Error("Error while deleting user authentication data", userAuth)
			rw.WriteHeader(http.StatusInternalServerError)
			rw.Write([]byte(`{"msg":"Error while deleting user authentication data"}`))
			return
		}
		if deck != "" {
			userLogin.Store(deck, false)
		}

		logger.Infoln(userAuth, "user logged out successfully")
		rw.WriteHeader(http.StatusOK)
		rw.Write([]byte(`{"msg":"user logged out successfully"}`))
		return

	})
}
