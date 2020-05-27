package db

import (
	"context"

	"github.com/google/uuid"
)

type Storer interface {
	ListTarget(context.Context) ([]Target, error)
	CreateTarget(context.Context, Target) (Target, error)
	UpdateTarget(context.Context, Target) (Target, error)
	ShowTarget(context.Context, uuid.UUID) (Target, error)
	DeleteTarget(context.Context, uuid.UUID) error
	ListTemplates(context.Context) ([]Template, error)
	CreateTemplate(context.Context, Template) (Template, error)
	UpdateTemplate(context.Context, Template) (Template, error)
	ShowTemplate(context.Context, uuid.UUID) (Template, error)
	DeleteTemplate(context.Context, uuid.UUID) error
}
