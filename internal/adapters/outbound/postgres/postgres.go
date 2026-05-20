package postgres

import (
	"app/internal/adapters/outbound/postgres/queries"
	"app/internal/core/domain"
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type RepositoryOrderPostgres struct {
	conn *pgxpool.Pool
	q    *queries.Queries
}

func NewRepositoryOrderPostgress(c *pgxpool.Pool) *RepositoryOrderPostgres {
	return &RepositoryOrderPostgres{
		conn: c,
		q:    queries.New(c),
	}
}

func (r *RepositoryOrderPostgres) CreateOrder(ctx context.Context, order *domain.Order) error {

	arg := queries.CreateOrderParams{
		OrderID: order.OrderID,
		Status:  order.Status,
		Price:   int32(order.Price),
	}

	// TO DO: Fazer retornar um *domain.Order
	_, err := r.q.CreateOrder(ctx, arg)
	if err != nil {
		return err
	}

	return nil
}
