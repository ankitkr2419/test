package db

import (
	"context"

	logger "github.com/sirupsen/logrus"
)

const (
	getTemplateListQuery = `SELECT * FROM templates
		ORDER BY name ASC`
)

type Template struct {
	ID   int    `db:"id" json:"id"`
	Name string `db:"name" json:"name"`
}

func (s *pgStore) ListTemplates(ctx context.Context) (t []Template, err error) {
	err = s.db.Select(&t, getTemplateListQuery)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error listing templates")
		return
	}
	return
}
