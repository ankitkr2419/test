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
	RecipeID    uuid.UUID     `json:"recipe_id" db:"recipe_id"`
	ShakerNo    int           `json:"shaker_no" db:"shaker_no"`
	CreatedAt   time.Time     `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time     `json:"updated_at" db:"updated_at"`
}

const (
	//get heating data with recipe_id
	getHeatingQuery = `SELECT h.*,p.recipe_id 
			FROM heating as h join processes as p 
			on h.id=p.id WHERE 
			h.id = $1`
)

func (s *pgStore) GetHeating(ctx context.Context, id uuid.UUID) (heating Heating, err error) {

	//create process record
	err = s.db.Get(&heating,
		getHeatingQuery,
		id,
	)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error creating heating")
		return
	}
	return

}
