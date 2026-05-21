package ports

import (
	"context"
)

type CreateOrderUseCase interface {
	CreateOrder(ctx context.Context, createOrderInput *CreateOrderInput) (*CreateOrderOutput, error)
}
