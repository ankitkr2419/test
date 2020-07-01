package db

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	logger "github.com/sirupsen/logrus"
)

const (
	getSampleQuery = `SELECT * FROM samples
		WHERE id = $1`
	createSampleQuery = `INSERT INTO samples (
		name)
		VALUES ($1) RETURNING id`

	findSampleQuery = `SELECT * from samples
		where name LIKE %s`
)

type Sample struct {
	ID   uuid.UUID `db:"id" json:"id"`
	Name string    `db:"name" json:"name" validate:"required"`
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
		logger.WithField("err", err.Error()).Error("Error in getting sample")
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
