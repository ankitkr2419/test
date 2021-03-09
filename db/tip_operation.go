package db

import (
	"context"
	"time"

	"github.com/google/uuid"
	logger "github.com/sirupsen/logrus"
)

type TipOps string

const (
	PickupTip  TipOps = "pickup"
	DiscardTip TipOps = "discard"
)

const (
	getTipOperationQuery = `SELECT id,
						type,
						position,
						process_id,
						created_at,
						updated_at
						FROM tip_operation
						WHERE process_id = $1`
	selectTipOperationQuery = `SELECT *
						FROM tip_operation`
	deleteTipOperationQuery = `DELETE FROM tip_operation
						WHERE process_id = $1`
	createTipOperationQuery = `INSERT INTO tip_operation (
						type,
						position,
						process_id,)
						VALUES ($1, $2, $3) RETURNING id`
	updateTipOperationQuery = `UPDATE tip_operation SET (
						type,
						position,
						updated_at) = ($1, $2, $3) WHERE WHERE process_id = $4`
)

type TipOperation struct {
	ID        uuid.UUID `db:"id" json:"id"`
	Type      TipOps    `db:"type" json:"type"`
	Position  int64     `db:"position" json:"position"`
	ProcessID uuid.UUID `db:"process_id" json:"process_id"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

func (s *pgStore) ShowTipOperation(ctx context.Context, id uuid.UUID) (dbTipOperation TipOperation, err error) {
	err = s.db.Get(&dbTipOperation, getTipOperationQuery, id)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error fetching tip operation")
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

func (s *pgStore) CreateTipOperation(ctx context.Context, t TipOperation) (createdTipOperation TipOperation, err error) {
	var lastInsertID uuid.UUID
	err = s.db.QueryRow(
		createTipOperationQuery,
		t.Type,
		t.Position,
		t.ProcessID,
	).Scan(&lastInsertID)

	if err != nil {
		logger.WithField("err", err.Error()).Error("Error creating TipOperation")
		return
	}

	err = s.db.Get(&createdTipOperation, getTipOperationQuery, lastInsertID)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error in getting TipOperation")
		return
	}
	return
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