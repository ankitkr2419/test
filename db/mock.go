package db

import (
	"context"

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

func (m *DBMockStore) CreateStage(ctx context.Context, stg Stage) (s Stage, err error) {
	args := m.Called(ctx, stg)
	return args.Get(0).(Stage), args.Error(1)
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

func (m *DBMockStore) UpsertTemplateTarget(ctx context.Context,tt []TemplateTarget,id uuid.UUID) (t []TemplateTarget, err error) {
	args := m.Called(ctx,tt, id)
	return args.Get(0).([]TemplateTarget), args.Error(1)
}

func (m *DBMockStore) ListTemplateTargets(ctx context.Context, id uuid.UUID) (t []TemplateTarget, err error) {
	args := m.Called(ctx, id)
	return args.Get(0).([]TemplateTarget), args.Error(1)
}