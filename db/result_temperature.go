package db

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	logger "github.com/sirupsen/logrus"
)

const (
	insertResultTempQuery = `INSERT INTO result_temperatures (
		experiment_id,
		temp,
		lid_temp,
		cycle)
		VALUES `

	getResultTempQuery = `SELECT * FROM result_temperatures
		WHERE
		experiment_id = $1`
)

// ResultTemperature stores temp as it increases
type ResultTemperature struct {
	ExperimentID uuid.UUID `db:"experiment_id" json:"experiment_id"`
	Temp         float32   `db:"temp" json:"temp"`
	LidTemp      float32   `db:"lid_temp" json:"lid_temp"`
	Cycle        uint16    `db:"cycle" json:"cycle"`
	CreatedAt    time.Time `db:"created_at" json:"created_at"`
}

func (s *pgStore) ListResultTemperature(ctx context.Context, experimentID uuid.UUID) (t []ResultTemperature, err error) {
	err = s.db.Select(&t, getResultTempQuery, experimentID)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error listing result temperature details")
		return
	}
	return
}

func (s *pgStore) InsertResultTemperature(ctx context.Context, r ResultTemperature) (err error) {

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

func makeInsertTempQuery(r ResultTemperature) (stmt string) {

	values := fmt.Sprintf("('%v', %v,%v,%v)", r.ExperimentID, r.Temp, r.LidTemp, r.Cycle)

	stmt = fmt.Sprintf(insertResultTempQuery + values)

	return
}
