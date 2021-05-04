package db

import (
	"context"
	"database/sql"
	"fmt"
	"mylab/cpagent/responses"

	"github.com/google/uuid"
	logger "github.com/sirupsen/logrus"
)

const(
	updateProcessNameQuery = `UPDATE processes SET name = $1 WHERE id = $2`

	getHighestSequenceNumberQuery = `SELECT max(sequence_num) FROM processes; `
)



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

func (s *pgStore) updateProcessName(ctx context.Context, id uuid.UUID, processType string, process interface{}) (err error) {
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


func (s *pgStore) getHighestSequenceNumber(ctx context.Context, tx *sql.Tx) (highestSeqNum int64, err error) {
	err = tx.QueryRow(
		getHighestSequenceNumberQuery,
	).Scan(&highestSeqNum)

	if err != nil {
		logger.WithField("err", err.Error()).Errorln(responses.ProcessHighestSeqNumFetchError)
	}

	logger.Infoln(responses.ProcessHighestSeqNumFetchSuccess, highestSeqNum)

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
		highestSeqNum, err := s.getHighestSequenceNumber(ctx, tx)
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

