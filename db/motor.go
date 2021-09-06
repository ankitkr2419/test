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
	updateMotorQuery = `update motors set
		deck = $1,
		number = $2,
		name = $3,
		ramp = $4,
		steps = $5,
		slow = $6,
		fast = $7, 
		updated_at = $8 where id = $9 `
	deleteMotorQuery = `DELETE FROM motors WHERE id = $1`
)

type Motor struct {
	ID        int       `db:"id" json:"id" validate:"lte=20,gte=0"`
	Deck      string    `db:"deck" json:"deck"`
	Number    int       `db:"number" json:"number" validate:"lte=10,gte=1"`
	Name      string    `db:"name" json:"name"`
	Ramp      int       `db:"ramp" json:"ramp" validate:"lte=3000,gte=1"`
	Steps     int       `db:"steps" json:"steps" validate:"lte=3000,gte=1"`
	Slow      int       `db:"slow" json:"slow" validate:"lte=9000,gte=100"`
	Fast      int       `db:"fast" json:"fast" validate:"lte=16000,gte=100"`
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

func (s *pgStore) ListMotors(ctx context.Context) (motors []Motor, err error) {
	err = s.db.Select(&motors, selectMotorsQuery)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error fetching motors data")
		return
	}

	return
}

func (s *pgStore) UpdateMotor(ctx context.Context, motor Motor) (err error) {
	_, err = s.db.Exec(updateMotorQuery, motor.Deck,
		motor.Number,
		motor.Name,
		motor.Ramp,
		motor.Steps,
		motor.Slow,
		motor.Fast,
		time.Now(),
		motor.ID)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error updating Motor data")
		return
	}
	return

}

func (s *pgStore) DeleteMotor(ctx context.Context, id int) (err error) {
	_, err = s.db.Exec(deleteMotorQuery, id)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error deleting Motor data")
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
