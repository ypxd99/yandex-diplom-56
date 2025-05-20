package service

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/ypxd99/yandex-diplom-56/internal/model"
	"github.com/ypxd99/yandex-diplom-56/internal/repository"
)

func (s *Service) Register(ctx context.Context, login, password string) (*model.User, error) {
	existingUser, err := s.repo.GetUserByLogin(ctx, login)
	if err != nil && !errors.Is(err, repository.ErrUserNotFound) {
		return nil, err
	}
	if existingUser != nil {
		return nil, repository.ErrUserExists
	}

	hashedPassword := hashPassword(password)

	user := &model.User{
		ID:       uuid.New(),
		Login:    login,
		Password: hashedPassword,
	}

	err = s.repo.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *Service) Login(ctx context.Context, login, password string) (*model.User, error) {
	user, err := s.repo.GetUserByLogin(ctx, login)
	if err != nil {
		if errors.Is(err, repository.ErrUserNotFound) {
			return nil, ErrInvalidCredentials
		}
		return nil, err
	}

	if user.Password != hashPassword(password) {
		return nil, ErrInvalidCredentials
	}

	return user, nil
}

func (s *Service) GetUserByID(ctx context.Context, id uuid.UUID) (*model.User, error) {
	user, err := s.repo.GetUserByID(ctx, id)
	if err != nil {
		if errors.Is(err, repository.ErrUserNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}
	return user, nil
}

func (s *Service) CreateOrder(ctx context.Context, userID uuid.UUID, number string) (*model.Order, error) {
	existingOrder, err := s.repo.GetOrderByNumber(ctx, number)
	if err != nil && !errors.Is(err, repository.ErrOrderNotFound) {
		return nil, err
	}
	if existingOrder != nil {
		if existingOrder.UserID == userID {
			return nil, ErrOrderAlreadyUploadedByThisUser
		}
		return nil, ErrOrderBelongsToAnotherUser
	}

	order := &model.Order{
		ID:         uuid.New(),
		UserID:     userID,
		Number:     number,
		Status:     model.OrderStatusNew,
		UploadedAt: time.Now(),
	}

	err = s.repo.CreateOrder(ctx, order)
	if err != nil {
		return nil, err
	}

	return order, nil
}

func (s *Service) GetUserOrders(ctx context.Context, userID uuid.UUID) ([]*model.Order, error) {
	return s.repo.GetUserOrders(ctx, userID)
}

func (s *Service) ProcessOrder(ctx context.Context, orderID uuid.UUID, status model.OrderStatus, accrual float64) error {
	order, err := s.repo.GetOrderByID(ctx, orderID)
	if err != nil {
		return err
	}

	err = s.repo.UpdateOrderStatus(ctx, orderID, status, accrual)
	if err != nil {
		return err
	}

	if status == model.OrderStatusProcessed && accrual > 0 {
		balance, err := s.repo.GetUserBalance(ctx, order.UserID)
		if err != nil {
			return err
		}

		balance.Current += accrual
		err = s.repo.UpdateUserBalance(ctx, order.UserID, balance.Current, balance.Withdrawn)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *Service) GetUserBalance(ctx context.Context, userID uuid.UUID) (*model.UserBalance, error) {
	return s.repo.GetUserBalance(ctx, userID)
}

func (s *Service) WithdrawBalance(ctx context.Context, userID uuid.UUID, orderNumber string, sum float64) error {
	balance, err := s.repo.GetUserBalance(ctx, userID)
	if err != nil {
		return err
	}

	if balance.Current < sum {
		return ErrInsufficientFunds
	}

	withdrawal := &model.Withdrawal{
		UserID:      userID,
		OrderNumber: orderNumber,
		Sum:         sum,
		ProcessedAt: time.Now(),
	}

	err = s.repo.CreateWithdrawal(ctx, withdrawal)
	if err != nil {
		return err
	}

	balance.Current -= sum
	balance.Withdrawn += sum
	err = s.repo.UpdateUserBalance(ctx, userID, balance.Current, balance.Withdrawn)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) GetUserWithdrawals(ctx context.Context, userID uuid.UUID) ([]*model.Withdrawal, error) {
	return s.repo.GetUserWithdrawals(ctx, userID)
}

func (s *Service) GetOrderByNumber(ctx context.Context, number string) (*model.Order, error) {
	order, err := s.repo.GetOrderByNumber(ctx, number)
	if err != nil {
		if err == repository.ErrOrderNotFound {
			return nil, ErrOrderNotFound
		}
		return nil, err
	}
	if order == nil {
		return nil, ErrOrderNotFound
	}
	return order, nil
}

func hashPassword(password string) string {
	hash := sha256.Sum256([]byte(password))
	return hex.EncodeToString(hash[:])
}
