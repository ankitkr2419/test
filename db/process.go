package db

import (
	"context"
	"database/sql"
	"fmt"
	"mylab/cpagent/responses"
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

	updateProcessNameQuery = `UPDATE processes SET name = $1 WHERE id = $2`

	getHighestSequenceNumberQuery = `SELECT max(sequence_num) FROM processes; `
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

func (s *pgStore) ListProcesses(ctx context.Context, id uuid.UUID) (dbProcess []Process, err error) {
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

func (s *pgStore) DuplicateProcess(ctx context.Context, processID uuid.UUID, process interface{}) (duplicate Process, err error) {

	var parent Process
	var highestSeqNum int64

	// if parent process exists only then create duplicate process
	parent, err = s.ShowProcess(ctx, processID)
	if err != nil {
		logger.WithField("err", err.Error()).Error(responses.ProcessFetchError)
		return
	}

	// Get highest sequence number
	// This sequence number is updation, we only need to get something unique
	err = s.db.QueryRow(
		getHighestSequenceNumberQuery,
	).Scan(&highestSeqNum)

	// get the highest sequence number for our process
	parent.SequenceNumber = highestSeqNum + 1

	duplicate, err = s.processOperation(ctx, "duplicate", parent.Type, process, parent)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error creating duplicate process")
		return
	}

	logger.Infoln("Duplicate process created in db : ", duplicate)
	return
}

func (s *pgStore) UpdateProcessName(ctx context.Context, id uuid.UUID, processType string, process interface{}) (err error) {
	processWithName, err := s.processOperation(ctx, "name", processType, process, Process{})
	if err != nil {
		err = fmt.Errorf("error in updating new process name")
		return
	}
	_, err = s.db.Exec(
		updateProcessNameQuery,
		processWithName.Name,
		id,
	)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error updating process name")
		return
	}
	return
}

