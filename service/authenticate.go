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
	logger "github.com/sirupsen/logrus"
)

type Claims struct {
	jwt.StandardClaims
	Role    string            `json:"role"`
	Deck    string            `json:"deck"`
	Payload map[string]string `json:"payload,omitempty"`
}

// Encodes information into token
func EncodeToken(userID string, role string, deck string, payload map[string]string) (string, error) {
	accessKey := config.GetSecretKey()
	token, tokenErr := generateToken(userID, role, deck, accessKey, payload)

	if tokenErr != nil {
		return "", fmt.Errorf("TOKEN ERROR: %v ", tokenErr)
	}
	return token, nil

}

func generateToken(userID, role, deck, accessKey string, payload map[string]string) (string, error) {
	tokenClaims := &Claims{
		Role:    role,
		Deck:    deck,
		Payload: payload,
		StandardClaims: jwt.StandardClaims{
			Subject:   userID,
			ExpiresAt: time.Now().Unix() + int64((time.Hour * 24).Seconds()),
			IssuedAt:  time.Now().Unix(),
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
func authenticate(next http.HandlerFunc, deps Dependencies) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {

		token := extractToken(req.Header.Get("Authorization"))
		if token != "" {
			decodedToken, err := decodeToken(token)
			if err != nil {
				res.WriteHeader(http.StatusUnauthorized)
				res.Write([]byte(`{"error":"error in decoding token"}`))
				return
			}

			_, err = getUser(decodedToken, deps)
			if err != nil {
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

func getUser(token jwt.MapClaims, deps Dependencies) (user db.User, err error) {
	username, ok := token["sub"].(string)
	if !ok {
		err = errors.New("failed to fetch user")
		return
	}
	user, err = deps.Store.ShowUser(context.Background(), username)
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

func authenticateAdmin(next http.HandlerFunc, deps Dependencies) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {

		token := extractToken(req.Header.Get("Authorization"))

		if token != "" {

			decodedToken, err := decodeToken(token)
			if err != nil {
				res.WriteHeader(http.StatusUnauthorized)
				res.Write([]byte(`{"error":"error in decoding token"}`))
				return
			}

			role, ok := decodedToken["role"].(string)
			if !ok {
				res.WriteHeader(http.StatusUnauthorized)
				res.Write([]byte(`{"error":"role not specified"}`))
				return
			}

			// TODO : add condition for engineer too when that flow is clear.
			if role != "admin" {
				res.WriteHeader(http.StatusForbidden)
				res.Write([]byte(`{"error":"action forbidden"}`))
				return
			}

			_, err = getUser(decodedToken, deps)
			if err != nil {
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
