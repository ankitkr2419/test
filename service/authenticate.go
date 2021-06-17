package service

import (
	"context"
	"fmt"
	"mylab/cpagent/config"
	"mylab/cpagent/db"
	"mylab/cpagent/responses"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	logger "github.com/sirupsen/logrus"
)

type Token struct {
	jwt.StandardClaims
	Role            string            `json:"role"`
	Deck            string            `json:"deck"`
	AuthID          uuid.UUID         `json:"auth_id"`
	ApplicationType string            `json:"app_type"`
	Payload         map[string]string `json:"payload,omitempty"`
}

// Encodes information into token
func EncodeToken(userID string, authID uuid.UUID, role string, deck string, application string, payload map[string]string) (string, error) {
	accessKey := config.GetSecretKey()
	token, tokenErr := generateToken(userID, authID, role, deck, accessKey, application, payload)

	if tokenErr != nil {
		return "", fmt.Errorf("TOKEN ERROR: %v ", tokenErr)
	}
	return token, nil

}

func generateToken(userID string, authID uuid.UUID, role, deck, accessKey string, application string, payload map[string]string) (string, error) {
	tokenClaims := &Token{
		Role:            role,
		Deck:            deck,
		AuthID:          authID,
		ApplicationType: application,
		Payload:         payload,
		StandardClaims: jwt.StandardClaims{
			Subject:  userID,
			IssuedAt: time.Now().Unix(),
		},
	}

	tokenInit := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims)
	tokenString, err := tokenInit.SignedString([]byte(accessKey))

	if err != nil {
		return "", err
	}
	return tokenString, nil

}

// DecodeToken decodes token and returns claims if token is valid
func decodeToken(token string) (jwt.MapClaims, error) {
	accessKey := config.GetSecretKey()
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(accessKey), nil
	})

	if err != nil {

		return nil, err
	}
	if parsedToken.Valid {
		claims := parsedToken.Claims.(jwt.MapClaims)
		return claims, nil
	}
	return nil, err

}

// Authenticate ... Authenticate token sent in the request
// if token is valid set userId in the context
func authenticate(next http.HandlerFunc, deps Dependencies, roles ...string) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {

		token := extractToken(req.Header.Get("Authorization"))
		if token != "" {
			vars := mux.Vars(req)
			deck := vars["deck"]

			user, err := getUserAuth(token, deck, deps, roles...)
			if err != nil {
				logger.WithField("err", err.Error()).Error(responses.UserUnauthorised)
				responseCodeAndMsg(res, http.StatusUnauthorized, ErrObj{Err: err.Error()})
				return
			}
			ctx := context.WithValue(req.Context(), contextKeyUsername, user.Username)
			ctx = context.WithValue(ctx, contextKeyUserAuthID, user.AuthID)
			next(res, req.WithContext(ctx))

		} else {
			logger.WithField("err", "TOKEN EMPTY").Error(responses.UserTokenEmptyError)
			responseCodeAndMsg(res, http.StatusUnauthorized, ErrObj{Err: responses.UserTokenEmptyError.Error()})
			return
		}
	}
}

func getUserAuth(token, deck string, deps Dependencies, roles ...string) (user db.UserAuth, err error) {

	var validRole bool
	decodedToken, err := decodeToken(token)
	if err != nil {
		logger.WithField("err", err.Error()).Error(responses.UserTokenDecodeError)
		err = responses.UserTokenDecodeError
		return
	}

	roleFromToken, ok := decodedToken["role"]
	if !ok {
		logger.WithField("err", "EMPTY ROLE").Error(responses.UserTokenRoleEmptyError)
		err = responses.UserTokenRoleEmptyError
		return
	}

	//validating role
	if len(roles) != 0 {
		for i := 0; i < len(roles); i++ {
			if roleFromToken == roles[i] {
				validRole = true
			}
		}
	} else {
		validRole = true
	}

	if !validRole {
		logger.WithField("err", "INVALID ROLE").Error(responses.UserTokenInvalidRoleError)
		err = responses.UserTokenInvalidRoleError
		return
	}

	//validate deck
	tokenDeck, ok := decodedToken["deck"].(string)
	if !ok {
		logger.WithField("err", "DECK TOKEN").Error(responses.UserTokenDeckError)
		err = responses.UserTokenDeckError
		return
	}

	// First if the deck is received from an api other than logout then for cloud user the deck
	// value is "" hence this condition becomes false and it will not validate deck.
	if deck != "" {

		// Since the logout api doesnot call authenticate method and call getUserAuth method
		// directly for fetching userAuth, so deck passed earlier was "" hence cross deck
		// validation could not be performed for logout.
		// So passing deck as cloudUser when the cloud user wants to logout and setting it to "" for cross deck validation.
		if deck == "cloudUser" {
			deck = ""
		}

		if tokenDeck != deck {
			logger.WithField("err", "CROSS DECK TOKEN").Error(responses.UserTokenCrossDeckError)
			err = responses.UserTokenCrossDeckError
			return
		}

	}

	//Validate Application Type
	app, ok := decodedToken["app_type"].(string)
	if !ok {
		logger.WithField("err", "APPLICATIONTYPE").Error(responses.UserTokenApplicationTypeError)
		err = responses.UserTokenApplicationTypeError
		return
	}

	if app == "" {
		if Application == None {
			logger.WithField("err", "APPLICATIONTYPE").Error(responses.UserTokenApplicationError)
			err = responses.UserTokenApplicationError
			return
		}
	}

	if app == RTPCR || app == Extraction || app == Combined {
		//proceed futher
		if app != Application {
			logger.WithField("err", "APPLICATIONTYPE").Error(responses.UserTokenAppMismatchError)
			err = responses.UserTokenAppMismatchError
			return
		}
	} else {

	}

	username, ok := decodedToken["sub"].(string)
	if !ok {
		logger.WithField("err", "USERNAME").Error(responses.UserTokenUsernameError)
		err = responses.UserTokenUsernameError
		return
	}
	id, ok := decodedToken["auth_id"].(string)
	if !ok {
		logger.WithField("err", "AUTHID").Error(responses.UserTokenAuthIdError)
		err = responses.UserTokenAuthIdError
		return
	}
	authID, err := uuid.Parse(id)
	if err != nil {
		logger.WithField("err", err.Error()).Error(responses.UserTokenAuthIdParseError)
		err = responses.UserTokenAuthIdParseError
		return
	}
	user, err = deps.Store.ShowUserAuth(context.Background(), username, authID)
	if err != nil {
		logger.WithField("err", err.Error()).Error(responses.UserAuthNotFoundError)
		err = responses.UserAuthNotFoundError
		return
	}
	return
}

func extractToken(tokenWithBearer string) string {

	if tokenWithBearer != "" {
		token := strings.Split(tokenWithBearer, "Bearer ")
		if token[1] != "" {
			return token[1]
		}
	}
	return ""

}
