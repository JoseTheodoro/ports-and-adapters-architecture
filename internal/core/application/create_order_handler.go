package application

import (
	"context"

	"app/internal/core/domain"
	"app/internal/core/ports"
)

type createOrderHandler struct {
	repo ports.OrderStore
}

func NewCreateOrderHandler(r ports.OrderStore) ports.OrderCreator {
	return &createOrderHandler{
		repo: r,
	}
}

func (c *createOrderHandler) CreateOrder(ctx context.Context, createOrderInput *ports.CreateOrderInput) (*ports.CreateOrderOutput, error) {

	order := domain.NewOrder(createOrderInput.Price, domain.CREATED)

	saved, err := c.repo.CreateOrder(ctx, order)
	if err != nil {
		return nil, err
	}

	orderOutPut := c.toOrderOutput(saved)

	return orderOutPut, nil
}

func (c *createOrderHandler) toOrderOutput(order *domain.Order) *ports.CreateOrderOutput {

	return &ports.CreateOrderOutput{
		OrderID:   order.OrderID,
		Price:     order.Price,
		Status:    order.Status,
		CreatedAt: order.CreatedAt,
		UpdatedAt: order.UpdatedAt,
	}

}
