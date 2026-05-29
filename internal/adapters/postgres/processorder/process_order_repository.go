package processorder

import (
	"app/db/queries"
	"app/internal/core/domain"
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	db *pgxpool.Pool
	q  *queries.Queries
}

func New(db *pgxpool.Pool) *Repository {
	return &Repository{
		db: db,
		q:  queries.New(db),
	}
}

func (r *Repository) FindByOrderID(ctx context.Context, orderID uuid.UUID) (*domain.Order, error) {

	row, err := r.q.FindByOrderID(ctx, orderID)
	if err != nil {
		return nil, err
	}

	order := r.toDomain(row)

	return order, nil

}

func (r *Repository) Process(ctx context.Context, order *domain.Order) error {
	params := queries.ProcessOrderParams{
		OrderID: order.OrderID,
		Status:  queries.OrderStatus(domain.PROCESSING),
	}
	return r.q.ProcessOrder(ctx, params)
}

func (r *Repository) toDomain(row queries.Order) *domain.Order {
	return &domain.Order{
		ID:        row.ID,
		OrderID:   row.OrderID,
		Price:     int64(row.Price),
		Status:    domain.OrderStatus(row.Status),
		CreatedAt: row.CreatedAt,
		UpdatedAt: row.UpdatedAt,
	}
}
