package processorder

import (
	"context"
	"fmt"
	"log/slog"
)

type Service struct {
	repo Repository
}

func New(r Repository) *Service {
	return &Service{
		repo: r,
	}
}

func (s *Service) Process(ctx context.Context, input *OrderProcessInput) error {
	slog.Info("processing order start", "input", input)

	order, err := s.repo.FindByOrderID(ctx, input.OrderID)
	if err != nil {
		return err
	}

	if err := order.Process(); err != nil {
		return fmt.Errorf("application process error: %w", err)
	}

	if err := s.repo.Process(ctx, order); err != nil {
		return fmt.Errorf("application process error: %w", err)
	}

	return nil
}
