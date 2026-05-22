package createorder

import (
	"time"

	"github.com/google/uuid"

	"app/internal/core/domain"
)

type OrderInput struct {
	Price int64
}

type OrderOutput struct {
	OrderID   uuid.UUID
	Price     int64
	Status    domain.OrderStatus
	CreatedAt time.Time
	UpdatedAt *time.Time
}
