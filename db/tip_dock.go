package db

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
	logger "github.com/sirupsen/logrus"
	"mylab/cpagent/responses"
)

type TipDock struct {
	ID        uuid.UUID `json:"id" db:"id"`
	Type      string    `json:"type" db:"type" validate:"required"`
	Position  int64     `json:"position" db:"position" validate:"required"`
	Height    float64   `json:"height" db:"height" validate:"required"`
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

	err = s.db.Get(&td, getTipDockQuery, pid)
	if err != nil {
		logger.WithField("err", err.Error()).Errorln(responses.TipDockingDBFetchError)
		return
	}
	return
}

func (s *pgStore) CreateTipDocking(ctx context.Context, ad TipDock, recipeID uuid.UUID) (createdAD TipDock, err error) {
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
			return
		}
		tx.Commit()
		createdAD, err = s.ShowTipDocking(ctx, createdAD.ProcessID)
		if err != nil {
			logger.Errorln(responses.TipDockingFetchError)
			return
		}
		logger.Infoln(responses.TipDockingCreateSuccess, createdAD)
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
	_, err = s.db.Exec(
		updateTipDockQuery,
		t.Type,
		t.Position,
		t.Height,
		time.Now(),
		t.ProcessID,
	)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error updating TipDocking")
		return
	}
	return
}
