package db

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/lib/pq"

	logger "github.com/sirupsen/logrus"
)

const (
	insertTipsTubesQuery1 = `INSERT INTO tips_and_tubes(
							id,
							name,
							type,
							allowed_positions,
							volume,
							height)
							VALUES %s `
	insertTipsTubesQuery2 = `ON CONFLICT DO NOTHING;`
	getTipsTubesQuery     = `SELECT *
							FROM tips_and_tubes`
)

type TipsTubes struct {
	ID               int64         `db:"id" json:"id"`
	Name             string        `db:"name" json:"name"`
	Type             string        `db:"type" json:"type"`
	AllowedPositions pq.Int64Array `db:"allowed_positions" json:"allowed_positions"`
	Volume           float64       `db:"volume" json:"volume"`
	Height           float64       `db:"height" json:"height"`
	CreatedAt        time.Time     `db:"created_at" json:"created_at"`
	UpdatedAt        time.Time     `db:"updated_at" json:"updated_at"`
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
		positions := "{"
		for _, pos := range t.AllowedPositions {
			positions = fmt.Sprintf("%s%v,", positions, pos)
		}
		positions = positions[:len(positions)-1] + "}"
		values = append(values, fmt.Sprintf("(%v, '%v', '%v', '%v', %v, %v)", t.ID, t.Name, t.Type, positions, t.Volume, t.Height))
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
