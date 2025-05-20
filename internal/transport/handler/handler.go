package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/ypxd99/yandex-diplom-56/internal/middleware"
	"github.com/ypxd99/yandex-diplom-56/internal/service"
	"github.com/ypxd99/yandex-diplom-56/util"
)

type Handler struct {
	service service.GophermartService
}

func InitHandler(service service.GophermartService) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRoutes(r *gin.Engine) {
	util.GetMetricsRoute(r)
	util.GetHealthcheckRoute(r)
	util.GetRouteList(r)

	r.Use(middleware.LoggingMiddleware())
	r.Use(middleware.GzipMiddleware())
	r.Use(middleware.AuthMiddleware())

	r.POST("/api/user/register", h.Register)
	r.POST("/api/user/login", h.Login)

	api := r.Group("/api/user")
	api.Use(middleware.RequireAuth())
	{
		api.POST("/orders", h.CreateOrder)
		api.GET("/orders", h.GetOrders)
		api.GET("/balance", h.GetBalance)
		api.POST("/balance/withdraw", h.WithdrawBalance)
		api.GET("/withdrawals", h.GetWithdrawals)
	}

	r.GET("/api/orders/:number", h.GetOrderAccrual)
}
