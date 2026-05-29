package rabbitmq

import (
	"app/internal/core/application/processorder"
	"encoding/json"
	"log/slog"
	"time"

	"github.com/google/uuid"
)

type OrderApprovedMessage struct {
	OrderID   uuid.UUID `json:"order_id"`
	Price     int64     `json:"price"`
	ApproveAt time.Time `json:"approved_at"`
}

func NewOrderApprovedMessage(body []byte) *OrderApprovedMessage {
	var msg OrderApprovedMessage
	if err := json.Unmarshal(body, &msg); err != nil {
		slog.Error("unmarshal order approved message", "error", err)
	}
	return &msg
}

func (m *OrderApprovedMessage) ToInput() *processorder.OrderProcessInput {
	input := &processorder.OrderProcessInput{
		OrderID:    m.OrderID,
		Price:      m.Price,
		ApprovedAt: m.ApproveAt,
	}
	return input
}
