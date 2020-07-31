package db

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	logger "github.com/sirupsen/logrus"
)

const (
	insertExpTempQuery = `INSERT INTO experiment_temperatures (
		experiment_id,
		temp,
		lid_temp,
		cycle)
		VALUES `

	getExpTempQuery = `SELECT * FROM experiment_temperatures
		WHERE
		experiment_id = $1
		ORDER BY created_at ASC`
)

// ExperimentTemperature stores temp as it increases
type ExperimentTemperature struct {
	ExperimentID uuid.UUID `db:"experiment_id" json:"experiment_id"`
	Temp         float32   `db:"temp" json:"temp"`
	LidTemp      float32   `db:"lid_temp" json:"lid_temp"`
	Cycle        uint16    `db:"cycle" json:"cycle"`
	CreatedAt    time.Time `db:"created_at" json:"created_at"`
	UpdatedAt    time.Time `db:"updated_at" json:"updated_at"`
}

func (s *pgStore) ListExperimentTemperature(ctx context.Context, experimentID uuid.UUID) (t []ExperimentTemperature, err error) {
	err = s.db.Select(&t, getExpTempQuery, experimentID)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error listing result temperature details")
		return
	}
	return
}

func (s *pgStore) InsertExperimentTemperature(ctx context.Context, r ExperimentTemperature) (err error) {

	stmt := makeInsertTempQuery(r)
	_, err = s.db.Exec(
		stmt,
	)
	if err != nil {
		logger.WithField("error in exec query", err.Error()).Error("Query Failed")
		return
	}

	return
}

func makeInsertTempQuery(r ExperimentTemperature) (stmt string) {

	values := fmt.Sprintf("('%v', %v,%v,%v)", r.ExperimentID, r.Temp, r.LidTemp, r.Cycle)

	stmt = fmt.Sprintf(insertExpTempQuery + values)

	return
}
