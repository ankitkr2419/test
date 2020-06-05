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
	ListStages(context.Context, uuid.UUID) ([]Stage, error)
	CreateStage(context.Context, Stage) (Stage, error)
	UpdateStage(context.Context, Stage) error
	ShowStage(context.Context, uuid.UUID) (Stage, error)
	DeleteStage(context.Context, uuid.UUID) error
	ListSteps(context.Context, uuid.UUID) ([]Step, error)
	CreateStep(context.Context, Step) (Step, error)
	UpdateStep(context.Context, Step) error
	ShowStep(context.Context, uuid.UUID) (Step, error)
	DeleteStep(context.Context, uuid.UUID) error
	ListTemplateTargets(context.Context, uuid.UUID) ([]TemplateTarget, error)
	// CreateTemplateTarget(context.Context, TemplateTarget) (TemplateTarget, error)
	// UpdateTemplateTarget(context.Context, TemplateTarget) error
	// ShowTemplateTarget(context.Context, uuid.UUID, uuid.UUID) (TemplateTarget, error)
	// DeleteTemplateTarget(context.Context, uuid.UUID, uuid.UUID) error
	UpsertTemplateTarget(context.Context, []TemplateTarget, uuid.UUID) ([]TemplateTarget, error)
}
