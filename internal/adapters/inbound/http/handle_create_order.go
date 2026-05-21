package http

import (
	"app/internal/adapters/inbound/http/dto"
	"app/internal/core/ports/inbound"
	"encoding/json"
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

	request, err := h.validate(r)
	if err != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message":"failed to validate order request"}`))
		return
	}

	createOrderInput, err := h.toOrderInput(request)
	if err != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message":"failed to order input"}`))
		return
	}

	orderOutput, err := h.creator.CreateOrder(r.Context(), createOrderInput)
	if err != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message":"failed to create order"}`))
		return
	}

	response := h.createResponse(orderOutput)

	if err := writeResponse(w, response); err != nil {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"message":"failed to create response order"}`))
		return
	}

}

func (h *HandleCreateOrder) validate(r *http.Request) (*dto.CreateOrderRequest, error) {
	var createOrderRequest = &dto.CreateOrderRequest{}

	// TODO: Implementar validação de borda  http.request (schema por exemplo)

	if err := json.NewDecoder(r.Body).Decode(createOrderRequest); err != nil {
		return nil, err
	}

	return createOrderRequest, nil
}

func (h *HandleCreateOrder) toOrderInput(r *dto.CreateOrderRequest) (*inbound.CreateOrderInput, error) {

	input := &inbound.CreateOrderInput{
		Price: r.Price,
	}

	return input, nil
}

func (h *HandleCreateOrder) createResponse(out *inbound.CreateOrderOutput) *dto.CreateOrderResponse {
	return &dto.CreateOrderResponse{
		OrderID:   out.OrderID,
		Price:     out.Price,
		Status:    string(out.Status),
		CreatedAt: out.CreatedAt,
		UpdatedAt: out.UpdatedAt,
	}
}

func writeResponse(w http.ResponseWriter, response *dto.CreateOrderResponse) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	return json.NewEncoder(w).Encode(response)
}
