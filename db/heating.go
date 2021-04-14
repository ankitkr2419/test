package db

import (
	"context"
	"time"

	"github.com/google/uuid"
	logger "github.com/sirupsen/logrus"
)

type Heating struct {
	ID          uuid.UUID `json:"id" db:"id"`
	Temperature float64   `json:"temperature" db:"temperature"`
	FollowTemp  bool      `json:"follow_temp" db:"follow_temp"`
	Duration    int64     `json:"duration" db:"duration"`
	ProcessID   uuid.UUID `json:"process_id" db:"process_id"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

const (
	getHeatingQuery    = `SELECT * FROM heating where process_id = $1`
	createHeatingQuery = `INSERT INTO heating (
		temperature,
		follow_temp,
		duration,
		process_id)
		VALUES ($1, $2, $3, $4) RETURNING id`
	updateHeatingQuery = `UPDATE heating SET (
		temperature,
		follow_temp,
		duration,
		updated_at) = 
		($1, $2, $3, $4) WHERE process_id = $5`
	deleteHeatingQuery = `DELETE FROM processes
	WHERE id = $1`
)

func (s *pgStore) ShowHeating(ctx context.Context, id uuid.UUID) (heating Heating, err error) {

	// get heating record
	err = s.db.Get(&heating,
		getHeatingQuery,
		id,
	)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error getting heating")
		return
	}
	return

}

func (s *pgStore) CreateHeating(ctx context.Context, h Heating) (createdHeating Heating, err error) {
	var lastInsertID uuid.UUID

	err = s.db.QueryRow(
		createHeatingQuery,
		h.Temperature,
		h.FollowTemp,
		h.Duration,
		h.ProcessID,
	).Scan(&lastInsertID)

	if err != nil {
		logger.WithField("err", err.Error()).Error("Error creating Heating")
		return
	}

	err = s.db.Get(&createdHeating, getHeatingQuery, h.ProcessID)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error in getting Heating")
		return
	}
	return
}

func (s *pgStore) UpdateHeating(ctx context.Context, ht Heating) (err error) {
	_, err = s.db.Exec(
		updateHeatingQuery,
		ht.Temperature,
		ht.FollowTemp,
		ht.Duration,
		time.Now(),
		ht.ProcessID,
	)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error updating heating")
		return
	}
	return
}

func (s *pgStore) DeleteHeating(ctx context.Context, id uuid.UUID) (err error) {
	_, err = s.db.Exec(deleteHeatingQuery, id)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error deleting heating")
		return
	}
	return
}
