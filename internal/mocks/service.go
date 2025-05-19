package mocks

import (
	"context"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"github.com/ypxd99/yandex-diplom-56/internal/model"
	"github.com/ypxd99/yandex-diplom-56/internal/service"
)

type MockGophermartService struct {
	mock.Mock
}

func (m *MockGophermartService) Register(ctx context.Context, login, password string) (*model.User, error) {
	args := m.Called(ctx, login, password)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.User), args.Error(1)
}

func (m *MockGophermartService) Login(ctx context.Context, login, password string) (*model.User, error) {
	args := m.Called(ctx, login, password)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.User), args.Error(1)
}

func (m *MockGophermartService) GetUserByID(ctx context.Context, id uuid.UUID) (*model.User, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.User), args.Error(1)
}

func (m *MockGophermartService) CreateOrder(ctx context.Context, userID uuid.UUID, number string) (*model.Order, error) {
	args := m.Called(ctx, userID, number)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.Order), args.Error(1)
}

func (m *MockGophermartService) GetUserOrders(ctx context.Context, userID uuid.UUID) ([]*model.Order, error) {
	args := m.Called(ctx, userID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*model.Order), args.Error(1)
}

func (m *MockGophermartService) ProcessOrder(ctx context.Context, orderID uuid.UUID, status model.OrderStatus, accrual float64) error {
	args := m.Called(ctx, orderID, status, accrual)
	return args.Error(0)
}

func (m *MockGophermartService) GetUserBalance(ctx context.Context, userID uuid.UUID) (*model.UserBalance, error) {
	args := m.Called(ctx, userID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.UserBalance), args.Error(1)
}

func (m *MockGophermartService) WithdrawBalance(ctx context.Context, userID uuid.UUID, orderNumber string, sum float64) error {
	args := m.Called(ctx, userID, orderNumber, sum)
	return args.Error(0)
}

func (m *MockGophermartService) GetUserWithdrawals(ctx context.Context, userID uuid.UUID) ([]*model.Withdrawal, error) {
	args := m.Called(ctx, userID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*model.Withdrawal), args.Error(1)
}

func (m *MockGophermartService) GetOrderByNumber(ctx context.Context, number string) (*model.Order, error) {
	args := m.Called(ctx, number)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.Order), args.Error(1)
}

var _ service.GophermartService = (*MockGophermartService)(nil)
