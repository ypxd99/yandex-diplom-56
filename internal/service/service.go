package service

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/ypxd99/yandex-diplom-56/internal/model"
	"github.com/ypxd99/yandex-diplom-56/internal/repository"
)

var (
	ErrUserNotFound                   = errors.New("user not found")
	ErrUserExists                     = errors.New("user already exists")
	ErrInvalidCredentials             = errors.New("invalid credentials")
	ErrOrderExists                    = errors.New("order already exists")
	ErrOrderAlreadyUploadedByThisUser = errors.New("order already uploaded by this user")
	ErrOrderNotFound                  = errors.New("order not found")
	ErrOrderBelongsToAnotherUser      = errors.New("order belongs to another user")
	ErrInsufficientFunds              = errors.New("insufficient funds")
)

type Service struct {
	repo repository.GophermartRepo
}

type GophermartService interface {
	Register(ctx context.Context, login, password string) (*model.User, error)
	Login(ctx context.Context, login, password string) (*model.User, error)
	GetUserByID(ctx context.Context, id uuid.UUID) (*model.User, error)

	CreateOrder(ctx context.Context, userID uuid.UUID, number string) (*model.Order, error)
	GetOrderByNumber(ctx context.Context, number string) (*model.Order, error)
	GetUserOrders(ctx context.Context, userID uuid.UUID) ([]*model.Order, error)
	ProcessOrder(ctx context.Context, orderID uuid.UUID, status model.OrderStatus, accrual float64) error

	GetUserBalance(ctx context.Context, userID uuid.UUID) (*model.UserBalance, error)
	WithdrawBalance(ctx context.Context, userID uuid.UUID, orderNumber string, sum float64) error
	GetUserWithdrawals(ctx context.Context, userID uuid.UUID) ([]*model.Withdrawal, error)
}

func InitService(repo repository.GophermartRepo) *Service {
	return &Service{repo: repo}
}
