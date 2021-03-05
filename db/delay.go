package db

import (
	"context"
	"time"

	"github.com/google/uuid"
	logger "github.com/sirupsen/logrus"
)

type Delay struct {
	ID        uuid.UUID     `db:"id" json:"id"`
	DelayTime time.Duration `db:"delay_time" json:"delay_time"`
	ProcessID uuid.UUID     `db:"distance" json:"distance"`
	CreatedAt time.Time     `db:"created_at" json:"created_at"`
	UpdatedAt time.Time     `db:"updated_at" json:"updated_at"`
}

const (
	//get delay data with recipe_id
	getDelayQuery = `SELECT * FROM delay where process_id = $1`
)

func (s *pgStore) ShowDelay(ctx context.Context, id uuid.UUID) (delay Delay, err error) {
	// get delay record
	err = s.db.Get(&delay,
		getDelayQuery,
		id,
	)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error getting heating")
		return
	}
	return
}
