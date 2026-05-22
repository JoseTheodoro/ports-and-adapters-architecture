package approveorder

import (
	"context"

	"app/internal/core/domain"
)

type Approver interface {
	Approve(ctx context.Context, approveOrderInput *OrderInput) (*OrderOutput, error)
}

type Repository interface {
	Approve(ctx context.Context, order *domain.Order) (*domain.Order, error)
}
