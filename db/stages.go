package db

import (
	"context"

	"github.com/google/uuid"
	logger "github.com/sirupsen/logrus"
)

const (
	createStageQuery = `INSERT INTO stages (
		name,
		type,
		repeat_count,
		template_id)
		VALUES ($1, $2, $3, $4) RETURNING id`

	getStageListQuery = `SELECT * FROM stages
		where template_id = $1`

	getStageQuery = `SELECT id,
		name,
		type,
		repeat_count,
		template_id
		FROM stages
		WHERE id = $1`

	updateStageQuery = `UPDATE stages SET (
		name,
		type,
		repeat_count,
		template_id =
		($1, $2, $3, $4) where id = $5`

	deleteStageQuery = `DELETE FROM stages WHERE id = $1`
)

type Stage struct {
	ID          uuid.UUID `db:"id" json:"id"`
	Name        string    `db:"name" json:"name"`
	Type        string    `db:"type" json:"type"`
	RepeatCount int       `db:"repeat_count" json:"repeat_count"`
	TemplateID  uuid.UUID `db:"template_id" json:"template_id"`
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
		stg.Name,
		stg.Type,
		stg.RepeatCount,
		stg.TemplateID,
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
		stg.Name,
		stg.Type,
		stg.RepeatCount,
		stg.TemplateID,
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
