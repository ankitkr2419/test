package db

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
	logger "github.com/sirupsen/logrus"
	"mylab/cpagent/responses"
)

type TipOps string

const (
	PickupTip  TipOps = "pickup"
	DiscardTip TipOps = "discard"
)

const (
	getTipOperationQuery = `SELECT *
						FROM tip_operation
						WHERE process_id = $1`
	selectTipOperationQuery = `SELECT *
						FROM tip_operation`
	deleteTipOperationQuery = `DELETE FROM tip_operation
						WHERE process_id = $1`
	createTipOperationQuery = `INSERT INTO tip_operation (
						type,
						position,
						process_id)
						VALUES ($1, $2, $3) RETURNING id`
	updateTipOperationQuery = `UPDATE tip_operation SET (
						type,
						position,
						updated_at) = ($1, $2, $3) WHERE process_id = $4`
)

type TipOperation struct {
	ID        uuid.UUID `db:"id" json:"id"`
	Type      TipOps    `db:"type" json:"type" validate:"required"`
	Position  int64     `db:"position" json:"position"`
	ProcessID uuid.UUID `db:"process_id" json:"process_id" validate:"required"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

func (s *pgStore) ShowTipOperation(ctx context.Context, id uuid.UUID) (dbTipOperation TipOperation, err error) {
	err = s.db.Get(&dbTipOperation, getTipOperationQuery, id)
	if err != nil {
		logger.WithField("err", err.Error()).Errorln(responses.TipOperationDBFetchError)
		return
	}
	return
}

func (s *pgStore) ListTipOperation(ctx context.Context) (dbTipOperation []TipOperation, err error) {
	err = s.db.Select(&dbTipOperation, selectTipOperationQuery)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error fetching tip operation")
		return
	}
	return
}

func (s *pgStore) CreateTipOperation(ctx context.Context, to TipOperation) (createdTO TipOperation, err error) {
	var tx *sql.Tx

	//update the process name before record creation
	err = s.UpdateProcessName(ctx, to.ProcessID, "TipOperation", to)
	if err != nil {
		logger.WithField("err", err.Error()).Errorln(responses.TipOperationUpdateNameError)
		return
	}

	tx, err = s.db.BeginTx(ctx, nil)
	if err != nil {
		logger.WithField("err:", err.Error()).Errorln(responses.TipOperationInitiateDBTxError)
		return TipOperation{}, err
	}

	createdTO, err = s.createTipOperation(ctx, to, tx)
	// failures are already logged
	// Commit the transaction else won't be able to Show

	// End the transaction in defer call
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		tx.Commit()
		createdTO, err = s.ShowTipOperation(ctx, createdTO.ProcessID)
		if err != nil {
			logger.Infoln("Error Creating Tip Operation process")
			return
		}
		logger.Infoln("Created Tip Operation Process: ", createdTO)
		return
	}()

	return
}

func (s *pgStore) createTipOperation(ctx context.Context, to TipOperation, tx *sql.Tx) (createdTipOperation TipOperation, err error) {

	var lastInsertID uuid.UUID

	err = tx.QueryRow(
		createTipOperationQuery,
		to.Type,
		to.Position,
		to.ProcessID,
	).Scan(&lastInsertID)

	if err != nil {
		logger.WithField("err", err.Error()).Errorln(responses.TipOperationDBCreateError)
		return
	}

	to.ID = lastInsertID
	return to, err
}

func (s *pgStore) DeleteTipOperation(ctx context.Context, id uuid.UUID) (err error) {
	_, err = s.db.Exec(deleteTipOperationQuery, id)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error deleting tip operation")
		return
	}
	return
}

func (s *pgStore) UpdateTipOperation(ctx context.Context, t TipOperation) (err error) {
	_, err = s.db.Exec(
		updateTipOperationQuery,
		t.Type,
		t.Position,
		time.Now(),
		t.ProcessID,
	)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error updating tip operation")
		return
	}
	return
}
