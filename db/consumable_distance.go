package db

import (
	"context"
	"fmt"
	"strings"

	logger "github.com/sirupsen/logrus"
)

const (
	insertConsDistaceQuery1 = `INSERT INTO consumable_distances(
							id,
							name,
							distance,
							description)
							VALUES %s `
	insertConsDistaceQuery2 = ` ON CONFLICT DO NOTHING;`
)

type ConsumableDistance struct {
	ID          int     `db:"id"`
	Name        string  `db:"name"`
	Distance    float64 `db:"distance"`
	Description string  `db:"description"`
}

func (s *pgStore) InsertConsumableDistance(ctx context.Context, consumabledistances []ConsumableDistance) (err error) {
	stmt := makeConsumableQuery(consumabledistances)

	_, err = s.db.Exec(
		stmt,
	)
	if err != nil {
		logger.WithField("error in exec query", err.Error()).Error("Query Failed")
		return
	}
	return
}

func makeConsumableQuery(consumabledistance []ConsumableDistance) string {
	values := make([]string, 0, len(consumabledistance))

	for _, c := range consumabledistance {
		values = append(values, fmt.Sprintf("(%v, '%v', %v, '%v') ", c.ID, c.Name, c.Distance, c.Description))
	}

	stmt := fmt.Sprintf(insertConsDistaceQuery1,
		strings.Join(values, ","))

	stmt += insertConsDistaceQuery2

	return stmt
}
