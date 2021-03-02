package db

import (
	"context"
	"time"

	"github.com/google/uuid"
	logger "github.com/sirupsen/logrus"
)

type AttachDetach struct {
	ID         uuid.UUID `db:"id" json:"id"`
	AttachType string    `db:"attach_type" json:"attach_type"`
	DetachType string    `db:"detach_type" json:"detach_type"`
	Height     float64   `db:"height" json:"height"`
	CreatedAt  time.Time `db:"created_at" json:"created_at"`
	UpdatedAt  time.Time `db:"updated_at" json:"updated_at"`
}

const (
	getAttachDetachQuery = `SELECT * FROM attach_detach where process_id = $1`
)

func (s *pgStore) ShowAttachDetach(ctx context.Context, processID uuid.UUID) (ad AttachDetach, err error) {

	err = s.db.Select(ad, getAttachDetachQuery, processID)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error fetching attach/detach operation")
		return
	}
	return
}
