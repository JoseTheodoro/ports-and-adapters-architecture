package http

import (
	"app/internal/core/application/approveorder"
	"errors"
	"net/http"

	"github.com/google/uuid"
)

type ApproveOrderHTTPHandler struct {
	service approveorder.Approver
}

func NewApproveOrderHTTPHandler(a approveorder.Approver) *ApproveOrderHTTPHandler {
	return &ApproveOrderHTTPHandler{service: a}
}

func (h *ApproveOrderHTTPHandler) ApproveOrder(w http.ResponseWriter, r *http.Request) {

	request := NewApproveOrderRequest(r)

	input, err := request.ToOrderInput()
	if err != nil {
		ToJSON(w, http.StatusBadRequest, map[string]string{"message": err.Error()})
		return
	}

	if err := h.service.Approve(r.Context(), input); err != nil {
		if errors.Is(err, approveorder.ErrOrderNotFound) {
			ToJSON(w, http.StatusNotFound, map[string]string{"message": err.Error()})
			return
		}
		if errors.Is(err, approveorder.ErrOrderAlredyApproved) {
			ToJSON(w, http.StatusOK, map[string]string{"message": err.Error()})
			return
		}
		ToJSON(w, http.StatusInternalServerError, map[string]string{"message": err.Error()})
		return
	}

	ToJSON(w, http.StatusOK, map[string]string{"message": "Order has been approved"})

}

type ApproveOrderRequest struct {
	OrderID string
}

func NewApproveOrderRequest(r *http.Request) *ApproveOrderRequest {
	id := r.PathValue("id")
	return &ApproveOrderRequest{OrderID: id}
}

func (a *ApproveOrderRequest) ToOrderInput() (*approveorder.OrderInput, error) {
	id, err := uuid.Parse(a.OrderID)
	if err != nil {
		return nil, ErrOrderIDInvalid
	}
	return &approveorder.OrderInput{
		OrderID: id,
	}, nil
}
