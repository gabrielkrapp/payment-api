package infra

import (
	"encoding/json"
	"net/http"
	pi "payment-api/internal/entity"
)

func PaymentHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Unsupported method", http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		Amount          int64  `json:"amount"`
		Currency        string `json:"currency"`
		PaymentMethodID string `json:"paymentMethodId"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Error decoding request body", http.StatusBadRequest)
		return
	}

	pi, err := pi.CreatePaymentIntent(req.Amount, req.Currency, req.PaymentMethodID)
	if err != nil {
		http.Error(w, "Failed to create payment intent", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(pi)
}
