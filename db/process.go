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
	go s.AddAuditLog(ctx, DBOperation, InitialisedState, ShowOperation, "", responses.ProcessInitialisedState)

	err = s.db.Get(&dbProcess, getProcessQuery, id)
	defer func() {
		if err != nil {
			go s.AddAuditLog(ctx, DBOperation, ErrorState, ShowOperation, "", err.Error())
		} else {
			go s.AddAuditLog(ctx, DBOperation, CompletedState, ShowOperation, "", responses.ProcessCompletedState)
		}
	}()
	if err != nil {
		logger.WithField("err", err.Error()).Errorln(responses.ProcessDBFetchError)
		return
	}
	return
}

func (s *pgStore) ListProcesses(ctx context.Context, id uuid.UUID) (dbProcess []Process, err error) {
	go s.AddAuditLog(ctx, DBOperation, InitialisedState, ShowOperation, "", responses.ProcessListInitialisedState)

	err = s.db.Select(&dbProcess, selectProcessQuery, id)
	defer func() {
		if err != nil {
			go s.AddAuditLog(ctx, DBOperation, ErrorState, ShowOperation, "", err.Error())
		} else {
			go s.AddAuditLog(ctx, DBOperation, CompletedState, ShowOperation, "", responses.ProcessListCompletedState)
		}
	}()
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error fetching process")
		return
	}
	return
}

func (s *pgStore) CreateProcess(ctx context.Context, p Process) (createdProcess Process, err error) {
	go s.AddAuditLog(ctx, DBOperation, InitialisedState, CreateOperation, "", responses.ProcessInitialisedState)

	var tx *sql.Tx

	tx, err = s.db.BeginTx(ctx, nil)
	if err != nil {
		logger.WithField("err:", err.Error()).Errorln(responses.ProcessInitiateDBTxError)
		return Process{}, err
	}

	createdProcess, err = s.createProcess(ctx, tx, p)
	// failures are already logged
	// Commit the transaction else won't be able to Show

	// End the transaction in defer call
	defer func() {
		if err != nil {
			tx.Rollback()
			go s.AddAuditLog(ctx, DBOperation, ErrorState, CreateOperation, "", err.Error())
			return
		}
		tx.Commit()
		createdProcess, err = s.ShowProcess(ctx, createdProcess.ID)
		logger.Infoln("Created Process: ", createdProcess)
		go s.AddAuditLog(ctx, DBOperation, CompletedState, CreateOperation, "", responses.ProcessCompletedState)
		return
	}()

	return
}

func (s *pgStore) DeleteProcess(ctx context.Context, id uuid.UUID) (err error) {
	go s.AddAuditLog(ctx, DBOperation, InitialisedState, CreateOperation, "", responses.ProcessInitialisedState)

	_, err = s.db.Exec(deleteProcessQuery, id)
	defer func() {
		if err != nil {
			go s.AddAuditLog(ctx, DBOperation, ErrorState, DeleteOperation, "", err.Error())
		} else {
			go s.AddAuditLog(ctx, DBOperation, CompletedState, DeleteOperation, "", responses.ProcessCompletedState)
		}
	}()
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error deleting process")
		return
	}
	return
}

func (s *pgStore) UpdateProcess(ctx context.Context, p Process) (err error) {
	go s.AddAuditLog(ctx, DBOperation, InitialisedState, UpdateOperation, "", responses.ProcessInitialisedState)

	_, err = s.db.Exec(
		updateProcessQuery,
		p.Name,
		p.SequenceNumber,
		time.Now(),
		p.ID,
	)
	defer func() {
		if err != nil {
			go s.AddAuditLog(ctx, DBOperation, ErrorState, UpdateOperation, "", err.Error())
		} else {
			go s.AddAuditLog(ctx, DBOperation, CompletedState, UpdateOperation, "", responses.ProcessCompletedState)
		}
	}()
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error updating process")
		return
	}
	return
}

func (s *pgStore) DuplicateProcess(ctx context.Context, processID uuid.UUID) (duplicateP Process, err error) {
	go s.AddAuditLog(ctx, DBOperation, InitialisedState, CreateOperation, "", responses.DuplicateProcessInitialisedState)

	var parent Process

	// if parent process exists only then create duplicate process
	parent, err = s.ShowProcess(ctx, processID)
	defer func() {
		if err != nil {
			go s.AddAuditLog(ctx, DBOperation, ErrorState, CreateOperation, "", err.Error())
		} else {
			go s.AddAuditLog(ctx, DBOperation, CompletedState, CreateOperation, "", responses.DuplicateProcessCompletedState)
		}
	}()

	if err != nil {
		logger.WithField("err", err.Error()).Error(responses.ProcessFetchError)
		return
	}

	var p interface{}
	// pass empty interface, just to use same method
	duplicateP, err = s.processOperation(ctx, duplicate, ProcessName(parent.Type), p, parent)
	if err != nil {
		// failure already logged
		return
	}

	logger.Infoln(responses.ProcessDuplicateCreateSuccess, duplicateP)
	return
}

//********************
//      ALGORITHM    *
//********************
// 1. initiate transaction
// 2. get count of processes of recipeID
// 3. if count is equal to len of sequenceArr, only then proceed
// 4. update each process by that sequence num + HSN in separate private method
// 5. subtract HSN from sequence num of process
// 6. end transaction
//
// 7. list processes
//
func (s *pgStore) RearrangeProcesses(ctx context.Context, recipeID uuid.UUID, sequenceArr []ProcessSequence) (processes []Process, err error) {
	go s.AddAuditLog(ctx, DBOperation, InitialisedState, UpdateOperation, "", responses.RearrangeProcessInitialisedState)

	var tx *sql.Tx

	// 1. initiate transaction
	tx, err = s.db.BeginTx(ctx, nil)
	if err != nil {
		logger.WithField("err:", err.Error()).Errorln(responses.ProcessInitiateDBTxError)
		return []Process{}, err
	}

	// End the transaction in defer call
	defer func() {
		// 6. end transaction
		if err != nil {
			tx.Rollback()
			go s.AddAuditLog(ctx, DBOperation, ErrorState, UpdateOperation, "", err.Error())
			return
		}
		tx.Commit()
		// 7. list processes
		processes, err = s.ListProcesses(ctx, recipeID)
		if err != nil {
			logger.Errorln(responses.ProcessRearrangeDBError, processes)
		}
		logger.Infoln(responses.ProcessRearrangeSuccess)
		go s.AddAuditLog(ctx, DBOperation, CompletedState, UpdateOperation, "", responses.RearrangeProcessCompletedState)
		return
	}()

	// 2. get count of processes of recipeID
	processCount, err := s.getProcessCount(ctx, tx, recipeID)
	if err != nil {
		// failure already logged
		return []Process{}, err
	}
	// processCount serves as the highest sequence number for our process when updating in db

	// 3. if count is equal to len of sequenceArr, only then proceed
	if len(sequenceArr) != int(processCount) {
		err = responses.ProcessCountDifferError
		logger.WithField("err", responses.ProcessCountDifferError).
			Errorln("length of sequence Arr:", len(sequenceArr), "process count: ", processCount)
		return
	}

	// 4. update each process by that sequence num + HSN in separate private method
	err = s.rearrangeProcessSequence(ctx, tx, sequenceArr, processCount)
	if err != nil {
		return
	}

	// 5. subtract HSN from sequence num of process
	err = s.subtractFromSequence(ctx, tx, recipeID, processCount)
	return
}
