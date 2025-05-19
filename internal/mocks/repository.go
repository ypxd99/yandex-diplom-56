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

func (_m *MockGophermartRepo) CleanUpTables(ctx context.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

func (_m *MockGophermartRepo) Close() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

func (_m *MockGophermartRepo) CreateOrder(ctx context.Context, order *model.Order) error {
	ret := _m.Called(ctx, order)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.Order) error); ok {
		r0 = rf(ctx, order)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

func (_m *MockGophermartRepo) CreateUser(ctx context.Context, user *model.User) error {
	ret := _m.Called(ctx, user)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.User) error); ok {
		r0 = rf(ctx, user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

func (_m *MockGophermartRepo) CreateUserBalance(ctx context.Context, balance *model.UserBalance) error {
	ret := _m.Called(ctx, balance)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.UserBalance) error); ok {
		r0 = rf(ctx, balance)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

func (_m *MockGophermartRepo) CreateWithdrawal(ctx context.Context, withdrawal *model.Withdrawal) error {
	ret := _m.Called(ctx, withdrawal)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.Withdrawal) error); ok {
		r0 = rf(ctx, withdrawal)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

func (_m *MockGophermartRepo) GetOrderByID(ctx context.Context, id uuid.UUID) (*model.Order, error) {
	ret := _m.Called(ctx, id)

	var r0 *model.Order
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) *model.Order); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Order)
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

func (_m *MockGophermartRepo) GetOrderByNumber(ctx context.Context, number string) (*model.Order, error) {
	ret := _m.Called(ctx, number)

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

func (_m *MockGophermartRepo) GetOrdersByStatus(ctx context.Context, status model.OrderStatus) ([]*model.Order, error) {
	ret := _m.Called(ctx, status)

	var r0 []*model.Order
	if rf, ok := ret.Get(0).(func(context.Context, model.OrderStatus) []*model.Order); ok {
		r0 = rf(ctx, status)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Order)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, model.OrderStatus) error); ok {
		r1 = rf(ctx, status)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *MockGophermartRepo) GetUserBalance(ctx context.Context, userID uuid.UUID) (*model.UserBalance, error) {
	ret := _m.Called(ctx, userID)

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

func (_m *MockGophermartRepo) GetUserByID(ctx context.Context, id uuid.UUID) (*model.User, error) {
	ret := _m.Called(ctx, id)

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

func (_m *MockGophermartRepo) GetUserByLogin(ctx context.Context, login string) (*model.User, error) {
	ret := _m.Called(ctx, login)

	var r0 *model.User
	if rf, ok := ret.Get(0).(func(context.Context, string) *model.User); ok {
		r0 = rf(ctx, login)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, login)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *MockGophermartRepo) GetUserOrders(ctx context.Context, userID uuid.UUID) ([]*model.Order, error) {
	ret := _m.Called(ctx, userID)

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

func (_m *MockGophermartRepo) GetUserWithdrawals(ctx context.Context, userID uuid.UUID) ([]*model.Withdrawal, error) {
	ret := _m.Called(ctx, userID)

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

func (_m *MockGophermartRepo) Status(ctx context.Context) (bool, error) {
	ret := _m.Called(ctx)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context) bool); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

func (_m *MockGophermartRepo) UpdateOrderStatus(ctx context.Context, orderID uuid.UUID, status model.OrderStatus, accrual float64) error {
	ret := _m.Called(ctx, orderID, status, accrual)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, model.OrderStatus, float64) error); ok {
		r0 = rf(ctx, orderID, status, accrual)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

func (_m *MockGophermartRepo) UpdateUserBalance(ctx context.Context, userID uuid.UUID, current float64, withdrawn float64) error {
	ret := _m.Called(ctx, userID, current, withdrawn)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID, float64, float64) error); ok {
		r0 = rf(ctx, userID, current, withdrawn)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

var _ repository.GophermartRepo = (*MockGophermartRepo)(nil)
