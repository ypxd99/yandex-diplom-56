package mocks

import (
	"context"

	"github.com/stretchr/testify/mock"
	"github.com/ypxd99/yandex-diplom-56/internal/repository"
)

type MockGophermartRepo struct {
	mock.Mock
}

func (m *MockGophermartRepo) Close() error {
	return nil
}

func (m *MockGophermartRepo) Status(ctx context.Context) (bool, error) {
	return true, nil
}

var _ repository.GophermartRepo = (*MockGophermartRepo)(nil)
