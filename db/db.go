package db

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type Storer interface {
	InsertDyes(context.Context, []Dye) ([]Dye, error)
	InsertTargets(context.Context, []Target) error
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
	UpdateStepCount(context.Context) error
	ListSteps(context.Context, uuid.UUID) ([]Step, error)
	CreateStep(context.Context, Step) (Step, error)
	UpdateStep(context.Context, Step) error
	ShowStep(context.Context, uuid.UUID) (Step, error)
	DeleteStep(context.Context, uuid.UUID) error
	ListTemplateTargets(context.Context, uuid.UUID) ([]TemplateTarget, error)
	UpsertTemplateTarget(context.Context, []TemplateTarget, uuid.UUID) ([]TemplateTarget, error)
	ListExperiments(context.Context) ([]Experiment, error)
	CreateExperiment(context.Context, Experiment) (Experiment, error)
	ShowExperiment(context.Context, uuid.UUID) (Experiment, error)
	ListExpTemplateTargets(context.Context, uuid.UUID) ([]ExpTemplateTarget, error)
	UpsertExpTemplateTarget(context.Context, []ExpTemplateTarget, uuid.UUID) ([]ExpTemplateTarget, error)
	CreateSample(context.Context, Sample) (Sample, error)
	FindSamples(context.Context, string) ([]Sample, error)
	ListWells(context.Context, uuid.UUID) ([]Well, error)
	UpsertWells(context.Context, []Well, uuid.UUID) ([]Well, error)
	ShowWell(context.Context, uuid.UUID) (Well, error)
	DeleteWell(context.Context, uuid.UUID) error
	GetWellTarget(context.Context, int32, uuid.UUID) ([]WellTarget, error)
	UpsertWellTargets(context.Context, []WellTarget, uuid.UUID, bool) ([]WellTarget, error)
	ListStageSteps(context.Context, uuid.UUID) ([]StageStep, error)
	UpdateStartTimeExperiments(context.Context, time.Time, uuid.UUID) error
	ListConfTargets(context.Context, uuid.UUID) ([]TargetDetails, error)
	InsertResult(context.Context, []Result) ([]Result, error)
	ListWellTargets(context.Context, uuid.UUID) ([]WellTarget, error)
	UpdateStopTimeExperiments(context.Context, time.Time, uuid.UUID) error
	GetResult(context.Context, uuid.UUID) ([]Result, error)
	UpdateColorWell(context.Context, string, uuid.UUID) error
	ListExperimentTemperature(context.Context, uuid.UUID) ([]ExperimentTemperature, error)
	InsertExperimentTemperature(context.Context, ExperimentTemperature) error
	ListNotification(context.Context, uuid.UUID) ([]Notification, error)
	InsertNotification(context.Context, Notification) error
	MarkNotificationasRead(context.Context, uuid.UUID) error
}
