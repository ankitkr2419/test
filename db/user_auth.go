package db

import (
	"context"

	"github.com/google/uuid"
	logger "github.com/sirupsen/logrus"
)

type UserAuth struct {
	AuthID   uuid.UUID `json:"auth_id" db:"auth_id"`
	Username string    `json:"username" db:"username"`
}

const (
	userAuthInsertQuery = `INSERT INTO user_auths (username) VALUES ($1) RETURNING auth_id`
	userAuthGetQuery    = `SELECT * FROM user_auths where username = $1 AND auth_id = $2`
	userAuthDeleteQuery = `DELETE FROM user_auths where username = $1 AND auth_id = $2`
)

func (s *pgStore) InsertUserAuths(ctx context.Context, username string) (authID uuid.UUID, err error) {

	err = s.db.QueryRow(
		userAuthInsertQuery,
		username,
	).Scan(&authID)

	if err != nil {
		logger.WithField("err", err.Error()).Error("Error creating user auth record")
		return
	}

	return
}

func (s *pgStore) ShowUserAuth(ctx context.Context, username string, authID uuid.UUID) (ua UserAuth, err error) {
	err = s.db.Get(&ua, userAuthGetQuery, username, authID)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error fetching user auth")
		return
	}

	return
}

func (s *pgStore) DeleteUserAuth(ctx context.Context, userAuth UserAuth) (err error) {
	_, err = s.db.Exec(userAuthDeleteQuery, userAuth.Username, userAuth.AuthID)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error deleting user auth")
		return
	}

	return
}
