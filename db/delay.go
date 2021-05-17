package db

import (
	"context"
	"database/sql"
	"time"

	"mylab/cpagent/responses"

	"github.com/google/uuid"
	logger "github.com/sirupsen/logrus"
)

type Delay struct {
	ID        uuid.UUID `db:"id" json:"id"`
	DelayTime int64     `db:"delay_time" json:"delay_time" validate:"required"`
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
			updated_at) = ($1, $2) WHERE process_id = $3`
)

func (s *pgStore) ShowDelay(ctx context.Context, id uuid.UUID) (delay Delay, err error) {
	go s.AddAuditLog(ctx, DBOperation, InitialisedState, ShowOperation, "", responses.DelayInitialisedState)

	// get delay record
	err = s.db.Get(&delay, getDelayQuery, id)
	defer func() {
		if err != nil {
			go s.AddAuditLog(ctx, DBOperation, InitialisedState, ShowOperation, "", err.Error())
		} else {
			go s.AddAuditLog(ctx, DBOperation, InitialisedState, ShowOperation, "", responses.DelayCompletedState)
		}
	}()
	if err != nil {
		logger.WithField("err", err.Error()).Errorln(responses.DelayDBFetchError)
		return
	}
	return
}

func (s *pgStore) CreateDelay(ctx context.Context, ad Delay, recipeID uuid.UUID) (createdAD Delay, err error) {
	go s.AddAuditLog(ctx, DBOperation, InitialisedState, CreateOperation, "", responses.DelayInitialisedState)

	var tx *sql.Tx
	tx, err = s.db.BeginTx(ctx, nil)
	if err != nil {
		logger.WithField("err:", err.Error()).Errorln(responses.DelayInitiateDBTxError)
		return
	}

	defer func() {
		if err != nil {
			tx.Rollback()
			logger.Errorln(responses.DelayCreateError)
			go s.AddAuditLog(ctx, DBOperation, InitialisedState, CreateOperation, "", err.Error())
			return
		}
		tx.Commit()
		createdAD, err = s.ShowDelay(ctx, createdAD.ProcessID)
		if err != nil {
			logger.Errorln(responses.DelayFetchError)
			return
		}
		logger.Infoln(responses.DelayCreateSuccess, createdAD)
		go s.AddAuditLog(ctx, DBOperation, InitialisedState, CreateOperation, "", responses.DelayCompletedState)
		return
	}()

	// Get highest sequence number
	// NOTE: failure already logged in internal calls

	highestSeqNum, err := s.getProcessCount(ctx, tx, recipeID)
	if err != nil {
		return
	}

	process, err := s.processOperation(ctx, name, DelayProcess, ad, Process{})
	if err != nil {
		return
	}
	// process has only a valid name
	process.SequenceNumber = highestSeqNum + 1
	process.Type = string(DelayProcess)
	process.RecipeID = recipeID

	// create the process
	process, err = s.createProcess(ctx, tx, process)
	if err != nil {
		return
	}

	ad.ProcessID = process.ID
	createdAD, err = s.createDelay(ctx, tx, ad)
	return
}

func (s *pgStore) createDelay(ctx context.Context, tx *sql.Tx, d Delay) (createdD Delay, err error) {
	go s.AddAuditLog(ctx, DBOperation, InitialisedState, UpdateOperation, "", responses.DelayInitialisedState)

	var lastInsertID uuid.UUID

	err = tx.QueryRow(
		createDelayQuery,
		d.DelayTime,
		d.ProcessID,
	).Scan(&lastInsertID)

	defer func() {
		if err != nil {
			go s.AddAuditLog(ctx, DBOperation, InitialisedState, UpdateOperation, "", err.Error())
		} else {
			go s.AddAuditLog(ctx, DBOperation, InitialisedState, UpdateOperation, "", responses.DelayCompletedState)
		}
	}()

	if err != nil {
		logger.WithField("err", err.Error()).Errorln(responses.DelayDBCreateError)
		return
	}

	d.ID = lastInsertID
	return d, err
}

func (s *pgStore) UpdateDelay(ctx context.Context, d Delay) (err error) {
	_, err = s.db.Exec(
		updateDelayQuery,
		d.DelayTime,
		time.Now(),
		d.ProcessID,
	)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error updating delay")
		return
	}
	return
}
