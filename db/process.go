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

func (s *pgStore) createProcess(ctx context.Context, p Process, tx *sql.Tx) (cp Process, err error) {

	var lastInsertID uuid.UUID

	err = tx.QueryRow(
		createProcessQuery,
		p.Name,
		p.Type,
		p.RecipeID,
		p.SequenceNumber,
	).Scan(&lastInsertID)

	if err != nil {
		logger.WithField("err", err.Error()).Errorln(responses.ProcessDBCreateError)
		return
	}

	p.ID = lastInsertID
	return p, err
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

func (s *pgStore) UpdateProcessName(ctx context.Context, id uuid.UUID, processType string, process interface{}) (err error) {
	// pass Process{} just to keep dame method call
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

	if operation == "duplicate" {
		// In a transaction insert entry into Process table
		// and then into its type table
		tx, err = s.db.BeginTx(ctx, nil)
		if err != nil {
			logger.WithField("err:", err.Error()).Errorln(responses.ProcessInitiateDBTxError)
			return Process{}, err
		}

		// End the transaction in defer call
		defer func() {
			if err != nil {
				logger.WithField("err", err.Error()).Errorln(responses.ProcessDuplicateCreationError)
				tx.Rollback()
				return
			}
			tx.Commit()

			pr, err = s.ShowProcess(ctx, pr.ID)
			if err != nil {
				logger.Infoln("Error Duplicating process")
				return
			}
			logger.Infoln("Process Duplicated => ", pr)

		}()

		// Get highest sequence number
		// This sequence number is updation, we only need to get something unique
		highestSeqNum, err := s.getHighestSequenceNumber(ctx)
		if err != nil {
			// failure already logged
			return Process{}, err
		}
		// get the highest sequence number for our process
		parent.SequenceNumber = highestSeqNum + 1

		// Insert parent entry into Process table
		pr, err = s.createProcess(ctx, parent, tx)
		if err != nil {
			// failure will be logged in defer before rollback
			return Process{}, err
		}
		// This creation is logged after commit
	}

	// handle interface conversion failures
	defer func() {
		if r := recover(); r != nil {
			err = responses.InvalidInterfaceConversionError
		}
	}()

	switch processType {
	case "Piercing":
		var pi Piercing
		if operation == "name" {
			pi = process.(Piercing)
			pr.Name = fmt.Sprintf("Piercing_%s", pi.Type)
			return
		}

		// Fetch Piercing process
		pi, err = s.ShowPiercing(ctx, parent.ID)
		if err != nil {
			return
		}

		pi.ProcessID = pr.ID
		// Create Piercing process
		pi, err = s.createPiercing(ctx, pi, tx)
		if err != nil {
			logger.WithField("err", err.Error()).Errorln(responses.PiercingDuplicateCreateError)
			return
		}

		logger.Infoln("Piercing Process Duplication in Progress => ", pi)
		return

	case "TipOperation":
		var to TipOperation
		if operation == "name" {
			to = process.(TipOperation)
			pr.Name = fmt.Sprintf("Tip_Operation_%s", to.Type)
			return
		}

		// Fetch TipOperation process
		to, err = s.ShowTipOperation(ctx, parent.ID)
		if err != nil {
			return
		}

		to.ProcessID = pr.ID
		// Create TipOperation process
		to, err = s.createTipOperation(ctx, to, tx)
		if err != nil {
			logger.WithField("err", err.Error()).Errorln(responses.TipOperationDuplicateCreateError)
			return
		}

		logger.Infoln("Tip Operation Process Duplication in Progress => ", to)
		return

	case "TipDocking":
		var td TipDock
		if operation == "name" {
			td = process.(TipDock)
			pr.Name = fmt.Sprintf("Tip_Docking_%s", td.Type)
			return
		}

		// Fetch TipDocking process
		td, err = s.ShowTipDocking(ctx, parent.ID)
		if err != nil {
			return
		}

		td.ProcessID = pr.ID
		// Create TipDocking process
		td, err = s.createTipDocking(ctx, td, tx)
		if err != nil {
			logger.WithField("err", err.Error()).Errorln(responses.TipDockingDuplicateCreateError)
			return
		}

		logger.Infoln("Tip Docking Process Duplication in Progress  => ", td)
		return

	case "AspireDispense":
		var ad AspireDispense
		if operation == "name" {
			ad = process.(AspireDispense)
			pr.Name = fmt.Sprintf("Aspire_Dispense_%s", ad.Category)
			return
		}

		// Fetch AspireDispense process
		ad, err = s.ShowAspireDispense(ctx, parent.ID)
		if err != nil {
			return
		}

		ad.ProcessID = pr.ID
		// Create AspireDispense process
		ad, err = s.createAspireDispense(ctx, ad, tx)
		if err != nil {
			logger.WithField("err", err.Error()).Errorln(responses.AspireDispenseDuplicateCreateError)
			return
		}

		logger.Infoln("Aspire Dispense Process Duplication in Progress  => ", ad)
		return

	case "Heating":
		var ht Heating
		if operation == "name" {
			pr.Name = "Heating"
			return
		}

		// Fetch Heating process
		ht, err = s.ShowHeating(ctx, parent.ID)
		if err != nil {
			return
		}

		ht.ProcessID = pr.ID
		// Create Heating process
		ht, err = s.createHeating(ctx, ht, tx)
		if err != nil {
			logger.WithField("err", err.Error()).Errorln(responses.HeatingDuplicateCreateError)
			return
		}

		logger.Infoln("Heating Process Duplication in Progress  => ", ht)
		return

	case "Shaking":
		var sk Shaker
		if operation == "name" {
			sk = process.(Shaker)
			if sk.WithTemp {
				pr.Name = "Shaking_With_temperature"
				return
			}
			pr.Name = "Shaking_Without_temperature"
			return
		}

		// Fetch Shaking process
		sk, err = s.ShowShaking(ctx, parent.ID)
		if err != nil {
			return
		}

		sk.ProcessID = pr.ID
		// Create Shaking process
		sk, err = s.createShaking(ctx, sk, tx)
		if err != nil {
			logger.WithField("err", err.Error()).Errorln(responses.ShakingDuplicateCreateError)
			return
		}

		logger.Infoln("Shaking Process Duplication in Progress  => ", sk)
		return

	case "AttachDetach":
		var ad AttachDetach
		if operation == "name" {
			ad = process.(AttachDetach)
			pr.Name = fmt.Sprintf("Magnet_%s", ad.Operation)
			return
		}

		// Fetch AttachDetach process
		ad, err = s.ShowAttachDetach(ctx, parent.ID)
		if err != nil {
			return
		}

		ad.ProcessID = pr.ID
		// Create AttachDetach process
		ad, err = s.createAttachDetach(ctx, ad, tx)
		if err != nil {
			logger.WithField("err", err.Error()).Errorln(responses.AttachDetachDuplicateCreateError)
			return
		}

		logger.Infoln("AttachDetach Process Duplication in Progress  => ", ad)
		return

	case "Delay":
		var dl Delay
		if operation == "name" {
			pr.Name = "Delay"
			return
		}
		// Fetch Delay process
		dl, err = s.ShowDelay(ctx, parent.ID)
		if err != nil {
			return
		}

		dl.ProcessID = pr.ID
		// Create Delay process
		dl, err = s.createDelay(ctx, dl, tx)
		if err != nil {
			logger.WithField("err", err.Error()).Errorln(responses.DelayDuplicateCreateError)
			return
		}
		logger.Infoln("Delay Process Duplication in Progress  => ", dl)
		return
	default:
		return Process{}, responses.ProcessTypeInvalid
	}
}

func (s *pgStore) getHighestSequenceNumber(ctx context.Context) (highestSeqNum int64, err error) {
	err = s.db.QueryRow(
		getHighestSequenceNumberQuery,
	).Scan(&highestSeqNum)

	if err != nil {
		logger.WithField("err", err.Error()).Errorln(responses.ProcessHighestSeqNumFetchError)
	}

	logger.Infoln(responses.ProcessHighestSeqNumFetchSuccess, highestSeqNum)

	return
}
