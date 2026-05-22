package createorder

import (
	"context"

	"app/internal/core/domain"
)

type Creator interface {
	Create(ctx context.Context, createOrderInput *OrderInput) (*OrderOutput, error)
}

type Repository interface {
	Create(ctx context.Context, order *domain.Order) (*domain.Order, error)
}
