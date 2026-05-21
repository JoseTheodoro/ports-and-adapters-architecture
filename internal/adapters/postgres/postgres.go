package postgres

import (
	queries2 "app/db/queries"
	"app/internal/core/domain"
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type RepositoryOrderPostgres struct {
	conn *pgxpool.Pool
	q    *queries2.Queries
}

func NewOrderRepositoryPostgress(c *pgxpool.Pool) *RepositoryOrderPostgres {
	return &RepositoryOrderPostgres{
		conn: c,
		q:    queries2.New(c),
	}
}

func (r *RepositoryOrderPostgres) CreateOrder(ctx context.Context, order *domain.Order) (*domain.Order, error) {

	arg := queries2.CreateOrderParams{
		OrderID: order.OrderID,
		Status:  string(order.Status),
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

func toDomain(o queries2.Order) *domain.Order {

	return &domain.Order{
		ID:        o.ID,
		OrderID:   o.OrderID,
		Price:     int64(o.Price),
		Status:    domain.OrderStatus(o.Status),
		CreatedAt: o.CreatedAt,
		UpdatedAt: o.UpdatedAt,
	}
}
