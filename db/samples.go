package db

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	logger "github.com/sirupsen/logrus"
)

const (
	createSampleQuery = `INSERT INTO samples (
		name)
		VALUES ($1) RETURNING id`

	getSamplesListQuery = `SELECT * FROM samples`

	getSampleQuery = `SELECT id,
		name
		FROM samples WHERE id = $1`

	updateSampleQuery = `UPDATE samples SET (
		name)
		= ($1)
		where id = $2`

	deleteSampleQuery = `DELETE FROM samples WHERE id = $1`

	findSampleQuery = `SELECT * from samples
		where name LIKE %s`
)

type Sample struct {
	ID   uuid.UUID `db:"id" json:"id"`
	Name string    `db:"name" json:"name" validate:"required"`
}

func (s *pgStore) ListSamples(ctx context.Context) (samples []Sample, err error) {
	err = s.db.Select(&samples, getSamplesListQuery)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error listing samples")
		return
	}

	return
}

func (s *pgStore) CreateSample(ctx context.Context, sp Sample) (createdSample Sample, err error) {
	var lastInsertID uuid.UUID
	err = s.db.QueryRow(
		createSampleQuery,
		sp.Name,
	).Scan(&lastInsertID)

	if err != nil {
		logger.WithField("err", err.Error()).Error("Error creating Sample")
		return
	}

	err = s.db.Get(&createdSample, getSampleQuery, lastInsertID)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error in getting Stage")
		return
	}
	return
}

func (s *pgStore) UpdateSample(ctx context.Context, sp Sample) (err error) {

	_, err = s.db.Exec(
		updateSampleQuery,
		sp.Name,
		sp.ID,
	)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error updating Stage")
		return
	}

	return
}

func (s *pgStore) ShowSample(ctx context.Context, id uuid.UUID) (dbSample Sample, err error) {
	err = s.db.Get(&dbSample, getSampleQuery, id)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error fetching Sample")
		return
	}

	return
}

func (s *pgStore) DeleteSample(ctx context.Context, id uuid.UUID) (err error) {

	_, err = s.db.Exec(
		deleteSampleQuery,
		id,
	)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error deleting Sample")
		return
	}

	return
}

func (s *pgStore) FindSamples(ctx context.Context, searchText string) (samples []Sample, err error) {

	// appending search text with % padding to match the text
	findSampleQuery := fmt.Sprintf(findSampleQuery, "'%"+searchText+"%'")

	err = s.db.Select(&samples, findSampleQuery)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error listing samples")
		return
	}

	return
}
