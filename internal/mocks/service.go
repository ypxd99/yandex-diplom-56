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

func (m *MockGophermartService) CreateOrder(ctx context.Context, userID uuid.UUID, number string) (*model.Order, error) {
	ret := m.Called(ctx, userID, number)

	var r0 *model.Order
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, string) *model.Order); ok {
		r0 = rf(ctx, userID, number)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Order)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID, string) error); ok {
		r1 = rf(ctx, userID, number)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (m *MockGophermartService) GetOrderByNumber(ctx context.Context, number string) (*model.Order, error) {
	ret := m.Called(ctx, number)

	var r0 *model.Order
	if rf, ok := ret.Get(0).(func(context.Context, string) *model.Order); ok {
		r0 = rf(ctx, number)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Order)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, number)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (m *MockGophermartService) GetUserBalance(ctx context.Context, userID uuid.UUID) (*model.UserBalance, error) {
	ret := m.Called(ctx, userID)

	var r0 *model.UserBalance
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) *model.UserBalance); ok {
		r0 = rf(ctx, userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.UserBalance)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(ctx, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (m *MockGophermartService) GetUserByID(ctx context.Context, id uuid.UUID) (*model.User, error) {
	ret := m.Called(ctx, id)

	var r0 *model.User
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) *model.User); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (m *MockGophermartService) GetUserOrders(ctx context.Context, userID uuid.UUID) ([]*model.Order, error) {
	ret := m.Called(ctx, userID)

	var r0 []*model.Order
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) []*model.Order); ok {
		r0 = rf(ctx, userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Order)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(ctx, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (m *MockGophermartService) GetUserWithdrawals(ctx context.Context, userID uuid.UUID) ([]*model.Withdrawal, error) {
	ret := m.Called(ctx, userID)

	var r0 []*model.Withdrawal
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) []*model.Withdrawal); ok {
		r0 = rf(ctx, userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Withdrawal)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(ctx, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (m *MockGophermartService) InitializeUserBalance(ctx context.Context, userID uuid.UUID) error {
	ret := m.Called(ctx, userID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) error); ok {
		r0 = rf(ctx, userID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

func (m *MockGophermartService) Login(ctx context.Context, login string, password string) (*model.User, error) {
	ret := m.Called(ctx, login, password)

	var r0 *model.User
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *model.User); ok {
		r0 = rf(ctx, login, password)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, login, password)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (m *MockGophermartService) ProcessOrder(ctx context.Context, orderID uuid.UUID, status model.OrderStatus, accrual float64) error {
	ret := m.Called(ctx, orderID, status, accrual)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, model.OrderStatus, float64) error); ok {
		r0 = rf(ctx, orderID, status, accrual)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

func (m *MockGophermartService) Register(ctx context.Context, login string, password string) (*model.User, error) {
	ret := m.Called(ctx, login, password)

	var r0 *model.User
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *model.User); ok {
		r0 = rf(ctx, login, password)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, login, password)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (m *MockGophermartService) WithdrawBalance(ctx context.Context, userID uuid.UUID, orderNumber string, sum float64) error {
	ret := m.Called(ctx, userID, orderNumber, sum)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, string, float64) error); ok {
		r0 = rf(ctx, userID, orderNumber, sum)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

var _ service.GophermartService = (*MockGophermartService)(nil)
