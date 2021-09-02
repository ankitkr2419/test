package db

import (
	"context"
	"fmt"
	"strings"
	"time"

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

	getAllConsDistanceQuery = `SELECT *
							FROM consumable_distances`
)

type ConsumableDistance struct {
	ID          int       `db:"id" json:"id"`
	Name        string    `db:"name" json:"name"`
	Distance    float64   `db:"distance" json:"distance"`
	Description string    `db:"description" json:"description"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
}

func (s *pgStore) InsertConsumableDistance(ctx context.Context, consumabledistances []ConsumableDistance) (err error) {
	stmt := makeConsumableQuery(consumabledistances)
	logger.Debugln("Consumables Query: ", stmt)
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

func (s *pgStore) ListConsDistances() (consumabledistance []ConsumableDistance, err error) {
	err = s.db.Select(&consumabledistance, getAllConsDistanceQuery)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error listing consumable distance details")
		return
	}
	return
}
