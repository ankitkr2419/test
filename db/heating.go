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
func (s *pgStore) CreateHeating(ctx context.Context, ad Heating, recipeID uuid.UUID) (createdAD Heating, err error) {
	var tx *sql.Tx
	tx, err = s.db.BeginTx(ctx, nil)
	if err != nil {
		logger.WithField("err:", err.Error()).Errorln(responses.HeatingInitiateDBTxError)
		return
	}

	defer func() {
		if err != nil {
			tx.Rollback()
			logger.Errorln(responses.HeatingCreateError)
			return
		}
		tx.Commit()
		createdAD, err = s.ShowHeating(ctx, createdAD.ProcessID)
		if err != nil {
			logger.Errorln(responses.HeatingFetchError)
			return
		}
		logger.Infoln(responses.HeatingCreateSuccess, createdAD)
		return
	}()

	// Get highest sequence number
	// NOTE: failure already logged in internal calls

	highestSeqNum, err := s.getProcessCount(ctx, tx, recipeID)
	if err != nil {
		return
	}
	
	process, err := s.processOperation(ctx, name, HeatingProcess, ad, Process{})
	if err != nil {
		return
	}
	// process has only a valid name
	process.SequenceNumber = highestSeqNum + 1
	process.Type = string(HeatingProcess)
	process.RecipeID = recipeID

	// create the process
	process, err = s.createProcess(ctx, tx, process)
	if err != nil {
		return
	}

	ad.ProcessID = process.ID
	createdAD, err = s.createHeating(ctx, tx, ad)
	return
}

func (s *pgStore) createHeating(ctx context.Context, tx *sql.Tx, h Heating) (createdH Heating, err error) {

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

	h.ID = lastInsertID
	return h, err
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
