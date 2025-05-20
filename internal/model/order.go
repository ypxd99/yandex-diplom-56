package model

import (
	"time"

	"github.com/google/uuid"
)

type OrderStatus string

const (
	OrderStatusNew        OrderStatus = "NEW"
	OrderStatusRegistered OrderStatus = "REGISTERED"
	OrderStatusProcessing OrderStatus = "PROCESSING"
	OrderStatusInvalid    OrderStatus = "INVALID"
	OrderStatusProcessed  OrderStatus = "PROCESSED"
)

type Order struct {
	ID         uuid.UUID   `bun:"id,pk"`
	UserID     uuid.UUID   `bun:"user_id,notnull"`
	Number     string      `bun:"number,unique,notnull"`
	Status     OrderStatus `bun:"status,notnull"`
	Accrual    float64     `bun:"accrual,notnull,default:0"`
	UploadedAt time.Time   `bun:"uploaded_at,notnull"`
}

type Withdrawal struct {
	ID          uuid.UUID `bun:"id,pk"`
	UserID      uuid.UUID `bun:"user_id,notnull"`
	OrderNumber string    `bun:"order_number,notnull"`
	Sum         float64   `bun:"sum,notnull"`
	ProcessedAt time.Time `bun:"processed_at,notnull"`
}
