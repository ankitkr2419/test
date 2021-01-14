package db

import (
	"context"
	"fmt"
	"strings"
	"time"

	logger "github.com/sirupsen/logrus"
)

const (
	insertTipsTubesQuery1 = `INSERT INTO tips_and_tubes(
							labware_id,
							consumable_distance_id,
							name,
							volume,
							height)
							VALUES %s `
	insertTipsTubesQuery2 = `ON CONFLICT DO NOTHING;`
	getTipsTubesQuery     = `SELECT *
							FROM tips_and_tubes`
)

type TipsTubes struct {
	LabwareID            int       `db:"labware_id" json:"labware_id"`
	ConsumabledistanceID int       `db:"consumable_distance_id" json:"consumable_distance_id"`
	Name                 string    `db:"name" json:"name"`
	Volume               float64   `db:"volume" json:"volume"`
	Height               float64   `db:"height" json:"height"`
	CreatedAt            time.Time `db:"created_at" json:"created_at"`
	UpdatedAt            time.Time `db:"updated_at" json:"updated_at"`
}

func (s *pgStore) InsertTipsTubes(ctx context.Context, tipstubes []TipsTubes) (err error) {
	stmt := makeTipsTubesQuery(tipstubes)

	_, err = s.db.Exec(
		stmt,
	)
	if err != nil {
		logger.WithField("error in exec query", err.Error()).Error("Query Failed")
		return
	}
	return
}

func makeTipsTubesQuery(tipstubes []TipsTubes) string {
	values := make([]string, 0, len(tipstubes))

	for _, t := range tipstubes {
		values = append(values, fmt.Sprintf("(%v, %v, '%v', %v, %v)", t.LabwareID, t.ConsumabledistanceID, t.Name, t.Volume, t.Height))
	}

	stmt := fmt.Sprintf(insertTipsTubesQuery1,
		strings.Join(values, ","))

	stmt += insertTipsTubesQuery2

	return stmt
}

func (s *pgStore) ListTipsTubes() (tipstubes []TipsTubes, err error) {
	err = s.db.Select(&tipstubes, getTipsTubesQuery)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error listing tipstubes details")
		return
	}
	return
}
