package db

import (
	"context"
	"database/sql"
	"time"

	"mylab/cpagent/responses"

	"github.com/google/uuid"

	logger "github.com/sirupsen/logrus"
)

type Category string

const (
	WS Category = "well_to_shaker"
	SW Category = "shaker_to_well"
	WW Category = "well_to_well"
	WD Category = "well_to_deck"
	DW Category = "deck_to_well"
	DD Category = "deck_to_deck"
	SD Category = "shaker_to_deck"
	DS Category = "deck_to_shaker"
)

const (
	getAspireDispenseQuery = `SELECT *
						FROM aspire_dispense
						WHERE process_id = $1`
	selectAspireDispenseQuery = `SELECT *
						FROM aspire_dispense`
	createAspireDispenseQuery = `INSERT INTO aspire_dispense (
						category,
						cartridge_type,
						source_position,
						aspire_height,
						aspire_mixing_volume,
						aspire_no_of_cycles,
						aspire_volume,
						aspire_air_volume,
						dispense_height,
						dispense_mixing_volume,
						dispense_no_of_cycles,
						destination_position,
						process_id)
						VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13) RETURNING id`
	updateAspireDispenseQuery = `UPDATE aspire_dispense SET (
						category,
						cartridge_type,
						source_position,
						aspire_height,
						aspire_mixing_volume,
						aspire_no_of_cycles,
						aspire_volume,
						aspire_air_volume,
						dispense_height,
						dispense_mixing_volume,
						dispense_no_of_cycles,
						destination_position,
						updated_at) = 
						($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13) WHERE process_id = $14`
)

type AspireDispense struct {
	ID                   uuid.UUID     `db:"id" json:"id"`
	Category             Category      `db:"category" json:"category"`
	CartridgeType        CartridgeType `db:"cartridge_type" json:"cartridge_type"`
	SourcePosition       int64         `db:"source_position" json:"source_position" validate:"lte=13"`
	AspireHeight         float64       `db:"aspire_height" json:"aspire_height" validate:"required,lte=60"`
	AspireMixingVolume   float64       `db:"aspire_mixing_volume" json:"aspire_mixing_volume" validate:"lte=700"`
	AspireNoOfCycles     int64         `db:"aspire_no_of_cycles" json:"aspire_no_of_cycles" validate:"lte=100"`
	AspireVolume         float64       `db:"aspire_volume" json:"aspire_volume" validate:"lte=700"`
	AspireAirVolume      float64       `db:"aspire_air_volume" json:"aspire_air_volume"`
	DispenseHeight       float64       `db:"dispense_height" json:"dispense_height" validate:"required,lte=60"`
	DispenseMixingVolume float64       `db:"dispense_mixing_volume" json:"dispense_mixing_volume" validate:"lte=1000"`
	DispenseNoOfCycles   int64         `db:"dispense_no_of_cycles" json:"dispense_no_of_cycles" validate:"lte=100"`
	DestinationPosition  int64         `db:"destination_position" json:"destination_position" validate:"lte=13"`
	ProcessID            uuid.UUID     `db:"process_id" json:"process_id"`
	CreatedAt            time.Time     `db:"created_at" json:"created_at"`
	UpdatedAt            time.Time     `db:"updated_at" json:"updated_at"`
}

func (s *pgStore) ShowAspireDispense(ctx context.Context, id uuid.UUID) (dbAspireDispense AspireDispense, err error) {
	// logging initialised db operation
	go s.AddAuditLog(ctx, DBOperation, InitialisedState, ShowOperation, "", responses.AspireDispenseInitialisedState)
	err = s.db.Get(&dbAspireDispense, getAspireDispenseQuery, id)

	//logging error if there is any otherwise logging success
	defer func() {
		if err != nil {
			go s.AddAuditLog(ctx, DBOperation, ErrorState, ShowOperation, "", err.Error())
		} else {
			go s.AddAuditLog(ctx, DBOperation, CompletedState, ShowOperation, "", responses.AspireDispenseCompletedState)
		}
	}()

	if err != nil {
		logger.WithField("err", err.Error()).Errorln(responses.AspireDispenseDBFetchError)
		return
	}

	return
}

func (s *pgStore) ListAspireDispense(ctx context.Context) (dbAspireDispense []AspireDispense, err error) {

	go s.AddAuditLog(ctx, DBOperation, InitialisedState, ShowOperation, "", responses.AspireDispenseListInitialisedState)

	err = s.db.Select(&dbAspireDispense, selectAspireDispenseQuery)

	//logging error if there is any otherwise logging success
	defer func() {
		if err != nil {
			go s.AddAuditLog(ctx, DBOperation, ErrorState, ShowOperation, "", err.Error())
		} else {
			go s.AddAuditLog(ctx, DBOperation, CompletedState, ShowOperation, "", responses.AspireDispenseListCompletedState)
		}
	}()

	if err != nil {
		logger.WithField("err", err.Error()).Error("Error fetching aspire dispense")
		return
	}
	return
}

