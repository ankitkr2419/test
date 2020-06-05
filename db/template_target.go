package db

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/google/uuid"
	logger "github.com/sirupsen/logrus"
)

const (
	getTempTargetListQuery = `SELECT * FROM template_targets
		ORDER BY name ASC
		template_id = $1`

	createTempTargetQuery = `INSERT INTO template_targets (
		template_id,
		target_id,
		threshold)
		VALUES ($1, $2, $3)`

	getTempTargetQuery = `SELECT template_id,
		target_id,
		threshold
		FROM template_targets WHERE template_id = $1 and target_id = $2`

	updateTempTargetQuery = `UPDATE template_targets SET
		threshold = $1
		WHERE template_id = $2 and target_id = $3`

	deleteTempTargetQuery = `DELETE FROM template_targets WHERE template_id = $1 and target_id = $2`

	upsertTempTargetQuery1 = `INSERT INTO template_targets (template_id,
		target_id,
		threshold)
		VALUES `

	upsertTempTargetQuery2 = `ON CONFLICT ON CONSTRAINT unq DO UPDATE
			SET threshold=excluded.threshold
			WHERE template_targets.threshold is distinct from excluded.threshold`
		)

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

// func (s *pgStore) CreateTemplateTarget(ctx context.Context, t TemplateTarget) (createdTT TemplateTarget, err error) {

// 	var id uuid.UUID

// 	err = s.db.QueryRow(createTempTargetQuery, t.TemplateID, t.TargetID, t.Threshold).Scan(&id)
// 	if err != nil {
// 		logger.WithField("err", err.Error()).Error("Error creating TemplateTarget")
// 		return
// 	}

// 	err = s.db.Get(&createdTT, getTempTargetQuery, id)
// 	if err != nil {
// 		logger.WithField("err", err.Error()).Error("Error in getting TemplateTarget")
// 		return
// 	}
// 	return
// }

// func (s *pgStore) UpdateTemplateTarget(ctx context.Context, t TemplateTarget) (err error) {
// 	result, err := s.db.Exec(
// 		updateTempTargetQuery,
// 		t.Threshold,
// 		t.TemplateID,
// 		t.TargetID,
// 	)

// 	if err != nil {
// 		logger.WithField("err", err.Error()).Error("Error updating TemplateTarget")
// 		return
// 	}

// 	c, _ := result.RowsAffected()
// 	// check row count as no error is returned when row not found for update
// 	if c == 0 {
// 		err = errors.New("Record Not Found")
// 		logger.WithField("err", err.Error()).Error("Error TemplateTarget not found")
// 	}

// 	return
// }

// func (s *pgStore) ShowTemplateTarget(ctx context.Context, tempID, targetID uuid.UUID) (dbTT TemplateTarget, err error) {
// 	err = s.db.Get(&dbTT, getTempTargetQuery, tempID, targetID)
// 	if err != nil {
// 		logger.WithField("err", err.Error()).Error("Error fetching TemplateTarget")
// 		return
// 	}

// 	return
// }

// func (s *pgStore) DeleteTemplateTarget(ctx context.Context, tempID, targetID uuid.UUID) (err error) {
// 	result, err := s.db.Exec(
// 		deleteTempTargetQuery,
// 		tempID,
// 		targetID,
// 	)
// 	if err != nil {
// 		logger.WithField("err", err.Error()).Error("Error deleting TemplateTarget")
// 		return
// 	}

// 	c, _ := result.RowsAffected()
// 	// check row count as no error is returned when row not found for update
// 	if c == 0 {
// 		err = errors.New("Record Not Found")
// 		logger.WithField("err", err.Error()).Error("Error Template not found")
// 	}

// 	return
// }

func (s *pgStore) UpsertTemplateTarget(ctx context.Context, t []TemplateTarget, temp_id uuid.UUID) (createdTT []TemplateTarget, err error) {

	stmt, args := updateQuery(t)

	logger.WithField("stmt",stmt).Info("log")
		logger.WithField("args",args).Info("log")

	_, err = s.db.Exec(stmt, args...)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error creating TemplateTarget")
		return
	}

	err = s.db.Select(&createdTT, getTempTargetListQuery, temp_id)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error listing template targets")
		return
	}
	return
}

func updateQuery(tt []TemplateTarget) (string, []interface{}) {

	values := make([]string, 0, len(tt))
	args := make([]interface{}, 0, len(tt)*3)

	for _, t := range tt {
		values = append(values, "(?, ?, ?)")
		args = append(args, t.TemplateID, t.TargetID, t.Threshold)
	}

	stmt := fmt.Sprintf(upsertTempTargetQuery1+`%s`,
		strings.Join(values, ","))

	stmt = replaceSQL(stmt, "(?, ?, ?)", len(values))

	stmt += upsertTempTargetQuery2
	return stmt, args
}

func replaceSQL(stmt, pattern string, len int) string {
	pattern += ","
	n := 0
	for strings.IndexByte(stmt, '?') != -1 {
		n++
		param := "$" + strconv.Itoa(n)
		stmt = strings.Replace(stmt, "?", param, 1)
	}
	stmt = strings.TrimSuffix(stmt, ",)")
	return stmt
}
