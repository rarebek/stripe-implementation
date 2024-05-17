package handlers

import (
	"encoding/json"
	"net/http"
	"restaurant-payment/models"
	"restaurant-payment/services"
)

type PaymentHandler struct {
	StripeService *services.StripeService
}

func NewPaymentHandler(stripeService *services.StripeService) *PaymentHandler {
	return &PaymentHandler{StripeService: stripeService}
}

func (h *PaymentHandler) HandlePayment(w http.ResponseWriter, r *http.Request) {
	var order models.Order
	if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
		http.Error(w, "Invalid order data", http.StatusBadRequest)
		return
	}

	token := r.Header.Get("Stripe-Token")
	if token == "" {
		http.Error(w, "Missing Stripe token", http.StatusBadRequest)
		return
	}

	charge, err := h.StripeService.Charge(order, token)
	if err != nil {
		http.Error(w, "Payment failed", http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"status": "success",
		"charge": charge,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
