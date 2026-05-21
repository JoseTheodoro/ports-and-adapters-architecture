package inbound

import (
	"app/internal/core/domain"
	"time"

	"github.com/google/uuid"
)

type CreateOrderOutput struct {
	OrderID   uuid.UUID
	Price     int64
	Status    domain.OrderStatus
	CreatedAt time.Time
	UpdatedAt *time.Time
}
