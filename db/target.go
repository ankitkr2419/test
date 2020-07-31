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

	insertTargetsQuery1 = `INSERT INTO targets(
				name,
				dye_id)
				VALUES %s `
	insertTargetsQuery2 = `ON CONFLICT DO NOTHING;`
)

type Target struct {
	ID    uuid.UUID `db:"id" json:"id"`
	Name  string    `db:"name" json:"name" validate:"required"`
	DyeID uuid.UUID `db:"dye_id" json:"dye_id" validate:"required"`
}

func (s *pgStore) ListTargets(ctx context.Context) (t []Target, err error) {
	err = s.db.Select(&t, getTargetListQuery)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error listing targets")
		return
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
