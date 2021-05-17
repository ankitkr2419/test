package db

import (
	"context"
	"database/sql"
	"time"

	"mylab/cpagent/responses"

	"github.com/google/uuid"
	logger "github.com/sirupsen/logrus"
)

type AttachDetach struct {
	ID            uuid.UUID `db:"id" json:"id"`
	Operation     string    `db:"operation" json:"operation"  validate:"required"`
	OperationType string    `db:"operation_type" json:"operation_type"`
	ProcessID     uuid.UUID `db:"process_id" json:"process_id"`
	CreatedAt     time.Time `db:"created_at" json:"created_at"`
	UpdatedAt     time.Time `db:"updated_at" json:"updated_at"`
}

const (
	getAttachDetachQuery    = `SELECT * FROM attach_detach where process_id = $1`
	createAttachDetachQuery = `INSERT INTO attach_detach (
		operation,
		operation_type,
		process_id)
		VALUES ($1, $2, $3) RETURNING id`
	updateAttachDetachQuery = `UPDATE attach_detach SET (
			operation,
			operation_type,
			updated_at) = 
			($1, $2, $3) WHERE process_id = $4`
)

func (s *pgStore) ShowAttachDetach(ctx context.Context, processID uuid.UUID) (ad AttachDetach, err error) {
	go s.AddAuditLog(ctx, DBOperation, InitialisedState, ShowOperation, "", responses.AttachDetachInitialisedState)

	err = s.db.Get(&ad, getAttachDetachQuery, processID)
	defer func() {
		if err != nil {
			go s.AddAuditLog(ctx, DBOperation, InitialisedState, ShowOperation, "", err.Error())
		} else {
			go s.AddAuditLog(ctx, DBOperation, InitialisedState, ShowOperation, "", responses.AttachDetachCompletedState)
		}
	}()
	if err != nil {
		logger.WithField("err", err.Error()).Errorln(responses.AttachDetachDBFetchError)
		return
	}
	return
}

func (s *pgStore) CreateAttachDetach(ctx context.Context, ad AttachDetach, recipeID uuid.UUID) (createdAD AttachDetach, err error) {
	go s.AddAuditLog(ctx, DBOperation, InitialisedState, CreateOperation, "", responses.AttachDetachInitialisedState)

	var tx *sql.Tx
	tx, err = s.db.BeginTx(ctx, nil)
	if err != nil {
		logger.WithField("err:", err.Error()).Errorln(responses.AttachDetachInitiateDBTxError)
		return
	}

	defer func() {
		if err != nil {
			tx.Rollback()
			logger.Errorln(responses.AttachDetachCreateError)
			go s.AddAuditLog(ctx, DBOperation, InitialisedState, CreateOperation, "", err.Error())
			return
		}
		tx.Commit()
		createdAD, err = s.ShowAttachDetach(ctx, createdAD.ProcessID)
		if err != nil {
			logger.Errorln(responses.AttachDetachFetchError)
			return
		}
		logger.Infoln(responses.AttachDetachCreateSuccess, createdAD)
		go s.AddAuditLog(ctx, DBOperation, InitialisedState, CreateOperation, "", responses.AttachDetachCompletedState)
		return
	}()

	// Get highest sequence number
	// NOTE: failure already logged in internal calls

	highestSeqNum, err := s.getProcessCount(ctx, tx, recipeID)
	if err != nil {
		return
	}

	process, err := s.processOperation(ctx, name, AttachDetachProcess, ad, Process{})
	if err != nil {
		return
	}
	// process has only a valid name
	process.SequenceNumber = highestSeqNum + 1
	process.Type = string(AttachDetachProcess)
	process.RecipeID = recipeID

	// create the process
	process, err = s.createProcess(ctx, tx, process)
	if err != nil {
		return
	}

	ad.ProcessID = process.ID
	createdAD, err = s.createAttachDetach(ctx, tx, ad)
	return
}

func (s *pgStore) createAttachDetach(ctx context.Context, tx *sql.Tx, ad AttachDetach) (createdAD AttachDetach, err error) {

	var lastInsertID uuid.UUID

	err = tx.QueryRow(
		createAttachDetachQuery,
		ad.Operation,
		ad.OperationType,
		ad.ProcessID,
	).Scan(&lastInsertID)

	if err != nil {
		logger.WithField("err", err.Error()).Errorln(responses.AttachDetachDBCreateError)
		return
	}

	ad.ID = lastInsertID
	return ad, err
}

func (s *pgStore) UpdateAttachDetach(ctx context.Context, a AttachDetach) (err error) {
	go s.AddAuditLog(ctx, DBOperation, InitialisedState, UpdateOperation, "", responses.AttachDetachInitialisedState)

	_, err = s.db.Exec(
		updateAttachDetachQuery,
		a.Operation,
		a.OperationType,
		time.Now(),
		a.ProcessID,
	)
	defer func() {
		if err != nil {
			go s.AddAuditLog(ctx, DBOperation, InitialisedState, UpdateOperation, "", err.Error())
		} else {
			go s.AddAuditLog(ctx, DBOperation, InitialisedState, UpdateOperation, "", responses.AttachDetachCompletedState)
		}
	}()
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error updating attach detach")
		return
	}
	return
}
