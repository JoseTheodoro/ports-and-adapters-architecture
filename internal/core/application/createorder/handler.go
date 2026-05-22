package createorder

import (
	"context"

	"app/internal/core/domain"
)

type createOrderHandler struct {
	repo Repository
}

func NewCreateOrderHandler(r Repository) Creator {
	return &createOrderHandler{
		repo: r,
	}
}

func (c *createOrderHandler) Create(ctx context.Context, createOrderInput *OrderInput) (*OrderOutput, error) {

	order := domain.NewOrder(createOrderInput.Price, domain.CREATED)

	saved, err := c.repo.Create(ctx, order)
	if err != nil {
		return nil, err
	}

	orderOutPut := c.toOrderOutput(saved)

	return orderOutPut, nil
}

func (c *createOrderHandler) toOrderOutput(order *domain.Order) *OrderOutput {

	return &OrderOutput{
		OrderID:   order.OrderID,
		Price:     order.Price,
		Status:    order.Status,
		CreatedAt: order.CreatedAt,
		UpdatedAt: order.UpdatedAt,
	}

}
