package main

import (
	"os"

	"github.com/stripe/stripe-go/v72"
)

func init() {
	stripe.Key = os.Getenv("STRIPE_SECRET_KEY")
}

func main() {

}
