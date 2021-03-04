package db

import (
	"context"
	"time"

	"github.com/google/uuid"
	logger "github.com/sirupsen/logrus"
)

type TipDock struct {
	ID        uuid.UUID `json:"id" db:"id"`
	Type      string    `json:"type" db:"type"`
	Position  int64     `json:"position" db:"position"`
	Height    float64   `json:"height" db:"height"`
	ProcessID uuid.UUID `json:"process_id" db:"process_id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

const (
	getTipDockQuery = `SELECT * FROM tip_dock where process_id = $1`
)

func (s *pgStore) ShowTipDocking(ctx context.Context, pid uuid.UUID) (td TipDock, err error) {

	err = s.db.Get(&td, getTipDockQuery, pid)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error listing tip docking details")
		return
	}
	return
}
