package service

import (
	"context"
	"mylab/cpagent/db"

	logger "github.com/sirupsen/logrus"
)

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
