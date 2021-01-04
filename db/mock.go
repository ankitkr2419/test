package db

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

type DBMockStore struct {
	mock.Mock
}

func (m *DBMockStore) ListTargets(ctx context.Context) (t []Target, err error) {
	args := m.Called(ctx)
	return args.Get(0).([]Target), args.Error(1)
}

func (m *DBMockStore) ListTemplates(ctx context.Context) (t []Template, err error) {
	args := m.Called(ctx)
	return args.Get(0).([]Template), args.Error(1)
}

func (m *DBMockStore) CreateTemplate(ctx context.Context, tp Template) (t Template, err error) {
	args := m.Called(ctx, tp)
	return args.Get(0).(Template), args.Error(1)
}

func (m *DBMockStore) UpdateTemplate(ctx context.Context, t Template) (err error) {
	args := m.Called(ctx, t)
	return args.Error(1)
}

func (m *DBMockStore) DeleteTemplate(ctx context.Context, id uuid.UUID) (err error) {
	args := m.Called(ctx, id)
	return args.Error(1)
}

func (m *DBMockStore) ShowTemplate(ctx context.Context, id uuid.UUID) (t Template, err error) {
	args := m.Called(ctx, id)
	return args.Get(0).(Template), args.Error(1)
}

func (m *DBMockStore) ListStages(ctx context.Context, t uuid.UUID) (s []Stage, err error) {
	args := m.Called(ctx, t)
	return args.Get(0).([]Stage), args.Error(1)
}

func (m *DBMockStore) CreateStages(ctx context.Context, stg []Stage) (s []Stage, err error) {
	args := m.Called(ctx, stg)
	return args.Get(0).([]Stage), args.Error(1)
}

func (m *DBMockStore) UpdateStage(ctx context.Context, s Stage) (err error) {
	args := m.Called(ctx, s)
	return args.Error(1)
}

func (m *DBMockStore) DeleteStage(ctx context.Context, id uuid.UUID) (err error) {
	args := m.Called(ctx, id)
	return args.Error(1)
}

func (m *DBMockStore) ShowStage(ctx context.Context, id uuid.UUID) (s Stage, err error) {
	args := m.Called(ctx, id)
	return args.Get(0).(Stage), args.Error(1)
}

func (m *DBMockStore) ListSteps(ctx context.Context, stg uuid.UUID) (s []Step, err error) {
	args := m.Called(ctx, stg)
	return args.Get(0).([]Step), args.Error(1)
}

func (m *DBMockStore) CreateStep(ctx context.Context, stg Step) (s Step, err error) {
	args := m.Called(ctx, stg)
	return args.Get(0).(Step), args.Error(1)
}

func (m *DBMockStore) UpdateStep(ctx context.Context, s Step) (err error) {
	args := m.Called(ctx, s)
	return args.Error(1)
}

func (m *DBMockStore) DeleteStep(ctx context.Context, id uuid.UUID) (err error) {
	args := m.Called(ctx, id)
	return args.Error(1)
}

func (m *DBMockStore) ShowStep(ctx context.Context, id uuid.UUID) (s Step, err error) {
	args := m.Called(ctx, id)
	return args.Get(0).(Step), args.Error(1)
}

func (m *DBMockStore) UpsertTemplateTarget(ctx context.Context, tt []TemplateTarget, id uuid.UUID) (t []TemplateTarget, err error) {
	args := m.Called(ctx, tt, id)
	return args.Get(0).([]TemplateTarget), args.Error(1)
}

func (m *DBMockStore) ListTemplateTargets(ctx context.Context, id uuid.UUID) (t []TemplateTarget, err error) {
	args := m.Called(ctx, id)
	return args.Get(0).([]TemplateTarget), args.Error(1)
}

func (m *DBMockStore) ListExperiments(ctx context.Context) (t []Experiment, err error) {
	args := m.Called(ctx)
	return args.Get(0).([]Experiment), args.Error(1)
}

func (m *DBMockStore) CreateExperiment(ctx context.Context, tp Experiment) (t Experiment, err error) {
	args := m.Called(ctx, tp)
	return args.Get(0).(Experiment), args.Error(1)
}

