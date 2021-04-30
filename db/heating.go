package db

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
	logger "github.com/sirupsen/logrus"
	"mylab/cpagent/responses"
)

type Heating struct {
	ID          uuid.UUID `json:"id" db:"id"`
	Temperature float64   `json:"temperature" db:"temperature" validate:"required,gte=20,lte=120"`
	FollowTemp  bool      `json:"follow_temp" db:"follow_temp"`
	Duration    int64     `json:"duration" db:"duration" validate:"required,gte=10,lte=3660"`
	ProcessID   uuid.UUID `json:"process_id" db:"process_id" validate:"required"`
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
)

func (s *pgStore) ShowHeating(ctx context.Context, id uuid.UUID) (heating Heating, err error) {

	// get heating record
	err = s.db.Get(&heating,
		getHeatingQuery,
		id,
	)
	if err != nil {
		logger.WithField("err", err.Error()).Errorln(responses.HeatingDBFetchError)
		return
	}
	return

}

func (s *pgStore) CreateHeating(ctx context.Context, h Heating) (createdHeating Heating, err error) {
	var tx *sql.Tx

	//update the process name before record creation
	err = s.UpdateProcessName(ctx, h.ProcessID, "Heating", h)
	if err != nil {
		logger.WithField("err", err.Error()).Errorln(responses.HeatingUpdateNameError)
		return
	}

	tx, err = s.db.BeginTx(ctx, nil)
	if err != nil {
		logger.WithField("err:", err.Error()).Errorln(responses.HeatingInitiateDBTxError)
		return Heating{}, err
	}

	// End the transaction in defer call
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		tx.Commit()
	}()

	createdHeating, err = s.createHeating(ctx, h, tx)
	// failures are already logged
	return
}

func (s *pgStore) createHeating(ctx context.Context, h Heating, tx *sql.Tx) (createdHeating Heating, err error) {

	var lastInsertID uuid.UUID

	err = tx.QueryRow(
		createHeatingQuery,
		h.Temperature,
		h.FollowTemp,
		h.Duration,
		h.ProcessID,
	).Scan(&lastInsertID)

	if err != nil {
		logger.WithField("err", err.Error()).Errorln(responses.HeatingDBCreateError)
		return
	}

	createdHeating, err = s.ShowHeating(ctx, h.ProcessID)
	// failures are already logged
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
