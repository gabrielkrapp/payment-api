package infra

import (
	"encoding/json"
	"net/http"
	"payment-api/internal/entity"

	"github.com/stripe/stripe-go"
)

func MakePaymentHandler(service entity.PaymentIntentService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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

		params := &stripe.PaymentIntentParams{
			Amount:        stripe.Int64(req.Amount),
			Currency:      stripe.String(req.Currency),
			PaymentMethod: stripe.String(req.PaymentMethodID),
			Confirm:       stripe.Bool(true),
		}

		pi, err := service.New(params)
		if err != nil {
			http.Error(w, "Failed to create payment intent", http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(pi)
	}
}
