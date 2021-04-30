package db

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
	logger "github.com/sirupsen/logrus"
	"mylab/cpagent/responses"
)

type AttachDetach struct {
	ID            uuid.UUID `db:"id" json:"id"`
	Operation     string    `db:"operation" json:"operation"  validate:"required"`
	OperationType string    `db:"operation_type" json:"operation_type"`
	ProcessID     uuid.UUID `db:"process_id" json:"process_id" validate:"required"`
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

	err = s.db.Get(&ad, getAttachDetachQuery, processID)
	if err != nil {
		logger.WithField("err", err.Error()).Errorln(responses.AttachDetachDBFetchError)
		return
	}
	return
}

func (s *pgStore) CreateAttachDetach(ctx context.Context, ad AttachDetach) (createdAttachDetach AttachDetach, err error) {
	var tx *sql.Tx

	//update the process name before record creation
	err = s.UpdateProcessName(ctx, ad.ProcessID, "AttachDetach", ad)
	if err != nil {
		logger.WithField("err", err.Error()).Errorln(responses.AttachDetachUpdateNameError)
		return
	}

	tx, err = s.db.BeginTx(ctx, nil)
	if err != nil {
		logger.WithField("err:", err.Error()).Errorln(responses.AttachDetachInitiateDBTxError)
		return AttachDetach{}, err
	}

	// End the transaction in defer call
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		tx.Commit()
	}()

	createdAttachDetach, err = s.createAttachDetach(ctx, ad, tx)
	// failures are already logged
	return
}

func (s *pgStore) createAttachDetach(ctx context.Context, ad AttachDetach, tx *sql.Tx) (createdAttachDetach AttachDetach, err error) {

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

	createdAttachDetach, err = s.ShowAttachDetach(ctx, lastInsertID)
	// failures are already logged
	return
}

func (s *pgStore) UpdateAttachDetach(ctx context.Context, a AttachDetach) (err error) {
	_, err = s.db.Exec(
		updateAttachDetachQuery,
		a.Operation,
		a.OperationType,
		time.Now(),
		a.ProcessID,
	)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error updating attach detach")
		return
	}
	return
}
