package db

import (
	"context"
	"fmt"
	"strings"

	logger "github.com/sirupsen/logrus"
)

const (
	insertLabwareQuery1 = `INSERT INTO labwares(
							id,
							name,
							description)
							VALUES %s `
	insertLabwareQuery2 = `ON CONFLICT DO NOTHING;`
)

type Labware struct {
	ID          int    `db:"id"`
	Name        string `db:"name"`
	Description string `db:"description"`
}

func (s *pgStore) InsertLabware(ctx context.Context, labwares []Labware) (err error) {
	stmt := makeLabwareQuery(labwares)

	_, err = s.db.Exec(
		stmt,
	)
	if err != nil {
		logger.WithField("error in exec query", err.Error()).Error("Query Failed")
		return
	}
	return
}

func makeLabwareQuery(labware []Labware) string {
	values := make([]string, 0, len(labware))

	for _, l := range labware {
		values = append(values, fmt.Sprintf("(%v, '%v', '%v')", l.ID, l.Name, l.Description))
	}

	stmt := fmt.Sprintf(insertLabwareQuery1,
		strings.Join(values, ","))

	stmt += insertLabwareQuery2

	return stmt
}
