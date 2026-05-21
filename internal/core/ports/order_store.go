package ports

import (
	"context"

	"app/internal/core/domain"
)

type OrderStore interface {
	CreateOrder(ctx context.Context, order *domain.Order) (*domain.Order, error)
}
