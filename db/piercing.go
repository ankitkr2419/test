package db

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
	logger "github.com/sirupsen/logrus"
)

type Discard string

const (
	at_pickup_passing Discard = "at_pickup_passing"
	at_discard_box    Discard = "at_discard_box"
)

const (
	getPiercingQuery = `SELECT id,
						cartridge_ids,
						discard,
						created_at,
						updated_at
						FROM piercing
						WHERE id = $1`
	selectPiercingQuery = `SELECT *
						FROM piercing`
	deletePiercingQuery = `DELETE FROM piercing
						WHERE id = $1`
	createPiercingQuery = `INSERT INTO piercing (
						id,
						cartridge_ids,
						discard)
						VALUES ($1, $2, $3) RETURNING id`
	updatePiercingQuery = `UPDATE piercing SET (
						cartridge_ids,
						discard,
						updated_at) = ($1, $2, $3) WHERE id = $4`
)

type Piercing struct {
	ID           uuid.UUID     `db:"id" json:"id"`
	CartridgeIDs pq.Int64Array `db:"cartridge_ids" json:"cartridge_ids"`
	Discard      Discard       `db:"discard" json:"discard"`
	CreatedAt    time.Time     `db:"created_at" json:"created_at"`
	UpdatedAt    time.Time     `db:"updated_at" json:"updated_at"`
}

func (s *pgStore) ShowPiercing(ctx context.Context, id uuid.UUID) (dbPiercing Piercing, err error) {
	err = s.db.Get(&dbPiercing, getPiercingQuery, id)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error fetching piercing")
		return
	}
	return
}

func (s *pgStore) ListPiercing(ctx context.Context) (dbPiercing []Piercing, err error) {
	err = s.db.Select(&dbPiercing, selectPiercingQuery)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error fetching piercing")
		return
	}
	return
}

func (s *pgStore) CreatePiercing(ctx context.Context, p Piercing) (createdPiercing Piercing, err error) {
	var lastInsertID uuid.UUID
	err = s.db.QueryRow(
		createPiercingQuery,
		p.ID,
		p.CartridgeIDs,
		p.Discard,
	).Scan(&lastInsertID)

	if err != nil {
		logger.WithField("err", err.Error()).Error("Error creating Piercing")
		return
	}

	err = s.db.Get(&createdPiercing, getPiercingQuery, lastInsertID)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error in getting Piercing")
		return
	}
	return
}

func (s *pgStore) DeletePiercing(ctx context.Context, id uuid.UUID) (err error) {
	_, err = s.db.Exec(deletePiercingQuery, id)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error deleting piercing")
		return
	}
	return
}

func (s *pgStore) UpdatePiercing(ctx context.Context, p Piercing) (err error) {
	_, err = s.db.Exec(
		updatePiercingQuery,
		p.CartridgeIDs,
		p.Discard,
		time.Now(),
		p.ID,
	)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error updating piercing")
		return
	}
	return
}
