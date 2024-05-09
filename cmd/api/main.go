package main

import (
	"log"
	"net/http"
	"payment-api/config"
	"payment-api/internal/api/entity"
	"payment-api/internal/api/infra"
)

func main() {
	config.InitStripe()

	server := config.NewHTTPServer(":8080")

	producer := config.NewKafkaProducer()
	defer producer.Close()

	paymentService := &entity.StripePaymentIntentService{}

	http.HandleFunc("/v1/payment", infra.MakePaymentHandler(paymentService, producer, "payment-intents"))

	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
