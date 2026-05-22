package http

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"

	"app/internal/core/application/createorder"
)

type CreateOrderHandler struct {
	creator createorder.Creator
}

func NewHandleCreateOrder(c createorder.Creator) *CreateOrderHandler {
	return &CreateOrderHandler{
		creator: c,
	}
}

func (h *CreateOrderHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {

	request := NewCreateOrderRequest(r)

	err := request.isValid()
	if err != nil {
		ToJSON(w, http.StatusBadRequest, map[string]string{"message": err.Error()})
		return
	}

	// TODO: toOrderInput will be request.ToOrderInput
	createOrderInput, err := h.toOrderInput(request)
	if err != nil {
		ToJSON(w, http.StatusInternalServerError, map[string]string{"message": err.Error()})
		return
	}

	orderOutput, err := h.creator.Create(r.Context(), createOrderInput)
	if err != nil {
		ToJSON(w, http.StatusInternalServerError, map[string]string{"message": err.Error()})
		return
	}
	ToJSON(w, http.StatusOK, NewCreateOrderResponse(orderOutput))
}

func (h *CreateOrderHandler) toOrderInput(r *CreateOrderRequest) (*createorder.OrderInput, error) {

	input := &createorder.OrderInput{
		Price: *r.Price,
	}

	return input, nil
}

type CreateOrderRequest struct {
	Price *int64 `json:"price"`
}

func NewCreateOrderRequest(r *http.Request) *CreateOrderRequest {
	var createOrderRequest = &CreateOrderRequest{}
	json.NewDecoder(r.Body).Decode(createOrderRequest)

	return createOrderRequest
}

func (r *CreateOrderRequest) isValid() error {

	if r.Price == nil {
		return ErrPriceIsRequired
	}

	return nil
}

type CreateOrderResponse struct {
	OrderID   uuid.UUID  `json:"order_id"`
	Price     int64      `json:"price"`
	Status    string     `json:"status"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

func NewCreateOrderResponse(out *createorder.OrderOutput) *CreateOrderResponse {
	return &CreateOrderResponse{
		OrderID:   out.OrderID,
		Price:     out.Price,
		Status:    string(out.Status),
		CreatedAt: out.CreatedAt,
		UpdatedAt: out.UpdatedAt,
	}
}
