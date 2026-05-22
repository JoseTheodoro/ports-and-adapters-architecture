package createorder

import (
	"context"

	"app/internal/core/domain"
)

type Repository interface {
	Create(ctx context.Context, order *domain.Order) (*domain.Order, error)
}

type Creator interface {
	Create(ctx context.Context, createOrderInput *OrderInput) (*OrderOutput, error)
}
