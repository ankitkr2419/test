package db

import (
	"context"

	"github.com/google/uuid"
	logger "github.com/sirupsen/logrus"
)

type UserAuth struct {
	AuthID   uuid.UUID `json:"auth_id"`
	Username string    `json:"username"`
}

const (
	userAuthInsertQuery = `INSERT INTO user_auths (username) VALUES ($1) RETURNING auth_id`
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
