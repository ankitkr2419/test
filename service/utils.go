package service

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"mylab/cpagent/db"
	"mylab/cpagent/plc"
	"mylab/cpagent/responses"
	"net/http"
	"sync"
	"time"

	"github.com/google/uuid"
	logger "github.com/sirupsen/logrus"
	"gopkg.in/go-playground/validator.v9"
)

const (
	contextKeyUsername          = "username"
	contextKeyUserAuthID        = "auth_id"
	blank                string = ""
	// NOTE: These are version specific
	maxDeckPosition  = 11
	minAspDisDeckPos = 6
	cartridge1Pos    = 8
	cartridge2Pos    = 10
	// Shaker Temp related ranges
	minShakerTempAllowed = 20
	maxShakerTempAllowed = 120
)

const (
	hold  = "hold"
	cycle = "cycle"
)

const (
	admin      = "admin"
	supervisor = "supervisor"
	engineer   = "engineer"
	operator   = "operator"
)

const (
	recipeC    = "recipe"
	processC   = "process"
	createC    = "create"
	deleteC    = "delete"
	updateC    = "update"
	duplicateC = "duplicate"
	rearrangeC = "rearrange"
)

var templateRunSuccess bool
var expStartTime time.Time

// TODO: Don't allow Template Update/Deletion if this Template is in Progress
var currentExpTemplate db.Template

var deckUserLogin sync.Map

// runNext will run the next step of process when set
var runNext, stepRunInProgress map[string]bool

// abortStepRun will help track if abort was performed after process completion
var nextStep, abortStepRun map[string]chan struct{}

func resetRunNext(deck string) {
	runNext[deck] = false
}

func setRunNext(deck string) {
	runNext[deck] = true
}

func resetStepRunInProgress(deck string) {
	stepRunInProgress[deck] = false
}

func setStepRunInProgress(deck string) {
	stepRunInProgress[deck] = true
}

func loadUtils() {
	deckUserLogin.Store(plc.DeckA, blank)
	deckUserLogin.Store(plc.DeckB, blank)
	runNext = map[string]bool{
		plc.DeckA: false,
		plc.DeckB: false,
	}

	stepRunInProgress = map[string]bool{
		plc.DeckA: false,
		plc.DeckB: false,
	}

	chanA := make(chan struct{}, 1)
	chanB := make(chan struct{}, 1)

	nextStep = map[string]chan struct{}{
		plc.DeckA: chanA,
		plc.DeckB: chanB,
	}

	chanC := make(chan struct{}, 1)
	chanD := make(chan struct{}, 1)

	abortStepRun = map[string]chan struct{}{
		plc.DeckA: chanC,
		plc.DeckB: chanD,
	}
}

type ErrObj struct {
	Err  string `json:"err"`
	Deck string `json:"deck,omitempty"`
}

type MsgObj struct {
	Msg  string `json:"msg"`
	Deck string `json:"deck,omitempty"`
}

func Validate(i interface{}) (valid bool, respBytes []byte) {

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
	logger.WithField("err", respBytes)
	rw.WriteHeader(http.StatusBadRequest)
	rw.Write(respBytes)
	return
}

func responseCodeAndMsg(rw http.ResponseWriter, statusCode int, msg interface{}) {
	rw.Header().Add("Content-Type", "application/json")
	rw.WriteHeader(statusCode)
	respBytes, err := json.Marshal(msg)
	if err != nil {
		logger.WithField("err", err.Error()).Error(responses.DataMarshallingError)
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write(responses.DataMarshallingError)
		return
	}
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

func LoadAllServiceFuncs(s db.Storer) (err error) {

	// Delete Unfinished Templates
	err = s.DeleteUnfinishedTemplates(context.Background())
	if err != nil {
		logger.WithField("err", err.Error()).Error("Cleanup of unfinished templates failed")
		return
	}

	// Create a default supervisor
	supervisor := db.User{
		Username: "supervisor",
		Password: MD5Hash("supervisor"),
		Role:     "supervisor",
	}

	// Add Default supervisor user to DB
	err = s.InsertUser(context.Background(), supervisor)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Setup Default Supervisor failed")
		return
	}

	// Create a default main user
	mainUser := db.User{
		Username: "main",
		Password: MD5Hash("main"),
		Role:     "admin",
	}

	// Add Default main user to DB
	err = s.InsertUser(context.Background(), mainUser)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Setup Default Main User failed")
		return
	}

	// Create a default operator user
	operatorUser := db.User{
		Username: "operator",
		Password: MD5Hash("operator"),
		Role:     "operator",
	}

	// Add Default operator user to DB
	err = s.InsertUser(context.Background(), operatorUser)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Setup Default Operator User failed")
		return
	}

	// Create a default engineer user
	engUser := db.User{
		Username: "engineer",
		Password: MD5Hash("engineer"),
		Role:     "engineer",
	}

	// Add Default engineer user to DB
	err = s.InsertUser(context.Background(), engUser)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Setup Default Engineer User failed")
		return
	}

	logger.Info("Default users added")

	loadUtils()
	return nil
}

func ValidateDyeTarget(wc db.WellConfig, deps Dependencies) (valid bool, msg string) {

	valid = true
	//max number of dyes 6. TODO take from config
	dyes := make(map[string]bool, 6)
	for _, v := range wc.Targets {
		dye, err := deps.Store.ListTargetDye(context.Background(), v)
		if err != nil {
			return false, err.Error()
		}
		for dyeKey := range dyes {
			if dyeKey == dye {
				return false, "invalid configuration for targets"
			}
		}
		dyes[dye] = true
	}
	return
}
