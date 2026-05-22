package approveorder

import (
	"app/internal/core/domain"
	"context"
	"errors"
)

type Service struct {
	repo Repository
}

func New(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

func (s *Service) Approve(ctx context.Context, input *OrderInput) error {

	// find and fill Order
	order, err := s.repo.FindByOrderID(ctx, input.OrderID)
	if err != nil {
		if errors.Is(err, domain.ErrOrderNotFound) {
			return ErrOrderNotFound
		}
		return err
	}
	// Approve, domain intrinsic validation
	if err := order.Approve(); err != nil {
		if errors.Is(err, domain.ErrOrderAlreadyApproved) {
			return ErrOrderAlredyApproved
		}
		return err
	}

	if err := s.repo.Approve(ctx, order); err != nil {
		return err
	}
	return nil
}

func (s *Service) toDomain(i *OrderInput) *domain.Order {

	return &domain.Order{
		OrderID: i.OrderID,
	}
}
