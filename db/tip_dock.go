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
	ProcessID uuid.UUID `json:"process_id" db:"process_id" validate:"required"`
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

func (s *pgStore) CreateTipDocking(ctx context.Context, td TipDock) (createdTD TipDock, err error) {
	var tx *sql.Tx

	//update the process name before record creation
	err = s.updateProcessName(ctx, td.ProcessID, "TipDocking", td)
	if err != nil {
		logger.WithField("err", err.Error()).Errorln(responses.TipDockingUpdateNameError)
		return
	}

	tx, err = s.db.BeginTx(ctx, nil)
	if err != nil {
		logger.WithField("err:", err.Error()).Errorln(responses.TipDockingInitiateDBTxError)
		return TipDock{}, err
	}

	createdTD, err = s.createTipDocking(ctx, tx, td)
	// failures are already logged
	// Commit the transaction else won't be able to Show

	// End the transaction in defer call
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		tx.Commit()
		createdTD, err = s.ShowTipDocking(ctx, createdTD.ProcessID)
		if err != nil {
			logger.Infoln("Error Creating Tip Docking process")
			return
		}
		logger.Infoln("Created Tip Docking Process: ", createdTD)
		return
	}()

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
