package http

import (
	"time"

	"github.com/google/uuid"

	"app/internal/core/ports"
)

type CreateOrderResponse struct {
	OrderID   uuid.UUID  `json:"order_id"`
	Price     int64      `json:"price"`
	Status    string     `json:"status"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

func NewCreateOrderResponse(out *ports.CreateOrderOutput) *CreateOrderResponse {
	return &CreateOrderResponse{
		OrderID:   out.OrderID,
		Price:     out.Price,
		Status:    string(out.Status),
		CreatedAt: out.CreatedAt,
		UpdatedAt: out.UpdatedAt,
	}
}
