package db

import (
	"context"
	"errors"
	"mylab/cpagent/config"
	"time"

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
		where stage_id = $1 ORDER BY created_at ASC`

	getStepQuery = `SELECT id,
		stage_id,
		ramp_rate,
		target_temp,
		hold_time,
		data_capture,
		created_at,
        updated_at
		FROM steps WHERE id = $1`

	updateStepQuery = `UPDATE steps SET (
		stage_id,
		ramp_rate,
		target_temp,
		hold_time,
		data_capture,
                updated_at) =
		($1, $2, $3, $4, $5,$6) where id = $7`

	deleteStepQuery = `DELETE FROM steps WHERE id = $1`
)

type Step struct {
	ID                uuid.UUID `db:"id" json:"step_id"`
	StageID           uuid.UUID `db:"stage_id" json:"stage_id"`
	RampRate          float32   `db:"ramp_rate" json:"ramp_rate"`
	TargetTemperature float32   `db:"target_temp" json:"target_temp"`
	HoldTime          int32     `db:"hold_time" json:"hold_time"`
	DataCapture       bool      `db:"data_capture" json:"data_capture"`
	CreatedAt         time.Time `db:"created_at" json:"step_created_at"`
	UpdatedAt         time.Time `db:"updated_at" json:"step_updated_at"`
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

	if st.DataCapture && (st.HoldTime < int32(config.GetCycleTime())) {
		err = errors.New("invalid step with invalid hold time")
		return
	}
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

	if st.DataCapture && (st.HoldTime < int32(config.GetCycleTime())) {
		err = errors.New("invalid step with invalid hold time")
		return
	}
	_, err = s.db.Exec(
		updateStepQuery,
		st.StageID,
		st.RampRate,
		st.TargetTemperature,
		st.HoldTime,
		st.DataCapture,
		time.Now(),
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
