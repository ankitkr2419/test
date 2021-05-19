package db

import (
	"context"
	"time"

	"database/sql"
	"mylab/cpagent/responses"

	"github.com/google/uuid"
	logger "github.com/sirupsen/logrus"
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
	go s.AddAuditLog(ctx, DBOperation, InitialisedState, ShowOperation, "", responses.ShakingInitialisedState)

	err = s.db.Get(&shaker,
		getShakerQuery,
		shakerID,
	)
	defer func() {
		if err != nil {
			go s.AddAuditLog(ctx, DBOperation, InitialisedState, ShowOperation, "", err.Error())
		} else {
			go s.AddAuditLog(ctx, DBOperation, CompletedState, ShowOperation, "", responses.ShakingCompletedState)
		}
	}()
	if err != nil {
		logger.WithField("err", err.Error()).Errorln(responses.ShakingDBFetchError)
		return
	}

	return
}

func (s *pgStore) CreateShaking(ctx context.Context, ad Shaker, recipeID uuid.UUID) (createdAD Shaker, err error) {
	go s.AddAuditLog(ctx, DBOperation, InitialisedState, CreateOperation, "", responses.ShakingInitialisedState)

	var tx *sql.Tx
	tx, err = s.db.BeginTx(ctx, nil)
	if err != nil {
		logger.WithField("err:", err.Error()).Errorln(responses.ShakingInitiateDBTxError)
		return
	}

	defer func() {
		if err != nil {
			tx.Rollback()
			logger.Errorln(responses.ShakingCreateError)
			go s.AddAuditLog(ctx, DBOperation, ErrorState, CreateOperation, "", err.Error())
			return
		}
		tx.Commit()
		createdAD, err = s.ShowShaking(ctx, createdAD.ProcessID)
		if err != nil {
			logger.Errorln(responses.ShakingFetchError)
			return
		}
		logger.Infoln(responses.ShakingCreateSuccess, createdAD)
		go s.AddAuditLog(ctx, DBOperation, CompletedState, CreateOperation, "", responses.ShakingCompletedState)

		return
	}()

	// Get highest sequence number
	// NOTE: failure already logged in internal calls

	highestSeqNum, err := s.getProcessCount(ctx, tx, recipeID)
	if err != nil {
		return
	}

	process, err := s.processOperation(ctx, name, ShakingProcess, ad, Process{})
	if err != nil {
		return
	}
	// process has only a valid name
	process.SequenceNumber = highestSeqNum + 1
	process.Type = string(ShakingProcess)
	process.RecipeID = recipeID

	// create the process
	process, err = s.createProcess(ctx, tx, process)
	if err != nil {
		return
	}

	ad.ProcessID = process.ID
	createdAD, err = s.createShaking(ctx, tx, ad)
	return
}

func (s *pgStore) createShaking(ctx context.Context, tx *sql.Tx, sh Shaker) (createdSh Shaker, err error) {

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

	sh.ID = lastInsertID
	return sh, err
}

func (s *pgStore) UpdateShaking(ctx context.Context, sh Shaker) (err error) {
	go s.AddAuditLog(ctx, DBOperation, InitialisedState, UpdateOperation, "", responses.ShakingInitialisedState)

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
	defer func() {
		if err != nil {
			go s.AddAuditLog(ctx, DBOperation, ErrorState, UpdateOperation, "", err.Error())
		} else {
			go s.AddAuditLog(ctx, DBOperation, CompletedState, UpdateOperation, "", responses.ShakingCompletedState)
		}
	}()
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error updating shaking")
		return
	}
	return
}
