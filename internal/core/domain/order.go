package domain

import (
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type OrderStatus string

var ErrOrderStatusInvalid = errors.New("order has a status invalid to change")
var ErrOrderPriceInvalid = errors.New("order must have a price grather than zero")
var ErrOrderNotFound = errors.New("order not found")
var ErrOrderAlreadyApproved = errors.New("order already approved")

const (
	CREATED    OrderStatus = "CREATED"
	PENDING    OrderStatus = "PENDING"
	APPROVED   OrderStatus = "APPROVED"
	PROCESSING OrderStatus = "PROCESSING"
	SHIPPED    OrderStatus = "SHIPPED"
	DELIVERED  OrderStatus = "DELIVERED"
	REJECTED   OrderStatus = "REJECTED"
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

func (o *Order) Approve() error {
	if o.Status == APPROVED {
		return ErrOrderAlreadyApproved
	}

	if o.Status != CREATED {
		return ErrOrderStatusInvalid
	}

	if o.Price <= 0 {
		return ErrOrderPriceInvalid
	}

	o.Status = APPROVED

	return nil
}

func (o *Order) Process() error {

	if o.Status != APPROVED {
		return fmt.Errorf("order status must have approved to processing")
	}

	if o.CreatedAt.Before(time.Now().AddDate(0, 0, -30)) == true {
		return fmt.Errorf("order has been expired")
	}

	o.Status = PROCESSING

	return nil
}
