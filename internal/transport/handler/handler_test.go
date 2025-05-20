package handler_test

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/ypxd99/yandex-diplom-56/internal/mocks"
	"github.com/ypxd99/yandex-diplom-56/internal/model"
	"github.com/ypxd99/yandex-diplom-56/internal/repository"
	"github.com/ypxd99/yandex-diplom-56/internal/service"
	"github.com/ypxd99/yandex-diplom-56/internal/transport/handler"
	"github.com/ypxd99/yandex-diplom-56/util"
)

func setupTestRouter() (*mocks.MockGophermartRepo, *gin.Engine) {
	gin.SetMode(gin.TestMode)
	util.InitLogger(util.GetConfig().Logger)
	mockRepo := new(mocks.MockGophermartRepo)
	svc := service.InitService(mockRepo)
	h := handler.InitHandler(svc)
	router := gin.New()
	h.InitRoutes(router)
	return mockRepo, router
}

func hashPassword(password string) string {
	hash := sha256.Sum256([]byte(password))
	return hex.EncodeToString(hash[:])
}

func TestHandler_Register(t *testing.T) {
	mockRepo, router := setupTestRouter()
	defer func() {
		mockRepo.ExpectedCalls = nil
		mockRepo.Calls = nil
	}()

	t.Run("successful registration", func(t *testing.T) {
		login := "testuser"
		password := "password123"

		mockRepo.On("GetUserByLogin", mock.Anything, login).Return(nil, repository.ErrUserNotFound).Once()
		mockRepo.On("CreateUser", mock.Anything, mock.MatchedBy(func(user *model.User) bool {
			return user.Login == login && user.Password == hashPassword(password)
		})).Return(nil).Once()

		reqBody := map[string]string{
			"login":    login,
			"password": password,
		}
		jsonBody, _ := json.Marshal(reqBody)

		req := httptest.NewRequest("POST", "/api/user/register", bytes.NewBuffer(jsonBody))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		mockRepo.AssertExpectations(t)
	})

	t.Run("invalid request", func(t *testing.T) {
		reqBody := map[string]string{
			"login": "testuser",
		}
		jsonBody, _ := json.Marshal(reqBody)

		req := httptest.NewRequest("POST", "/api/user/register", bytes.NewBuffer(jsonBody))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	})
}

func TestHandler_Login(t *testing.T) {
	mockRepo, router := setupTestRouter()
	defer func() {
		mockRepo.ExpectedCalls = nil
		mockRepo.Calls = nil
	}()

	t.Run("successful login", func(t *testing.T) {
		login := "testuser"
		password := "password123"
		userID := uuid.New()

		hash := hashPassword(password)
		mockRepo.On("GetUserByLogin", mock.Anything, login).Return(&model.User{
			ID:       userID,
			Login:    login,
			Password: hash,
		}, nil).Once()

		reqBody := map[string]string{
			"login":    login,
			"password": password,
		}
		jsonBody, _ := json.Marshal(reqBody)

		req := httptest.NewRequest("POST", "/api/user/login", bytes.NewBuffer(jsonBody))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		mockRepo.AssertExpectations(t)
	})

	t.Run("invalid credentials", func(t *testing.T) {
		login := "testuser"
		password := "wrongpassword"

		hash := hashPassword("password123")
		mockRepo.On("GetUserByLogin", mock.Anything, login).Return(&model.User{
			ID:       uuid.New(),
			Login:    login,
			Password: hash,
		}, nil).Once()

		reqBody := map[string]string{
			"login":    login,
			"password": password,
		}
		jsonBody, _ := json.Marshal(reqBody)

		req := httptest.NewRequest("POST", "/api/user/login", bytes.NewBuffer(jsonBody))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()

		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
		mockRepo.AssertExpectations(t)
	})
}
