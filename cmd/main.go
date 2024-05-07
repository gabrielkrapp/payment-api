package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"payment-api/internal/entity"
	"payment-api/internal/infra"

	"github.com/IBM/sarama"
	"github.com/stripe/stripe-go/v72"
)

func init() {
	stripe.Key = os.Getenv("STRIPE_SECRET_KEY")
	if stripe.Key == "" {
		log.Fatal("STRIPE_SECRET_KEY is not set")
	}
}

func main() {
	server := &http.Server{
		Addr:           ":8080",
		Handler:        nil,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	brokers := []string{"localhost:9092"}
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5
	config.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		log.Fatalf("Failed to create Kafka producer: %v", err)
	}
	defer producer.Close()

	paymentService := &entity.StripePaymentIntentService{}

	http.HandleFunc("/payment", infra.MakePaymentHandler(paymentService, producer, "payment-intents"))

	err = server.ListenAndServe()
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
