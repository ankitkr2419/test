package db

import (
	"context"
	"database/sql"
	"time"

	"mylab/cpagent/responses"

	"github.com/google/uuid"
	logger "github.com/sirupsen/logrus"
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
	ProcessID uuid.UUID `db:"process_id" json:"process_id"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

func (s *pgStore) ShowTipOperation(ctx context.Context, id uuid.UUID) (dbTipOperation TipOperation, err error) {
	go s.AddAuditLog(ctx, DBOperation, InitialisedState, ShowOperation, "", responses.TipOperationInitialisedState)

	err = s.db.Get(&dbTipOperation, getTipOperationQuery, id)
	defer func() {
		if err != nil {
			go s.AddAuditLog(ctx, DBOperation, ErrorState, ShowOperation, "", err.Error())
		} else {
			go s.AddAuditLog(ctx, DBOperation, CompletedState, ShowOperation, "", responses.TipOperationCompletedState)
		}
	}()
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
	go s.AddAuditLog(ctx, DBOperation, InitialisedState, CreateOperation, "", responses.TipOperationInitialisedState)

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
			go s.AddAuditLog(ctx, DBOperation, ErrorState, CreateOperation, "", err.Error())
			return
		}
		tx.Commit()
		createdAD, err = s.ShowTipOperation(ctx, createdAD.ProcessID)
		if err != nil {
			logger.Errorln(responses.TipOperationFetchError)
			return
		}
		logger.Infoln(responses.TipOperationCreateSuccess, createdAD)
		go s.AddAuditLog(ctx, DBOperation, CompletedState, CreateOperation, "", responses.TipOperationCompletedState)
		return
	}()

	// Get highest sequence number
	// NOTE: failure already logged in internal calls

	highestSeqNum, err := s.getProcessCount(ctx, tx, recipeID)
	if err != nil {
		return
	}

	process, err := s.processOperation(ctx, name, TipOperationProcess, ad, Process{})
	if err != nil {
		return
	}
	// process has only a valid name
	process.SequenceNumber = highestSeqNum + 1
	process.Type = string(TipOperationProcess)
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
	go s.AddAuditLog(ctx, DBOperation, InitialisedState, UpdateOperation, "", responses.TipOperationInitialisedState)

	_, err = s.db.Exec(
		updateTipOperationQuery,
		t.Type,
		t.Position,
		time.Now(),
		t.ProcessID,
	)
	defer func() {
		if err != nil {
			go s.AddAuditLog(ctx, DBOperation, ErrorState, UpdateOperation, "", err.Error())
		} else {
			go s.AddAuditLog(ctx, DBOperation, CompletedState, UpdateOperation, "", responses.TipOperationCompletedState)
		}
	}()

	if err != nil {
		logger.WithField("err", err.Error()).Error("Error updating tip operation")
		return
	}
	return
}
