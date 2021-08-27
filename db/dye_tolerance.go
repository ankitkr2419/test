package db

import (
	"context"
	"fmt"
	"strings"

	"github.com/google/uuid"
	logger "github.com/sirupsen/logrus"
)

type DyeWellTolerance struct {
	ID            uuid.UUID `db:"id" json:"id"`
	DyeID         uuid.UUID `db:"dye_id" json:"dye_id"`
	WellNo        int       `db:"well_no" json:"well_no"`
	KitID         string    `db:"kit_id" json:"kit_id"`
	OpticalResult float64   `db:"optical_result" json:"optical_result"`
}

const (
	insertDyeWellToleranceQuery1 = ` INSERT INTO dye_well_tolerance(
		dye_id,
		well_no,
		kit_id,
		optical_result)
		VALUES %s`
	insertDyeWellToleranceQuery2 = ` ON CONFLICT ( dye_id , kit_id ) DO UPDATE
	SET optical_result = excluded.optical_result                                                                             
	where dye_well_tolerance.dye_id = excluded.dye_id AND dye_well_tolerance. kit_id = excluded. kit_id ;`
)

func (s *pgStore) InsertDyeWellTolerance(ctx context.Context, dwTolerance []DyeWellTolerance) (result DyeWellTolerance, err error) {
	stmt := makeDyeWellToleranceQuery(dwTolerance)

	logger.Debugln("tolerance Query: ", stmt)
	_, err = s.db.Exec(
		stmt,
	)
	if err != nil {
		logger.WithField("error in exec query", err.Error()).Error("Query Failed")
		return
	}
	return
}

func makeDyeWellToleranceQuery(dwTolerance []DyeWellTolerance) string {
	values := make([]string, 0, len(dwTolerance))

	for _, c := range dwTolerance {
		values = append(values, fmt.Sprintf("('%v', %v,'%v',%v) ", c.DyeID, c.WellNo, c.KitID, c.OpticalResult))
	}

	stmt := fmt.Sprintf(insertDyeWellToleranceQuery1,
		strings.Join(values, ","))

	stmt += insertConsDistaceQuery2

	return stmt
}
