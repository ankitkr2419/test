package db

import (
	"context"
	"errors"
	"fmt"
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

func (s *pgStore) UpdateProcessName(ctx context.Context, id uuid.UUID, processType string, process interface{}) (err error) {
	processName, err := getProcessName(processType, process)
	if err != nil {
		err = fmt.Errorf("error in updating new process name")
		return
	}
	_, err = s.db.Exec(
		updateProcessNameQuery,
		processName,
		id,
	)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error updating process name")
		return
	}
	return
}

func getProcessName(processType string, process interface{}) (processName string, err error) {

	switch processType {
	case "Piercing":
		piercing := process.(Piercing)
		processName = fmt.Sprintf("Piercing_%s", piercing.Type)
		return

	case "TipOperation":
		tipOpr := process.(TipOperation)
		processName = fmt.Sprintf("Tip_Operation_%s", tipOpr.Type)
		return

	case "TipDocking":
		tipDock := process.(TipDock)
		processName = fmt.Sprintf("Tip_Docking_%s", tipDock.Type)
		return

	case "AspireDispense":
		aspDis := process.(AspireDispense)
		processName = fmt.Sprintf("Aspire_Dispense_%s", aspDis.Category)
		return

	case "Heating":
		processName = fmt.Sprintf("Heating")
		return

	case "Shaking":
		shaking := process.(Shaker)
		if shaking.WithTemp {
			processName = fmt.Sprintf("Shaking_With_temperature")
			return
		}
		processName = fmt.Sprintf("Shaking_Without_temperature")
		return

	case "AttachDetach":
		atDet := process.(AttachDetach)
		processName = fmt.Sprintf("Magnet_%s", atDet.Operation)
		return

	case "Delay":
		processName = fmt.Sprintf("Delay")
		return

	default:
		return "", errors.New("wrong process type")
	}
}
