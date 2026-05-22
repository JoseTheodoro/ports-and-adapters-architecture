package createorder

import (
	"context"

	"app/internal/core/domain"
)

type Workflow struct {
	repo Repository
}

func New(r Repository) Runner {
	return &Workflow{
		repo: r,
	}
}

func (c *Workflow) Run(ctx context.Context, createOrderInput *OrderInput) (*OrderOutput, error) {

	order := domain.NewOrder(createOrderInput.Price, domain.CREATED)

	saved, err := c.repo.Create(ctx, order)
	if err != nil {
		return nil, err
	}

	orderOutPut := c.toOrderOutput(saved)

	return orderOutPut, nil
}

func (c *Workflow) toOrderOutput(order *domain.Order) *OrderOutput {

	return &OrderOutput{
		OrderID:   order.OrderID,
		Price:     order.Price,
		Status:    order.Status,
		CreatedAt: order.CreatedAt,
		UpdatedAt: order.UpdatedAt,
	}

}
