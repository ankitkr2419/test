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
	getPiercingQuery = `SELECT *
						FROM piercing
						WHERE process_id = $1`
	selectPiercingQuery = `SELECT *
						FROM piercing`
	deletePiercingQuery = `DELETE FROM piercing
						WHERE process_id = $1`
	createPiercingQuery = `INSERT INTO piercing (
						type,
						cartridge_wells,
						discard,
						process_id)
						VALUES ($1, $2, $3, $4) RETURNING id`
	updatePiercingQuery = `UPDATE piercing SET (
						type,
						cartridge_wells,
						discard,
						updated_at) = ($1, $2, $3, $4) WHERE process_id = $5`
)

type Piercing struct {
	ID             uuid.UUID     `db:"id" json:"id"`
	Type           CartridgeType `db:"type" json:"type"`
	CartridgeWells pq.Int64Array `db:"cartridge_wells" json:"cartridge_wells"`
	Discard        Discard       `db:"discard" json:"discard"`
	ProcessID      uuid.UUID     `db:"process_id" json:"process_id"`
	CreatedAt      time.Time     `db:"created_at" json:"created_at"`
	UpdatedAt      time.Time     `db:"updated_at" json:"updated_at"`
}

func (s *pgStore) ShowPiercing(ctx context.Context, processID uuid.UUID) (dbPiercing Piercing, err error) {
	err = s.db.Get(&dbPiercing, getPiercingQuery, processID)
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
		p.Type,
		p.CartridgeWells,
		p.Discard,
		p.ProcessID,
	).Scan(&lastInsertID)

	if err != nil {
		logger.WithField("err", err.Error()).Error("Error creating Piercing")
		return
	}

	err = s.db.Get(&createdPiercing, getPiercingQuery, p.ProcessID)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error in getting Piercing")
		return
	}
	return
}

func (s *pgStore) UpdatePiercing(ctx context.Context, p Piercing) (err error) {
	_, err = s.db.Exec(
		updatePiercingQuery,
		p.Type,
		p.CartridgeWells,
		p.Discard,
		time.Now(),
		p.ProcessID,
	)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error updating piercing")
		return
	}
	return
}
