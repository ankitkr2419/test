package db

import (
	"context"
	"time"

	"github.com/google/uuid"
	logger "github.com/sirupsen/logrus"
)

type Heating struct {
	ID          uuid.UUID     `json:"id" db:"id"`
	Temperature float64       `json:"temperature" db:"temperature"`
	FollowTemp  bool          `json:"follow_temp" db:"follow_temp"`
	Duration    time.Duration `json:"duration" db:"duration"`
	ProcessID   uuid.UUID     `json:"process_id" db:"process_id"`
	CreatedAt   time.Time     `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time     `json:"updated_at" db:"updated_at"`
}

const (
	//get heating data with recipe_id
	getHeatingQuery = `SELECT * FROM heating where process_id = $1`
)

func (s *pgStore) GetHeating(ctx context.Context, id uuid.UUID) (heating Heating, err error) {

	// get heating record
	err = s.db.Get(&heating,
		getHeatingQuery,
		id,
	)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error getting heating")
		return
	}
	return

}
