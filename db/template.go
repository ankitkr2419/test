package db

import (
	"context"
	"errors"

	"github.com/google/uuid"
	logger "github.com/sirupsen/logrus"
)

const (
	getTemplateListQuery = `SELECT * FROM templates
		ORDER BY name ASC`

	getPublishedTemplateListQuery = `SELECT * FROM templates
		ORDER BY name ASC
		WHERE publish = true`

	createTemplateQuery = `INSERT INTO templates (
		name,
		description)
		VALUES ($1, $2) RETURNING id`

	getTemplateQuery = `SELECT *
		FROM templates WHERE id = $1`

	updateTemplateQuery = `UPDATE templates SET
		name = $1,
		description = $2
		where id = $3 AND publish = false`

	deleteTemplateQuery = `DELETE FROM templates WHERE id = $1`

	publishTempQuery = `UPDATE templates SET
	publish = true
	where id = $1`
)

type Template struct {
	ID          uuid.UUID `db:"id" json:"id"`
	Name        string    `db:"name" json:"name" validate:"required"`
	Description string    `db:"description" json:"description" validate:"required"`
	Publish     bool      `db:"publish json:"publish"`
}
type ErrorResponse struct {
	Code    string            `json:"code"`
	Message string            `json:"message"`
	Fields  map[string]string `json:"fields"`
}

// ValidateTemplate to publish
func ValidateTemplate(targets []TemplateTarget, steps []StageStep) (errorResponse map[string]ErrorResponse, valid bool) {
	validationErrors := make(map[string]string)

	if len(targets) == 0 {
		validationErrors["targets"] = "No targets added"
	}

	if len(steps) == 0 {
		validationErrors["steps"] = "No steps/stage added"
	} else {
		var holdsteps []StageStep
		var cyclesteps []StageStep
		var cycles uint16
		for _, s := range steps {
			switch s.Type {
			case "hold":
				holdsteps = append(holdsteps, s)
			case "cycle":
				cyclesteps = append(cyclesteps, s)
				cycles = s.RepeatCount
			}
		}
		if len(holdsteps) == 0 {
			validationErrors["holdstep"] = "No holdstep added"
		}
		if len(cyclesteps) == 0 {
			validationErrors["cyclestep"] = "No cyclestep added"
		}
		if cycles < 5 && cycles > 100 {
			validationErrors["repeatCount"] = "Invalid reapeat_count in cycle stage"
		}

	}
	if len(validationErrors) == 0 {
		valid = true
		return
	}

	errorResponse = map[string]ErrorResponse{"error": ErrorResponse{
		Code:    "invalid_data",
		Message: "Please provide valid template data",
		Fields:  validationErrors,
	},
	}

	return
}

func (s *pgStore) ListTemplates(ctx context.Context) (t []Template, err error) {
	err = s.db.Select(&t, getTemplateListQuery)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error listing templates")
		return
	}
	return
}

func (s *pgStore) ListPublishedTemplates(ctx context.Context) (t []Template, err error) {
	err = s.db.Select(&t, getTemplateListQuery)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error listing templates")
		return
	}
	return
}

func (s *pgStore) CreateTemplate(ctx context.Context, t Template) (createdTemp Template, err error) {

	var id uuid.UUID

	err = s.db.QueryRow(createTemplateQuery, t.Name, t.Description).Scan(&id)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error creating Template")
		return
	}

	err = s.db.Get(&createdTemp, getTemplateQuery, id)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error in getting Template")
		return
	}
	return
}

func (s *pgStore) UpdateTemplate(ctx context.Context, t Template) (err error) {
	result, err := s.db.Exec(
		updateTemplateQuery,
		t.Name,
		t.Description,
		t.ID,
	)

	if err != nil {
		logger.WithField("err", err.Error()).Error("Error updating Template")
		return
	}

	c, _ := result.RowsAffected()
	// check row count as no error is returned when row not found for update
	if c == 0 {
		err = errors.New("Record Not Found")
		logger.WithField("err", err.Error()).Error("Error Template not found")
	}

	return
}

func (s *pgStore) ShowTemplate(ctx context.Context, id uuid.UUID) (dbTemp Template, err error) {
	err = s.db.Get(&dbTemp, getTemplateQuery, id)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error fetching Template")
		return
	}

	return
}

func (s *pgStore) DeleteTemplate(ctx context.Context, id uuid.UUID) (err error) {

	// TBD: add delete cascade here
	result, err := s.db.Exec(
		deleteTemplateQuery,
		id,
	)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error deleting Template")
		return
	}

	c, _ := result.RowsAffected()
	// check row count as no error is returned when row not found for update
	if c == 0 {
		err = errors.New("Record Not Found")
		logger.WithField("err", err.Error()).Error("Error Template not found")
	}

	return
}

func (s *pgStore) PublishTemplate(ctx context.Context, id uuid.UUID) (err error) {

	_, err = s.db.Exec(
		publishTempQuery,
		id,
	)

	if err != nil {
		logger.WithField("err", err.Error()).Error("Error publishing Template")
		return
	}

	return
}
