package model

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id" db:"id"`
	Login     string    `json:"login" db:"login"`
	Password  string    `json:"-" db:"password"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

type UserBalance struct {
	UserID    uuid.UUID `bun:"user_id,pk"`
	Current   float64   `bun:"current,notnull"`
	Withdrawn float64   `bun:"withdrawn,notnull"`
}
