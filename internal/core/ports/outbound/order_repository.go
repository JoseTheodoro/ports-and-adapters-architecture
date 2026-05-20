package outbound

import (
	"app/internal/core/domain"
	"context"
)

type OrderRepository interface {
	CreateOrder(ctx context.Context, order *domain.Order) error
}
