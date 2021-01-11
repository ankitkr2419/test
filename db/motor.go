package db

import (
	"context"
	"fmt"
	"strings"
	"time"

	logger "github.com/sirupsen/logrus"
)

const (
	insertMotorQuery1 = `INSERT INTO motors(
							id,
							deck,
							number,
							name,
							ramp,
							steps,
							slow,
							fast)
							VALUES %s `
	insertMotorQuery2 = `ON CONFLICT DO NOTHING;`
	selectMotorsQuery = `SELECT * 
						FROM motors`
)

type Motor struct {
	ID        int       `db:"id" json:"id"`
	Deck      string    `db:"deck" json:"deck"`
	Number    int       `db:"number" json:"number"`
	Name      string    `db:"name" json:"name"`
	Ramp      int       `db:"ramp" json:"ramp"`
	Steps     int       `db:"steps" json:"steps"`
	Slow      int       `db:"slow" json:"slow"`
	Fast      int       `db:"fast" json:"fast"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

func (s *pgStore) InsertMotor(ctx context.Context, motors []Motor) (err error) {
	stmt := makeMotorQuery(motors)

	_, err = s.db.Exec(
		stmt,
	)
	if err != nil {
		logger.WithField("error in exec query", err.Error()).Error("Query Failed")
		return
	}
	return
}

func makeMotorQuery(motor []Motor) string {
	values := make([]string, 0, len(motor))

	for _, m := range motor {
		values = append(values, fmt.Sprintf("(%v, '%v', %v, '%v', %v, %v, %v, %v)", m.ID, m.Deck, m.Number, m.Name, m.Ramp, m.Steps, m.Slow, m.Fast))
	}

	stmt := fmt.Sprintf(insertMotorQuery1,
		strings.Join(values, ","))

	stmt += insertMotorQuery2

	return stmt
}

func (s *pgStore) GetAllMotors() (motor []Motor, err error) {
	err = s.db.Select(&motor, selectMotorsQuery)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error listing well details")
		return
	}
	return
}
