package db

import (
	"context"
	"time"

	"github.com/google/uuid"
	logger "github.com/sirupsen/logrus"
)

type AttachDetach struct {
	ID            uuid.UUID `db:"id" json:"id"`
	Operation     string    `db:"operation" json:"operation"  validate:"required"`
	OperationType string    `db:"operation_type" json:"operation_type"`
	ProcessID     uuid.UUID `db:"process_id" json:"process_id" validate:"required"`
	CreatedAt     time.Time `db:"created_at" json:"created_at"`
	UpdatedAt     time.Time `db:"updated_at" json:"updated_at"`
}

const (
	getAttachDetachQuery    = `SELECT * FROM attach_detach where process_id = $1`
	createAttachDetachQuery = `INSERT INTO attach_detach (
		operation,
		operation_type,
		process_id)
		VALUES ($1, $2, $3) RETURNING id`
	updateAttachDetachQuery = `UPDATE attach_detach SET (
			operation,
			operation_type,
			updated_at) = 
			($1, $2, $3) WHERE process_id = $4`
)

func (s *pgStore) ShowAttachDetach(ctx context.Context, processID uuid.UUID) (ad AttachDetach, err error) {

	err = s.db.Get(&ad, getAttachDetachQuery, processID)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error fetching attach/detach operation")
		return
	}
	return
}

func (s *pgStore) CreateAttachDetach(ctx context.Context, a AttachDetach) (createdAttachDetach AttachDetach, err error) {
	var lastInsertID uuid.UUID

	err = s.UpdateProcessName(ctx, a.ProcessID, "AttachDetach", a)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error in updating attach detach process name")
		return
	}

	err = s.db.QueryRow(
		createAttachDetachQuery,
		a.Operation,
		a.OperationType,
		a.ProcessID,
	).Scan(&lastInsertID)

	if err != nil {
		logger.WithField("err", err.Error()).Error("Error creating AttachDetach")
		return
	}

	err = s.db.Get(&createdAttachDetach, getAttachDetachQuery, a.ProcessID)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error in getting AttachDetach")
		return
	}

	return
}

func (s *pgStore) UpdateAttachDetach(ctx context.Context, a AttachDetach) (err error) {
	_, err = s.db.Exec(
		updateAttachDetachQuery,
		a.Operation,
		a.OperationType,
		time.Now(),
		a.ProcessID,
	)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error updating attach detach")
		return
	}
	return
}
