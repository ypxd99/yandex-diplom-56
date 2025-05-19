package repository

import "context"

type GophermartRepo interface {
	Close() error
	Status(ctx context.Context) (bool, error)
}
