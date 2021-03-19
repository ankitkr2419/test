package db

import (
	"context"
	"time"

	"github.com/google/uuid"
	logger "github.com/sirupsen/logrus"
)

type Shaking struct {
	ID          uuid.UUID     `json:"id" db:"id"`
	WithTemp    bool          `json:"with_temp" db:"with_temp"`
	Temperature float64       `json:"temperature" db:"temperature"`
	FollowTemp  bool          `json:"follow_temp" db:"follow_temp"`
	ProcessID   uuid.UUID     `json:"process_id" db:"process_id"`
	RPM1        int64         `json:"rpm_1" db:"rpm_1"`
	RPM2        int64         `json:"rpm_2" db:"rpm_2"`
	Time1       time.Duration `json:"time_1" db:"time_1"`
	Time2       time.Duration `json:"time_2" db:"time_2"`
	CreatedAt   time.Time     `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time     `json:"updated_at" db:"updated_at"`
}

const (
	getShakingQuery    = `SELECT * FROM shaking where process_id = $1`
	createShakingQuery = `INSERT INTO shaking (
		with_temp,
		temperature,
		follow_temp,
		rpm_1,
		rpm_2,
		time_1,
		time_2,
		process_id)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id`
)

func (s *pgStore) ShowShaking(ctx context.Context, shakingID uuid.UUID) (shaking Shaking, err error) {

	err = s.db.Get(&shaking, getShakingQuery, shakingID)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error getting shaking data")
		return
	}

	return
}

func (s *pgStore) CreateShaking(ctx context.Context, sh Shaking) (createdShaking Shaking, err error) {
	var lastInsertID uuid.UUID

	err = s.db.QueryRow(
		createShakingQuery,
		sh.WithTemp,
		sh.Temperature,
		sh.FollowTemp,
		sh.RPM1,
		sh.RPM2,
		sh.Time1,
		sh.Time2,
		sh.ProcessID,
	).Scan(&lastInsertID)

	if err != nil {
		logger.WithField("err", err.Error()).Error("Error creating Shaking")
		return
	}

	err = s.db.Get(&createdShaking, getShakingQuery, sh.ProcessID)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error in getting Shaking")
		return
	}
	return
}