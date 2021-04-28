package service

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"mylab/cpagent/db"
	"net/http"

	"github.com/google/uuid"
	logger "github.com/sirupsen/logrus"
	"gopkg.in/go-playground/validator.v9"
)

const (
	hold  = "hold"
	cycle = "cycle"
)


var stepRun  map[string]bool
var nextStep map[string]chan struct{}

func resetStepRun(deck string){
	stepRun[deck] = false
}

func setStepRun(deck string){
	stepRun[deck] = true
}

func LoadUtils(){
	stepRun = map[string]bool{
		"A": false,
		"B": false,
	}

	chanA := make(chan struct{}, 1)
	chanB := make(chan struct{}, 1)

	nextStep = map[string]chan struct{}{
		"A": chanA,
		"B": chanB,
	}
}

func validate(i interface{}) (valid bool, respBytes []byte) {

	fieldErrors := make(map[string]string)

	v := validator.New()
	err := v.Struct(i)

	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			fieldErrors[e.Namespace()] = e.Tag()
			fieldErrors["error"] = "invalid value for field " + e.Field()

			logger.WithFields(logger.Fields{
				"field": e.Namespace(),
				"tag":   e.Tag(),
				"error": "invalid value",
			}).Error("Validation Error")
		}

		respBytes, err = json.Marshal(fieldErrors)
		if err != nil {
			logger.WithField("err", err.Error()).Error("Error marshaling data")
			return
		}
		return
	}

	valid = true
	return
}

func parseUUID(s string) (validUUID uuid.UUID, err error) {

	validUUID, err = uuid.Parse(s)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error id key invalid")
		return
	}

	return
}

func isvalidID(id uuid.UUID) bool {
	// 00000000-0000-0000-0000-000000000000 is default value of uuid
	if id.String() == "00000000-0000-0000-0000-000000000000" {
		return false
	}
	return true
}

func responseBadRequest(rw http.ResponseWriter, respBytes []byte) {
	rw.Header().Add("Content-Type", "application/json")
	rw.WriteHeader(http.StatusBadRequest)
	rw.Write(respBytes)
	return
}

// LogNotification add log for notification

func LogNotification(deps Dependencies, msg string) {

	n := db.Notification{
		Message: msg,
	}

	err := deps.Store.InsertNotification(context.Background(), n)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error in Log Notification")
		return
	}
	return
}

//MD5Hash return hashed string
func MD5Hash(s string) string {
	// hash string to [16]uint8
	hash := md5.Sum([]byte(s))

	return hex.EncodeToString(hash[:])
}
