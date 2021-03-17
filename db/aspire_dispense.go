package db

import (
	"context"
	"time"

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
	deleteAspireDispenseQuery = `DELETE FROM aspire_dispense
						WHERE process_id = $1`
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
						dispense_volume,
						dispense_blow_volume,
						destination_position,
						process_id)
						VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15) RETURNING id`
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
						dispense_volume,
						dispense_blow_volume,
						destination_position,
						updated_at) = 
						($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15) WHERE process_id = $16`
)

type AspireDispense struct {
	ID                   uuid.UUID     `db:"id" json:"id"`
	Category             Category      `db:"category" json:"category"`
	CartridgeType        CartridgeType `db:"cartridge_type" json:"cartridge_type"`
	SourcePosition       int64         `db:"source_position" json:"source_position"`
	AspireHeight         float64       `db:"aspire_height" json:"aspire_height"`
	AspireMixingVolume   float64       `db:"aspire_mixing_volume" json:"aspire_mixing_volume"`
	AspireNoOfCycles     int64         `db:"aspire_no_of_cycles" json:"aspire_no_of_cycles"`
	AspireVolume         float64       `db:"aspire_volume" json:"aspire_volume"`
	AspireAirVolume      float64       `db:"aspire_air_volume" json:"aspire_air_volume"`
	DispenseHeight       float64       `db:"dispense_height" json:"dispense_height"`
	DispenseMixingVolume float64       `db:"dispense_mixing_volume" json:"dispense_mixing_volume"`
	DispenseNoOfCycles   int64         `db:"dispense_no_of_cycles" json:"dispense_no_of_cycles"`
	DispenseVolume       float64       `db:"dispense_volume" json:"dispense_volume"`
	DispenseBlowVolume   float64       `db:"dispense_blow_volume" json:"dispense_blow_volume"`
	DestinationPosition  int64         `db:"destination_position" json:"destination_position"`
	ProcessID            uuid.UUID     `db:"process_id" json:"process_id"`
	CreatedAt            time.Time     `db:"created_at" json:"created_at"`
	UpdatedAt            time.Time     `db:"updated_at" json:"updated_at"`
}

func (s *pgStore) ShowAspireDispense(ctx context.Context, id uuid.UUID) (dbAspireDispense AspireDispense, err error) {
	err = s.db.Get(&dbAspireDispense, getAspireDispenseQuery, id)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error fetching aspire dispense")
		return
	}
	return
}

func (s *pgStore) ListAspireDispense(ctx context.Context) (dbAspireDispense []AspireDispense, err error) {
	err = s.db.Select(&dbAspireDispense, selectAspireDispenseQuery)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error fetching aspire dispense")
		return
	}
	return
}

func (s *pgStore) CreateAspireDispense(ctx context.Context, ad AspireDispense) (createdAspireDispense AspireDispense, err error) {
	var lastInsertID uuid.UUID
	err = s.db.QueryRow(
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
		ad.DispenseVolume,
		ad.DispenseBlowVolume,
		ad.DestinationPosition,
		ad.ProcessID,
	).Scan(&lastInsertID)

	if err != nil {
		logger.WithField("err", err.Error()).Error("Error creating aspire dispense")
		return
	}

	err = s.db.Get(&createdAspireDispense, getAspireDispenseQuery, ad.ProcessID)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error in getting aspire dispense")
		return
	}
	return
}

func (s *pgStore) DeleteAspireDispense(ctx context.Context, id uuid.UUID) (err error) {
	_, err = s.db.Exec(deleteAspireDispenseQuery, id)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error deleting aspire dispense")
		return
	}
	return
}

func (s *pgStore) UpdateAspireDispense(ctx context.Context, ad AspireDispense) (err error) {
	_, err = s.db.Exec(
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
		ad.DispenseVolume,
		ad.DispenseBlowVolume,
		ad.DestinationPosition,
		time.Now(),
		ad.ID,
	)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error updating aspire dispense")
		return
	}
	return
}
