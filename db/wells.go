package db

import (
	"context"
	"fmt"
	"strings"

	"github.com/google/uuid"
	logger "github.com/sirupsen/logrus"
)

const (
	getWellsListQuery = `SELECT * FROM wells
		WHERE experiment_id = $1`

	getWellQuery = `SELECT id,
		position,
		experiment_id,
		sample_id,
		task,
		color_codes
		FROM wells WHERE id = $1`

	deleteWellQuery = `DELETE FROM wells WHERE id = $1`

	upsertWellQuery1 = `INSERT INTO wells (
		position,
		experiment_id,
		sample_id,
		task)
		VALUES %s`

	upsertWellQuery2 = ` ON CONFLICT (position, experiment_id) DO UPDATE
			SET
			sample_id = excluded.sample_id,
			task = excluded.task
			WHERE wells.position = excluded.position AND wells.experiment_id = excluded.experiment_id`
)

type Well struct {
	ID           uuid.UUID    `db:"id" json:"id"`
	Position     int          `db:"position" json:"position"validate:"required"`
	ExperimentID uuid.UUID    `db:"experiment_id" json:"experiment_id"validate:"required"`
	SampleID     uuid.UUID    `db:"sample_id" json:"sample_id"validate:"required"`
	Task         string       `db:"task" json:"task"validate:"required"`
	ColorCode    string       `db:"color_code" json:"color_code"`
	Targets      []WellTarget `json:"targets" validate:"required"`
}

func (s *pgStore) ListWells(ctx context.Context, experimentID uuid.UUID) (Wells []Well, err error) {
	err = s.db.Select(&Wells, getWellsListQuery, experimentID)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error listing Wells")
		return
	}

	return
}

func (s *pgStore) ShowWell(ctx context.Context, id uuid.UUID) (dbWell Well, err error) {
	err = s.db.Get(&dbWell, getWellQuery, id)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error fetching Well")
		return
	}

	return
}

func (s *pgStore) DeleteWell(ctx context.Context, id uuid.UUID) (err error) {

	_, err = s.db.Exec(
		deleteWellQuery,
		id,
	)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error deleting Well")
		return
	}

	return
}

func (s *pgStore) UpsertWells(ctx context.Context, Wells []Well, experimentID uuid.UUID) (DBWells []Well, err error) {

	stmt := makeWellQuery(Wells)

	_, err = s.db.Exec(
		stmt,
	)
	if err != nil {
		logger.WithField("error in exec query", err.Error()).Error("Query Failed")
		return
	}

	err = s.db.Select(&DBWells, getWellsListQuery, experimentID)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error listing Wells")
		return
	}

	return
}

// prepare bulk insert query statement
func makeWellQuery(Wells []Well) string {

	values := make([]string, 0, len(Wells))

	for _, t := range Wells {
		values = append(values, fmt.Sprintf("(%v,'%v', '%v', '%v')", t.Position, t.ExperimentID, t.SampleID, t.Task))
	}

	stmt := fmt.Sprintf(upsertWellQuery1,
		strings.Join(values, ","))

	stmt += upsertWellQuery2

	return stmt
}
