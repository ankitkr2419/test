package db

import (
	"context"
	"fmt"
	"strings"

	"github.com/google/uuid"
	logger "github.com/sirupsen/logrus"
)

const (
	getTargetListQuery = `SELECT targets.* FROM targets
			ORDER BY name ASC`
	getTargetListNameQuery = `SELECT * FROM targets
	where name = $1 LIMIT 1`

	insertTargetsQuery1 = `INSERT INTO targets(
				name,
				dye_id)
				VALUES %s `
	insertTargetsQuery2 = `ON CONFLICT DO NOTHING;`

	fetchTargetDyeQuery          = `SELECT d.Name as dye FROM targets as t ,dyes as d WHERE t.dye_id = d.id AND t.id = $1`
	upsertExpTargThresholdQuery1 = `INSERT INTO exp_target_threshold (
		exp_id,
		target_id,
		threshold)
		VALUES %s`

	upsertExpTargThresholdQuery2 = ` ON CONFLICT (exp_id, target_id,threshold) DO UPDATE                           
	SET threshold=excluded.threshold                                                                            
	where exp_target_threshold.exp_id = excluded.exp_id and exp_target_threshold.target_id = excluded.target_id`
)

type Target struct {
	ID    uuid.UUID `db:"id" json:"id"`
	Name  string    `db:"name" json:"name" validate:"required"`
	DyeID uuid.UUID `db:"dye_id" json:"dye_id" validate:"required"`
}
type ExpTargetThreshold struct {
	ExperimentID uuid.UUID `db:"exp_id" json:"exp_id" validate:"required"`
	TargetID     uuid.UUID `db:"target_id" json:"target_id" validate:"required"`
	Threshold    float32   `db:"threshold" json:"threshold" validate:"required"`
}

func (s *pgStore) ListTargetDye(ctx context.Context, targetID uuid.UUID) (dye string, err error) {
	err = s.db.Get(&dye, fetchTargetDyeQuery, targetID)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error listing targets dye")
		return
	}
	return
}
func (s *pgStore) ListTargets(ctx context.Context) (t []Target, err error) {
	err = s.db.Select(&t, getTargetListQuery)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error listing targets")
		return
	}

	return
}

func (s *pgStore) GetTargetByName(ctx context.Context, name string) (t Target, err error) {
	err = s.db.Get(&t, getTargetListNameQuery, name)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error listing targets by name")
		return
	}
	return
}
func (s *pgStore) UpsertTargetThreshold(ctx context.Context, tt []ExpTargetThreshold) (err error) {

	for _, v := range tt {
		stmt := makeTargetThresoldQuery(v)
		logger.Infof(stmt)

		_, err = s.db.Exec(stmt)
		if err != nil {
			logger.WithField("err", err.Error()).Error("Error upserting target threshold")
			return
		}
	}
	return
}
func (s *pgStore) InsertTargets(ctx context.Context, Targets []Target) (err error) {

	stmt := makeTargetQuery(Targets)

	_, err = s.db.Exec(
		stmt,
	)
	if err != nil {
		logger.WithField("error in exec query", err.Error()).Error("Query Failed")
		return
	}

	return
}

// prepare bulk insert query statement
func makeTargetQuery(Targets []Target) string {

	values := make([]string, 0, len(Targets))

	for _, t := range Targets {
		values = append(values, fmt.Sprintf("('%v', '%v')", t.Name, t.DyeID))
	}

	stmt := fmt.Sprintf(insertTargetsQuery1,
		strings.Join(values, ","))

	stmt += insertTargetsQuery2

	return stmt
}

func makeTargetThresoldQuery(tt ExpTargetThreshold) string {

	values := make([]string, 0, 1)

	values = append(values, fmt.Sprintf("('%v','%v', %v)", tt.ExperimentID, tt.TargetID, tt.Threshold))

	stmt := fmt.Sprintf(upsertExpTargThresholdQuery1,
		strings.Join(values, ","))

	stmt += upsertExpTargThresholdQuery2
	return stmt
}
