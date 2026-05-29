package approveorder

import (
	"app/internal/core/domain"
	"context"
	"errors"
	"time"
)

type Service struct {
	repo Repository
	pub  Publisher
}

func New(r Repository, p Publisher) *Service {
	return &Service{
		repo: r,
		pub:  p,
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

	// publish event OrderApproved
	orderApprovedEvent := &domain.OrderApprovedEvent{
		OrderID:    order.OrderID,
		Price:      order.Price,
		ApprovedAt: time.Now(),
	}

	if err := s.pub.PublishOrderApproved(ctx, orderApprovedEvent); err != nil {
		return err
	}

	return nil
}

func (s *Service) toDomain(i *OrderInput) *domain.Order {

	return &domain.Order{
		OrderID: i.OrderID,
	}
}
