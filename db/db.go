package db

import (
	"context"

	"github.com/google/uuid"
)

type Storer interface {
	ListTargets(context.Context) ([]Target, error)
	ListTemplates(context.Context) ([]Template, error)
	CreateTemplate(context.Context, Template) (Template, error)
	UpdateTemplate(context.Context, Template) error
	ShowTemplate(context.Context, uuid.UUID) (Template, error)
	DeleteTemplate(context.Context, uuid.UUID) error
}
