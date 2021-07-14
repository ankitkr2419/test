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

	getUserQuery = `SELECT * from users WHERE username = $1`

	insertUsersQuery1 = `INSERT INTO users (
				username,
				password,
			    role)
				VALUES  `
	insertUsersQuery2 = ` ON CONFLICT DO NOTHING;`

	updateUsersQuery = `UPDATE users SET username=$1, password=$2, role= $3 where username=$4 ` + insertUsersQuery2
)

type User struct {
	Username  string    `db:"username" json:"username" validate:"required"`
	Password  string    `db:"password" json:"password" validate:"required"`
	Role      string    `db:"role" json:"role"`
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

	values := fmt.Sprintf("('%v','%v','%v')", u.Username, u.Password, u.Role)

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

func (s *pgStore) UpdateUser(ctx context.Context, u User, oldName string) (err error) {

	r, err := s.db.Exec(
		updateUsersQuery,
		u.Username,
		u.Password,	
		u.Role,
		oldName,	
	)
	if err != nil {
		logger.WithField("error in exec query", err.Error()).Errorln("Query Failed for Update User")
		return
	}

	c, _ := r.RowsAffected()
	// check row count as no error is returned when row not found for update
	if c == 0 {
		err = errors.New("Record Couldn't be Updated")
		logger.WithField("err", err.Error()).Error("Possible Duplicate User")
		return
	}

	return
}

func (s *pgStore) ShowUser(ctx context.Context, username string) (user User, err error) {
	err = s.db.Get(&user, getUserQuery, username)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error fetching user")
		return
	}
	return
}
