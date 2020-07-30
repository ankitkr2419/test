package db

import (
	"context"
	"errors"
	"fmt"
	"time"

	logger "github.com/sirupsen/logrus"
)

const (
	validateUserQuery = `SELECT * FROM users as u
	        WHERE u.username = $1 AND u.password = $2`

	insertUsersQuery1 = `INSERT INTO users (
				username,
				password)
				VALUES  `
	insertUsersQuery2 = ` ON CONFLICT DO NOTHING;`
)

type User struct {
	Username  string    `db:"username" json:"username" validate:"required"`
	Password  string    `db:"password" json:"password" validate:"required"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

func (s *pgStore) ValidateUser(ctx context.Context, u User) (err error) {

	r, err := s.db.Exec(
		validateUserQuery,
		u.Username,
		u.Password,
	)
	if err != nil {
		logger.WithField("error in exec query", err.Error()).Error("Query Failed")
		return
	}

	c, _ := r.RowsAffected()
	// check row count as no error is returned when row not found for update
	if c == 0 {
		err = errors.New("Record Not Found")
		logger.WithField("err", err.Error()).Error("Error User not found")
		return
	}

	return
}

func (s *pgStore) InsertUser(ctx context.Context, u User) (err error) {

	values := fmt.Sprintf("('%v','%v')", u.Username, u.Password)

	stmt := fmt.Sprintf(insertUsersQuery1 + values + insertUsersQuery2)

	_, err = s.db.Exec(
		stmt,
	)
	if err != nil {
		logger.WithField("error in exec query", err.Error()).Error("Query Failed")
		return
	}

	return
}
