package order

import (
	dto2 "app/internal/adapters/inbound/http/order/dto"
	"app/internal/core/ports/inbound/order"
	"app/internal/core/ports/inbound/order/dto"
	"encoding/json"
	"fmt"
	"net/http"
)

type HandleCreateOrder struct {
	creator order.CreateOrderUseCase
}

func NewHandleCreateOrder(c order.CreateOrderUseCase) *HandleCreateOrder {
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

func (h *HandleCreateOrder) validate(r *http.Request) (*dto2.CreateOrderRequest, error) {
	var createOrderRequest = &dto2.CreateOrderRequest{}

	// TODO: Implementar validação de borda  http.request (schema por exemplo)

	if err := json.NewDecoder(r.Body).Decode(createOrderRequest); err != nil {
		return nil, err
	}

	return createOrderRequest, nil
}

func (h *HandleCreateOrder) toOrderInput(r *dto2.CreateOrderRequest) (*dto.CreateOrderInput, error) {

	input := &dto.CreateOrderInput{
		Price: r.Price,
	}

	return input, nil
}

func (h *HandleCreateOrder) createResponse(out *dto.CreateOrderOutput) *dto2.CreateOrderResponse {
	return &dto2.CreateOrderResponse{
		OrderID:   out.OrderID,
		Price:     out.Price,
		Status:    string(out.Status),
		CreatedAt: out.CreatedAt,
		UpdatedAt: out.UpdatedAt,
	}
}

func writeResponse(w http.ResponseWriter, response *dto2.CreateOrderResponse) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	return json.NewEncoder(w).Encode(response)
}
