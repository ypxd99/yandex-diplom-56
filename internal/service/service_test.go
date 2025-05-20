package service_test

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/ypxd99/yandex-diplom-56/internal/mocks"
	"github.com/ypxd99/yandex-diplom-56/internal/model"
	"github.com/ypxd99/yandex-diplom-56/internal/repository"
	"github.com/ypxd99/yandex-diplom-56/internal/service"
)

func hashPassword(password string) string {
	hash := sha256.Sum256([]byte(password))
	return hex.EncodeToString(hash[:])
}

func TestService_UserOperations(t *testing.T) {
	mockRepo := new(mocks.MockGophermartRepo)
	svc := service.InitService(mockRepo)

	t.Run("register user", func(t *testing.T) {
		login := "testuser"
		password := "password123"

		mockRepo.On("GetUserByLogin", mock.Anything, login).Return(nil, repository.ErrUserNotFound).Once()
		mockRepo.On("CreateUser", mock.Anything, mock.MatchedBy(func(user *model.User) bool {
			return user.Login == login
		})).Return(nil).Once()

		user, err := svc.Register(context.Background(), login, password)
		assert.NoError(t, err)
		assert.NotNil(t, user)
		assert.Equal(t, login, user.Login)
	})

	t.Run("register existing user", func(t *testing.T) {
		login := "existinguser"
		password := "password123"

		mockRepo.On("GetUserByLogin", mock.Anything, login).Return(&model.User{
			ID:    uuid.New(),
			Login: login,
		}, nil).Once()

		user, err := svc.Register(context.Background(), login, password)
		assert.Error(t, err)
		assert.Equal(t, service.ErrUserExists, err)
		assert.Nil(t, user)
	})

	t.Run("login user", func(t *testing.T) {
		login := "testuser"
		password := "password123"
		userID := uuid.New()

		hash := hashPassword(password)

		mockRepo.On("GetUserByLogin", mock.Anything, login).Return(&model.User{
			ID:       userID,
			Login:    login,
			Password: hash,
		}, nil).Once()

		user, err := svc.Login(context.Background(), login, password)
		assert.NoError(t, err)
		assert.NotNil(t, user)
		assert.Equal(t, userID, user.ID)
	})

	t.Run("login with invalid credentials", func(t *testing.T) {
		login := "testuser"
		password := "wrongpassword"

		hash := hashPassword("password123")

		mockRepo.On("GetUserByLogin", mock.Anything, login).Return(&model.User{
			ID:       uuid.New(),
			Login:    login,
			Password: hash,
		}, nil).Once()

		user, err := svc.Login(context.Background(), login, password)
		assert.Error(t, err)
		assert.Equal(t, service.ErrInvalidCredentials, err)
		assert.Nil(t, user)
	})
}

func TestService_OrderOperations(t *testing.T) {
	mockRepo := new(mocks.MockGophermartRepo)
	svc := service.InitService(mockRepo)

	t.Run("create order", func(t *testing.T) {
		userID := uuid.New()
		orderNumber := "12345678903"

		mockRepo.On("GetOrderByNumber", mock.Anything, orderNumber).Return(nil, repository.ErrOrderNotFound).Once()
		mockRepo.On("CreateOrder", mock.Anything, mock.MatchedBy(func(order *model.Order) bool {
			return order.UserID == userID && order.Number == orderNumber && order.Status == model.OrderStatusNew
		})).Return(nil).Once()

		order, err := svc.CreateOrder(context.Background(), userID, orderNumber)
		assert.NoError(t, err)
		assert.NotNil(t, order)
		assert.Equal(t, userID, order.UserID)
		assert.Equal(t, orderNumber, order.Number)
		assert.Equal(t, model.OrderStatusNew, order.Status)
	})

	t.Run("create existing order", func(t *testing.T) {
		userID := uuid.New()
		orderNumber := "12345678903"

		existingOrder := &model.Order{
			ID:         uuid.New(),
			UserID:     userID,
			Number:     orderNumber,
			Status:     model.OrderStatusNew,
			UploadedAt: time.Now(),
		}

		mockRepo.On("GetOrderByNumber", mock.Anything, orderNumber).Return(existingOrder, nil).Once()

		order, err := svc.CreateOrder(context.Background(), userID, orderNumber)
		assert.Error(t, err)
		assert.Equal(t, service.ErrOrderAlreadyUploadedByThisUser, err)
		assert.Nil(t, order)
	})

	t.Run("get user orders", func(t *testing.T) {
		userID := uuid.New()
		orders := []*model.Order{
			{
				ID:         uuid.New(),
				UserID:     userID,
				Number:     "12345678903",
				Status:     model.OrderStatusNew,
				UploadedAt: time.Now(),
			},
		}

		mockRepo.On("GetUserOrders", mock.Anything, userID).Return(orders, nil).Once()

		result, err := svc.GetUserOrders(context.Background(), userID)
		assert.NoError(t, err)
		assert.Equal(t, orders, result)
	})
}
