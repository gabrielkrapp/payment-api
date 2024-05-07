package config

import (
	"log"
	"os"

	"github.com/stripe/stripe-go/v72"
)

func InitStripe() {
	stripe.Key = os.Getenv("STRIPE_SECRET_KEY")
	if stripe.Key == "" {
		log.Fatal("STRIPE_SECRET_KEY is not set")
	}
}
