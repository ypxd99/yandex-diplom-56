package postgres_test

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/ypxd99/yandex-diplom-56/internal/model"
	"github.com/ypxd99/yandex-diplom-56/internal/repository/postgres"
)

func setupTestDB(t *testing.T) *postgres.PostgresRepo {
	db, err := postgres.Connect(context.Background())
	assert.NoError(t, err)

	err = db.CleanupTables(context.Background())
	assert.NoError(t, err)

	return db
}

func TestPostgresRepo_UserOperations(t *testing.T) {
	db := setupTestDB(t)

	t.Run("create and get user", func(t *testing.T) {
		user := &model.User{
			ID:       uuid.New(),
			Login:    "testuser",
			Password: "hashedpassword",
		}

		err := db.CreateUser(context.Background(), user)
		assert.NoError(t, err)

		retrievedUser, err := db.GetUserByLogin(context.Background(), user.Login)
		assert.NoError(t, err)
		assert.NotNil(t, retrievedUser)
		assert.Equal(t, user.ID, retrievedUser.ID)
		assert.Equal(t, user.Login, retrievedUser.Login)
		assert.Equal(t, user.Password, retrievedUser.Password)
	})

	t.Run("get non-existent user", func(t *testing.T) {
		user, err := db.GetUserByLogin(context.Background(), "nonexistent")
		assert.NoError(t, err)
		assert.Nil(t, user)
	})
}

func TestPostgresRepo_OrderOperations(t *testing.T) {
	db := setupTestDB(t)

	t.Run("create and get order", func(t *testing.T) {
		user := &model.User{
			ID:       uuid.New(),
			Login:    "testuser",
			Password: "hashedpassword",
		}
		err := db.CreateUser(context.Background(), user)
		assert.NoError(t, err)

		order := &model.Order{
			ID:         uuid.New(),
			UserID:     user.ID,
			Number:     "12345678903",
			Status:     model.OrderStatusNew,
			Accrual:    0,
			UploadedAt: time.Now(),
		}

		err = db.CreateOrder(context.Background(), order)
		assert.NoError(t, err)

		retrievedOrder, err := db.GetOrderByNumber(context.Background(), order.Number)
		assert.NoError(t, err)
		assert.NotNil(t, retrievedOrder)
		assert.Equal(t, order.ID, retrievedOrder.ID)
		assert.Equal(t, order.UserID, retrievedOrder.UserID)
		assert.Equal(t, order.Number, retrievedOrder.Number)
		assert.Equal(t, order.Status, retrievedOrder.Status)
		assert.Equal(t, order.Accrual, retrievedOrder.Accrual)
	})

	t.Run("get user orders", func(t *testing.T) {
		user := &model.User{
			ID:       uuid.New(),
			Login:    "testuser2",
			Password: "hashedpassword",
		}
		err := db.CreateUser(context.Background(), user)
		assert.NoError(t, err)

		orders := []*model.Order{
			{
				ID:         uuid.New(),
				UserID:     user.ID,
				Number:     "12345678906",
				Status:     model.OrderStatusNew,
				Accrual:    0,
				UploadedAt: time.Now(),
			},
			{
				ID:         uuid.New(),
				UserID:     user.ID,
				Number:     "12345678907",
				Status:     model.OrderStatusProcessing,
				Accrual:    100,
				UploadedAt: time.Now(),
			},
		}

		for _, order := range orders {
			err = db.CreateOrder(context.Background(), order)
			assert.NoError(t, err)
		}

		retrievedOrders, err := db.GetUserOrders(context.Background(), user.ID)
		assert.NoError(t, err)
		assert.Len(t, retrievedOrders, 2)
	})

	t.Run("update order status", func(t *testing.T) {
		user := &model.User{
			ID:       uuid.New(),
			Login:    "testuser3",
			Password: "hashedpassword",
		}
		err := db.CreateUser(context.Background(), user)
		assert.NoError(t, err)

		order := &model.Order{
			ID:         uuid.New(),
			UserID:     user.ID,
			Number:     "12345678905",
			Status:     model.OrderStatusNew,
			Accrual:    0,
			UploadedAt: time.Now(),
		}

		err = db.CreateOrder(context.Background(), order)
		assert.NoError(t, err)

		err = db.UpdateOrderStatus(context.Background(), order.ID, model.OrderStatusProcessed, 500)
		assert.NoError(t, err)

		retrievedOrder, err := db.GetOrderByNumber(context.Background(), order.Number)
		assert.NoError(t, err)
		assert.NotNil(t, retrievedOrder)
		assert.Equal(t, model.OrderStatusProcessed, retrievedOrder.Status)
		assert.Equal(t, float64(500), retrievedOrder.Accrual)
	})
}
