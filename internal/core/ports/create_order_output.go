package ports

import (
	"time"

	"github.com/google/uuid"

	"app/internal/core/domain"
)

type CreateOrderOutput struct {
	OrderID   uuid.UUID
	Price     int64
	Status    domain.OrderStatus
	CreatedAt time.Time
	UpdatedAt *time.Time
}
