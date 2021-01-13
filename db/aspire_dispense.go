package db

import (
	"context"
	"time"

	"github.com/google/uuid"
	logger "github.com/sirupsen/logrus"
)

type Category string

const (
	well_to_shaker Category = "well_to_shaker"
	shaker_to_well Category = "shaker_to_well"
	well_to_well   Category = "well_to_well"
)

const (
	getAspireDispenseQuery = `SELECT id,
						category,
						well_no_source,
						aspire_height,
						aspire_mixing_volume,
						aspire_no_of_cycles,
						aspire_volume,
						dispense_height,
						dispense_mixing_volume,
						dispense_no_of_cycles,
						dispense_vol,
						dispense_blow,
						well_to_destination,
						created_at,
						updated_at
						FROM aspire_dispense
						WHERE id = $1`
	selectAspireDispenseQuery = `SELECT *
						FROM aspire_dispense`
	deleteAspireDispenseQuery = `DELETE FROM aspire_dispense
						WHERE id = $1`
	createAspireDispenseQuery = `INSERT INTO aspire_dispense (
						id,
						category,
						well_no_source,
						aspire_height,
						aspire_mixing_volume,
						aspire_no_of_cycles,
						aspire_volume,
						dispense_height,
						dispense_mixing_volume,
						dispense_no_of_cycles,
						dispense_vol,
						dispense_blow,
						well_to_destination)
						VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13) RETURNING id`
	updateAspireDispenseQuery = `UPDATE aspire_dispense SET (
						category,
						well_no_source,
						aspire_height,
						aspire_mixing_volume,
						aspire_no_of_cycles,
						aspire_volume,
						dispense_height,
						dispense_mixing_volume,
						dispense_no_of_cycles,
						dispense_vol,
						dispense_blow,
						well_to_destination,
						updated_at) = 
						($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13) WHERE id = $14`
)

type AspireDispense struct {
	ID                 uuid.UUID `db:"id" json:"id"`
	Category           Category  `db:"category" json:"category"`
	WellNoSource       int64     `db:"well_no_source" json:"well_no_source"`
	AspireHeight       float64   `db:"aspire_height" json:"aspire_height"`
	AspireMixingVolume float64   `db:"aspire_mixing_volume" json:"aspire_mixing_volume"`
	AspireNoOfCycles   int64     `db:"aspire_no_of_cycles" json:"aspire_no_of_cycles"`
	AspireVolume       float64   `db:"aspire_volume" json:"aspire_volume"`
	DispenseHeight     float64   `db:"dispense_height" json:"dispense_height"`
	DispenseMixingVol  float64   `db:"dispense_mixing_volume" json:"dispense_mixing_volume"`
	DispenseNoOfCycles int64     `db:"dispense_no_of_cycles" json:"dispense_no_of_cycles"`
	DispenseVol        float64   `db:"dispense_vol" json:"dispense_vol"`
	DispenseBlow       float64   `db:"dispense_blow" json:"dispense_blow"`
	WellToDestination  int64     `db:"well_to_destination" json:"well_to_destination"`
	CreatedAt          time.Time `db:"created_at" json:"created_at"`
	UpdatedAt          time.Time `db:"updated_at" json:"updated_at"`
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
		ad.ID,
		ad.Category,
		ad.WellNoSource,
		ad.AspireHeight,
		ad.AspireMixingVolume,
		ad.AspireNoOfCycles,
		ad.AspireVolume,
		ad.DispenseHeight,
		ad.DispenseMixingVol,
		ad.DispenseNoOfCycles,
		ad.DispenseVol,
		ad.DispenseBlow,
		ad.WellToDestination,
	).Scan(&lastInsertID)

	if err != nil {
		logger.WithField("err", err.Error()).Error("Error creating aspire dispense")
		return
	}

	err = s.db.Get(&createdAspireDispense, getAspireDispenseQuery, lastInsertID)
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
		ad.WellNoSource,
		ad.AspireHeight,
		ad.AspireMixingVolume,
		ad.AspireNoOfCycles,
		ad.AspireVolume,
		ad.DispenseHeight,
		ad.DispenseMixingVol,
		ad.DispenseNoOfCycles,
		ad.DispenseVol,
		ad.DispenseBlow,
		ad.WellToDestination,
		time.Now(),
		ad.ID,
	)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error updating aspire dispense")
		return
	}
	return
}
