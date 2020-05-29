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
	Name        string    `db:"name" json:"name" validate:"required"`
	Description string    `db:"description" json:"description" validate:"required"`
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

	var id uuid.UUID

	err = s.db.QueryRow(createTemplateQuery, t.Name, t.Description).Scan(&id)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error creating Template")
		return
	}

	logger.WithField("id", id).Info("Inserted ID")

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
