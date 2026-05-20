package inbound

import "context"

type CreatorOrder interface {
	CreateOrder(ctx context.Context) error
}
