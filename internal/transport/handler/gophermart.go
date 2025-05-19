package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ypxd99/yandex-diplom-56/internal/middleware"
	"github.com/ypxd99/yandex-diplom-56/internal/model"
	"github.com/ypxd99/yandex-diplom-56/internal/service"
	"github.com/ypxd99/yandex-diplom-56/util"
)

func (h *Handler) Register(c *gin.Context) {
	var req struct {
		Login    string `json:"login" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	user, err := h.service.Register(c.Request.Context(), req.Login, req.Password)
	if err != nil {
		if err == service.ErrUserExists {
			c.JSON(http.StatusConflict, gin.H{"error": "User already exists"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	token, err := middleware.GenerateToken(user.ID, []byte(util.GetConfig().Auth.SecretKey))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func (h *Handler) Login(c *gin.Context) {
	var req struct {
		Login    string `json:"login" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	user, err := h.service.Login(c.Request.Context(), req.Login, req.Password)
	if err != nil {
		if err == service.ErrInvalidCredentials {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	token, err := middleware.GenerateToken(user.ID, []byte(util.GetConfig().Auth.SecretKey))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func (h *Handler) CreateOrder(c *gin.Context) {
	userID, err := middleware.GetUserID(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	contentType := c.GetHeader("Content-Type")
	if contentType != "text/plain" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Content-Type must be text/plain"})
		return
	}

	number, err := c.GetRawData()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	orderNumber := string(number)
	if !isValidLuhn(orderNumber) {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Invalid order number"})
		return
	}

	_, err = h.service.CreateOrder(c.Request.Context(), userID, orderNumber)
	if err != nil {
		if err == service.ErrOrderAlreadyUploadedByThisUser {
			c.Status(http.StatusOK)
			return
		}
		if err == service.ErrOrderBelongsToAnotherUser {
			c.JSON(http.StatusConflict, gin.H{"error": "Order already exists for another user"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.Status(http.StatusAccepted)
}

func isValidLuhn(number string) bool {
	sum := 0
	alternate := false

	for i := len(number) - 1; i >= 0; i-- {
		digit := int(number[i] - '0')
		if digit < 0 || digit > 9 {
			return false
		}

		if alternate {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}

		sum += digit
		alternate = !alternate
	}

	return sum%10 == 0
}

func (h *Handler) GetOrders(c *gin.Context) {
	userID, err := middleware.GetUserID(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	orders, err := h.service.GetUserOrders(c.Request.Context(), userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	if len(orders) == 0 {
		c.Header("Content-Type", "application/json")
		c.Status(http.StatusNoContent)
		return
	}

	type OrderResponse struct {
		Number     string  `json:"number"`
		Status     string  `json:"status"`
		Accrual    float64 `json:"accrual,omitempty"`
		UploadedAt string  `json:"uploaded_at"`
	}

	response := make([]OrderResponse, len(orders))
	for i, order := range orders {
		response[i] = OrderResponse{
			Number:     order.Number,
			Status:     string(order.Status),
			Accrual:    order.Accrual,
			UploadedAt: order.UploadedAt.Format(time.RFC3339),
		}
	}

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, response)
}

func (h *Handler) GetBalance(c *gin.Context) {
	userID, err := middleware.GetUserID(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	balance, err := h.service.GetUserBalance(c.Request.Context(), userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	if balance.Current < 0 {
		balance.Current = 0
	}
	if balance.Withdrawn < 0 {
		balance.Withdrawn = 0
	}

	c.JSON(http.StatusOK, balance)
}

func (h *Handler) WithdrawBalance(c *gin.Context) {
	userID, err := middleware.GetUserID(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var req struct {
		Order string  `json:"order" binding:"required"`
		Sum   float64 `json:"sum" binding:"required"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	if !isValidLuhn(req.Order) {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Invalid order number"})
		return
	}

	if req.Sum <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Sum must be positive"})
		return
	}

	err = h.service.WithdrawBalance(c.Request.Context(), userID, req.Order, req.Sum)
	if err != nil {
		if err == service.ErrInsufficientFunds {
			c.JSON(http.StatusPaymentRequired, gin.H{"error": "Insufficient funds"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.Status(http.StatusOK)
}

func (h *Handler) GetWithdrawals(c *gin.Context) {
	userID, err := middleware.GetUserID(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	withdrawals, err := h.service.GetUserWithdrawals(c.Request.Context(), userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	if len(withdrawals) == 0 {
		c.Status(http.StatusNoContent)
		return
	}

	type WithdrawalResponse struct {
		Order       string  `json:"order"`
		Sum         float64 `json:"sum"`
		ProcessedAt string  `json:"processed_at"`
	}

	response := make([]WithdrawalResponse, len(withdrawals))
	for i, w := range withdrawals {
		response[i] = WithdrawalResponse{
			Order:       w.OrderNumber,
			Sum:         w.Sum,
			ProcessedAt: w.ProcessedAt.Format(time.RFC3339),
		}
	}

	c.JSON(http.StatusOK, response)
}

func (h *Handler) GetOrderAccrual(c *gin.Context) {
	userID, err := middleware.GetUserID(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	number := c.Param("number")
	if !isValidLuhn(number) {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": "Invalid order number"})
		return
	}

	order, err := h.service.GetOrderByNumber(c.Request.Context(), number)
	if err != nil {
		if err == service.ErrOrderNotFound {
			c.Status(http.StatusNoContent)
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	if order.UserID != userID {
		c.Status(http.StatusNoContent)
		return
	}

	response := gin.H{
		"order":  order.Number,
		"status": string(order.Status),
	}

	if order.Status == model.OrderStatusProcessed {
		response["accrual"] = order.Accrual
	}

	c.JSON(http.StatusOK, response)
}
