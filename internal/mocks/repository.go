package mocks

import (
	"context"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"github.com/ypxd99/yandex-diplom-56/internal/model"
	"github.com/ypxd99/yandex-diplom-56/internal/repository"
)

type MockGophermartRepo struct {
	mock.Mock
}

func (m *MockGophermartRepo) Close() error {
	args := m.Called()
	return args.Error(0)
}

func (m *MockGophermartRepo) Status(ctx context.Context) (bool, error) {
	args := m.Called(ctx)
	return args.Bool(0), args.Error(1)
}

func (m *MockGophermartRepo) CreateUser(ctx context.Context, user *model.User) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}

func (m *MockGophermartRepo) GetUserByLogin(ctx context.Context, login string) (*model.User, error) {
	args := m.Called(ctx, login)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.User), args.Error(1)
}

func (m *MockGophermartRepo) GetUserByID(ctx context.Context, id uuid.UUID) (*model.User, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.User), args.Error(1)
}

func (m *MockGophermartRepo) CreateOrder(ctx context.Context, order *model.Order) error {
	args := m.Called(ctx, order)
	return args.Error(0)
}

func (m *MockGophermartRepo) GetOrderByNumber(ctx context.Context, number string) (*model.Order, error) {
	args := m.Called(ctx, number)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.Order), args.Error(1)
}

func (m *MockGophermartRepo) GetOrderByID(ctx context.Context, id uuid.UUID) (*model.Order, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.Order), args.Error(1)
}

func (m *MockGophermartRepo) GetUserOrders(ctx context.Context, userID uuid.UUID) ([]*model.Order, error) {
	args := m.Called(ctx, userID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*model.Order), args.Error(1)
}

func (m *MockGophermartRepo) UpdateOrderStatus(ctx context.Context, orderID uuid.UUID, status model.OrderStatus, accrual float64) error {
	args := m.Called(ctx, orderID, status, accrual)
	return args.Error(0)
}

func (m *MockGophermartRepo) GetUserBalance(ctx context.Context, userID uuid.UUID) (*model.UserBalance, error) {
	args := m.Called(ctx, userID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*model.UserBalance), args.Error(1)
}

func (m *MockGophermartRepo) UpdateUserBalance(ctx context.Context, userID uuid.UUID, current, withdrawn float64) error {
	args := m.Called(ctx, userID, current, withdrawn)
	return args.Error(0)
}

func (m *MockGophermartRepo) CreateWithdrawal(ctx context.Context, withdrawal *model.Withdrawal) error {
	args := m.Called(ctx, withdrawal)
	return args.Error(0)
}

func (m *MockGophermartRepo) GetUserWithdrawals(ctx context.Context, userID uuid.UUID) ([]*model.Withdrawal, error) {
	args := m.Called(ctx, userID)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*model.Withdrawal), args.Error(1)
}

func (m *MockGophermartRepo) CleanUpTables(ctx context.Context) error {
	args := m.Called(ctx)
	return args.Error(0)
}

var _ repository.GophermartRepo = (*MockGophermartRepo)(nil)
