package db

import (
	"context"
	"time"

	"github.com/google/uuid"
	logger "github.com/sirupsen/logrus"
)

const (
	insertNotificationQuery = `INSERT INTO notifications (
		message,
		read)
		VALUES ($1,$2) `

	getNotificationQuery = `SELECT * FROM notifications
		WHERE
		read = false`

	markasReadQuery = `UPDATE notifications
		SET read = true
		WHERE id  = $1`
)

// Notification stores all the messages
type Notification struct {
	ID        uuid.UUID `db:"id" json:"id"`
	Message   string    `db:"message" json:"message"`
	Read      bool      `db:"read" json:"read"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

func (s *pgStore) ListNotification(ctx context.Context, experimentID uuid.UUID) (t []Notification, err error) {
	err = s.db.Select(&t, getNotificationQuery)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error listing result temperature details")
		return
	}
	return
}

func (s *pgStore) InsertNotification(ctx context.Context, n Notification) (err error) {

	_, err = s.db.Exec(
		insertNotificationQuery,
		n.Message,
		false,
	)
	if err != nil {
		logger.WithField("error in notification exec query", err.Error()).Error("Query Failed")
		return
	}

	return
}

func (s *pgStore) MarkNotificationasRead(ctx context.Context, id uuid.UUID) (err error) {

	_, err = s.db.Exec(
		markasReadQuery,
		id,
	)
	if err != nil {
		logger.WithField("error in notification exec query", err.Error()).Error("Query Failed")
		return
	}

	return
}
