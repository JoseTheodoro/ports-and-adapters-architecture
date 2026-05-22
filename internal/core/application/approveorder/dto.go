package approveorder

import "github.com/google/uuid"

type OrderInput struct {
	OrderID uuid.UUID
}

type OrderOutput struct{}