func (m *DBMockStore) ShowExperiment(ctx context.Context, id uuid.UUID) (t Experiment, err error) {
	args := m.Called(ctx, id)
	return args.Get(0).(Experiment), args.Error(1)
}

func (m *DBMockStore) UpsertExpTemplateTarget(ctx context.Context, tt []ExpTemplateTarget, id uuid.UUID) (t []ExpTemplateTarget, err error) {
	args := m.Called(ctx, tt, id)
	return args.Get(0).([]ExpTemplateTarget), args.Error(1)
}

func (m *DBMockStore) ListExpTemplateTargets(ctx context.Context, id uuid.UUID) (t []ExpTemplateTarget, err error) {
	args := m.Called(ctx, id)
	return args.Get(0).([]ExpTemplateTarget), args.Error(1)
}

func (m *DBMockStore) InsertDyes(ctx context.Context, d []Dye) (dyes []Dye, err error) {
	args := m.Called(ctx, d)
	return args.Get(0).([]Dye), args.Error(1)
}

func (m *DBMockStore) InsertTargets(ctx context.Context, t []Target) (err error) {
	args := m.Called(ctx, t)
	return args.Error(1)
}

func (m *DBMockStore) ListSamples(ctx context.Context) (s []Sample, err error) {
	args := m.Called(ctx)
	return args.Get(0).([]Sample), args.Error(1)
}

func (m *DBMockStore) CreateSample(ctx context.Context, stg Sample) (s Sample, err error) {
	args := m.Called(ctx, stg)
	return args.Get(0).(Sample), args.Error(1)
}

func (m *DBMockStore) UpdateSample(ctx context.Context, s Sample) (err error) {
	args := m.Called(ctx, s)
	return args.Error(1)
}

func (m *DBMockStore) DeleteSample(ctx context.Context, id uuid.UUID) (err error) {
	args := m.Called(ctx, id)
	return args.Error(1)
}

func (m *DBMockStore) ShowSample(ctx context.Context, id uuid.UUID) (s Sample, err error) {
	args := m.Called(ctx, id)
	return args.Get(0).(Sample), args.Error(1)
}

func (m *DBMockStore) FindSamples(ctx context.Context, text string) (s []Sample, err error) {
	args := m.Called(ctx, text)
	return args.Get(0).([]Sample), args.Error(1)
}

func (m *DBMockStore) ListWells(ctx context.Context, experimentID uuid.UUID) (w []Well, err error) {
	args := m.Called(ctx, experimentID)
	return args.Get(0).([]Well), args.Error(1)
}

func (m *DBMockStore) UpsertWells(ctx context.Context, w []Well, id uuid.UUID) (wdb []Well, err error) {
	args := m.Called(ctx, w, id)
	return args.Get(0).([]Well), args.Error(1)
}

func (m *DBMockStore) DeleteWell(ctx context.Context, id uuid.UUID) (err error) {
	args := m.Called(ctx, id)
	return args.Error(1)
}

func (m *DBMockStore) ShowWell(ctx context.Context, id uuid.UUID) (w Well, err error) {
	args := m.Called(ctx, id)
	return args.Get(0).(Well), args.Error(1)
}

func (m *DBMockStore) UpdateStepCount(ctx context.Context) (err error) {
	args := m.Called(ctx)
	return args.Error(1)
}

func (m *DBMockStore) GetWellTarget(ctx context.Context, i int32, wellID uuid.UUID) (w []WellTarget, err error) {
	args := m.Called(ctx, i, wellID)
	return args.Get(0).([]WellTarget), args.Error(1)
}

func (m *DBMockStore) UpsertWellTargets(ctx context.Context, w []WellTarget, id uuid.UUID, s bool) (wt []WellTarget, err error) {
	args := m.Called(ctx, w, id, s)
	return args.Get(0).([]WellTarget), args.Error(1)
}
func (m *DBMockStore) ListStageSteps(ctx context.Context, id uuid.UUID) (s []StageStep, err error) {
	args := m.Called(ctx, id)
	return args.Get(0).([]StageStep), args.Error(1)
}
func (m *DBMockStore) UpdateStartTimeExperiments(ctx context.Context, t time.Time, id uuid.UUID, i uint16) (err error) {
	args := m.Called(ctx, t, id, i)
	return args.Error(1)
}

