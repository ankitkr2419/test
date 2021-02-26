package db

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	logger "github.com/sirupsen/logrus"
)

type Shaker struct {
	ID          uuid.UUID     `json:"id" db:"id"`
	WithTemp    bool          `json:"with_temp" db:"with_temp"`
	Temperature int           `json:"temperature" db:"temperature"`
	FollowTemp  bool          `json:"follow_temp" db:"follow_temp"`
	ProcessID   uuid.UUID     `json:"process_id" db:"process_id"`
	Rpm1        int           `json:"rpm_1" db:"rpm_1"`
	Rpm2        int           `json:"rpm_2" db:"rpm_2"`
	Time1       time.Duration `json:"time_1" db:"time_1"`
	Time2       time.Duration `json:"time_2" db:"time_2"`
	RecipeID    uuid.UUID     `json:"recipe_id" db:"recipe_id"`
	CreatedAt   time.Time     `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time     `json:"updated_at" db:"updated_at"`
}

const (
	getShakerQuery = `SELECT s.*,p.recipe_id 
	FROM shaking as s join processes as p 
	on s.process_id = p.id WHERE 
	p.recipe_id = $1`
)

func (s *pgStore) ShowShaking(ctx context.Context, shakerID uuid.UUID) (shaker Shaker, err error) {

	err = s.db.Get(&shaker,
		getShakerQuery,
		shakerID,
	)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error getting shaking")
		return
	}

	fmt.Printf("shaker %v", shaker)
	return
}
