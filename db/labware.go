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
	getallLabwares      = `SELECT *
						FROM labwares`
)

type Labware struct {
	ID          int    `db:"id" json:"id"`
	Name        string `db:"name" json:"name"`
	Description string `db:"description" json:"description"`
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

func (s *pgStore) GetAllLabwares() (labwares []Labware, err error) {
	err = s.db.Select(&labwares, getallLabwares)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error listing labware details")
		return
	}
	return
}
