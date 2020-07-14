package db

import (
	"context"
	"time"

	"github.com/google/uuid"
	logger "github.com/sirupsen/logrus"
)

const (
	createStageQuery = `INSERT INTO stages (
		type,
		repeat_count,
		template_id,
		step_count)
		VALUES ($1, $2, $3, $4) RETURNING id`

	getStageListQuery = `SELECT * FROM stages
		where template_id = $1
		ORDER BY created_at ASC`

	getStageQuery = `SELECT id,
		type,
		repeat_count,
		template_id,
		step_count,
		created_at,
        updated_at
		FROM stages
		WHERE id = $1`

	updateStageQuery = `UPDATE stages SET (
		type,
		repeat_count,
		template_id,
		updated_at) =
		($1, $2, $3,$4) where id = $5`

	deleteStageQuery = `DELETE FROM stages WHERE id = $1`

	updateStepCountQuery = `UPDATE stages
		SET step_count = subquery.no_of_steps
			FROM (
    			SELECT count(stage_id) AS no_of_steps, stage_id
    			FROM   steps
    			GROUP  BY stage_id
    		) AS subquery
		Where stages.id = subquery.stage_id`

	getStageStepQuery = `SELECT *
		FROM stages , steps
		WHERE
		stages.id = steps.stage_id
		AND stages.template_id = $1
		ORDER BY steps.created_at ASC`
)

type Stage struct {
	ID          uuid.UUID `db:"id" json:"id"`
	Type        string    `db:"type" json:"type"  validate:"required"`
	RepeatCount uint16    `db:"repeat_count" json:"repeat_count"`
	TemplateID  uuid.UUID `db:"template_id" json:"template_id" validate:"required"`
	StepCount   int       `db:"step_count" json:"step_count"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
}

// StageStep used to get data for particular template
type StageStep struct {
	Stage
	Step
}

func (s *pgStore) ListStages(ctx context.Context, template_id uuid.UUID) (stgs []Stage, err error) {
	err = s.db.Select(&stgs, getStageListQuery, template_id)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error listing stages")
		return
	}

	return
}

func (s *pgStore) CreateStage(ctx context.Context, stg Stage) (createdStage Stage, err error) {
	var lastInsertId uuid.UUID
	err = s.db.QueryRow(
		createStageQuery,
		stg.Type,
		stg.RepeatCount,
		stg.TemplateID,
		0, //initial step count = 0
	).Scan(&lastInsertId)

	if err != nil {
		logger.WithField("err", err.Error()).Error("Error creating Stage")
		return
	}

	err = s.db.Get(&createdStage, getStageQuery, lastInsertId)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error in getting Stage")
		return
	}
	return
}

func (s *pgStore) UpdateStage(ctx context.Context, stg Stage) (err error) {
	_, err = s.db.Exec(
		updateStageQuery,
		stg.Type,
		stg.RepeatCount,
		stg.TemplateID,
		time.Now(),
		stg.ID,
	)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error updating Stage")
		return
	}
	return
}

func (s *pgStore) ShowStage(ctx context.Context, id uuid.UUID) (dbStage Stage, err error) {
	err = s.db.Get(&dbStage, getStageQuery, id)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error fetching Stage")
		return
	}

	return
}

func (s *pgStore) DeleteStage(ctx context.Context, id uuid.UUID) (err error) {

	// added delete cascade
	_, err = s.db.Exec(
		deleteStageQuery,
		id,
	)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error deleting Stage")
		return
	}

	return
}

func (s *pgStore) UpdateStepCount(ctx context.Context) (err error) {
	_, err = s.db.Exec(
		updateStepCountQuery,
	)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error updating Stage")
		return
	}
	return
}

func (s *pgStore) ListStageSteps(ctx context.Context, templateID uuid.UUID) (ss []StageStep, err error) {
	err = s.db.Select(&ss, getStageStepQuery, templateID)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error listing stage steps")
		return
	}

	return
}
