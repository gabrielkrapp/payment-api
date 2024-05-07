package infra

import (
	"encoding/json"
	"log"
	"net/http"
	"payment-api/internal/entity"

	"github.com/IBM/sarama"
	"github.com/stripe/stripe-go"

	"payment-api/pkg/kafka"
)

func MakePaymentHandler(service entity.PaymentIntentService, producer sarama.SyncProducer, topic string) http.HandlerFunc {
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

		message, err := json.Marshal(pi)
		if err != nil {
			log.Printf("Failed to marshal payment intent: %v", err)
		} else {
			kafka.KafkaProducer(producer, topic, message)
		}

		err = json.NewEncoder(w).Encode(pi)
		if err != nil {
			log.Printf("Failed to encode data: %v", err)
			http.Error(w, "Failed to encode response", http.StatusInternalServerError)
			return
		}
	}
}
