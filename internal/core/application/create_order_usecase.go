package application

import (
	"app/internal/core/domain"
	"app/internal/core/ports/inbound"
	"app/internal/core/ports/outbound"
	"context"
)

type CreateOrderUseCase struct {
	repo outbound.OrderRepository
}

func NewCreateOrderUseCase(r outbound.OrderRepository) *CreateOrderUseCase {
	return &CreateOrderUseCase{
		repo: r,
	}
}

func (c *CreateOrderUseCase) CreateOrder(ctx context.Context, createOrderInput *inbound.CreateOrderInput) (*inbound.CreateOrderOutput, error) {

	order := domain.NewOrder(createOrderInput.Price, domain.CREATED)

	saved, err := c.repo.CreateOrder(ctx, order)
	if err != nil {
		return nil, err
	}

	orderOutPut := c.toOrderOutput(saved)

	return orderOutPut, nil
}

func (c *CreateOrderUseCase) toOrderOutput(order *domain.Order) *inbound.CreateOrderOutput {

	return &inbound.CreateOrderOutput{
		OrderID:   order.OrderID,
		Price:     order.Price,
		Status:    order.Status,
		CreatedAt: order.CreatedAt,
		UpdatedAt: order.UpdatedAt,
	}

}