func (m *DBMockStore) ListConfTargets(ctx context.Context, id uuid.UUID) (t []TargetDetails, err error) {
	args := m.Called(ctx, id)
	return args.Get(0).([]TargetDetails), args.Error(1)
}
func (m *DBMockStore) InsertResult(ctx context.Context, r []Result) (dbR []Result, err error) {
	args := m.Called(ctx, r)
	return args.Get(0).([]Result), args.Error(1)
}

func (m *DBMockStore) ListWellTargets(ctx context.Context, wellID uuid.UUID) (w []WellTarget, err error) {
	args := m.Called(ctx, wellID)
	return args.Get(0).([]WellTarget), args.Error(1)
}

func (m *DBMockStore) UpdateStopTimeExperiments(ctx context.Context, t time.Time, id uuid.UUID, s string) (err error) {
	args := m.Called(ctx, t, id, s)
	return args.Error(1)
}

func (m *DBMockStore) GetResult(ctx context.Context, id uuid.UUID) (result []Result, err error) {
	args := m.Called(ctx, id)
	return args.Get(0).([]Result), args.Error(1)
}

func (m *DBMockStore) UpdateColorWell(ctx context.Context, s string, id uuid.UUID) (err error) {
	args := m.Called(ctx, s, id)
	return args.Error(1)
}

func (m *DBMockStore) PublishTemplate(ctx context.Context, id uuid.UUID) (err error) {
	args := m.Called(ctx, id)
	return args.Error(1)
}

func (m *DBMockStore) ListPublishedTemplates(ctx context.Context) (t []Template, err error) {
	args := m.Called(ctx)
	return args.Get(0).([]Template), args.Error(1)
}

func (m *DBMockStore) ListExperimentTemperature(ctx context.Context, id uuid.UUID) (result []ExperimentTemperature, err error) {
	args := m.Called(ctx, id)
	return args.Get(0).([]ExperimentTemperature), args.Error(1)
}

func (m *DBMockStore) InsertExperimentTemperature(ctx context.Context, r ExperimentTemperature) (err error) {
	args := m.Called(ctx, r)
	return args.Error(1)
}

func (m *DBMockStore) ListNotification(ctx context.Context, id uuid.UUID) (result []Notification, err error) {
	args := m.Called(ctx, id)
	return args.Get(0).([]Notification), args.Error(1)
}

func (m *DBMockStore) InsertNotification(ctx context.Context, r Notification) (err error) {
	args := m.Called(ctx, r)
	return args.Error(1)
}

func (m *DBMockStore) MarkNotificationasRead(ctx context.Context, id uuid.UUID) (err error) {
	args := m.Called(ctx, id)
	return args.Error(1)
}

func (m *DBMockStore) InsertUser(ctx context.Context, u User) (err error) {
	args := m.Called(ctx, u)
	return args.Error(1)
}

func (m *DBMockStore) ValidateUser(ctx context.Context, u User) (err error) {
	args := m.Called(ctx, u)
	return args.Error(1)
}

func (m *DBMockStore) CheckIfICTargetAdded(ctx context.Context, id uuid.UUID) (WarnResponse, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(WarnResponse), args.Error(1)
}

func (m *DBMockStore) InsertMotor(ctx context.Context, motors []Motor) (err error) {
	args := m.Called(ctx, motors)
	return args.Error(1)
}

func (m *DBMockStore) InsertConsumableDistance(ctx context.Context, c []ConsumableDistance) (err error) {
	args := m.Called(ctx, c)
	return args.Error(1)
}

func (m *DBMockStore) InsertLabware(ctx context.Context, l []Labware) (err error) {
	args := m.Called(ctx, l)
	return args.Error(1)
}

func (m *DBMockStore) InsertTipsTubes(ctx context.Context, t []TipsTubes) (err error) {
	args := m.Called(ctx, t)
	return args.Error(1)
}

func (m *DBMockStore) InsertCartridge(ctx context.Context, c []Cartridge) (err error) {
	args := m.Called(ctx, c)
	return args.Error(1)
}
