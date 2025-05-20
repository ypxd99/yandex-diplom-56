package service

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/ypxd99/yandex-diplom-56/internal/model"
	"github.com/ypxd99/yandex-diplom-56/util"
)

type OrderAccrual struct {
	Order   string            `json:"order"`
	Status  model.OrderStatus `json:"status"`
	Accrual float64           `json:"accrual,omitempty"`
}

func (s *Service) StartAccrualWorker(ctx context.Context) {
	logger := util.GetLogger()
	logger.Info("Starting accrual worker")

	go func() {
		ticker := time.NewTicker(5 * time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-ctx.Done():
				logger.Info("Accrual worker stopped")
				return
			case <-ticker.C:
				s.processNewOrders(ctx)
			}
		}
	}()
}

func (s *Service) processNewOrders(ctx context.Context) {
	logger := util.GetLogger()
	orders, err := s.repo.GetOrdersByStatus(ctx, model.OrderStatusNew)
	if err != nil {
		logger.Errorf("Failed to get new orders: %v", err)
		return
	}

	var wg sync.WaitGroup
	for _, order := range orders {
		wg.Add(1)

		go func(order *model.Order) {
			defer wg.Done()

			err := s.repo.UpdateOrderStatus(ctx, order.ID, model.OrderStatusProcessing, 0)
			if err != nil {
				logger.Errorf("Failed to update order status to PROCESSING: %v", err)
				return
			}

			accrual, err := s.checkAccrualSystem(ctx, order.Number)
			if err != nil {
				logger.Errorf("Failed to check accrual system: %v", err)
				return
			}

			if accrual == nil {
				return
			}

			err = s.ProcessOrder(ctx, order.ID, accrual.Status, accrual.Accrual)
			if err != nil {
				logger.Errorf("Failed to process order: %v", err)
			}
		}(order)
	}

	processingOrders, err := s.repo.GetOrdersByStatus(ctx, model.OrderStatusProcessing)
	if err != nil {
		logger.Errorf("Failed to get processing orders: %v", err)
		return
	}

	for _, order := range processingOrders {
		wg.Add(1)

		go func(order *model.Order) {
			defer wg.Done()

			accrual, err := s.checkAccrualSystem(ctx, order.Number)
			if err != nil {
				logger.Errorf("Failed to check accrual system: %v", err)
				return
			}

			if accrual == nil {
				return
			}

			if accrual.Status != model.OrderStatusProcessing {
				err = s.ProcessOrder(ctx, order.ID, accrual.Status, accrual.Accrual)
				if err != nil {
					logger.Errorf("Failed to process order: %v", err)
				}
			}
		}(order)
	}

	wg.Wait()
}

func (s *Service) checkAccrualSystem(ctx context.Context, orderNumber string) (*OrderAccrual, error) {
	cfg := util.GetConfig()
	url := fmt.Sprintf("%s/api/orders/%s", cfg.Accrual.AccrualSystemAddress, orderNumber)

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusTooManyRequests {
		retryAfter := resp.Header.Get("Retry-After")
		if retryAfter != "" {
			seconds, err := time.ParseDuration(retryAfter + "s")
			if err == nil {
				time.Sleep(seconds)
			}
		}
		return nil, fmt.Errorf("rate limited by accrual system")
	}

	if resp.StatusCode == http.StatusNoContent {
		return nil, nil
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("accrual system returned status %d", resp.StatusCode)
	}

	var accrual OrderAccrual
	err = json.NewDecoder(resp.Body).Decode(&accrual)
	if err != nil {
		return nil, err
	}

	return &accrual, nil
}

func (s *Service) InitializeUserBalance(ctx context.Context, userID uuid.UUID) error {
	balance := &model.UserBalance{
		UserID:    userID,
		Current:   0,
		Withdrawn: 0,
	}
	return s.repo.CreateUserBalance(ctx, balance)
}
