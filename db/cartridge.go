package db

import (
	"context"
	"fmt"
	"strings"
	"time"

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
	insertCartridgeQuery2   = `ON CONFLICT DO NOTHING;`
	selectAllCartridgeQuery = `SELECT *
							FROM cartridges`
)

type Cartridge struct {
	ID          int       `db:"id" json:"id"`
	LabwareID   int       `db:"labware_id" json:"labware_id"`
	Type        string    `db:"type" json:"type"`
	Description string    `db:"description" json:"description"`
	WellNum     int       `db:"well_num" json:"well_num"`
	Distance    float64   `db:"distance" json:"distance"`
	Height      float64   `db:"height" json:"height"`
	Volume      float64   `db:"volume" json:"volume"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
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

func (s *pgStore) ListCartridges() (cartridge []Cartridge, err error) {
	err = s.db.Select(&cartridge, selectAllCartridgeQuery)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error listing cartridge details")
		return
	}
	return
}
