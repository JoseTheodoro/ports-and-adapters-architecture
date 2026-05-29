package processorder

import (
	"time"

	"github.com/google/uuid"
)

type OrderProcessInput struct {
	OrderID    uuid.UUID
	Price      int64
	ApprovedAt time.Time
}
