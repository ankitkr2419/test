package db

import (
	"context"
	"fmt"
	"strings"

	"github.com/google/uuid"
	logger "github.com/sirupsen/logrus"
)

const (
	getWellsListQuery = `SELECT
		wells.id,wells.position,wells.experiment_id,wells.sample_id,wells.task,wells.color_code,
		samples.name AS sample_name
		FROM wells ,samples
		WHERE
		wells.sample_id = samples.id
		AND wells.experiment_id = $1;`

	getWellsByIDQuery = `SELECT
		wells.id,wells.position,wells.experiment_id,wells.sample_id,wells.task,wells.color_code,
		samples.name AS sample_name
		FROM wells ,samples
		WHERE
		wells.sample_id = samples.id
		AND wells.id IN (%s);`

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
		task,
		color_code)
		VALUES %s`

	upsertWellQuery2 = ` ON CONFLICT (position, experiment_id) DO UPDATE
			SET
			sample_id = excluded.sample_id,
			task = excluded.task
			WHERE wells.position = excluded.position AND wells.experiment_id = excluded.experiment_id
			RETURNING id`
)

type WellConfig struct {
	Position []int32     `json:"position" validate:"required"`
	Sample   Sample      `json:"sample" validate:"required"`
	Task     string      `json:"task" validate:"required"`
	Targets  []uuid.UUID `json:"targets" validate:"required"`
}

type Well struct {
	ID           uuid.UUID    `db:"id" json:"id"`
	Position     int32        `db:"position" json:"position" validate:"required"`
	ExperimentID uuid.UUID    `db:"experiment_id" json:"experiment_id"validate:"required"`
	SampleID     uuid.UUID    `db:"sample_id" json:"sample_id"validate:"required"`
	Task         string       `db:"task" json:"task"validate:"required"`
	ColorCode    string       `db:"color_code" json:"color_code"`
	Targets      []WellTarget `json:"targets" validate:"required"`
	SampleName   string       `db:"sample_name" json:"sample_name"`
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

	wellIDs := make([]uuid.UUID, 0)

	// bulk insert do not return last inserted ids for all rows so,exec with loop

	for _, w := range Wells {

		stmt := makeWellQuery(w)

		var lastInsertID uuid.UUID

		err = s.db.QueryRow(
			stmt,
		).Scan(&lastInsertID)
		if err != nil {
			logger.WithField("error in exec query", err.Error()).Error("Query Failed")
			return
		}

		wellIDs = append(wellIDs, lastInsertID)
	}

	stmt := getWellsQuery(wellIDs)

	err = s.db.Select(&DBWells, stmt)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error listing Wells")
		return
	}

	return
}

// prepare bulk insert query statement
func makeWellQuery(w Well) string {

	values := make([]string, 0, 1)

	values = append(values, fmt.Sprintf("(%v,'%v', '%v', '%v','%v')", w.Position, w.ExperimentID, w.SampleID, w.Task, w.ColorCode))

	stmt := fmt.Sprintf(upsertWellQuery1,
		strings.Join(values, ","))

	stmt += upsertWellQuery2

	return stmt
}

func getWellsQuery(id []uuid.UUID) string {

	values := make([]string, 0, len(id))

	for _, i := range id {
		values = append(values, fmt.Sprintf("'%v'", i))
	}

	stmt := fmt.Sprintf(getWellsByIDQuery,
		strings.Join(values, ","))

	return stmt
}
