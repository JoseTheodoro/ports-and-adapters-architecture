package approveorder

import (
	"context"

	"app/internal/core/domain"

	"github.com/google/uuid"
)

type Approver interface {
	Approve(ctx context.Context, approveOrderInput *OrderInput) error
}

type Repository interface {
	Approve(ctx context.Context, order *domain.Order) error
	FindByOrderID(ctx context.Context, orderID uuid.UUID) (*domain.Order, error)
}
