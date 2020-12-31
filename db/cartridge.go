package db

import (
	"context"
	"fmt"
	"strings"

	logger "github.com/sirupsen/logrus"
)

const (
	insertCartridgeQuery1 = `INSERT INTO cartridges(
							id,
							labware_id,
							type,
							description,
							well_num,
							distance,
							height,
							volume)
							VALUES %s `
	insertCartridgeQuery2 = `ON CONFLICT DO NOTHING;`
)

type Cartridge struct {
	ID          int     `db:"id"`
	LabwareID   int     `db:"labware_id"`
	Type        string  `db:"type"`
	Description string  `db:"description"`
	WellNum     int     `db:"wells_num"`
	Distance    float64 `db:"distance"`
	Height      float64 `db:"height"`
	Volume      float64 `db:"volume"`
}

func (s *pgStore) InsertCartridge(ctx context.Context, cartridges []Cartridge) (err error) {
	stmt := makeCartridgeQuery(cartridges)

	_, err = s.db.Exec(
		stmt,
	)
	if err != nil {
		logger.WithField("error in exec query", err.Error()).Error("Query Failed")
		return
	}
	return
}

func makeCartridgeQuery(cartridge []Cartridge) string {
	values := make([]string, 0, len(cartridge))

	for _, c := range cartridge {
		values = append(values, fmt.Sprintf("(%v, %v, '%v', '%v', %v, %v,  %v, %v)", c.ID, c.LabwareID, c.Type, c.Description, c.WellNum, c.Distance, c.Height, c.Volume))
	}

	stmt := fmt.Sprintf(insertCartridgeQuery1,
		strings.Join(values, ","))

	stmt += insertCartridgeQuery2

	return stmt
}
