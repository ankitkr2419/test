package db

import (
	"context"

	"github.com/google/uuid"
	logger "github.com/sirupsen/logrus"
)

const (
	getTemplateListQuery = `SELECT * FROM templates
		ORDER BY name ASC`

	createTemplateQuery = `INSERT INTO templates (
		name,
		description)
		VALUES ($1, $2) RETURNING id`

	getTemplateQuery = `SELECT id,
		name,
		description
		FROM templates WHERE id = $1`

	updateTemplateQuery = `UPDATE templates SET
		name = $1,
		description = $2
		where id = $3`

	deleteTemplateQuery = `DELETE FROM templates WHERE id = $1`
)

type Template struct {
	ID          uuid.UUID `db:"id" json:"id"`
	Name        string    `db:"name" json:"name"`
	Description string    `db:"description" json:"description"`
}

type myuuid uuid.UUID

func (t *myuuid) Scan(v interface{}) error {
	id, err := uuid.Parse(string(v.([]byte)))
	if err != nil {
		return err
	}
	*t = myuuid(id)
	return nil
}

func (t *Template) Validate() (errorResponse map[string]ErrorResponse, valid bool) {
	fieldErrors := make(map[string]string)

	if t.Name == "" {
		fieldErrors["name"] = "Can't be blank"
	}
	if t.Description == "" {
		fieldErrors["description"] = "Can't be blank"
	}

	if len(fieldErrors) == 0 {
		valid = true
		return
	}

	errorResponse = map[string]ErrorResponse{"error": ErrorResponse{
		Code:    "invalid_data",
		Message: "Please provide valid target data",
		Fields:  fieldErrors,
	},
	}

	//TODO: Ask what other validations are expected

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

func (s *pgStore) CreateTemplate(ctx context.Context, t Template) (createdTemp Template, err error) {
	row, err := s.db.Query(createTemplateQuery, t.Name, t.Description)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error creating Template")
		return
	}
	defer row.Close()
	var id uuid.UUID

	for row.Next() {
		err = row.Scan((*myuuid)(&id))
		if err != nil {
			logger.WithField("err", err.Error()).Error("Error in scanning rows")
		}
	}

	err = s.db.Get(&createdTemp, getTemplateQuery, id)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error in getting Template")
		return
	}
	return
}

func (s *pgStore) UpdateTemplate(ctx context.Context, t Template) (updatedTemp Template, err error) {
	var dbTemp Template
	err = s.db.Get(&dbTemp, getTemplateQuery, t.ID)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error in getting Template")
		return
	}

	_, err = s.db.Exec(
		updateTemplateQuery,
		t.Name,
		t.Description,
		t.ID,
	)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error updating Template")
		return
	}

	err = s.db.Get(&updatedTemp, getTemplateQuery, t.ID)

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
	_, err = s.db.Exec(
		deleteTemplateQuery,
		id,
	)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error deleting Template")
		return
	}

	return
}
