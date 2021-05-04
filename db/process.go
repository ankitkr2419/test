package db

import (
	"context"
	"database/sql"
	"mylab/cpagent/responses"
	"time"

	"github.com/google/uuid"
	logger "github.com/sirupsen/logrus"
)

const (
	getProcessQuery = `SELECT *
						FROM processes
						WHERE id = $1`
	selectProcessQuery = `SELECT *
						FROM processes where recipe_id = $1 ORDER BY sequence_num`
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
						sequence_num,
						updated_at)
						VALUES ($1, $2, $3) WHERE id = $4`
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
		logger.WithField("err", err.Error()).Errorln(responses.ProcessDBFetchError)
		return
	}
	return
}

func (s *pgStore) ListProcesses(ctx context.Context, id uuid.UUID) (dbProcess []Process, err error) {
	err = s.db.Select(&dbProcess, selectProcessQuery, id)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error fetching process")
		return
	}
	return
}

func (s *pgStore) CreateProcess(ctx context.Context, p Process) (createdProcess Process, err error) {
	var tx *sql.Tx

	tx, err = s.db.BeginTx(ctx, nil)
	if err != nil {
		logger.WithField("err:", err.Error()).Errorln(responses.ProcessInitiateDBTxError)
		return Process{}, err
	}

	createdProcess, err = s.createProcess(ctx, p, tx)
	// failures are already logged
	// Commit the transaction else won't be able to Show

	// End the transaction in defer call
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		tx.Commit()
		createdProcess, err = s.ShowProcess(ctx, createdProcess.ID)
		logger.Infoln("Created Process: ", createdProcess)
		return
	}()

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

func (s *pgStore) DuplicateProcess(ctx context.Context, processID uuid.UUID) (duplicate Process, err error) {

	var parent Process

	// if parent process exists only then create duplicate process
	parent, err = s.ShowProcess(ctx, processID)
	if err != nil {
		logger.WithField("err", err.Error()).Error(responses.ProcessFetchError)
		return
	}

	var p interface{}
	// pass empty interface, just to use same method
	duplicate, err = s.processOperation(ctx, "duplicate", parent.Type, p, parent)
	if err != nil {
		// failure already logged
		return
	}

	logger.Infoln(responses.ProcessDuplicateCreateSuccess, duplicate)
	return
}

func (s *pgStore) RearrangeProcesses(ctx context.Context, id uuid.UUID, sequenceArr []ProcessSequence) (processes []Process, err error){
	return
}


