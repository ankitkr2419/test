package db

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
	logger "github.com/sirupsen/logrus"
	"mylab/cpagent/responses"
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
	Type           CartridgeType `db:"type" json:"type" validate:"required"`
	CartridgeWells pq.Int64Array `db:"cartridge_wells" json:"cartridge_wells" validate:"required"`
	Discard        Discard       `db:"discard" json:"discard"`
	ProcessID      uuid.UUID     `db:"process_id" json:"process_id" validate:"required"`
	CreatedAt      time.Time     `db:"created_at" json:"created_at"`
	UpdatedAt      time.Time     `db:"updated_at" json:"updated_at"`
}

func (s *pgStore) ShowPiercing(ctx context.Context, processID uuid.UUID) (dbPiercing Piercing, err error) {
	err = s.db.Get(&dbPiercing, getPiercingQuery, processID)
	if err != nil {
		logger.WithField("err", err.Error()).Errorln(responses.PiercingDBFetchError)
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

func (s *pgStore) CreatePiercing(ctx context.Context, pi Piercing) (createdPi Piercing, err error) {
	var tx *sql.Tx

	//update the process name before record creation
	err = s.UpdateProcessName(ctx, pi.ProcessID, "Piercing", pi)
	if err != nil {
		logger.WithField("err", err.Error()).Errorln(responses.PiercingUpdateNameError)
		return
	}

	tx, err = s.db.BeginTx(ctx, nil)
	if err != nil {
		logger.WithField("err:", err.Error()).Errorln(responses.PiercingInitiateDBTxError)
		return Piercing{}, err
	}

	createdPi, err = s.createPiercing(ctx, pi, tx)
	// failures are already logged
	// Commit the transaction else won't be able to Show

	// End the transaction in defer call
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		tx.Commit()
		createdPi, err = s.ShowPiercing(ctx, createdPi.ProcessID)
		if err != nil {
			logger.Infoln("Error Creating Piercing process")
			return
		}
		logger.Infoln("Created Piercing Process: ", createdPi)
		return
	}()

	return
}

func (s *pgStore) createPiercing(ctx context.Context, pi Piercing, tx *sql.Tx) (createdPiercing Piercing, err error) {

	var lastInsertID uuid.UUID

	err = s.db.QueryRow(
		createPiercingQuery,
		pi.Type,
		pi.CartridgeWells,
		pi.Discard,
		pi.ProcessID,
	).Scan(&lastInsertID)

	if err != nil {
		logger.WithField("err", err.Error()).Errorln(responses.PiercingDBCreateError)
		return
	}

	pi.ID = lastInsertID
	return pi, err
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
