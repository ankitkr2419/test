package db

import (
	"context"
	"time"

	"database/sql"
	"github.com/google/uuid"
	logger "github.com/sirupsen/logrus"
	"mylab/cpagent/responses"
)

type Shaker struct {
	ID          uuid.UUID `json:"id" db:"id"`
	WithTemp    bool      `json:"with_temp" db:"with_temp"`
	Temperature float64   `json:"temperature" db:"temperature" validate:"required_with=WithTemp,gte=20,lte=120"`
	FollowTemp  bool      `json:"follow_temp" db:"follow_temp"`
	ProcessID   uuid.UUID `json:"process_id" db:"process_id" validate:"required"`
	RPM1        int64     `json:"rpm_1" db:"rpm_1" validate:"required"`
	RPM2        int64     `json:"rpm_2" db:"rpm_2"`
	Time1       int64     `json:"time_1" db:"time_1" validate:"required"`
	Time2       int64     `json:"time_2" db:"time_2"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

const (
	getShakerQuery     = `SELECT * FROM shaking where process_id = $1`
	createShakingQuery = `INSERT INTO shaking (
		with_temp,
		temperature,
		follow_temp,
		rpm_1,
		rpm_2,
		time_1,
		time_2,
		process_id)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id`

	updateShakingQuery = `UPDATE shaking SET (
		with_temp,
		temperature,
		follow_temp,
		rpm_1,
		rpm_2,
		time_1,
		time_2,
		updated_at) = 
		($1, $2, $3, $4, $5, $6, $7, $8) WHERE process_id = $9`
)

func (s *pgStore) ShowShaking(ctx context.Context, shakerID uuid.UUID) (shaker Shaker, err error) {

	err = s.db.Get(&shaker,
		getShakerQuery,
		shakerID,
	)
	if err != nil {
		logger.WithField("err", err.Error()).Errorln(responses.ShakingDBFetchError)
		return
	}

	return
}

func (s *pgStore) CreateShaking(ctx context.Context, sh Shaker) (createdShaking Shaker, err error) {
	var tx *sql.Tx

	//update the process name before record creation
	err = s.UpdateProcessName(ctx, sh.ProcessID, "Shaking", sh)
	if err != nil {
		logger.WithField("err", err.Error()).Errorln(responses.ShakingUpdateNameError)
		return
	}

	tx, err = s.db.BeginTx(ctx, nil)
	if err != nil {
		logger.WithField("err:", err.Error()).Errorln(responses.ShakingInitiateDBTxError)
		return Shaker{}, err
	}

	// End the transaction in defer call
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		tx.Commit()
	}()

	createdShaking, err = s.createShaking(ctx, sh, tx)
	// failures are already logged
	return
}

func (s *pgStore) createShaking(ctx context.Context, sh Shaker, tx *sql.Tx) (createdShaking Shaker, err error) {

	var lastInsertID uuid.UUID

	err = tx.QueryRow(
		createShakingQuery,
		sh.WithTemp,
		sh.Temperature,
		sh.FollowTemp,
		sh.RPM1,
		sh.RPM2,
		sh.Time1,
		sh.Time2,
		sh.ProcessID,
	).Scan(&lastInsertID)

	if err != nil {
		logger.WithField("err", err.Error()).Errorln(responses.ShakingDBCreateError)
		return
	}

	createdShaking, err = s.ShowShaking(ctx, lastInsertID)
	// failures are already logged
	return
}

func (s *pgStore) UpdateShaking(ctx context.Context, sh Shaker) (err error) {
	_, err = s.db.Exec(
		updateShakingQuery,
		sh.WithTemp,
		sh.Temperature,
		sh.FollowTemp,
		sh.RPM1,
		sh.RPM2,
		sh.Time1,
		sh.Time2,
		time.Now(),
		sh.ProcessID,
	)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error updating shaking")
		return
	}
	return
}
