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
		WHEREwell_position = $1 AND experiment_id =$2`

	getWellTargetsQuery = `SELECT
		well_position,
		target_id,
		ct
		FROM well_targets WHERE well_position = $1 AND experiment_id =$2`

	deleteWellTargetQuery = `DELETE FROM well_targets WHERE experiment_id = $1 AND well_position IN (%s)`

	insertWellTargetQuery = `INSERT INTO well_targets (
		experiment_id,
		well_position,
		target_id)
		VALUES %s `
	insertWellTargetQuery2 = `ON CONFLICT (well_position, experiment_id,target_id) DO UPDATE
			SET
			ct = excluded.ct
			WHERE well_targets.position = excluded.position AND well_targets.experiment_id = excluded.experiment_id AND 				well_targets.target_id = excluded.target_id AND well_targets.CT != ''`

	getWellTargetListQuery = `SELECT wt.experiment_id,
		wt.well_position,
		wt.target_id,
		wt.ct,
		t.name as target_name
		FROM well_targets as wt , targets as t
		WHERE
		wt.target_id = t.id
		AND
		wt.experiment_id = $1
		AND
		wt.well_position IN (%s) `
	getWellTargetExpListQuery = `SELECT wt.experiment_id,
		wt.well_position,
		wt.target_id,
		wt.ct,
		t.name as target_name
		FROM well_targets as wt , targets as t
		WHERE
		wt.target_id = t.id
		AND
		wt.experiment_id = $1`
)

type WellTarget struct {
	ExperimentID uuid.UUID `db:"experiment_id" json:"experiment_id"`
	WellPosition int32     `db:"well_position" json:"well_position"`
	TargetID     uuid.UUID `db:"target_id" json:"target_id" validate:"required"`
	TargetName   string    `db:"target_name" json:"target_name"`
	CT           string    `db:"ct" json:"ct"`
}

func (s *pgStore) GetWellTarget(ctx context.Context, wellposition int32, experimentID uuid.UUID) (WellTargets []WellTarget, err error) {
	err = s.db.Select(&WellTargets, getWellTargetsListQuery, wellposition, experimentID)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error listing WellTargets")
		return
	}

	return
}

func (s *pgStore) UpsertWellTargets(ctx context.Context, WellTargets []WellTarget, experimentID uuid.UUID) (targets []WellTarget, err error) {

	stmt := makeWellTargetQuery(WellTargets, experimentID)
	// delstmt := WellTargetQuery(WellTargets, deleteWellTargetQuery)
	getstmt := WellTargetQuery(WellTargets, getWellTargetListQuery)
	fmt.Println(stmt)
	//fmt.Println(delstmt)
	fmt.Println(getstmt)
	tx, err := s.db.Begin()
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error in creating transaction")
		return
	}
	/*_, err = tx.Exec(
		delstmt,experimentID,
	)
	if err != nil {
		tx.Rollback()
		logger.WithField("err", err.Error()).Error("Error deleting previous well targets")
		return
	}*/

	_, err = tx.Exec(
		stmt,
	)
	if err != nil {
		tx.Rollback()
		logger.WithField("error in exec query", err.Error()).Error("Query Failed")
		return
	}

	rows, err := tx.Query(
		getstmt, experimentID,
	)

	for rows.Next() {
		var r WellTarget
		if err := rows.Scan(&r.ExperimentID, &r.WellPosition, &r.TargetID, &r.CT, &r.TargetName); err != nil {
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

func (s *pgStore) ListWellTargets(ctx context.Context, experimentID uuid.UUID) (WellTargets []WellTarget, err error) {

	err = s.db.Select(&WellTargets, getWellTargetExpListQuery, experimentID)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error listing WellTargets")
		return
	}

	return
}

// prepare bulk insert query statement
func makeWellTargetQuery(WellTargets []WellTarget, experimentID uuid.UUID) string {

	values := make([]string, 0, len(WellTargets))

	for _, t := range WellTargets {
		values = append(values, fmt.Sprintf("('%v',%v,'%v')", experimentID, t.WellPosition, t.TargetID))
	}

	stmt := fmt.Sprintf(insertWellTargetQuery,
		strings.Join(values, ","))

	stmt += insertWellTargetQuery2

	return stmt
}

// prepare bulk delete/select query statement
func wellTargetQuery(WellTargets []WellTarget, q string) string {

	values := make([]string, 0, len(WellTargets))

	for _, t := range WellTargets {
		values = append(values, fmt.Sprintf("%v", t.WellPosition))
	}

	stmt := fmt.Sprintf(q,
		strings.Join(values, ","))

	return stmt
}

// prepare bulk insert query statement
func getWellTargetQuery(WellTargets []int32, q string) string {

	values := make([]string, 0, len(WellTargets))

	for _, t := range WellTargets {
		values = append(values, fmt.Sprintf("%v", t))
	}

	stmt := fmt.Sprintf(q,
		strings.Join(values, ","))

	return stmt
}
