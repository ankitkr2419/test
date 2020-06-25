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

	upsertExpTempTargetQuery = `INSERT INTO experiment_template_targets (
		experiment_id,
		template_id,
		target_id,
		threshold)
		VALUES `

	deleteExpTempTargetsQuery = `DELETE FROM experiment_template_targets
		WHERE
		experiment_template_targets.experiment_id = $1
		AND
		experiment_template_targets.template_id = $2 `
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

	tx, err := s.db.Begin()
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error in creating transaction")
		return
	}
	_, err = tx.Exec(
		deleteExpTempTargetsQuery,
		t[0].ExperimentID,
		t[0].TemplateID,
	)
	if err != nil {
		tx.Rollback()
		logger.WithField("err", err.Error()).Error("Error deleting previous targets")
		return
	}

	_, err = tx.Exec(
		stmt,
	)
	if err != nil {
		tx.Rollback()
		logger.WithField("error in exec query", err.Error()).Error("Query Failed")
		return
	}

	tx.Commit()

	err = s.db.Select(&createdETT, getExpTempTargetListQuery, ExperimentID)
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

	stmt := fmt.Sprintf(upsertExpTempTargetQuery+" %s",
		strings.Join(values, ","))

	return stmt
}
