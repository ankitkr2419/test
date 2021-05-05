package service

import (
	"context"
	"errors"
	"fmt"
	"mylab/cpagent/config"
	"mylab/cpagent/db"
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
	Role    string            `json:"role"`
	Deck    string            `json:"deck"`
	AuthID  uuid.UUID         `json:"auth_id"`
	Payload map[string]string `json:"payload,omitempty"`
}

// Encodes information into token
func EncodeToken(userID string, authID uuid.UUID, role string, deck string, payload map[string]string) (string, error) {
	accessKey := config.GetSecretKey()
	token, tokenErr := generateToken(userID, authID, role, deck, accessKey, payload)

	if tokenErr != nil {
		return "", fmt.Errorf("TOKEN ERROR: %v ", tokenErr)
	}
	return token, nil

}

func generateToken(userID string, authID uuid.UUID, role, deck, accessKey string, payload map[string]string) (string, error) {
	tokenClaims := &Token{
		Role:    role,
		Deck:    deck,
		AuthID:  authID,
		Payload: payload,
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
func authenticate(next http.HandlerFunc, deps Dependencies, role ...string) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {

		token := extractToken(req.Header.Get("Authorization"))
		if token != "" {
			vars := mux.Vars(req)
			deck := vars["deck"]
			_, err := getUserAuth(token, deck, deps, role...)
			if err != nil {
				logger.Errorln("error in authorizing user :", err.Error())
				res.WriteHeader(http.StatusUnauthorized)
				res.Write([]byte(`{"error":"unauthorised user"}`))
				return
			}
			next(res, req)

		} else {
			res.WriteHeader(http.StatusUnauthorized)
			res.Write([]byte(`{"error":"unauthorised access"}`))
			return
		}
	}
}

func getUserAuth(token, deck string, deps Dependencies, role ...string) (user db.UserAuth, err error) {

	var validRole bool
	decodedToken, err := decodeToken(token)
	if err != nil {
		logger.Errorln("decoding token error", err.Error())
		err = errors.New("failed to decode token")
		return
	}

	roleFromToken, ok := decodedToken["role"]
	if !ok {
		logger.Errorln("failed to fetch role error")
		err = errors.New("failed to fetch role")
		return
	}

	//validating role
	if len(role) != 0 {
		for i := 0; i < len(role); i++ {
			if roleFromToken == role[i] {
				validRole = true
			}
		}
	} else {
		if roleFromToken == operator {
			validRole = true
		}
	}
	if !validRole {
		logger.Errorln("invalid role")
		err = errors.New("invalid role")
		return
	}

	//validate deck
	tokenDeck, ok := decodedToken["deck"].(string)
	if !ok {
		logger.Errorln("failed to fetch deck error")
		err = errors.New("failed to fetch deck")
		return
	}

	if tokenDeck != deck {
		logger.Errorln("invalid token for deck error")
		err = errors.New("wrong token for deck")
		return
	}

	if deck != "" {
		value, ok := userLogin.Load(deck)
		if !ok {
			logger.Errorln("invalid deck name error")
			err = errors.New("invalid deck name")
			return
		}
		if value.(bool) == false {
			logger.Errorln("deck logged out error")
			err = fmt.Errorf(`"error":"already logged out of deck %s"`, deck)
			return
		}
	}

	username, ok := decodedToken["sub"].(string)
	if !ok {
		logger.Errorln("username error")
		err = errors.New("failed to fetch user")
		return
	}
	id, ok := decodedToken["auth_id"].(string)
	if !ok {
		logger.Errorln("authID error")
		err = errors.New("failed to fetch user")
		return
	}
	authID, err := uuid.Parse(id)
	if err != nil {
		logger.Errorln("authID parse error")
		err = errors.New("failed to fetch user auth")
		return
	}
	user, err = deps.Store.ShowUserAuth(context.Background(), username, authID)
	if err != nil {
		logger.Errorln("user not found")
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
