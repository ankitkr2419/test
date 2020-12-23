package db

import (
	"context"
	"fmt"
	"strings"

	logger "github.com/sirupsen/logrus"
)

const (
	insertTipsTubesQuery1 = `INSERT INTO tips_and_tubes(
							labware_id,
							name,
							volume,
							height)
							VALUES %s `
	insertTipsTubesQuery2 = `ON CONFLICT DO NOTHING;`
)

type TipsTubes struct {
	LabwareID int     `db:"labware_id"`
	Name      string  `db:"name"`
	Volume    float64 `db:"volume"`
	Height    float64 `db:"height"`
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
		values = append(values, fmt.Sprintf("(%v, '%v', %v, %v)", t.LabwareID, t.Name, t.Volume, t.Height))
	}

	stmt := fmt.Sprintf(insertTipsTubesQuery1,
		strings.Join(values, ","))

	stmt += insertTipsTubesQuery2

	return stmt
}
