package db

import (
	"context"

	"github.com/stretchr/testify/mock"
)

type DBMockStore struct {
	mock.Mock
}

func (m *DBMockStore) ListTarget(ctx context.Context) (t []Target, err error) {
	args := m.Called(ctx)
	return args.Get(0).([]Target), args.Error(1)
}

func (m *DBMockStore) CreateTarget(ctx context.Context, tar Target) (t Target, err error) {
	args := m.Called(ctx, tar)
	return args.Get(0).(Target), args.Error(1)
}
