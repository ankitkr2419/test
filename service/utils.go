package service

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	logger "github.com/sirupsen/logrus"
	"gopkg.in/go-playground/validator.v9"
)

func validate(i interface{}) (valid bool, respBytes []byte) {

	fieldErrors := make(map[string]string)

	v := validator.New()
	err := v.Struct(i)

	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			fieldErrors[e.Namespace()] = e.Tag()

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
