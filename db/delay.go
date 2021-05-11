package db

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
	logger "github.com/sirupsen/logrus"
	"mylab/cpagent/responses"
)

type Delay struct {
	ID        uuid.UUID `db:"id" json:"id"`
	DelayTime int64     `db:"delay_time" json:"delay_time" validate:"required"`
	ProcessID uuid.UUID `db:"process_id" json:"process_id" validate:"required"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

const (
	getDelayQuery    = `SELECT * FROM delay where process_id = $1`
	createDelayQuery = `INSERT INTO delay (
		delay_time,
		process_id)
		VALUES ($1, $2) RETURNING id`
	updateDelayQuery = `UPDATE delay SET (
			delay_time,
			updated_at) = ($1, $2) WHERE process_id = $3`
)

func (s *pgStore) ShowDelay(ctx context.Context, id uuid.UUID) (delay Delay, err error) {
	// get delay record
	err = s.db.Get(&delay,
		getDelayQuery,
		id,
	)
	if err != nil {
		logger.WithField("err", err.Error()).Errorln(responses.DelayDBFetchError)
		return
	}
	return
}

func (s *pgStore) CreateDelay(ctx context.Context, d Delay) (createdD Delay, err error) {
	var tx *sql.Tx

	//update the process name before record creation
	err = s.updateProcessName(ctx, d.ProcessID, "Delay", d)
	if err != nil {
		logger.WithField("err", err.Error()).Errorln(responses.DelayUpdateNameError)
		return
	}

	tx, err = s.db.BeginTx(ctx, nil)
	if err != nil {
		logger.WithField("err:", err.Error()).Errorln(responses.DelayInitiateDBTxError)
		return Delay{}, err
	}

	createdD, err = s.createDelay(ctx, tx, d)
	// failures are already logged
	// Commit the transaction else won't be able to Show

	// End the transaction in defer call
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}
		tx.Commit()
		createdD, err = s.ShowDelay(ctx, createdD.ProcessID)
		if err != nil {
			logger.Infoln("Error Creating Delay process")
			return
		}
		logger.Infoln("Created Delay Process: ", createdD)
		return
	}()

	return
}

func (s *pgStore) createDelay(ctx context.Context, tx *sql.Tx, d Delay) (createdD Delay, err error) {

	var lastInsertID uuid.UUID

	err = tx.QueryRow(
		createDelayQuery,
		d.DelayTime,
		d.ProcessID,
	).Scan(&lastInsertID)

	if err != nil {
		logger.WithField("err", err.Error()).Errorln(responses.DelayDBCreateError)
		return
	}

	d.ID = lastInsertID
	return d, err
}

func (s *pgStore) UpdateDelay(ctx context.Context, d Delay) (err error) {
	_, err = s.db.Exec(
		updateDelayQuery,
		d.DelayTime,
		time.Now(),
		d.ProcessID,
	)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error updating delay")
		return
	}
	return
}
