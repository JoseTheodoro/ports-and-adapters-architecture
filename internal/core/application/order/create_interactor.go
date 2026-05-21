package order

import (
	"app/internal/core/domain"
	orderport "app/internal/core/ports/inbound/order"
	"app/internal/core/ports/inbound/order/dto"
	"app/internal/core/ports/outbound"
	"context"
)

type createOrderInteractor struct {
	repo outbound.OrderRepository
}

func NewCreateOrderInteractor(r outbound.OrderRepository) orderport.CreateOrderUseCase {
	return &createOrderInteractor{
		repo: r,
	}
}

func (c *createOrderInteractor) CreateOrder(ctx context.Context, createOrderInput *dto.CreateOrderInput) (*dto.CreateOrderOutput, error) {

	order := domain.NewOrder(createOrderInput.Price, domain.CREATED)

	saved, err := c.repo.CreateOrder(ctx, order)
	if err != nil {
		return nil, err
	}

	orderOutPut := c.toOrderOutput(saved)

	return orderOutPut, nil
}

func (c *createOrderInteractor) toOrderOutput(order *domain.Order) *dto.CreateOrderOutput {

	return &dto.CreateOrderOutput{
		OrderID:   order.OrderID,
		Price:     order.Price,
		Status:    order.Status,
		CreatedAt: order.CreatedAt,
		UpdatedAt: order.UpdatedAt,
	}

}
