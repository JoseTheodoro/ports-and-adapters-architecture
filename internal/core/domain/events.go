package domain

import (
	"time"

	"github.com/google/uuid"
)

type OrderApprovedEvent struct {
	OrderID    uuid.UUID `json:"order_id"`
	Price      int64     `json:"price"`
	ApprovedAt time.Time `json:"approved_at"`
}
