package db

import (
	"context"
	"database/sql"
	"time"

	"mylab/cpagent/responses"

	"github.com/google/uuid"
	logger "github.com/sirupsen/logrus"
)

type TipDock struct {
	ID       uuid.UUID `json:"id" db:"id"`
	Type     string    `json:"type" db:"type" validate:"required"`
	Position int64     `json:"position" db:"position" validate:"required,lte=13"`
	// since we are not considering the height of labware so we keep the maximum height as 25mm.
	Height    float64   `json:"height" db:"height" validate:"required,lte=25"`
	ProcessID uuid.UUID `json:"process_id" db:"process_id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

const (
	getTipDockQuery    = `SELECT * FROM tip_docking where process_id = $1`
	createTipDockQuery = `INSERT INTO tip_docking (
		type,
		position,
		height,
		process_id)
		VALUES ($1, $2, $3, $4) RETURNING id`
	updateTipDockQuery = `UPDATE tip_docking SET (
			type,
			position,
			height,
			updated_at) = 
			($1, $2, $3,$4) WHERE process_id = $5`
)

func (s *pgStore) ShowTipDocking(ctx context.Context, pid uuid.UUID) (td TipDock, err error) {
	go s.AddAuditLog(ctx, DBOperation, InitialisedState, ShowOperation, "", responses.TipDockingInitialisedState)

	err = s.db.Get(&td, getTipDockQuery, pid)
	defer func() {
		if err != nil {
			go s.AddAuditLog(ctx, DBOperation, ErrorState, ShowOperation, "", err.Error())
		} else {
			go s.AddAuditLog(ctx, DBOperation, CompletedState, ShowOperation, "", responses.TipDockingCompletedState)
		}
	}()
	if err != nil {
		logger.WithField("err", err.Error()).Errorln(responses.TipDockingDBFetchError)
		return
	}
	return
}

func (s *pgStore) CreateTipDocking(ctx context.Context, ad TipDock, recipeID uuid.UUID) (createdAD TipDock, err error) {
	go s.AddAuditLog(ctx, DBOperation, InitialisedState, CreateOperation, "", responses.TipDockingInitialisedState)

	var tx *sql.Tx
	tx, err = s.db.BeginTx(ctx, nil)
	if err != nil {
		logger.WithField("err:", err.Error()).Errorln(responses.TipDockingInitiateDBTxError)
		return
	}

	defer func() {
		if err != nil {
			tx.Rollback()
			logger.Errorln(responses.TipDockingCreateError)
			go s.AddAuditLog(ctx, DBOperation, ErrorState, CreateOperation, "", err.Error())
			return
		}
		tx.Commit()
		createdAD, err = s.ShowTipDocking(ctx, createdAD.ProcessID)
		if err != nil {
			logger.Errorln(responses.TipDockingFetchError)
			return
		}
		logger.Infoln(responses.TipDockingCreateSuccess, createdAD)
		go s.AddAuditLog(ctx, DBOperation, CompletedState, CreateOperation, "", responses.TipDockingCompletedState)
		return
	}()

	// Get highest sequence number
	// NOTE: failure already logged in internal calls

	highestSeqNum, err := s.getProcessCount(ctx, tx, recipeID)
	if err != nil {
		return
	}

	process, err := s.processOperation(ctx, name, TipDockingProcess, ad, Process{})
	if err != nil {
		return
	}
	// process has only a valid name
	process.SequenceNumber = highestSeqNum + 1
	process.Type = TipDockingProcess
	process.RecipeID = recipeID

	// create the process
	process, err = s.createProcess(ctx, tx, process)
	if err != nil {
		return
	}

	ad.ProcessID = process.ID
	createdAD, err = s.createTipDocking(ctx, tx, ad)
	return
}

func (s *pgStore) createTipDocking(ctx context.Context, tx *sql.Tx, t TipDock) (createdTD TipDock, err error) {

	var lastInsertID uuid.UUID

	err = tx.QueryRow(
		createTipDockQuery,
		t.Type,
		t.Position,
		t.Height,
		t.ProcessID,
	).Scan(&lastInsertID)

	if err != nil {
		logger.WithField("err", err.Error()).Errorln(responses.TipDockingDBCreateError)
		return
	}

	t.ID = lastInsertID
	return t, err
}

func (s *pgStore) UpdateTipDock(ctx context.Context, t TipDock) (err error) {
	go s.AddAuditLog(ctx, DBOperation, InitialisedState, UpdateOperation, "", responses.TipDockingInitialisedState)
	var tx *sql.Tx
	tx, err = s.db.BeginTx(ctx, nil)
	if err != nil {
		logger.WithField("err:", err.Error()).Errorln(responses.TipDockingInitiateDBTxError)
		return
	}

	defer func() {
		if err != nil {
			tx.Rollback()
			logger.Errorln(responses.TipDockingUpdateError)
			go s.AddAuditLog(ctx, DBOperation, ErrorState, UpdateOperation, "", err.Error())
			return
		}
		tx.Commit()

		logger.Infoln(responses.TipDockingUpdateSuccess)
		go s.AddAuditLog(ctx, DBOperation, CompletedState, UpdateOperation, "", responses.TipDockingCompletedState)
		return
	}()
	err = s.updateProcessName(ctx, tx, t.ProcessID, TipDockingProcess, t)
	if err != nil {
		logger.WithField("err:", err.Error()).Errorln(responses.AspireDispenseUpdateNameError)
		return
	}

	result, err := tx.Exec(
		updateTipDockQuery,
		t.Type,
		t.Position,
		t.Height,
		time.Now(),
		t.ProcessID,
	)
	defer func() {
		if err != nil {
			go s.AddAuditLog(ctx, DBOperation, ErrorState, UpdateOperation, "", err.Error())
		} else {
			go s.AddAuditLog(ctx, DBOperation, CompletedState, UpdateOperation, "", responses.TipDockingCompletedState)
		}
	}()
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error updating TipDocking")
		return
	}

	c, _ := result.RowsAffected()
	// check row count as no error is returned when row not found for update
	if c == 0 {
		return responses.ProcessIDInvalidError
	}

	return
}
