package db

import (
	"context"
	"fmt"
	"strings"

	"github.com/google/uuid"
	logger "github.com/sirupsen/logrus"
)

const (
	getWellTargetsListQuery = `SELECT * FROM well_targets
		WHERE well = $1`

	getWellTargetsQuery = `SELECT
		well_id,
		target_id,
		ct
		FROM well_targets WHERE well_id = $1`

	deleteWellTargetQuery = `DELETE FROM well_targets WHERE well_id IN (%s)`

	insertWellTargetQuery = `INSERT INTO well_targets (
		well_id,
		target_id)
		VALUES %s`
)

type WellTarget struct {
	WellID   uuid.UUID `db:"well_id" json:"well_id"`
	TargetID uuid.UUID `db:"target_id" json:"target_id" validate:"required"`
	CT       string    `db:"ct" json:"ct"`
}

func (s *pgStore) ListWellTargets(ctx context.Context, wellID uuid.UUID) (WellTargets []WellTarget, err error) {
	err = s.db.Select(&WellTargets, getWellTargetsListQuery, wellID)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error listing WellTargets")
		return
	}

	return
}

func (s *pgStore) UpsertWellTargets(ctx context.Context, WellTargets []WellTarget) (err error) {

	stmt := makeWellTargetQuery(WellTargets)
	delstmt := delWellTargetQuery(WellTargets)

	tx, err := s.db.Begin()
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error in creating transaction")
		return
	}
	_, err = tx.Exec(
		delstmt,
	)
	if err != nil {
		tx.Rollback()
		logger.WithField("err", err.Error()).Error("Error deleting previous well targets")
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
	return
}

// prepare bulk insert query statement
func makeWellTargetQuery(WellTargets []WellTarget) string {

	values := make([]string, 0, len(WellTargets))

	for _, t := range WellTargets {
		values = append(values, fmt.Sprintf("('%v','%v')", t.WellID, t.TargetID))
	}

	stmt := fmt.Sprintf(insertWellTargetQuery,
		strings.Join(values, ","))

	return stmt
}

// prepare bulk insert query statement
func delWellTargetQuery(WellTargets []WellTarget) string {

	values := make([]string, 0, len(WellTargets))

	for _, t := range WellTargets {
		values = append(values, fmt.Sprintf("'%v'", t.WellID))
	}

	stmt := fmt.Sprintf(deleteWellTargetQuery,
		strings.Join(values, ","))

	return stmt
}
