package application

import (
	"app/internal/core/domain"
	"app/internal/core/ports/outbound"
	"context"

	"github.com/google/uuid"
)

type CreateOrderUseCase struct {
	repo outbound.OrderRepository
}

func NewCreateOrderUseCase(r outbound.OrderRepository) *CreateOrderUseCase {
	return &CreateOrderUseCase{
		repo: r,
	}
}

func (c *CreateOrderUseCase) CreateOrder(ctx context.Context) error {

	// TO DO: Receber CreateOrderRequest ou CreateOrderInput
	order := &domain.Order{
		OrderID: uuid.New(),
		Status:  "CREATED",
		Price:   1500,
	}

	if err := c.repo.CreateOrder(ctx, order); err != nil {
		return err
	}

	return nil
}
