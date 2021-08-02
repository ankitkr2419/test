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
	return args.Error(0)
}

func (m *DBMockStore) DeleteTemplate(ctx context.Context, id uuid.UUID) (err error) {
	args := m.Called(ctx, id)
	return args.Error(0)
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
	return args.Error(0)
}

func (m *DBMockStore) DeleteStage(ctx context.Context, id uuid.UUID) (err error) {
	args := m.Called(ctx, id)
	return args.Error(0)
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
	return args.Error(0)
}

func (m *DBMockStore) DeleteStep(ctx context.Context, id uuid.UUID) (err error) {
	args := m.Called(ctx, id)
	return args.Error(0)
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
	return args.Error(0)
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
	return args.Error(0)
}

func (m *DBMockStore) DeleteSample(ctx context.Context, id uuid.UUID) (err error) {
	args := m.Called(ctx, id)
	return args.Error(0)
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
	return args.Error(0)
}

func (m *DBMockStore) ShowWell(ctx context.Context, id uuid.UUID) (w Well, err error) {
	args := m.Called(ctx, id)
	return args.Get(0).(Well), args.Error(1)
}

func (m *DBMockStore) UpdateStepCount(ctx context.Context) (err error) {
	args := m.Called(ctx)
	return args.Error(0)
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

func (m *DBMockStore) UpdateStartTimeExperiments(ctx context.Context, t time.Time, experimentID uuid.UUID, repeatCycle uint16, state string) (err error) {
	args := m.Called(ctx, t, experimentID, repeatCycle, state)
	return args.Error(0)
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
	return args.Error(0)
}

func (m *DBMockStore) GetResult(ctx context.Context, id uuid.UUID) (result []Result, err error) {
	args := m.Called(ctx, id)
	return args.Get(0).([]Result), args.Error(1)
}

func (m *DBMockStore) UpdateColorWell(ctx context.Context, s string, id uuid.UUID) (err error) {
	args := m.Called(ctx, s, id)
	return args.Error(0)
}

func (m *DBMockStore) PublishTemplate(ctx context.Context, id uuid.UUID) (err error) {
	args := m.Called(ctx, id)
	return args.Error(0)
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
	return args.Error(0)
}

func (m *DBMockStore) ListNotification(ctx context.Context, id uuid.UUID) (result []Notification, err error) {
	args := m.Called(ctx, id)
	return args.Get(0).([]Notification), args.Error(1)
}

func (m *DBMockStore) InsertNotification(ctx context.Context, r Notification) (err error) {
	args := m.Called(ctx, r)
	return args.Error(0)
}

func (m *DBMockStore) MarkNotificationasRead(ctx context.Context, id uuid.UUID) (err error) {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *DBMockStore) InsertUser(ctx context.Context, u User) (err error) {
	args := m.Called(ctx, u)
	return args.Error(0)
}

func (m *DBMockStore) ValidateUser(ctx context.Context, u User) (us User, err error) {
	args := m.Called(ctx, u)
	return args.Get(0).(User), args.Error(1)
}

func (m *DBMockStore) CheckIfICTargetAdded(ctx context.Context, id uuid.UUID) (WarnResponse, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(WarnResponse), args.Error(1)
}

func (m *DBMockStore) InsertMotor(ctx context.Context, motors []Motor) (err error) {
	args := m.Called(ctx, motors)
	return args.Error(0)
}

func (m *DBMockStore) InsertConsumableDistance(ctx context.Context, c []ConsumableDistance) (err error) {
	args := m.Called(ctx, c)
	return args.Error(0)
}

func (m *DBMockStore) InsertTipsTubes(ctx context.Context, t []TipsTubes) (err error) {
	args := m.Called(ctx, t)
	return args.Error(0)
}

func (m *DBMockStore) InsertCartridge(ctx context.Context, c []Cartridge, w []CartridgeWells) (err error) {
	args := m.Called(ctx, c, w)
	return args.Error(0)
}

func (m *DBMockStore) ListCartridges(ctx context.Context) (c []Cartridge, err error) {
	args := m.Called(ctx)
	return args.Get(0).([]Cartridge), args.Error(1)
}

func (m *DBMockStore) ListTipsTubes(ttype string) (t []TipsTubes, err error) {
	args := m.Called(ttype)
	return args.Get(0).([]TipsTubes), args.Error(1)
}

func (m *DBMockStore) ListConsDistances() (c []ConsumableDistance, err error) {
	args := m.Called()
	return args.Get(0).([]ConsumableDistance), args.Error(1)
}

func (m *DBMockStore) ListMotors(ctx context.Context) (motor []Motor, err error) {
	args := m.Called(ctx)
	return args.Get(0).([]Motor), args.Error(1)
}

func (m *DBMockStore) ShowPiercing(ctx context.Context, id uuid.UUID) (p Piercing, err error) {
	args := m.Called(ctx, id)
	return args.Get(0).(Piercing), args.Error(1)
}

func (m *DBMockStore) ListPiercing(ctx context.Context) (p []Piercing, err error) {
	args := m.Called(ctx)
	return args.Get(0).([]Piercing), args.Error(1)
}

func (m *DBMockStore) CreatePiercing(ctx context.Context, p Piercing, id uuid.UUID) (pi Piercing, err error) {
	args := m.Called(ctx, p, id)
	return args.Get(0).(Piercing), args.Error(1)
}

func (m *DBMockStore) UpdatePiercing(ctx context.Context, p Piercing) (err error) {
	args := m.Called(ctx, p)
	return args.Error(0)
}

func (m *DBMockStore) ShowRecipe(ctx context.Context, id uuid.UUID) (p Recipe, err error) {
	args := m.Called(ctx, id)
	return args.Get(0).(Recipe), args.Error(1)
}

func (m *DBMockStore) ListRecipes(ctx context.Context) (p []Recipe, err error) {
	args := m.Called(ctx)
	return args.Get(0).([]Recipe), args.Error(1)
}

func (m *DBMockStore) CreateRecipe(ctx context.Context, p Recipe) (pi Recipe, err error) {
	args := m.Called(ctx, p)
	return args.Get(0).(Recipe), args.Error(1)
}

func (m *DBMockStore) DeleteRecipe(ctx context.Context, id uuid.UUID) (err error) {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *DBMockStore) UpdateRecipe(ctx context.Context, p Recipe) (err error) {
	args := m.Called(ctx, p)
	return args.Error(0)
}

func (m *DBMockStore) ShowProcess(ctx context.Context, id uuid.UUID) (p Process, err error) {
	args := m.Called(ctx, id)
	return args.Get(0).(Process), args.Error(1)
}

func (m *DBMockStore) ListProcesses(ctx context.Context, id uuid.UUID) (p []Process, err error) {
	args := m.Called(ctx, id)
	return args.Get(0).([]Process), args.Error(1)
}

func (m *DBMockStore) CreateProcess(ctx context.Context, p Process) (pi Process, err error) {
	args := m.Called(ctx, p)
	return args.Get(0).(Process), args.Error(1)
}

func (m *DBMockStore) DeleteProcess(ctx context.Context, id uuid.UUID) (err error) {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *DBMockStore) UpdateProcess(ctx context.Context, p Process) (err error) {
	args := m.Called(ctx, p)
	return args.Error(0)
}

func (m *DBMockStore) ListCartridgeWells() (cw []CartridgeWells, err error) {
	args := m.Called()
	return args.Get(0).([]CartridgeWells), args.Error(1)
}

func (m *DBMockStore) ShowAspireDispense(ctx context.Context, id uuid.UUID) (ad AspireDispense, err error) {
	args := m.Called(ctx, id)
	return args.Get(0).(AspireDispense), args.Error(1)
}

func (m *DBMockStore) ListAspireDispense(ctx context.Context) (ad []AspireDispense, err error) {
	args := m.Called(ctx)
	return args.Get(0).([]AspireDispense), args.Error(1)
}

func (m *DBMockStore) CreateAspireDispense(ctx context.Context, ad AspireDispense, id uuid.UUID) (a AspireDispense, err error) {
	args := m.Called(ctx, ad, id)
	return args.Get(0).(AspireDispense), args.Error(1)
}

func (m *DBMockStore) UpdateAspireDispense(ctx context.Context, ad AspireDispense) (err error) {
	args := m.Called(ctx, ad)
	return args.Error(0)
}

func (m *DBMockStore) UpdateAttachDetach(ctx context.Context, ad AttachDetach) (err error) {
	args := m.Called(ctx, ad)
	return args.Error(0)
}

func (m *DBMockStore) UpdateShaking(ctx context.Context, sh Shaker) (err error) {
	args := m.Called(ctx, sh)
	return args.Error(0)
}

func (m *DBMockStore) UpdateTipDock(ctx context.Context, td TipDock) (err error) {
	args := m.Called(ctx, td)
	return args.Error(0)
}

func (m *DBMockStore) UpdateHeating(ctx context.Context, ht Heating) (err error) {
	args := m.Called(ctx, ht)
	return args.Error(0)
}

func (m *DBMockStore) CreateAttachDetach(ctx context.Context, ad AttachDetach, id uuid.UUID) (createdAttachDetach AttachDetach, err error) {
	args := m.Called(ctx, ad, id)
	return args.Get(0).(AttachDetach), args.Error(1)
}

func (m *DBMockStore) CreateDelay(ctx context.Context, d Delay, id uuid.UUID) (createdDelay Delay, err error) {
	args := m.Called(ctx, d, id)
	return args.Get(0).(Delay), args.Error(1)
}

func (m *DBMockStore) CreateHeating(ctx context.Context, h Heating, id uuid.UUID) (createdHeating Heating, err error) {
	args := m.Called(ctx, h, id)
	return args.Get(0).(Heating), args.Error(1)
}

func (m *DBMockStore) CreateShaking(ctx context.Context, sh Shaker, id uuid.UUID) (createdShaking Shaker, err error) {
	args := m.Called(ctx, sh, id)
	return args.Get(0).(Shaker), args.Error(1)
}

func (m *DBMockStore) CreateTipDocking(ctx context.Context, t TipDock, id uuid.UUID) (createdTipDocking TipDock, err error) {
	args := m.Called(ctx, t, id)
	return args.Get(0).(TipDock), args.Error(1)
}

func (m *DBMockStore) CreateTipOperation(ctx context.Context, t TipOperation, id uuid.UUID) (createdTipOperation TipOperation, err error) {
	args := m.Called(ctx, t, id)
	return args.Get(0).(TipOperation), args.Error(1)
}

func (m *DBMockStore) ShowAttachDetach(ctx context.Context, processID uuid.UUID) (ad AttachDetach, err error) {
	args := m.Called(ctx, processID)
	return args.Get(0).(AttachDetach), args.Error(1)
}

func (m *DBMockStore) ShowDelay(ctx context.Context, processID uuid.UUID) (createdDelay Delay, err error) {
	args := m.Called(ctx, processID)
	return args.Get(0).(Delay), args.Error(1)
}

func (m *DBMockStore) ShowHeating(ctx context.Context, id uuid.UUID) (heating Heating, err error) {
	args := m.Called(ctx, id)
	return args.Get(0).(Heating), args.Error(1)
}

func (m *DBMockStore) ShowShaking(ctx context.Context, id uuid.UUID) (shaking Shaker, err error) {
	args := m.Called(ctx, id)
	return args.Get(0).(Shaker), args.Error(1)
}

func (m *DBMockStore) ShowTipDocking(ctx context.Context, pid uuid.UUID) (td TipDock, err error) {
	args := m.Called(ctx, pid)
	return args.Get(0).(TipDock), args.Error(1)
}

func (m *DBMockStore) ShowTipOperation(ctx context.Context, id uuid.UUID) (dbTipOperation TipOperation, err error) {
	args := m.Called(ctx, id)
	return args.Get(0).(TipOperation), args.Error(1)
}

func (m *DBMockStore) ShowTip(id int64) (tip TipsTubes, err error) {
	args := m.Called(id)
	return args.Get(0).(TipsTubes), args.Error(1)
}

func (m *DBMockStore) DuplicateProcess(ctx context.Context, id uuid.UUID) (Process, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(Process), args.Error(1)
}

func (m *DBMockStore) RearrangeProcesses(ctx context.Context, id uuid.UUID, ps []ProcessSequence) ([]Process, error) {
	args := m.Called(ctx, id, ps)
	return args.Get(0).([]Process), args.Error(1)
}

func (m *DBMockStore) ShowUser(ctx context.Context, username string) (user User, err error) {
	args := m.Called(ctx, username)
	return args.Get(0).(User), args.Error(1)
}

func (m *DBMockStore) UpdateDelay(ctx context.Context, d Delay) (err error) {
	args := m.Called(ctx, d)
	return args.Error(0)
}

func (m *DBMockStore) UpdateTipOperation(ctx context.Context, t TipOperation) (err error) {
	args := m.Called(ctx, t)
	return args.Error(0)
}

func (m *DBMockStore) InsertUserAuths(ctx context.Context, username string) (authID uuid.UUID, err error) {
	args := m.Called(ctx, username)
	return args.Get(0).(uuid.UUID), args.Error(1)
}

func (m *DBMockStore) ShowUserAuth(ctx context.Context, username string, authID uuid.UUID) (ua UserAuth, err error) {
	args := m.Called(ctx, username, authID)
	return args.Get(0).(UserAuth), args.Error(1)
}

func (m *DBMockStore) DeleteUserAuth(ctx context.Context, userAuth UserAuth) (err error) {
	args := m.Called(ctx, userAuth)
	return args.Error(0)
}

func (m *DBMockStore) InsertAuditLog(ctx context.Context, al AuditLog) (err error) {
	args := m.Called(ctx, al)
	return args.Error(0)
}

func (m *DBMockStore) ShowAuditLog(ctx context.Context) (al AuditLog, err error) {
	args := m.Called(ctx)
	return args.Get(0).(AuditLog), args.Error(1)
}

func (m *DBMockStore) ListTipsTubesByPosition(ctx context.Context, ttype string, position int64) (tipstubes []TipsTubes, err error) {
	args := m.Called(ctx, ttype, position)
	return args.Get(0).([]TipsTubes), args.Error(1)
}

func (m *DBMockStore) AddAuditLog(ctx context.Context, activity ActivityType, state StateType, oprType OperationType, deck, description string) (err error) {
	args := m.Called(ctx, activity, state, oprType, deck, description)
	return args.Error(0)
}

func (m *DBMockStore) GetTargetByName(ctx context.Context, name string) (t Target, err error) {
	args := m.Called(ctx, name)
	return args.Get(0).(Target), args.Error(1)
}

func (m *DBMockStore) DeleteCartridge(ctx context.Context, id int64) (err error) {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *DBMockStore) DeleteMotor(ctx context.Context, id int) (err error) {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *DBMockStore) DeleteTipTube(ctx context.Context, id int64) (err error) {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *DBMockStore) FinishTemplate(ctx context.Context, id uuid.UUID) (err error) {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func (m *DBMockStore) ListFinishedTemplates(ctx context.Context) (t []Template, err error) {
	args := m.Called(ctx)
	return args.Get(0).([]Template), args.Error(1)
}

func (m *DBMockStore) UpdateEstimatedTime(ctx context.Context, id uuid.UUID, et int64) (err error) {
	args := m.Called(ctx, id, et)
	return args.Error(0)
}

func (m *DBMockStore) UpdateMotor(ctx context.Context, motor Motor) (err error) {
	args := m.Called(ctx, motor)
	return args.Error(0)
}

func (m *DBMockStore) UpdateUser(ctx context.Context, u User, oldName string) (err error) {
	args := m.Called(ctx, u, oldName)
	return args.Error(0)
}

func (m *DBMockStore) DeleteUnfinishedTemplates(ctx context.Context) (err error) {
	args := m.Called(ctx)
	return args.Error(0)
}
