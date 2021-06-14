package service

import (
	"fmt"
	"context"
	"mylab/cpagent/db"
	"mylab/cpagent/responses"
	"net/http"

	logger "github.com/sirupsen/logrus"
)

const (
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
	fmt.Printf("\nVersion\t\t: %v \nUser\t\t: %v \nMachine\t\t: %v \nBranch\t\t: %v \nCommitID\t: %v \nBuilt\t\t: %v\n",
		Version, User, Machine, Branch, CommitID, BuiltOn)
}
