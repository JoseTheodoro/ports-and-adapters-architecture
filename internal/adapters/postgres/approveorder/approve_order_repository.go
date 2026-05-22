package approveorder

import (
	"app/db/queries"
	"app/internal/core/domain"
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
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
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, domain.ErrOrderNotFound
		}
		return nil, err
	}

	return r.toDomain(row), nil
}

func (r *Repository) Approve(ctx context.Context, order *domain.Order) error {

	arg := queries.ApproveOrderParams{
		Status:  queries.OrderStatus(order.Status),
		OrderID: order.OrderID,
	}
	return r.q.ApproveOrder(ctx, arg)
}

func (r *Repository) toDomain(o queries.Order) *domain.Order {
	return &domain.Order{
		ID:        o.ID,
		OrderID:   o.OrderID,
		Price:     int64(o.Price),
		Status:    domain.OrderStatus(o.Status),
		CreatedAt: o.CreatedAt,
		UpdatedAt: o.UpdatedAt,
	}
}
