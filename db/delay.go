package db

import (
	"context"
	"time"

	"github.com/google/uuid"
	logger "github.com/sirupsen/logrus"
)

type Delay struct {
	ID        uuid.UUID `db:"id" json:"id"`
	DelayTime int64     `db:"delay_time" json:"delay_time"`
	ProcessID uuid.UUID `db:"process_id" json:"process_id"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

const (
	getDelayQuery    = `SELECT * FROM delay where process_id = $1`
	createDelayQuery = `INSERT INTO delay (
		delay_time,
		process_id)
		VALUES ($1, $2) RETURNING id`
	updateDelayQuery = `UPDATE delay SET (
			delay_time,
			updated_at) = ($1, $2) WHERE WHERE process_id = $3`
)

func (s *pgStore) ShowDelay(ctx context.Context, id uuid.UUID) (delay Delay, err error) {
	// get delay record
	err = s.db.Get(&delay,
		getDelayQuery,
		id,
	)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error getting delay process")
		return
	}
	return
}

func (s *pgStore) CreateDelay(ctx context.Context, d Delay) (createdDelay Delay, err error) {
	var lastInsertID uuid.UUID
	err = s.db.QueryRow(
		createDelayQuery,
		d.DelayTime,
		d.ProcessID,
	).Scan(&lastInsertID)

	if err != nil {
		logger.WithField("err", err.Error()).Error("Error creating Delay")
		return
	}

	err = s.db.Get(&createdDelay, getDelayQuery, d.ProcessID)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error in getting Delay")
		return
	}
	return
}

func (s *pgStore) UpdateDelay(ctx context.Context, d Delay) (err error) {
	_, err = s.db.Exec(
		updateDelayQuery,
		d.DelayTime,
		time.Now(),
	)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error updating delay")
		return
	}
	return
}
