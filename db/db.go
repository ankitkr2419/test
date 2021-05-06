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
	CreateStages(context.Context, []Stage) ([]Stage, error)
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
	UpdateStartTimeExperiments(context.Context, time.Time, uuid.UUID, uint16) error
	ListConfTargets(context.Context, uuid.UUID) ([]TargetDetails, error)
	InsertResult(context.Context, []Result) ([]Result, error)
	ListWellTargets(context.Context, uuid.UUID) ([]WellTarget, error)
	UpdateStopTimeExperiments(context.Context, time.Time, uuid.UUID, string) error
	GetResult(context.Context, uuid.UUID) ([]Result, error)
	UpdateColorWell(context.Context, string, uuid.UUID) error
	PublishTemplate(context.Context, uuid.UUID) error
	ListPublishedTemplates(context.Context) ([]Template, error)
	ListExperimentTemperature(context.Context, uuid.UUID) ([]ExperimentTemperature, error)
	InsertExperimentTemperature(context.Context, ExperimentTemperature) error
	ListNotification(context.Context, uuid.UUID) ([]Notification, error)
	InsertNotification(context.Context, Notification) error
	MarkNotificationasRead(context.Context, uuid.UUID) error
	InsertUser(context.Context, User) error
	ValidateUser(context.Context, User) error
	CheckIfICTargetAdded(context.Context, uuid.UUID) (WarnResponse, error)
	InsertMotor(context.Context, []Motor) error
	InsertConsumableDistance(context.Context, []ConsumableDistance) error
	InsertTipsTubes(context.Context, []TipsTubes) error
	InsertCartridge(context.Context, []Cartridge, []CartridgeWells) error
	ListMotors() ([]Motor, error)
	ListConsDistances() ([]ConsumableDistance, error)
	ListTipsTubes(ttype string) (tipstubes []TipsTubes, err error)
	ShowTip(id int64) (TipsTubes, error)
	ListCartridges() ([]Cartridge, error)
	ListCartridgeWells() ([]CartridgeWells, error)
	ShowPiercing(context.Context, uuid.UUID) (Piercing, error)
	ShowTipOperation(context.Context, uuid.UUID) (TipOperation, error)
	CreateTipOperation(context.Context, TipOperation) (TipOperation, error)
	ShowRecipe(context.Context, uuid.UUID) (Recipe, error)
	CreateRecipe(context.Context, Recipe) (Recipe, error)
	DeleteRecipe(context.Context, uuid.UUID) error
	UpdateRecipe(context.Context, Recipe) error
	ListRecipes(context.Context) ([]Recipe, error)
	// ListProcesses by Recipe ID
	ListProcesses(context.Context, uuid.UUID) ([]Process, error)
	ShowProcess(context.Context, uuid.UUID) (Process, error)
	CreateProcess(context.Context, Process) (Process, error)
	DeleteProcess(context.Context, uuid.UUID) error
	UpdateProcess(context.Context, Process) error
	DuplicateProcess(context.Context, uuid.UUID) (Process, error)
	ListPiercing(context.Context) ([]Piercing, error)
	CreatePiercing(context.Context, Piercing) (Piercing, error)
	UpdatePiercing(context.Context, Piercing) error
	ShowAspireDispense(context.Context, uuid.UUID) (AspireDispense, error)
	ListAspireDispense(context.Context) ([]AspireDispense, error)
	CreateAspireDispense(context.Context, AspireDispense) (AspireDispense, error)
	UpdateAspireDispense(context.Context, AspireDispense) error
	ShowTipDocking(context.Context, uuid.UUID) (TipDock, error)
	CreateTipDocking(context.Context, TipDock) (TipDock, error)
	ShowHeating(ctx context.Context, id uuid.UUID) (heating Heating, err error)
	CreateHeating(context.Context, Heating) (Heating, error)
	CreateAttachDetach(context.Context, AttachDetach) (AttachDetach, error)
	ShowAttachDetach(ctx context.Context, processID uuid.UUID) (AttachDetach, error)
	CreateDelay(context.Context, Delay) (Delay, error)
	ShowDelay(ctx context.Context, id uuid.UUID) (delay Delay, err error)
	CreateShaking(context.Context, Shaker) (Shaker, error)
	ShowShaking(ctx context.Context, id uuid.UUID) (shaking Shaker, err error)
	UpdateTipOperation(ctx context.Context, t TipOperation) (err error)
	UpdateDelay(ctx context.Context, d Delay) (err error)
	UpdateShaking(ctx context.Context, sh Shaker) (err error)
	UpdateAttachDetach(ctx context.Context, a AttachDetach) (err error)
	UpdateTipDock(ctx context.Context, t TipDock) (err error)
	UpdateHeating(ctx context.Context, ht Heating) (err error)
	UpdateProcessName(ctx context.Context, id uuid.UUID, processType string, process interface{}) (err error)
	ShowUser(ctx context.Context, username string) (user User, err error)
}
