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
type Discard string

const (
	PickupTip  TipOps = "pickup"
	DiscardTip TipOps = "discard"

	at_pickup_passing Discard = "at_pickup_passing"
	at_discard_box    Discard = "at_discard_box"
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
						discard,
						process_id)
						VALUES ($1, $2, $3, $4) RETURNING id`
	updateTipOperationQuery = `UPDATE tip_operation SET (
						type,
						position,
						discard,
						updated_at) = ($1, $2, $3, $4) WHERE process_id = $5`
)

type TipOperation struct {
	ID       uuid.UUID `db:"id" json:"id"`
	Type      TipOps    `db:"type" json:"type" validate:"required"`
	Position  int64     `db:"position" json:"position,omitempty"`
	Discard   Discard   `db:"discard" json:"discard,omitempty"`
	ProcessID uuid.UUID `db:"process_id" json:"process_id"`
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

func (s *pgStore) CreateTipOperation(ctx context.Context, ad TipOperation, recipeID uuid.UUID) (createdAD TipOperation, err error) {
	var tx *sql.Tx
	tx, err = s.db.BeginTx(ctx, nil)
	if err != nil {
		logger.WithField("err:", err.Error()).Errorln(responses.TipOperationInitiateDBTxError)
		return
	}

	defer func() {
		if err != nil {
			tx.Rollback()
			logger.Errorln(responses.TipOperationCreateError)
			return
		}
		tx.Commit()
		createdAD, err = s.ShowTipOperation(ctx, createdAD.ProcessID)
		if err != nil {
			logger.Errorln(responses.TipOperationFetchError)
			return
		}
		logger.Infoln(responses.TipOperationCreateSuccess, createdAD)
		return
	}()

	// Get highest sequence number
	// NOTE: failure already logged in internal calls

	highestSeqNum, err := s.getProcessCount(ctx, tx, recipeID)
	if err != nil {
		return
	}
	
	var opsType ProcessType
	if ad.Type == PickupTip{
		opsType = TipPickupProcess
	} else{
		opsType = TipDiscardProcess
	}

	process, err := s.processOperation(ctx, name, opsType , ad, Process{})
	if err != nil {
		return
	}
	// process has only a valid name
	process.SequenceNumber = highestSeqNum + 1
	process.Type = opsType
	process.RecipeID = recipeID

	// create the process
	process, err = s.createProcess(ctx, tx, process)
	if err != nil {
		return
	}

	ad.ProcessID = process.ID
	createdAD, err = s.createTipOperation(ctx, tx, ad)
	return
}

func (s *pgStore) createTipOperation(ctx context.Context, tx *sql.Tx, to TipOperation) (createdTipOperation TipOperation, err error) {

	var lastInsertID uuid.UUID

	err = tx.QueryRow(
		createTipOperationQuery,
		to.Type,
		to.Position,
		to.Discard,
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
		t.Discard,
		time.Now(),
		t.ProcessID,
	)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error updating tip operation")
		return
	}
	return
}