func (s *pgStore) CreateAspireDispense(ctx context.Context, ad AspireDispense, recipeID uuid.UUID) (createdAD AspireDispense, err error) {
	// logging initialised db operation
	go s.AddAuditLog(ctx, DBOperation, InitialisedState, CreateOperation, "", responses.AspireDispenseInitialisedState)

	var tx *sql.Tx
	tx, err = s.db.BeginTx(ctx, nil)
	if err != nil {
		logger.WithField("err:", err.Error()).Errorln(responses.AspireDispenseInitiateDBTxError)
		return
	}

	defer func() {
		if err != nil {
			tx.Rollback()
			logger.Errorln(responses.AspireDispenseCreateError)
			go s.AddAuditLog(ctx, DBOperation, ErrorState, CreateOperation, "", err.Error())
			return
		}
		tx.Commit()
		createdAD, err = s.ShowAspireDispense(ctx, createdAD.ProcessID)
		if err != nil {
			logger.Errorln(responses.AspireDispenseFetchError)
			return
		}
		logger.Infoln(responses.AspireDispenseCreateSuccess, createdAD)
		go s.AddAuditLog(ctx, DBOperation, CompletedState, CreateOperation, "", responses.AspireDispenseCompletedState)
		return
	}()

	// Get highest sequence number
	// NOTE: failure already logged in internal calls

	highestSeqNum, err := s.getProcessCount(ctx, tx, recipeID)
	if err != nil {
		return
	}

	process, err := s.processOperation(ctx, name, AspireDispenseProcess, ad, Process{})
	if err != nil {
		return
	}
	// process has only a valid name
	process.SequenceNumber = highestSeqNum + 1
	process.Type = AspireDispenseProcess
	process.RecipeID = recipeID

	// create the process
	process, err = s.createProcess(ctx, tx, process)
	if err != nil {
		return
	}

	ad.ProcessID = process.ID
	createdAD, err = s.createAspireDispense(ctx, tx, ad)
	return
}

func (s *pgStore) createAspireDispense(ctx context.Context, tx *sql.Tx, ad AspireDispense) (a AspireDispense, err error) {

	var lastInsertID uuid.UUID

	err = tx.QueryRow(
		createAspireDispenseQuery,
		ad.Category,
		ad.CartridgeType,
		ad.SourcePosition,
		ad.AspireHeight,
		ad.AspireMixingVolume,
		ad.AspireNoOfCycles,
		ad.AspireVolume,
		ad.AspireAirVolume,
		ad.DispenseHeight,
		ad.DispenseMixingVolume,
		ad.DispenseNoOfCycles,
		ad.DestinationPosition,
		ad.ProcessID,
	).Scan(&lastInsertID)

	if err != nil {
		logger.WithField("err", err.Error()).Errorln(responses.AspireDispenseDBCreateError)
		return
	}

	ad.ID = lastInsertID
	return ad, err
}

func (s *pgStore) UpdateAspireDispense(ctx context.Context, ad AspireDispense) (err error) {
	// logging initialised db operation
	go s.AddAuditLog(ctx, DBOperation, InitialisedState, UpdateOperation, "", responses.AspireDispenseInitialisedState)

	var tx *sql.Tx
	tx, err = s.db.BeginTx(ctx, nil)
	if err != nil {
		logger.WithField("err:", err.Error()).Errorln(responses.AspireDispenseInitiateDBTxError)
		return
	}

	defer func() {
		if err != nil {
			tx.Rollback()
			logger.Errorln(responses.AspireDispenseUpdateError)
			go s.AddAuditLog(ctx, DBOperation, ErrorState, UpdateOperation, "", err.Error())
			return
		}
		tx.Commit()
		logger.Infoln(responses.AspireDispenseUpdateSuccess)
		go s.AddAuditLog(ctx, DBOperation, CompletedState, UpdateOperation, "", responses.AspireDispenseCompletedState)
		return
	}()

	err = s.updateProcessName(ctx, tx, ad.ProcessID, AspireDispenseProcess, ad)
	if err != nil {
		logger.WithField("err:", err.Error()).Errorln(responses.AspireDispenseUpdateNameError)
		return
	}

	result, err := tx.Exec(
		updateAspireDispenseQuery,
		ad.Category,
		ad.CartridgeType,
		ad.SourcePosition,
		ad.AspireHeight,
		ad.AspireMixingVolume,
		ad.AspireNoOfCycles,
		ad.AspireVolume,
		ad.AspireAirVolume,
		ad.DispenseHeight,
		ad.DispenseMixingVolume,
		ad.DispenseNoOfCycles,
		ad.DestinationPosition,
		time.Now(),
		ad.ProcessID,
	)

	if err != nil {
		logger.WithField("err", err.Error()).Error("Error updating aspire dispense")
		return
	}

	c, _ := result.RowsAffected()
	// check row count as no error is returned when row not found for update
	if c == 0 {
		return responses.ProcessIDInvalidError
	}

	return
}
