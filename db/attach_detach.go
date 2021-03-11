package db

import (
	"context"
	"time"

	"github.com/google/uuid"
	logger "github.com/sirupsen/logrus"
)

type AttachDetach struct {
	ID            uuid.UUID `db:"id" json:"id"`
	Operation     string    `db:"operation" json:"operation"`
	OperationType string    `db:"operation_type" json:"operation_type"`
	ProcessID     uuid.UUID `db:"process_id" json:"process_id"`
	CreatedAt     time.Time `db:"created_at" json:"created_at"`
	UpdatedAt     time.Time `db:"updated_at" json:"updated_at"`
}

const (
	getAttachDetachQuery = `SELECT id, operation, operation_type, process_id, created_at, updated_at FROM attach_detach where process_id = $1`
)

func (s *pgStore) ShowAttachDetach(ctx context.Context, processID uuid.UUID) (ad AttachDetach, err error) {

	err = s.db.Get(&ad, getAttachDetachQuery, processID)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error fetching attach/detach operation")
		return
	}
	return
}
