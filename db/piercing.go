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
	ProcessID      uuid.UUID     `db:"process_id" json:"process_id"`
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

func (s *pgStore) CreatePiercing(ctx context.Context, ad Piercing, recipeID uuid.UUID) (createdAD Piercing, err error) {
	var tx *sql.Tx
	tx, err = s.db.BeginTx(ctx, nil)
	if err != nil {
		logger.WithField("err:", err.Error()).Errorln(responses.PiercingInitiateDBTxError)
		return
	}

	defer func() {
		if err != nil {
			tx.Rollback()
			logger.Errorln(responses.PiercingCreateError)
			return
		}
		tx.Commit()
		createdAD, err = s.ShowPiercing(ctx, createdAD.ProcessID)
		if err != nil {
			logger.Errorln(responses.PiercingFetchError)
			return
		}
		logger.Infoln(responses.PiercingCreateSuccess, createdAD)
		return
	}()

	// Get highest sequence number
	// NOTE: failure already logged in internal calls

	highestSeqNum, err := s.getProcessCount(ctx, tx, recipeID)
	if err != nil {
		return
	}
	
	process, err := s.processOperation(ctx, name, PiercingProcess, ad, Process{})
	if err != nil {
		return
	}
	// process has only a valid name
	process.SequenceNumber = highestSeqNum + 1
	process.Type = string(PiercingProcess)
	process.RecipeID = recipeID

	// create the process
	process, err = s.createProcess(ctx, tx, process)
	if err != nil {
		return
	}

	ad.ProcessID = process.ID
	createdAD, err = s.createPiercing(ctx, tx, ad)
	return
}

func (s *pgStore) createPiercing(ctx context.Context, tx *sql.Tx, pi Piercing) (createdPiercing Piercing, err error) {

	var lastInsertID uuid.UUID

	err = tx.QueryRow(
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
