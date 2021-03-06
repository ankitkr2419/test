package db

import (
	"context"
	"errors"
	"fmt"
	"mylab/cpagent/config"
	"strings"

	"github.com/google/uuid"
	logger "github.com/sirupsen/logrus"
)

const (
	getTempTargetListQuery = `SELECT * FROM template_targets
		where template_id = $1`

	upsertTempTargetQuery = `INSERT INTO template_targets (
		template_id,
		target_id,
		threshold)
		VALUES `

	deleteTempTargetsQuery = `DELETE FROM template_targets
		where template_id = $1`

	selectICTarget = `SELECT tt.*,d.position 
		FROM template_targets as tt,
		targets as t ,
		dyes as d 
		WHERE tt.target_id = t.id AND t.dye_id = d.id AND 
		tt.template_id = $1 AND d.position = $2`
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

	tx, err := s.db.Begin()
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error in creating transaction")
		return
	}
	_, err = tx.Exec(
		deleteTempTargetsQuery,
		temp_id,
	)
	if err != nil {
		tx.Rollback()
		logger.WithField("err", err.Error()).Error("Error deleting previous template targets")
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

	stmt := fmt.Sprintf(upsertTempTargetQuery+" %s",
		strings.Join(values, ","))

	return stmt
}

func (s *pgStore) CheckIfICTargetAdded(ctx context.Context, temp_id uuid.UUID) (resp WarnResponse, err error) {
	result, err := s.db.Exec(
		selectICTarget,
		temp_id,
		config.GetICPosition(),
	)

	if err != nil {
		logger.WithField("err", err.Error()).Error("Error getting IC target")
		return
	}

	c, _ := result.RowsAffected()
	// check row count as no error is returned when row not found for update
	if c == 0 {
		resp.Code = "Warning"
		resp.Message = "Absence of Internal Control"
		err = errors.New("Record Not Found")
		logger.WithField("err", err.Error()).Error("Error IC target not found")
		return
	}

	return
}
