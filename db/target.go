package db

import (
	"context"

	logger "github.com/sirupsen/logrus"
)

const(
	createTargetQuery = `INSERT INTO target (
		name,
		dye_id,
		well_id,
		template_id,
		ct)
		VALUES ($1, $2, $3, $4) RETURNING id`

	getTargetListQuery = `SELECT * FROM target
		ORDER BY name ASC`

	getTargetQuery = `SELECT id,
		name,
		dye_id,
		well_id,
		template_id,
		ct FROM target WHERE id = $1`
)
type ErrorResponse struct {
	Code 	string `json:"code"`
	Message	string `json:"message"`
	Fields map[string]string `json:"fields"`
}
type Target struct {
	ID           int     `db:"id" json:"id"`
	Name         string  `db:"name" json:"name"`
	Dye_ID       int     `db:"dye_id" json:"dye_id"`
	Well_ID      int     `db:"well_id" json:"well_id"`
	Templated_ID int     `db:"templated_id" json:"template_id"`
	CT           float64 `db:"ct" json:"ct"`
}
func (t *Target) Validate() (errorResponse map[string]ErrorResponse, valid bool) {
	fieldErrors := make(map[string]string)

	if t.Name == "" {
		fieldErrors["name"] = "Can't be blank"
	}

		if len(fieldErrors) == 0 {
		valid = true
		return
	}

	errorResponse = map[string]ErrorResponse{"error": ErrorResponse{
			Code: "invalid_data",
			Message: "Please provide valid target data",
			Fields: fieldErrors,
		},
	}

	//TODO: Ask what other validations are expected

	return
}

func (s *pgStore) ListTarget(ctx context.Context) (t []Target, err error) {
	err = s.db.Select(&t, getTargetListQuery)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error listing users")
		return
	}

	return
}

func (s *pgStore) CreateTarget(ctx context.Context, t Target) (createdTarget Target, err error) {
	lastInsertId := 0
	err = s.db.QueryRow(
		createTargetQuery,
		t.Name,
		t.Dye_ID,
		t.Well_ID,
		t.Templated_ID,
		t.CT,
	).Scan(&lastInsertId)

	if err != nil {
		logger.WithField("err", err.Error()).Error("Error creating Target")
		return
	}

	err = s.db.Get(&createdTarget, getTargetQuery, lastInsertId)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error in getting Target")
		return
	}
	return
}