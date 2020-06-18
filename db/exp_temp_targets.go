package db

import (
	"context"
	"fmt"
	"strings"

	"github.com/google/uuid"
	logger "github.com/sirupsen/logrus"
)

const (
	getExpTempTargetListQuery = `SELECT * FROM experiment_template_targets
		where experiment_id = $1`

	upsertExpTempTargetQuery1 = `INSERT INTO experiment_template_targets (
		expriment_id,
		template_id,
		target_id,
		threshold)
		VALUES `

	upsertExpTempTargetQuery2 = ` ON CONFLICT (experiment_id, template_id, target_id) DO UPDATE
		SET threshold=excluded.threshold
		WHERE
		experiment_template_targets.experiment_id = excluded.experiment_id
		AND
		experiment_template_targets.template_id = excluded.template_id
		AND
		experiment_template_targets.target_id = excluded.target_id`
)

//ExpTemplateTarget is used to store target mapped to template with respect to experiment
type ExpTemplateTarget struct {
	ExperimentID uuid.UUID `db:"experiment_id" json:"experiment_id" validate:"required"`
	TemplateID   uuid.UUID `db:"template_id" json:"template_id" validate:"required"`
	TargetID     uuid.UUID `db:"target_id" json:"target_id" validate:"required"`
	Threshold    float64   `db:"threshold" json:"threshold" validate:"required"`
}

func (s *pgStore) ListExpTemplateTargets(ctx context.Context, experimentID uuid.UUID) (t []ExpTemplateTarget, err error) {
	err = s.db.Select(&t, getExpTempTargetListQuery, experimentID)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error listing experiment template targets")
		return
	}
	return
}

func (s *pgStore) UpsertExpTemplateTarget(ctx context.Context, t []ExpTemplateTarget, ExperimentID uuid.UUID) (createdETT []ExpTemplateTarget, err error) {

	stmt := makeInsertQuery(t)

	_, err = s.db.Exec(
		stmt,
	)
	if err != nil {
		logger.WithField("error in exec query", err.Error()).Error("Query Failed")
		return
	}

	err = s.db.Select(&createdETT, getTempTargetListQuery, ExperimentID)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error listing experiment template targets")
		return
	}
	return
}

// prepare bulk insert query statement
func makeInsertQuery(tt []ExpTemplateTarget) string {

	values := make([]string, 0, len(tt))

	for _, t := range tt {
		// single quotes used to insert uuid
		values = append(values, fmt.Sprintf("('%v','%v', '%v',%v)", t.ExperimentID, t.TemplateID, t.TargetID, t.Threshold))
	}

	stmt := fmt.Sprintf(upsertExpTempTargetQuery1+" %s",
		strings.Join(values, ","))

	stmt += upsertExpTempTargetQuery2
	return stmt
}
