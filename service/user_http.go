package service

import (
	"encoding/json"
	"fmt"
	"mylab/cpagent/db"
	"mylab/cpagent/plc"
	"mylab/cpagent/responses"
	"net/http"

	"github.com/gorilla/mux"

	logger "github.com/sirupsen/logrus"
)

func validateUserHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		var u db.User
		var token string
		var valueI interface{}
		var ok bool

		vars := mux.Vars(req)
		deck := vars["deck"]

		err := json.NewDecoder(req.Body).Decode(&u)
		if err != nil {
			responseCodeAndMsg(rw, http.StatusBadRequest, ErrObj{Err: responses.UserDecodeError.Error(), Deck: deck})
			return
		}

		if deck == blank {
			goto skipDeckUserCheck
		}

		valueI, ok = deckUserLogin.Load(deck)
		if !ok {
			logger.WithField("err", "DECK ERROR").Error(responses.UserInvalidDeckError)
			responseCodeAndMsg(rw, http.StatusBadRequest, ErrObj{Err: responses.UserInvalidDeckError.Error(), Deck: deck})
			return
		}
		if valueI.(string) != u.Username && valueI.(string) != blank {
			loggedInInfo := fmt.Errorf("%v. %v user already logged in.", responses.UserDeckLoginError, valueI)
			logger.WithField("err", "DECK ERROR").Error(loggedInInfo)
			responseCodeAndMsg(rw, http.StatusForbidden, ErrObj{Err: loggedInInfo.Error(), Deck: deck})
			return
		}

	skipDeckUserCheck:

		valid, respBytes := Validate(u)
		if !valid {
			responseBadRequest(rw, respBytes)
			return
		}

		//hash password to validate
		u.Password = MD5Hash(u.Password)

		// Getting back user along with his role
		u, err = deps.Store.ValidateUser(req.Context(), u)
		if err != nil || u.Role == blank {
			if err == nil {
				err = responses.UserInvalidError
			}
			logger.WithField("err", err.Error()).Error(responses.UserInvalidError)
			responseCodeAndMsg(rw, http.StatusInternalServerError, ErrObj{Err: responses.UserInvalidError.Error(), Deck: deck})
			return
		}

		//create a new user_auth record
		authID, err := deps.Store.InsertUserAuths(req.Context(), u.Username)
		if err != nil {
			logger.WithField("err", err.Error()).Error(responses.UserAuthError)
			responseCodeAndMsg(rw, http.StatusInternalServerError, ErrObj{Err: responses.UserAuthError.Error(), Deck: deck})
			return
		}

		if deck != blank {
			token, err = EncodeToken(u.Username, authID, u.Role, deck, Application, map[string]string{})
			deckUserLogin.Store(deck, u.Username)
		} else {
			token, err = EncodeToken(u.Username, authID, u.Role, blank, Application, map[string]string{})
		}

		if err != nil {
			logger.WithField("err", err.Error()).Error(responses.UserTokenEncodeError)
			responseCodeAndMsg(rw, http.StatusInternalServerError, ErrObj{Err: responses.UserTokenEncodeError.Error(), Deck: deck})
			return
		}
		response := map[string]string{
			"msg":   fmt.Sprintf(`%s logged in successfully`, u.Role),
			"token": token,
			"role":  u.Role,
		}

		logger.WithFields(logger.Fields{
			"Username": u.Username,
			"Role":     u.Role,
			"Deck":     deck,
		}).Infoln("User logged in successfully")

		checkEngOrAdminLoggedOnDeck(deps, deck, u)

		logger.Infoln(responses.UserLoginSuccess)
		responseCodeAndMsg(rw, http.StatusOK, response)

		reloadPLCFuncsExceptUtils(deps, deck)

	})
}

func reloadPLCFuncsExceptUtils(deps Dependencies, deck string) {
	if deck != blank && (Application == Extraction || Application == Combined) && !deps.PlcDeck[anotherDeck(deck)].IsRunInProgress() {
		logger.Infoln("Reloading all the PLC Funcs")
		go plc.LoadAllPLCFuncsExceptUtils(deps.Store)
	}
}

func anotherDeck(deck string) string {
	if deck == plc.DeckA {
		return plc.DeckB
	}
	return plc.DeckA
}

func checkEngOrAdminLoggedOnDeck(deps Dependencies, deck string, u db.User) {
	if deck != blank && (u.Role == admin || u.Role == engineer) && (Application == Combined || Application == Extraction) {
		deps.PlcDeck[deck].SetEngineerOrAdminLogged(true)
	} else if deck != blank && (Application == Combined || Application == Extraction) {
		deps.PlcDeck[deck].SetEngineerOrAdminLogged(false)
	}
	return
}

func createUserHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		var u db.User
		err := json.NewDecoder(req.Body).Decode(&u)
		if err != nil {
			responseCodeAndMsg(rw, http.StatusBadRequest, ErrObj{Err: responses.UserDecodeError.Error()})
			return
		}

		valid, respBytes := Validate(u)
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

		valid, respBytes := Validate(u)
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
		var ok bool
		var userI interface{}

		// if the user is a deck user then only validate that if the user is logged out.
		// otherwise set the deck to cloud user

		if deck == blank {
			deck = "cloudUser"
			goto skipDeckUserCheck
		}

		userI, ok = deckUserLogin.Load(deck)
		if !ok {
			logger.WithField("err", "DECK TOKEN").Error(responses.UserInvalidDeckError)
			responseCodeAndMsg(rw, http.StatusForbidden, ErrObj{Err: responses.UserInvalidDeckError.Error()})
			return
		}

		if userI.(string) == blank {
			logger.WithField("err", "DECK ALREADY LOGGED OUT").Error(responses.UserTokenLoggedOutDeckError)
			responseCodeAndMsg(rw, http.StatusForbidden, ErrObj{Err: responses.UserTokenLoggedOutDeckError.Error()})
			return
		}

	skipDeckUserCheck:
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
		if deck == plc.DeckA || deck == plc.DeckB {
			deckUserLogin.Store(deck, blank)
			if Application == Combined || Application == Extraction {
				deps.PlcDeck[deck].SetEngineerOrAdminLogged(false)
			}
		}
		logger.WithFields(logger.Fields{
			"Deck": deck,
		}).Infoln("User logged out successfully")

		responseCodeAndMsg(rw, http.StatusOK, MsgObj{Msg: responses.UserLogoutSuccess})
		return

	})
}

func deleteUserHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		vars := mux.Vars(req)
		username := vars["username"]

		currentUser := req.Context().Value(db.ContextKeyUsername).(string)
		if currentUser == username {
			logger.WithField("err", "SAMEUSERERROR").Error(responses.SameUserDeleteError)
			responseCodeAndMsg(rw, http.StatusForbidden, ErrObj{Err: responses.SameUserDeleteError.Error()})
			return
		}
		err := deps.Store.DeleteUser(req.Context(), username)
		if err != nil {
			logger.WithField("err", err.Error())
			responseCodeAndMsg(rw, http.StatusInternalServerError, ErrObj{Err: responses.UserDeleteError.Error()})
			return
		}

		logger.Infoln(responses.UserDeleteSuccess)
		responseCodeAndMsg(rw, http.StatusOK, MsgObj{Msg: responses.UserDeleteSuccess})
	})
}
