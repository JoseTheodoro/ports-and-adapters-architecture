package http

import (
	"encoding/json"
	"net/http"

	"app/internal/core/ports"
)

type HandleCreateOrder struct {
	creator ports.CreateOrderUseCase
}

func NewHandleCreateOrder(c ports.CreateOrderUseCase) *HandleCreateOrder {
	return &HandleCreateOrder{
		creator: c,
	}
}

func (h *HandleCreateOrder) CreateOrder(w http.ResponseWriter, r *http.Request) {

	request, err := h.validate(r)

	if err != nil {
		ToJSON(w, http.StatusBadRequest, map[string]string{"message": err.Error()})
		return
	}

	createOrderInput, err := h.toOrderInput(request)
	if err != nil {
		ToJSON(w, http.StatusInternalServerError, map[string]string{"message": err.Error()})
		return
	}

	orderOutput, err := h.creator.CreateOrder(r.Context(), createOrderInput)
	if err != nil {
		ToJSON(w, http.StatusInternalServerError, map[string]string{"message": err.Error()})
		return
	}
	ToJSON(w, http.StatusOK, NewCreateOrderResponse(orderOutput))
}

func (h *HandleCreateOrder) validate(r *http.Request) (*CreateOrderRequest, error) {
	var createOrderRequest = &CreateOrderRequest{}

	// TODO: Implementar validação de borda  http.request (schema por exemplo)

	if err := json.NewDecoder(r.Body).Decode(createOrderRequest); err != nil {
		return nil, err
	}

	return createOrderRequest, nil
}

func (h *HandleCreateOrder) toOrderInput(r *CreateOrderRequest) (*ports.CreateOrderInput, error) {

	input := &ports.CreateOrderInput{
		Price: r.Price,
	}

	return input, nil
}
