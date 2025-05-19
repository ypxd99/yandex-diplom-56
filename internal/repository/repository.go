package repository

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/ypxd99/yandex-diplom-56/internal/model"
)

var (
	ErrUserNotFound              = errors.New("user not found")
	ErrUserExists                = errors.New("user already exists")
	ErrOrderNotFound             = errors.New("order not found")
	ErrOrderBelongsToAnotherUser = errors.New("order belongs to another user")
)

type GophermartRepo interface {
	Close() error
	Status(ctx context.Context) (bool, error)
	CleanUpTables(ctx context.Context) error

	CreateUser(ctx context.Context, user *model.User) error
	GetUserByLogin(ctx context.Context, login string) (*model.User, error)
	GetUserByID(ctx context.Context, id uuid.UUID) (*model.User, error)

	CreateOrder(ctx context.Context, order *model.Order) error
	GetOrderByNumber(ctx context.Context, number string) (*model.Order, error)
	GetOrderByID(ctx context.Context, id uuid.UUID) (*model.Order, error)
	GetUserOrders(ctx context.Context, userID uuid.UUID) ([]*model.Order, error)
	UpdateOrderStatus(ctx context.Context, orderID uuid.UUID, status model.OrderStatus, accrual float64) error

	GetUserBalance(ctx context.Context, userID uuid.UUID) (*model.UserBalance, error)
	UpdateUserBalance(ctx context.Context, userID uuid.UUID, current, withdrawn float64) error

	CreateWithdrawal(ctx context.Context, withdrawal *model.Withdrawal) error
	GetUserWithdrawals(ctx context.Context, userID uuid.UUID) ([]*model.Withdrawal, error)
}
