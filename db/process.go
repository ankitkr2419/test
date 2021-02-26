package db

import (
	"context"
	"time"

	"github.com/google/uuid"
	logger "github.com/sirupsen/logrus"
)

const (
	getProcessQuery = `SELECT id,
						name,
						type,
						recipe_id,
						sequence_num,
						created_at,
						updated_at
						FROM processes
						WHERE id = $1`
	selectProcessQuery = `SELECT *
						FROM processes where recipe_id = $1 `
	deleteProcessQuery = `DELETE FROM processes
						WHERE id = $1`
	createProcessQuery = `INSERT INTO processes (
						name,
						type,
						recipe_id,
						sequence_num)
						VALUES ($1, $2, $3, $4) RETURNING id`
	updateProcessQuery = `UPDATE processes SET (
						name,
						type,
						recipe_id,
						sequence_num,
						updated_at)
						VALUES ($1, $2, $3, $4, $5) WHERE id = $6`
)

type Process struct {
	ID             uuid.UUID `db:"id" json:"id"`
	Name           string    `db:"name" json:"name"`
	Type           string    `db:"type" json:"type"`
	RecipeID       uuid.UUID `db:"recipe_id" json:"recipe_id"`
	SequenceNumber int64     `db:"sequence_num" json:"sequence_num"`
	CreatedAt      time.Time `db:"created_at" json:"created_at"`
	UpdatedAt      time.Time `db:"updated_at" json:"updated_at"`
}

func (s *pgStore) ShowProcess(ctx context.Context, id uuid.UUID) (dbProcess Process, err error) {
	err = s.db.Get(&dbProcess, getProcessQuery, id)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error fetching process")
		return
	}
	return
}

func (s *pgStore) ListProcess(ctx context.Context, id uuid.UUID) (dbProcess []Process, err error) {
	err = s.db.Select(&dbProcess, selectProcessQuery, id)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error fetching process")
		return
	}
	return
}

func (s *pgStore) CreateProcess(ctx context.Context, p Process) (createdProcess Process, err error) {
	var lastInsertID uuid.UUID
	err = s.db.QueryRow(
		createProcessQuery,
		p.Name,
		p.Type,
		p.RecipeID,
		p.SequenceNumber,
	).Scan(&lastInsertID)

	if err != nil {
		logger.WithField("err", err.Error()).Error("Error creating Process")
		return
	}

	err = s.db.Get(&createdProcess, getProcessQuery, lastInsertID)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error in getting Process")
		return
	}
	return
}

func (s *pgStore) DeleteProcess(ctx context.Context, id uuid.UUID) (err error) {
	_, err = s.db.Exec(deleteProcessQuery, id)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error deleting process")
		return
	}
	return
}

func (s *pgStore) UpdateProcess(ctx context.Context, p Process) (err error) {
	_, err = s.db.Exec(
		updateProcessQuery,
		p.Name,
		p.Type,
		p.RecipeID,
		p.SequenceNumber,
		time.Now(),
		p.ID,
	)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error updating process")
		return
	}
	return
}
