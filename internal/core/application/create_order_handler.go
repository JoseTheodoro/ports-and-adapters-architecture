package application

import (
	"app/internal/core/domain"
	"app/internal/core/ports"
	"context"
)

type createOrderInteractor struct {
	repo ports.OrderRepository
}

func NewCreateOrderInteractor(r ports.OrderRepository) ports.CreateOrderUseCase {
	return &createOrderInteractor{
		repo: r,
	}
}

func (c *createOrderInteractor) CreateOrder(ctx context.Context, createOrderInput *ports.CreateOrderInput) (*ports.CreateOrderOutput, error) {

	order := domain.NewOrder(createOrderInput.Price, domain.CREATED)

	saved, err := c.repo.CreateOrder(ctx, order)
	if err != nil {
		return nil, err
	}

	orderOutPut := c.toOrderOutput(saved)

	return orderOutPut, nil
}

func (c *createOrderInteractor) toOrderOutput(order *domain.Order) *ports.CreateOrderOutput {

	return &ports.CreateOrderOutput{
		OrderID:   order.OrderID,
		Price:     order.Price,
		Status:    order.Status,
		CreatedAt: order.CreatedAt,
		UpdatedAt: order.UpdatedAt,
	}

}
