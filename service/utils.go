package service

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
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

func getProcessName(processType string, process interface{}) (processName string, err error) {

	switch processType {
	case "Piercing":
		piercing := process.(db.Piercing)
		processName = fmt.Sprintf("Piercing_%s", piercing.Type)
		return

	case "TipOperation":
		tipOpr := process.(db.TipOperation)
		processName = fmt.Sprintf("Tip_Operation_%s", tipOpr.Type)
		return

	case "TipDocking":
		tipDock := process.(db.TipDock)
		processName = fmt.Sprintf("Tip_Docking_%s", tipDock.Type)
		return

	case "AspireDispense":
		aspDis := process.(db.AspireDispense)
		processName = fmt.Sprintf("Aspire_Dispense_%s", aspDis.Category)
		return

	case "Heating":
		processName = fmt.Sprintf("Heating")
		return

	case "Shaking":
		shaking := process.(db.Shaker)
		if shaking.WithTemp {
			processName = fmt.Sprintf("Shaking_With_temperature")
			return
		}
		processName = fmt.Sprintf("Shaking_Without_temperature")
		return

	case "AttachDetach":
		atDet := process.(db.AttachDetach)
		processName = fmt.Sprintf("Magnet_%s", atDet.Operation)
		return

	case "Delay":
		processName = fmt.Sprintf("Delay")
		return

	default:
		return "", errors.New("wrong process type")
	}
}

func updateProcessName(ctx context.Context, deps Dependencies, processID uuid.UUID, processType string, process interface{}) (err error) {

	processName, err := getProcessName(processType, process)
	if err != nil {
		err = fmt.Errorf("error in creating new process name")
		return
	}

	err = deps.Store.UpdateProcessName(ctx, processID, processName)
	if err != nil {
		err = fmt.Errorf("error in updating process name")
		return
	}
	return
}
