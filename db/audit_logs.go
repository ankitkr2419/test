package db

import (
	"context"
	"mylab/cpagent/responses"
	"time"

	"github.com/google/uuid"
	logger "github.com/sirupsen/logrus"
)

type AuditLog struct {
	ID           uuid.UUID `db:"id" json:"id"`
	Username     string    `db:"username" json:"username" validate:"required"`
	ActivityType string    `db:"activity_type" json:"activity_type"  `
	StateType    string    `db:"state_type" json:"state_type"`
	Deck         string    `db:"deck" json:"deck"`
	Description  string    `db:"description" json:"description"`
	CreatedAt    time.Time `db:"created_at" json:"created_at"`
	UpdatedAt    time.Time `db:"updated_at" json:"updated_at"`
}

const (
	auditLogsInsertQuery = `INSERT INTO audit_logs(
		username,
		activity_type,
		state_type,
		deck,
		description,
	) VALUES ($1,$2,$3,$4,$5) RETURNING id`

	auditLogsSelectQuery = `SELECT * FROM audit_logs`
)

func (s *pgStore) InsertAuditLog(ctx context.Context, al AuditLog) (err error) {
	var lastInsertedID uuid.UUID
	err = s.db.QueryRow(auditLogsInsertQuery,
		al.Username,
		al.ActivityType,
		al.StateType,
		al.Deck,
		al.Description).Scan(&lastInsertedID)

	if err != nil {
		logger.WithField("err", err.Error()).Errorln(responses.ProcessDBFetchError)
		return
	}
	return
}

func (s *pgStore) ShowAuditLog(ctx context.Context) (al AuditLog, err error) {
	err = s.db.Select(&al, auditLogsSelectQuery)

	if err != nil {
		logger.WithField("err", err.Error()).Errorln(responses.ProcessDBFetchError)
		return
	}
	return
}
