package db

import (
	"context"

	"github.com/google/uuid"
	logger "github.com/sirupsen/logrus"
)

const (
	createStepQuery = `INSERT INTO steps (
		stage_id,
		ramp_rate,
		target_temp,
		hold_time,
		data_capture)
		VALUES ($1, $2, $3, $4 ,$5) RETURNING id`

	getStepsListQuery = `SELECT * FROM steps
		where stage_id = $1`

	getStepQuery = `SELECT id,
		stage_id,
		ramp_rate,
		target_temp,
		hold_time,
		data_capture
		FROM steps WHERE id = $1`

	updateStepQuery = `UPDATE steps SET (
		stage_id,
		ramp_rate,
		target_temp,
		hold_time,
		data_capture) =
		($1, $2, $3, $4, $5) where id = $6`

	deleteStepQuery = `DELETE FROM steps WHERE id = $1`
)

type Step struct {
	ID                uuid.UUID `db:"id" json:"id"`
	StageID           uuid.UUID `db:"stage_id" json:"stage_id"`
	RampRate          float64   `db:"ramp_rate" json:"ramp_rate"`
	TargetTemperature float64   `db:"target_temp" json:"target_temp"`
	HoldTime          int       `db:"hold_time" json:"hold_time"`
	DataCapture       bool      `db:"data_capture" json:"data_capture"`
}

func (s *pgStore) ListSteps(ctx context.Context, stgID uuid.UUID) (steps []Step, err error) {
	err = s.db.Select(&steps, getStepsListQuery, stgID)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error listing steps")
		return
	}

	return
}

func (s *pgStore) CreateStep(ctx context.Context, st Step) (createdStep Step, err error) {
	var lastInsertID uuid.UUID
	err = s.db.QueryRow(
		createStepQuery,
		st.StageID,
		st.RampRate,
		st.TargetTemperature,
		st.HoldTime,
		st.DataCapture,
	).Scan(&lastInsertID)

	if err != nil {
		logger.WithField("err", err.Error()).Error("Error creating Step")
		return
	}

	err = s.db.Get(&createdStep, getStepQuery, lastInsertID)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error in getting Stage")
		return
	}
	return
}

func (s *pgStore) UpdateStep(ctx context.Context, st Step) (err error) {

	_, err = s.db.Exec(
		updateStepQuery,
		st.StageID,
		st.RampRate,
		st.TargetTemperature,
		st.HoldTime,
		st.DataCapture,
		st.ID,
	)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error updating Stage")
		return
	}

	return
}

func (s *pgStore) ShowStep(ctx context.Context, id uuid.UUID) (dbStep Step, err error) {
	err = s.db.Get(&dbStep, getStepQuery, id)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error fetching Step")
		return
	}

	return
}

func (s *pgStore) DeleteStep(ctx context.Context, id uuid.UUID) (err error) {

	// added delete cascade
	_, err = s.db.Exec(
		deleteStepQuery,
		id,
	)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error deleting Step")
		return
	}

	return
}
