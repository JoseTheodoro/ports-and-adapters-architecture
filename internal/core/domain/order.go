package domain

import (
	"time"

	"github.com/google/uuid"
)

type Order struct {
	ID        int64
	OrderID   uuid.UUID
	Price     int64
	Status    string
	CreatedAt time.Time
	UpdatedAt *time.Time
}
