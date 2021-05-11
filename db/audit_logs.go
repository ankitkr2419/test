package db

import (
	"context"
	"mylab/cpagent/responses"
	"time"

	"github.com/google/uuid"
	logger "github.com/sirupsen/logrus"
)

type ActivityType string
type StateType string
type OperationType string

const (
	ApiOperation     ActivityType  = "api"
	DBOperation      ActivityType  = "db"
	MachineOperation ActivityType  = "machine"
	InitialisedState StateType     = "initialised"
	CompletedState   StateType     = "completed"
	AbortedState     StateType     = "aborted"
	PausedState      StateType     = "paused"
	ResumedState     StateType     = "resumed"
	ErrorState       StateType     = "error"
	CreateOperation  OperationType = "create"
	ShowOperation    OperationType = "show"
	UpdateOperation  OperationType = "update"
	DeleteOperation  OperationType = "delete"
)

type AuditLog struct {
	ID           uuid.UUID     `db:"id" json:"id"`
	Username     string        `db:"username" json:"username" validate:"required"`
	ActivityType ActivityType  `db:"activity_type" json:"activity_type"  `
	StateType    StateType     `db:"state_type" json:"state_type"`
	Operation    OperationType `db:"operation_type" json:"operation_type"`
	Deck         string        `db:"deck" json:"deck"`
	Description  string        `db:"description" json:"description"`
	CreatedAt    time.Time     `db:"created_at" json:"created_at"`
	UpdatedAt    time.Time     `db:"updated_at" json:"updated_at"`
}

const (
	auditLogsInsertQuery = `INSERT INTO audit_logs(
		username,
		activity_type,
		state_type,
		operation_type,
		deck,
		description,
	) VALUES ($1,$2,$3,$4,$5,$6) RETURNING id`

	auditLogsSelectQuery = `SELECT * FROM audit_logs`
)

func (s *pgStore) InsertAuditLog(ctx context.Context, al AuditLog) (err error) {
	var lastInsertedID uuid.UUID
	err = s.db.QueryRow(auditLogsInsertQuery,
		al.Username,
		al.ActivityType,
		al.StateType,
		al.Operation,
		al.Deck,
		al.Description).Scan(&lastInsertedID)

	if err != nil {
		logger.WithField("err", err.Error()).Errorln(responses.AuditLogDBCreateError)
		return
	}
	return
}

func (s *pgStore) ShowAuditLog(ctx context.Context) (al AuditLog, err error) {
	err = s.db.Select(&al, auditLogsSelectQuery)

	if err != nil {
		logger.WithField("err", err.Error()).Errorln(responses.AuditLogDBShowError)
		return
	}
	return
}

func (s *pgStore) AddAuditLog(ctx context.Context, activity ActivityType, state StateType,
	oprType OperationType, deck, description, username string) (err error) {

	log := AuditLog{
		Username:     username,
		ActivityType: activity,
		StateType:    state,
		Deck:         deck,
		Operation:    oprType,
		Description:  description,
	}

	err = s.InsertAuditLog(ctx, log)
	if err != nil {
		logger.WithField("err", err.Error()).Error(responses.AuditLogCreateError)
		return
	}

	return
}
