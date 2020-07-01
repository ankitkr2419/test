package db

import (
	"context"
	"fmt"
	"strings"

	"github.com/google/uuid"
	logger "github.com/sirupsen/logrus"
)

const (
	getwellsConfigured = `SELECT
             experiments.id as experiment_id,template_targets.template_id,template_targets.target_id,template_targets.threshold,dyes.position as dye_position,targets.name as target_name
			FROM
			experiments
            INNER JOIN template_targets ON template_targets.template_id = experiments.template_id
			INNER JOIN targets ON targets.id = template_targets.target_id
			INNER JOIN dyes ON dyes.id = targets.dye_id
			WHERE experiments.id = $1`

	insertResultQuery = `INSERT INTO results (
		experiment_id,
		target_id,
		well_position,
		cycle,
		f_Value)
		 VALUES %s`

	getResultListQuery = `SELECT * FROM results
		WHERE
		experiment_id = $1
		AND
		cycle = $2`
)

type Result struct {
	ExperimentID uuid.UUID `db:"experiment_id" json:"experiment_id"`
	WellPosition int32     `db:"well_id" json:"well_id"`
	TargetID     uuid.UUID `db:"target_id" json:"target_id"`
	Cycle        uint16    `db:"cycle" json:"cycle"`
	FValue       uint16    `db:"f_value" json:"f_value"`
}

type TargetDetails struct {
	ExperimentID uuid.UUID `db:"experiment_id" json:"experiment_id"`
	TemplateID   uuid.UUID `db:"template_id" json:"template_id"`
	TargetID     uuid.UUID `db:"target_id" json:"target_id"`
	Threshold    float32   `db:"threshold" json:"threshold"`
	TargetName   string    `db:"target_name" json:"target_name"`
	DyePosition  int32     `db:"dye_position" json:"dye_position"`
}

func (s *pgStore) ListConfTargets(ctx context.Context, experimentID uuid.UUID) (w []TargetDetails, err error) {
	err = s.db.Select(&w, getwellsConfigured, experimentID)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error listing well details")
		return
	}
	return
}

func (s *pgStore) InsertResult(ctx context.Context, r []Result) (err error) {
	stmt := makeResultQuery(r)

	_, err = s.db.Exec(
		stmt,
	)
	if err != nil {
		logger.WithField("error in exec query", err.Error()).Error("Query Failed")
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
