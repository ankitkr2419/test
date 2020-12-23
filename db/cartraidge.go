package db

import (
	"context"
	"fmt"
	"strings"

	logger "github.com/sirupsen/logrus"
)

const (
	insertCartraidgeQuery1 = `INSERT INTO cartridges(
							labware_id,
							type,
							description,
							wells,
							distances,
							heights,
							volumes)
							VALUES %s `
	insertCartraidgeQuery2 = `ON CONFLICT DO NOTHING;`
)

type Cartraidge struct {
	LabwareID   int       `db:"labware_id"`
	Type        string    `db:"type"`
	Description string    `db:"description"`
	Wells       int       `db:"wells"`
	Distances   []float64 `db:"distances"`
	Heights     []float64 `db:"heights"`
	Volumes     []float64 `db:"volumes"`
}

func (s *pgStore) InsertCartraidge(ctx context.Context, cartraidges []Cartraidge) (err error) {
	stmt := makeCartraidgeQuery(cartraidges)

	_, err = s.db.Exec(
		stmt,
	)
	if err != nil {
		logger.WithField("error in exec query", err.Error()).Error("Query Failed")
		return
	}
	return
}

func makeCartraidgeQuery(cartraidge []Cartraidge) string {
	values := make([]string, 0, len(cartraidge))

	for _, c := range cartraidge {
		distances, heights, volumes := []string{}, []string{}, []string{}
		for i := 0; i < c.Wells; i++ {
			distances = append(distances, fmt.Sprintf("%v", c.Distances[i]))
			heights = append(heights, fmt.Sprintf("%v", c.Heights[i]))
			volumes = append(volumes, fmt.Sprintf("%v", c.Volumes[i]))
		}

		d := "{" + strings.Join(distances, ",") + "}"
		h := "{" + strings.Join(heights, ",") + "}"
		v := "{" + strings.Join(volumes, ",") + "}"

		values = append(values, fmt.Sprintf("(%v, '%v', '%v', %v, '%v', '%v', '%v')", c.LabwareID, c.Type, c.Description, c.Wells, d, h, v))
	}

	stmt := fmt.Sprintf(insertCartraidgeQuery1,
		strings.Join(values, ","))

	stmt += insertCartraidgeQuery2

	return stmt
}
