package http

import (
	"app/internal/core/ports/inbound"
	"fmt"
	"net/http"
)

type HandleCreateOrder struct {
	creator inbound.CreatorOrder
}

func NewHandleCreateOrder(c inbound.CreatorOrder) *HandleCreateOrder {
	return &HandleCreateOrder{
		creator: c,
	}
}

func (h *HandleCreateOrder) CreateOrder(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Handling at adapters inbound http: Requesting Application Layer")

	// TO DO: Implementar um CreateOrderRequest ou CreateOrderInput
	if err := h.creator.CreateOrder(r.Context()); err != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message":"failed to create order"}`))
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message":"order has been successfully created"}`))

}
