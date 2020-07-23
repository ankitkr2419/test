package db

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	logger "github.com/sirupsen/logrus"
)

const (
	getwellsConfigured = `SELECT
			e.id as experiment_id,
			ett.template_id,ett.target_id,ett.threshold,
			d.position as dye_position,
			t.name as target_name
			FROM
			experiments e
            INNER JOIN experiment_template_targets ett ON ett.experiment_id = e.id
			INNER JOIN targets t ON t.id = ett.target_id
			INNER JOIN dyes d ON d.id = t.dye_id
			WHERE e.id = $1`

	insertResultQuery = `INSERT INTO results (
		experiment_id,
		target_id,
		well_position,
		cycle,
		f_value)
		 VALUES %s`

	getResultListQuery = `SELECT * FROM results
		WHERE
		experiment_id = $1
		AND
		cycle = $2`

	getResultWellTargetsQuery = `SELECT
		r.experiment_id,r.target_id,w.id as well_d,r.f_value,e.threshold
			FROM
			results r
			INNER JOIN experiment_template_targets e ON e.target_id = r.target_id
			AND r.experiment_id = e.experiment_id
            INNER JOIN wells w ON r.experiment_id = w.experiment_id
			AND r.well_position = w.position
			WHERE r.experiment_id = $1 AND r.cycle = $2`

	getAllCyclesResultQuery = `
		SELECT
 		r.experiment_id,t.template_id,r.target_id,t.threshold,r.cycle,r.f_value,r.well_position
		FROM results as r , experiment_template_targets as t
		WHERE r.experiment_id = t.experiment_id AND r.target_id = t.target_id
		AND r.experiment_id = $1
		ORDER BY created_at ASC`
)

type Result struct {
	ExperimentID uuid.UUID `db:"experiment_id" json:"experiment_id"`
	TemplateID   uuid.UUID `db:"template_id" json:"template_id"`
	WellPosition int32     `db:"well_position" json:"well_position"`
	TargetID     uuid.UUID `db:"target_id" json:"target_id"`
	Cycle        uint16    `db:"cycle" json:"cycle"`
	FValue       uint16    `db:"f_value" json:"f_value"`
	Threshold    float32   `db:"threshold" json:"threshold"`
}

type TargetDetails struct {
	ExperimentID uuid.UUID `db:"experiment_id" json:"experiment_id"`
	TemplateID   uuid.UUID `db:"template_id" json:"template_id"`
	TargetID     uuid.UUID `db:"target_id" json:"target_id"`
	Threshold    float32   `db:"threshold" json:"threshold"`
	TargetName   string    `db:"target_name" json:"target_name"`
	DyePosition  int32     `db:"dye_position" json:"dye_position"`
}

type WellTargetResults struct {
	WellTarget
	Cycle     uint16    `db:"cycle" json:"cycle"`
	FValue    uint16    `db:"f_value" json:"f_value"`
	Threshold float32   `db:"threshold" json:"threshold"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

func (s *pgStore) ListConfTargets(ctx context.Context, experimentID uuid.UUID) (w []TargetDetails, err error) {
	err = s.db.Select(&w, getwellsConfigured, experimentID)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error listing well details")
		return
	}
	return
}

func (s *pgStore) ListWellTargetsResult(ctx context.Context, r Result) (w []WellTargetResults, err error) {
	err = s.db.Select(&w, getResultWellTargetsQuery, r.ExperimentID, r.Cycle)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error listing well details")
		return
	}
	return
}

func (s *pgStore) GetResult(ctx context.Context, ExperimentID uuid.UUID) (rDB []Result, err error) {
	err = s.db.Select(&rDB, getAllCyclesResultQuery, ExperimentID)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error listing well details")
		return
	}

	return
}

func (s *pgStore) InsertResult(ctx context.Context, r []Result) (rDB []Result, err error) {
	stmt := makeResultQuery(r)

	_, err = s.db.Exec(
		stmt,
	)
	if err != nil {
		logger.WithField("error in exec query", err.Error()).Error("Query Failed")
		return
	}

	err = s.db.Select(&rDB, getAllCyclesResultQuery, r[0].ExperimentID)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error listing well details")
		return
	}

	return
}

func makeResultQuery(results []Result) string {

	values := make([]string, 0, len(results))

	for _, r := range results {

		values = append(values, fmt.Sprintf("('%v', '%v', %v,%v,%v)", r.ExperimentID, r.TargetID, r.WellPosition, r.Cycle, r.FValue))
	}

	stmt := fmt.Sprintf(insertResultQuery,
		strings.Join(values, ","))

	return stmt
}
