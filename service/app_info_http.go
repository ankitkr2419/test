package service

import (
	"context"
	"fmt"
	"mylab/cpagent/db"
	"mylab/cpagent/plc"
	"mylab/cpagent/responses"
	"net/http"
	"os"

	logger "github.com/sirupsen/logrus"
)

const (
	None       = "none"
	Combined   = "combined"
	RTPCR      = "rtpcr"
	Extraction = "extraction"
)

// TODO: Set Application variable in main via CLI
// variables for Binary Build info
var Version, Application, User, Machine, CommitID, Branch, BuiltOn string

func appInfoHandler(deps Dependencies) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		ctx := context.WithValue(req.Context(), contextKeyUsername, "main")

		appInfo := struct {
			Application string `json:"app"`
			Version     string `json:"version"`
			User        string `json:"user"`
			Machine     string `json:"machine"`
			CommitID    string `json:"commit_id"`
			Branch      string `json:"branch"`
			BuiltOn     string `json:"built_on"`
		}{
			Application: Application,
			Version:     Version,
			User:        User,
			Machine:     Machine,
			CommitID:    CommitID,
			Branch:      Branch,
			BuiltOn:     BuiltOn,
		}

		logger.Infoln(responses.AppInfoFetch, appInfo)
		responseCodeAndMsg(rw, http.StatusOK, appInfo)
		go deps.Store.AddAuditLog(ctx, db.ApiOperation, db.CompletedState, db.ShowOperation, "", responses.AppInfoRequested)
	})
}

// NOTE: Application doesn't make sense below as it is set at run time only
func PrintBinaryInfo() {
	fmt.Printf("\nApplication\t: %v \nVersion\t\t: %v \nUser\t\t: %v \nMachine\t\t: %v \nBranch\t\t: %v \nCommitID\t: %v \nBuilt\t\t: %v\n",
		Application, Version, User, Machine, Branch, CommitID, BuiltOn)
}

func ShutDownGracefully(deps Dependencies) (err error) {
	var err1, err2, err3, err4 error
	// We received an interrupt signal, shut down.
	logger.Warnln("\n..............................................\n----Application shutting down gracefully ----|\n.............................................|")
	if Application == Combined || Application == RTPCR {
		err1 = deps.Tec.ReachRoomTemp()
		if err1 != nil {
			logger.Errorln("Couldn't reach the room temp!")
			err = err1
		}
		err2 = deps.Plc.SwitchOffLidTemp()
		if err2 != nil {
			logger.Errorln("Couldn't Switch off the Lid!")
			err = fmt.Errorf("%v\n%v", err, err2)
		}
	}
	if Application == Combined || Application == Extraction {
		_, err3 = deps.PlcDeck[plc.DeckA].SwitchOffAllCoils()
		if err3 != nil {
			logger.Errorln("Couldn't switch off Deck A motor!")
			err = fmt.Errorf("%v\n%v", err, err3)
		}

		_, err4 = deps.PlcDeck[plc.DeckB].SwitchOffAllCoils()
		if err4 != nil {
			logger.Errorln("Couldn't switch off Deck B motor!")
			err = fmt.Errorf("%v\n%v", err, err4)
		}
	}
	if err != nil {
		logger.Errorln("Shutdown graceful error: ", err)
		os.Exit(-1)
	}

	logger.Warnln("\n...........................................\n-------Graceful Shutdown complete --------|\n..........................................|")

	os.Exit(0)
	return
}
