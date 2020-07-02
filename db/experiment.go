package db

import (
	"context"
	"time"

	"github.com/google/uuid"
	logger "github.com/sirupsen/logrus"
)

const (
	getExperimentListQuery = `SELECT * FROM experiments`

	createExperimentQuery = `INSERT INTO experiments (
		description,
		template_id,
		operator_name,
 		start_time,
 		end_time)
		VALUES ($1, $2,$3, $4,$5) RETURNING id`

	getExperimentQuery = `SELECT e.id,
		e.description,
		e.template_id,
		e.operator_name,
		e.start_time,
		e.end_time,
        t.name as template_name
		FROM experiments as e,templates as t
		WHERE t.id = e.template_id AND e.id = $1`

	updateStartTimeQuery = `UPDATE experiments
		SET start_time = $1
		WHERE id = $2`
)

type Experiment struct {
	ID           uuid.UUID `db:"id" json:"id"`
	Description  string    `db:"description" json:"description" validate:"required"`
	TemplateID   uuid.UUID `db:"template_id" json:"template_id" validate:"required"`
	TemplateName string    `db:"template_name" json:"template_name"`
	OperatorName string    `db:"operator_name" json:"operator_name"`
	StartTime    time.Time `db:"start_time" json:"start_time"`
	EndTime      time.Time `db:"end_time" json:"end_time"`
}

func (s *pgStore) ListExperiments(ctx context.Context) (e []Experiment, err error) {
	err = s.db.Select(e, getExperimentListQuery)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error listing experiments")
		return
	}
	return
}

func (s *pgStore) CreateExperiment(ctx context.Context, e Experiment) (createdTemp Experiment, err error) {

	var id uuid.UUID

	err = s.db.QueryRow(
		createExperimentQuery,
		e.Description,
		e.TemplateID,
		e.OperatorName,
		e.StartTime,
		e.EndTime).Scan(&id)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error creating Experiment")
		return
	}

	err = s.db.Get(&createdTemp, getExperimentQuery, id)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error in getting Experiment")
		return
	}
	return
}

func (s *pgStore) ShowExperiment(ctx context.Context, id uuid.UUID) (e Experiment, err error) {
	err = s.db.Get(&e, getExperimentQuery, id)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error fetching Experiment")
		return
	}

	return
}

func (s *pgStore) UpdateStartTimeExperiments(ctx context.Context, t time.Time, experimentID uuid.UUID) (err error) {
	_, err = s.db.Exec(
		updateStartTimeQuery,
		t,
		experimentID,
	)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error updating experiments")
		return
	}
	return
}
