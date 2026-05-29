package processorder

import (
	"context"

	"github.com/google/uuid"

	"app/internal/core/domain"
)

type Processor interface {
	Process(context.Context, *OrderProcessInput) error
}

type Repository interface {
	Process(context.Context, *domain.Order) error
	FindByOrderID(context.Context, uuid.UUID) (*domain.Order, error)
}
