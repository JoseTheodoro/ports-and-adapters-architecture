package dto

import (
	"time"

	"github.com/google/uuid"
)

type CreateOrderResponse struct {
	OrderID   uuid.UUID  `json:"order_id"`
	Price     int64      `json:"price"`
	Status    string     `json:"status"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}
