package db

import (
	"context"
	"errors"
	"fmt"
	"time"

	logger "github.com/sirupsen/logrus"
)

const (
	validateUserQuery = `SELECT role FROM users
	        WHERE username = $1 AND password = $2`

	getUserQuery = `SELECT * from users WHERE username = $1`

	insertUsersQuery1 = `INSERT INTO users (
				username,
				password,
			    role)
				VALUES  `
	insertUsersQuery2 = ` ON CONFLICT (username) DO UPDATE                           
	SET disabled = false ,
	role = excluded.role,
	password = excluded.password                                                                            
	where users.username = excluded.username;`

	updateUsersQuery = `UPDATE users SET username=$1, password=$2, role= $3 where username=$4 `
	deleteUserQuery  = `UPDATE users SET disabled = true where username = $1`
)

type User struct {
	Username  string    `db:"username" json:"username" validate:"required"`
	Password  string    `db:"password" json:"password" validate:"required"`
	Role      string    `db:"role" json:"role"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

func (s *pgStore) ValidateUser(ctx context.Context, u User) (User, error) {

	rows, err := s.db.Query(
		validateUserQuery,
		u.Username,
		u.Password,
	)

	if err != nil {
		logger.WithField("error in exec query", err.Error()).Error("Query Failed")
		return User{}, err
	}

	if rows.Next() {
		if err = rows.Scan(&u.Role); err != nil {
			logger.WithField("err", err.Error()).Error("Error getting user role details")
			return User{}, err
		}
	}

	return u, nil
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

func (s *pgStore) DeleteUser(ctx context.Context, username string) (err error) {
	_, err = s.db.Exec(
		deleteUserQuery,
		username,
	)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error deleting User")
		return
	}

	return
}
