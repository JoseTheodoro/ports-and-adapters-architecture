package domain

import (
	"time"

	"github.com/google/uuid"
)

type OrderStatus string

const (
	CREATED   OrderStatus = "CREATED"
	PENDING   OrderStatus = "PENDING"
	APPROVED  OrderStatus = "APPROVED"
	PROCSSING OrderStatus = "PROCESSING"
	SHIPPED   OrderStatus = "SHIPPED"
	DELIVERED OrderStatus = "DELIVERED"
	REJECTED  OrderStatus = "REJECTED"
)

type Order struct {
	ID        int64
	OrderID   uuid.UUID
	Price     int64
	Status    OrderStatus
	CreatedAt time.Time
	UpdatedAt *time.Time
}

func NewOrder(price int64, status OrderStatus) *Order {
	return &Order{
		OrderID: uuid.New(),
		Price:   price,
		Status:  status,
	}
}
