package service

import (
	"encoding/json"
	"fmt"
	"mylab/cpagent/db"
	"mylab/cpagent/responses"
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
			responseCodeAndMsg(rw, http.StatusBadRequest, ErrObj{Err: responses.UserDecodeError.Error()})
			return
		}

		vars := mux.Vars(req)
		deck := vars["deck"]

		if deck != "" {
			value, ok := userLogin.Load(deck)
			if !ok {
				logger.WithField("err", "DECK ERROR").Error(responses.UserInvalidDeckError)
				responseCodeAndMsg(rw, http.StatusBadRequest, ErrObj{Err: responses.UserInvalidDeckError.Error()})
				return
			}
			if value.(bool) == true {
				logger.WithField("err", "DECK ERROR").Error(responses.UserDeckLoginError)
				responseCodeAndMsg(rw, http.StatusForbidden, ErrObj{Err: responses.UserDeckLoginError.Error()})
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

		// Getting back user along with his role
		u, err = deps.Store.ValidateUser(req.Context(), u)
		if err != nil || u.Role == "" {
			if err == nil {
				err = responses.UserInvalidError
			}
			logger.WithField("err", err.Error()).Error(responses.UserInvalidError)
			responseCodeAndMsg(rw, http.StatusInternalServerError, ErrObj{Err: responses.UserInvalidError.Error()})
			return
		}

		//create a new user_auth record
		authID, err := deps.Store.InsertUserAuths(req.Context(), u.Username)
		if err != nil {
			logger.WithField("err", err.Error()).Error(responses.UserAuthError)
			responseCodeAndMsg(rw, http.StatusInternalServerError, ErrObj{Err: responses.UserAuthError.Error()})
			return
		}

		if deck != "" {
			token, err = EncodeToken(u.Username, authID, u.Role, deck, Application, map[string]string{})
			userLogin.Store(deck, true)
		} else {
			token, err = EncodeToken(u.Username, authID, u.Role, "", Application, map[string]string{})
		}

		if err != nil {
			logger.WithField("err", err.Error()).Error(responses.UserTokenEncodeError)
			responseCodeAndMsg(rw, http.StatusInternalServerError, ErrObj{Err: responses.UserTokenEncodeError.Error()})
		}
		response := map[string]string{
			"msg":   fmt.Sprintf(`%s logged in successfully`, u.Role),
			"token": token,
			"role":  u.Role,
		}

		if err != nil {
			logger.WithField("err", err.Error()).Error(responses.UserMarshallingError)
			responseCodeAndMsg(rw, http.StatusInternalServerError, ErrObj{Err: responses.UserMarshallingError.Error()})
			return
		}
		logger.Infoln(responses.UserLoginSuccess)
		responseCodeAndMsg(rw, http.StatusOK, response)
	})
}

func createUserHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		var u db.User
		err := json.NewDecoder(req.Body).Decode(&u)
		if err != nil {
			responseCodeAndMsg(rw, http.StatusBadRequest, ErrObj{Err: responses.UserDecodeError.Error()})
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
			logger.WithField("err", err.Error()).Error(responses.UserInsertError)
			responseCodeAndMsg(rw, http.StatusInternalServerError, ErrObj{Err: responses.UserInsertError.Error()})
			return
		}

		logger.Infoln(responses.UserCreateSuccess)
		responseCodeAndMsg(rw, http.StatusCreated, MsgObj{Msg: responses.UserCreateSuccess})
		return
	})
}

func updateUserHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		var u db.User

		vars := mux.Vars(req)
		oldU := vars["old_username"]

		err := json.NewDecoder(req.Body).Decode(&u)
		if err != nil {
			logger.WithField("err", err.Error()).Errorln(responses.UserDecodeError)
			responseCodeAndMsg(rw, http.StatusBadRequest, ErrObj{Err: responses.UserDecodeError.Error()})
			return
		}

		valid, respBytes := validate(u)
		if !valid {
			responseBadRequest(rw, respBytes)
			return
		}

		//hash password to validate
		u.Password = MD5Hash(u.Password)

		err = deps.Store.UpdateUser(req.Context(), u, oldU)
		if err != nil {
			logger.WithField("err", err.Error()).Error(responses.UserUpdateError)
			responseCodeAndMsg(rw, http.StatusInternalServerError, ErrObj{Err: responses.UserUpdateError.Error()})
			return
		}

		logger.Infoln(responses.UserUpdateSuccess)
		responseCodeAndMsg(rw, http.StatusOK, MsgObj{Msg: responses.UserUpdateSuccess})
		return
	})
}

func logoutUserHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		token := extractToken(req.Header.Get("Authorization"))
		vars := mux.Vars(req)
		deck := vars["deck"]
		validRoles := []string{admin, engineer, supervisor, operator}

		// if the user is a deck user then only validate that if the user is logged out.
		// otherwise set the deck to cloud user
		if deck != "" {
			value, ok := userLogin.Load(deck)
			if !ok {
				logger.WithField("err", "DECK TOKEN").Error(responses.UserInvalidDeckError)
				responseCodeAndMsg(rw, http.StatusForbidden, ErrObj{Err: responses.UserInvalidDeckError.Error()})

				return
			}
			if value.(bool) == false {
				logger.WithField("err", "DECK LOGGED OUT").Error(responses.UserTokenLoggedOutDeckError)
				responseCodeAndMsg(rw, http.StatusForbidden, ErrObj{Err: responses.UserTokenLoggedOutDeckError.Error()})

				return
			}

		} else {
			deck = "cloudUser"
		}

		userAuth, err := getUserAuth(token, deck, deps, Application, validRoles...)
		if err != nil {
			logger.WithField("err", err.Error()).Error(responses.UserAuthDataFetchError)
			responseCodeAndMsg(rw, http.StatusForbidden, ErrObj{Err: responses.UserAuthDataFetchError.Error()})
			return
		}

		err = deps.Store.DeleteUserAuth(req.Context(), userAuth)
		if err != nil {
			logger.WithField("err", err.Error()).Error(responses.UserAuthDataDeleteError)
			responseCodeAndMsg(rw, http.StatusInternalServerError, ErrObj{Err: responses.UserAuthDataDeleteError.Error()})
			return
		}
		if deck != "" {
			userLogin.Store(deck, false)
		}

		logger.Infoln(responses.UserLogoutSuccess)
		responseCodeAndMsg(rw, http.StatusOK, MsgObj{Msg: responses.UserLogoutSuccess})
		return

	})
}
