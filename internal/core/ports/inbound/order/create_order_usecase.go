package order

import (
	"app/internal/core/ports/inbound/order/dto"
	"context"
)

type CreateOrderUseCase interface {
	CreateOrder(ctx context.Context, createOrderInput *dto.CreateOrderInput) (*dto.CreateOrderOutput, error)
}
