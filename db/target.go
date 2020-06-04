package db

import (
	"context"

	"github.com/google/uuid"
	logger "github.com/sirupsen/logrus"
)

const (
	getTargetListQuery = `SELECT * FROM targets
		ORDER BY name ASC`
)

type Target struct {
	ID    uuid.UUID `db:"id" json:"id"`
	Name  string    `db:"name" json:"name" validate:"required"`
	DyeID uuid.UUID `db:"dye_id" json:"dye_id" validate:"required"`
}

func (s *pgStore) ListTargets(ctx context.Context) (t []Target, err error) {
	err = s.db.Select(&t, getTargetListQuery)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error listing targets")
		return
	}

	return
}
