package createorder

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"

	"app/db/queries"
	"app/internal/core/domain"
)

type Repository struct {
	conn *pgxpool.Pool
	q    *queries.Queries
}

func New(c *pgxpool.Pool) *Repository {
	return &Repository{
		conn: c,
		q:    queries.New(c),
	}
}

func (r *Repository) Create(ctx context.Context, order *domain.Order) (*domain.Order, error) {

	arg := queries.CreateOrderParams{
		OrderID: order.OrderID,
		Status:  queries.OrderStatus(order.Status),
		Price:   int32(order.Price),
	}

	row, err := r.q.CreateOrder(ctx, arg)
	if err != nil {
		return nil, err
	}

	// TODO: row deve ser traduzido para domain.order ou inbound.CreateOrderOutput?
	o := toDomain(row)

	return o, nil
}

func toDomain(o queries.Order) *domain.Order {

	return &domain.Order{
		ID:        o.ID,
		OrderID:   o.OrderID,
		Price:     int64(o.Price),
		Status:    domain.OrderStatus(o.Status),
		CreatedAt: o.CreatedAt,
		UpdatedAt: o.UpdatedAt,
	}
}
