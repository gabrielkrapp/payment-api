package main

import (
	"log"
	"net/http"
	"payment-api/config"
	"payment-api/internal/entity"
	"payment-api/internal/infra"
)

func main() {
	config.InitStripe()

	server := config.NewHTTPServer(":8080")

	brokers := []string{"localhost:9092"}
	producer := config.NewKafkaProducer(brokers)
	defer producer.Close()

	paymentService := &entity.StripePaymentIntentService{}

	http.HandleFunc("/payment", infra.MakePaymentHandler(paymentService, producer, "payment-intents"))

	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
