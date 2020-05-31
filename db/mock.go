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