func (s *pgStore) processOperation(ctx context.Context, operation string, processType string, process interface{}, parent Process) (pr Process, err error) {

	var tx *sql.Tx
	var lastInsertID uuid.UUID

	if operation == "duplicate" {
		// In a transaction insert entry into Process table
		// and then into its type table
		tx, err = s.db.BeginTx(ctx, nil)
		if err != nil {
			logger.WithField("err:", err.Error()).Error("Error while initiating transaction")
			return Process{}, err
		}

		defer func() {
			if err != nil {
				tx.Rollback()
				return
			}
			tx.Commit()
		}()

		// Insert entry into Process table
		err = tx.QueryRow(
			createProcessQuery,
			parent.Name,
			parent.Type,
			parent.RecipeID,
			parent.SequenceNumber,
		).Scan(&lastInsertID)

		if err != nil {
			logger.WithField("err", err.Error()).Error("Error creating duplicate Process")
			return Process{}, err
		}

		pr = parent
		pr.ID = lastInsertID
		// Avoiding complete fetch and only updating these fields
		pr.CreatedAt = time.Now()
		pr.UpdatedAt = time.Now()
		logger.Infoln("Process Duplicated => ", pr)

	}

	defer func() {
		if r := recover(); r != nil {
			err = responses.InvalidInterfaceConversionError
		}
	}()

	switch processType {
	case "Piercing":
		if operation == "name" {
			p := process.(Piercing)
			pr.Name = fmt.Sprintf("Piercing_%s", p.Type)
			return
		}

		// Create Piercing process
		p := process.(*Piercing)
		p.ProcessID = pr.ID

		err = tx.QueryRow(
			createPiercingQuery,
			p.Type,
			p.CartridgeWells,
			p.Discard,
			p.ProcessID,
		).Scan(&lastInsertID)

		if err != nil {
			logger.WithField("err", err.Error()).Error("Error creating duplicate piercing")
			return
		}

		p.ID = lastInsertID
		logger.Infoln("Piercing Process Duplicated => ", p)
		return

	case "TipOperation":
		if operation == "name" {
			t := process.(TipOperation)
			pr.Name = fmt.Sprintf("Tip_Operation_%s", t.Type)
			return
		}
		
		// Create TipOperation  process
		t := process.(*TipOperation)
		t.ProcessID = pr.ID

		err = tx.QueryRow(
			createTipOperationQuery,
			t.Type,
			t.Position,
			t.ProcessID,
		).Scan(&lastInsertID)

		if err != nil {
			logger.WithField("err", err.Error()).Error("Error creating duplicate tip operation")
			return
		}

		t.ID = lastInsertID
		logger.Infoln("Tip Operation Process Duplicated => ", t)
		return

	case "TipDocking":
		t := process.(TipDock)
		if operation == "name" {
			pr.Name = fmt.Sprintf("Tip_Docking_%s", t.Type)
			return
		}
		// Create TipDocking process
		t.ProcessID = pr.ID

		err = tx.QueryRow(
			createTipDockQuery,
			t.Type,
			t.Position,
			t.Height,
			t.ProcessID,
		).Scan(&lastInsertID)

		if err != nil {
			logger.WithField("err", err.Error()).Error("Error creating duplicate tip docking")
			return
		}

		t.ID = lastInsertID
		logger.Infoln("Tip Docking Process Duplicated => ", t)
		return

	case "AspireDispense":
		ad := process.(AspireDispense)
		if operation == "name" {
			pr.Name = fmt.Sprintf("Aspire_Dispense_%s", ad.Category)
			return
		}
		// Create AspireDispense process
		ad.ProcessID = pr.ID

		err = tx.QueryRow(
			createAspireDispenseQuery,
			ad.Category,
			ad.CartridgeType,
			ad.SourcePosition,
			ad.AspireHeight,
			ad.AspireMixingVolume,
			ad.AspireNoOfCycles,
			ad.AspireVolume,
			ad.AspireAirVolume,
			ad.DispenseHeight,
			ad.DispenseMixingVolume,
			ad.DispenseNoOfCycles,
			ad.DestinationPosition,
			ad.ProcessID,
		).Scan(&lastInsertID)

		if err != nil {
			logger.WithField("err", err.Error()).Error("Error creating duplicate aspire dispense")
			return
		}

		ad.ID = lastInsertID
		logger.Infoln("Aspire Dispense Process Duplicated => ", ad)
		return

	case "Heating":
		if operation == "name" {
			pr.Name = "Heating"
			return
		}
		// Create Heating process
		h := process.(Heating)
		h.ProcessID = pr.ID

		err = tx.QueryRow(
			createHeatingQuery,
			h.Temperature,
			h.FollowTemp,
			h.Duration,
			h.ProcessID,
		).Scan(&lastInsertID)

		if err != nil {
			logger.WithField("err", err.Error()).Error("Error creating duplicate heating")
			return
		}

		h.ID = lastInsertID
		logger.Infoln("Heating Process Duplicated => ", h)
		return

	case "Shaking":
		sh := process.(Shaker)
		if operation == "name" {
			if sh.WithTemp {
				pr.Name = "Shaking_With_temperature"
				return
			}
			pr.Name = "Shaking_Without_temperature"
			return
		}
		// Create Shaking process
		sh.ProcessID = pr.ID

		err = tx.QueryRow(
			createShakingQuery,
			sh.WithTemp,
			sh.Temperature,
			sh.FollowTemp,
			sh.RPM1,
			sh.RPM2,
			sh.Time1,
			sh.Time2,
			sh.ProcessID,
		).Scan(&lastInsertID)

		if err != nil {
			logger.WithField("err", err.Error()).Error("Error creating duplicate shaking")
			return
		}

		sh.ID = lastInsertID
		logger.Infoln("Shaking Process Duplicated => ", sh)
		return

	case "AttachDetach":
		ad := process.(AttachDetach)
		if operation == "name" {
			pr.Name = fmt.Sprintf("Magnet_%s", ad.Operation)
			return
		}
		// Create AttachDetach process
		ad.ProcessID = pr.ID

		err = tx.QueryRow(
			createAttachDetachQuery,
			ad.Operation,
			ad.OperationType,
			ad.ProcessID,
		).Scan(&lastInsertID)

		if err != nil {
			logger.WithField("err", err.Error()).Error("Error creating duplicate attach detach")
			return
		}

		ad.ID = lastInsertID
		logger.Infoln("AttachDetach Process Duplicated => ", ad)
		return

	case "Delay":
		if operation == "name" {
			pr.Name = "Delay"
			return
		}
		// Create Delay process
		d := process.(Delay)
		d.ProcessID = pr.ID

		err = tx.QueryRow(
			createDelayQuery,
			d.DelayTime,
			d.ProcessID,
		).Scan(&lastInsertID)

		if err != nil {
			logger.WithField("err", err.Error()).Error("Error creating duplicate delay")
			return
		}
		d.ID = lastInsertID
		logger.Infoln("Delay Process Duplicated => ", d)
		return
	default:
		return Process{}, responses.ProcessTypeInvalid
	}
}
