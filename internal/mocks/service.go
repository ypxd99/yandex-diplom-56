package mocks

import (
	"github.com/stretchr/testify/mock"
	"github.com/ypxd99/yandex-diplom-56/internal/service"
)

type MockGophermartService struct {
	mock.Mock
}

var _ service.GophermartService = (*MockGophermartService)(nil)
