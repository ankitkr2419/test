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
							FROM consumable_distances order by id`
	getConsDistanceQueryByDeck = `SELECT *
							FROM consumable_distances where id BETWEEN $1 AND $2 order by id`
	updateConsDistaceQuery1 = `UPDATE consumable_distances SET 
	distance = $1,
	description = $2 WHERE id = $3`
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

func (s *pgStore) UpdateConsumableDistance(ctx context.Context, c ConsumableDistance) (err error) {

	_, err = s.db.Exec(
		updateConsDistaceQuery1, c.Distance, c.Description, c.ID,
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

func (s *pgStore) ListConsDistancesDeck(ctx context.Context, min, max int64) (consumabledistance []ConsumableDistance, err error) {
	err = s.db.Select(&consumabledistance, getConsDistanceQueryByDeck, min, max)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error listing consumable distance details")
		return
	}
	return
}
