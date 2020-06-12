package db

import (
	"context"
	"fmt"
	"strings"

	"github.com/google/uuid"
	logger "github.com/sirupsen/logrus"
)

const (
	getTempTargetListQuery = `SELECT * FROM template_targets
		where template_id = $1`

	updateTempTargetQuery = `UPDATE template_targets SET
		threshold = $1
		WHERE template_id = $2 and target_id = $3`

	deleteTempTargetQuery = `DELETE FROM template_targets WHERE template_id = $1 and target_id = $2`

	upsertTempTargetQuery1 = `INSERT INTO template_targets (
		template_id,
		target_id,
		threshold)
		VALUES `

	upsertTempTargetQuery2 = ` ON CONFLICT (template_id, target_id) DO UPDATE
			SET threshold=excluded.threshold
			WHERE template_targets.template_id = excluded.template_id AND template_targets.target_id = excluded.target_id`
)

//TemplateTarget is used to store target mapped to template
type TemplateTarget struct {
	TemplateID uuid.UUID `db:"template_id" json:"template_id" validate:"required"`
	TargetID   uuid.UUID `db:"target_id" json:"target_id" validate:"required"`
	Threshold  float64   `db:"threshold" json:"threshold" validate:"required"`
}

func (s *pgStore) ListTemplateTargets(ctx context.Context, templateID uuid.UUID) (t []TemplateTarget, err error) {
	err = s.db.Select(&t, getTempTargetListQuery, templateID)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error listing template targets")
		return
	}
	return
}

func (s *pgStore) UpsertTemplateTarget(ctx context.Context, t []TemplateTarget, temp_id uuid.UUID) (createdTT []TemplateTarget, err error) {

	stmt := makeQuery(t)

	_, err = s.db.Exec(
		stmt,
	)
	if err != nil {
		logger.WithField("error in exec query", err.Error()).Error("Query Failed")
		return
	}

	err = s.db.Select(&createdTT, getTempTargetListQuery, temp_id)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error listing template targets")
		return
	}
	return
}

// prepare bulk insert query statement
func makeQuery(tt []TemplateTarget) string {

	values := make([]string, 0, len(tt))

	for _, t := range tt {
		// single quotes used to insert uuid
		values = append(values, fmt.Sprintf("('%v', '%v',%v)", t.TemplateID, t.TargetID, t.Threshold))
	}

	stmt := fmt.Sprintf(upsertTempTargetQuery1+" %s",
		strings.Join(values, ","))

	stmt += upsertTempTargetQuery2
	return stmt
}
