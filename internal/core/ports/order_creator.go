package ports

import (
	"context"
)

type OrderCreator interface {
	CreateOrder(ctx context.Context, createOrderInput *CreateOrderInput) (*CreateOrderOutput, error)
}
