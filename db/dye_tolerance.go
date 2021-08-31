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
	DyeID         uuid.UUID `db:"dye_id" json:"dye_id" validate:"required"`
	WellNo        int       `db:"well_no" json:"well_no"`
	KitID         string    `db:"kit_id" json:"kit_id" validate:"required,len=8"`
	OpticalResult float64   `db:"optical_result" json:"optical_result"`
	Valid         bool      `db:"valid" json:"valid"`
}

const (
	insertDyeWellToleranceQuery1 = ` INSERT INTO dye_well_tolerance(
		dye_id,
		well_no,
		kit_id,
		optical_result,
		valid)
		VALUES %s`
	insertDyeWellToleranceQuery2 = ` ON CONFLICT (dye_id,well_no) DO UPDATE
	SET kit_id = excluded.kit_id,
	optical_result = excluded.optical_result,
	valid = excluded.valid                                                                            
	where dye_well_tolerance.dye_id = excluded.dye_id AND dye_well_tolerance.well_no = excluded.well_no;`
)

func (s *pgStore) UpsertDyeWellTolerance(ctx context.Context, dwTolerance []DyeWellTolerance) (err error) {
	stmt := makeDyeWellToleranceQuery(dwTolerance)

	logger.Infoln("tolerance Query: ", stmt)
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
		values = append(values, fmt.Sprintf("('%v', %v,'%v',%v,%v) ", c.DyeID, c.WellNo, c.KitID, c.OpticalResult, c.Valid))
	}

	stmt := fmt.Sprintf(insertDyeWellToleranceQuery1,
		strings.Join(values, ","))

	stmt += insertDyeWellToleranceQuery2

	return stmt
}
