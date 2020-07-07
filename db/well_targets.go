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

	getWellTargetListQuery = `SELECT wt.well_id,
		wt.target_id,
		wt.ct,
		t.name as target_name
		FROM well_targets as wt , targets as t
		WHERE
		wt.target_id = t.id
		AND
		well_id IN (%s)`
)

type WellTarget struct {
	WellID     uuid.UUID `db:"well_id" json:"well_id"`
	TargetID   uuid.UUID `db:"target_id" json:"target_id" validate:"required"`
	TargetName string    `db:"target_name" json:"target_name"`
	CT         string    `db:"ct" json:"ct"`
}

func (s *pgStore) GetWellTarget(ctx context.Context, wellID uuid.UUID) (WellTargets []WellTarget, err error) {
	err = s.db.Select(&WellTargets, getWellTargetsListQuery, wellID)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error listing WellTargets")
		return
	}

	return
}

func (s *pgStore) UpsertWellTargets(ctx context.Context, WellTargets []WellTarget) (targets []WellTarget, err error) {

	stmt := makeWellTargetQuery(WellTargets)
	delstmt := WellTargetQuery(WellTargets, deleteWellTargetQuery)
	getstmt := WellTargetQuery(WellTargets, getWellTargetListQuery)

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

	rows,err := tx.Query(
		getstmt,
	)

       for rows.Next() {
		var r WellTarget
		if err := rows.Scan(&r.WellID,&r.TargetID,&r.CT,&r.TargetName); err != nil {
			logger.WithField("err", err.Error()).Error("Error getting new well targets")
		}
		targets = append(targets, r)
	}

	if err != nil {
		tx.Rollback()
		logger.WithField("err", err.Error()).Error("Error getting new well targets")
		return
	}

	tx.Commit()
	return
}

func (s *pgStore) ListWellTargets(ctx context.Context, wellID []uuid.UUID) (WellTargets []WellTarget, err error) {
	stmt := getWellTargetQuery(wellID, getWellTargetListQuery)

	err = s.db.Select(&WellTargets, stmt)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error listing WellTargets")
		return
	}

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

// prepare bulk delete/select query statement
func WellTargetQuery(WellTargets []WellTarget, q string) string {

	values := make([]string, 0, len(WellTargets))

	for _, t := range WellTargets {
		values = append(values, fmt.Sprintf("'%v'", t.WellID))
	}

	stmt := fmt.Sprintf(q,
		strings.Join(values, ","))

	return stmt
}

// prepare bulk insert query statement
func getWellTargetQuery(WellTargets []uuid.UUID, q string) string {

	values := make([]string, 0, len(WellTargets))

	for _, t := range WellTargets {
		values = append(values, fmt.Sprintf("'%v'", t))
	}

	stmt := fmt.Sprintf(q,
		strings.Join(values, ","))

	return stmt
}
