package main

import (
	"log"
	"net/http"
	"os"
	"time"

	route "payment-api/internal/infra"

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

	http.HandleFunc("/payment", route.PaymentHandler)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
