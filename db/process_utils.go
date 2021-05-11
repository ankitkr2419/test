package db

import (
	"context"
	"database/sql"
	"fmt"
	"mylab/cpagent/responses"

	"github.com/google/uuid"
	logger "github.com/sirupsen/logrus"
)

const (
	updateProcessNameQuery       = `UPDATE processes SET name = $1 WHERE id = $2;`
	getProcessCountQuery         = `SELECT count(*) FROM processes WHERE recipe_id = $1;`
	subtractProcessSequenceQuery = `UPDATE processes SET sequence_num = (sequence_num - $1) WHERE recipe_id = $2;`
	rearrangeSequenceQuery       = `UPDATE processes SET sequence_num = $1 WHERE id = $2`
)

func (s *pgStore) createProcess(ctx context.Context, tx *sql.Tx, p Process) (cp Process, err error) {

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

func (s *pgStore) getProcessCount(ctx context.Context, tx *sql.Tx, recipeID uuid.UUID) (processCount int64, err error) {
	err = tx.QueryRow(
		getProcessCountQuery,
		recipeID,
	).Scan(&processCount)

	if err != nil {
		logger.WithField("err", err.Error()).Errorln(responses.ProcessCountFetchError)
		return
	}
	logger.Infoln(responses.ProcessCountFetchSuccess, processCount)

	return
}

func (s *pgStore) rearrangeProcessSequence(ctx context.Context, tx *sql.Tx, sequenceArr []ProcessSequence, processCount int64) (err error) {

	for _, pr := range sequenceArr {
		_, err = tx.Exec(
			rearrangeSequenceQuery,
			pr.SequenceNumber+processCount,
			pr.ID,
		)

		if err != nil {
			logger.WithField("err", err.Error()).Errorln(responses.ProcessRearrangeDBError)
			return
		}
	}

	logger.Infoln(responses.ProcessRearrangeDBSuccess, processCount)
	return
}

//  subtract processCount from sequence num of process
func (s *pgStore) subtractFromSequence(ctx context.Context, tx *sql.Tx, recipeID uuid.UUID, processCount int64) (err error) {
	_, err = tx.Query(
		subtractProcessSequenceQuery,
		processCount,
		recipeID,
	)

	if err != nil {
		logger.WithField("err", err.Error()).Errorln(responses.ProcessSubtractError)
		return
	}

	logger.Infoln(responses.ProcessSubtractSuccess, processCount)
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
		highestSeqNum, err := s.getProcessCount(ctx, tx, parent.RecipeID)
		if err != nil {
			// failure already logged
			return Process{}, err
		}
		// get the highest sequence number for our process
		parent.SequenceNumber = highestSeqNum + 1

		// Insert parent entry into Process table
		pr, err = s.createProcess(ctx, tx, parent)
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
		pi, err = s.createPiercing(ctx, tx, pi)
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
		to, err = s.createTipOperation(ctx, tx, to)
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
		td, err = s.createTipDocking(ctx, tx, td)
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
		ad, err = s.createAspireDispense(ctx, tx, ad)
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
		ht, err = s.createHeating(ctx, tx, ht)
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
		sk, err = s.createShaking(ctx, tx, sk)
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
		ad, err = s.createAttachDetach(ctx, tx, ad)
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
		dl, err = s.createDelay(ctx, tx, dl)
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
