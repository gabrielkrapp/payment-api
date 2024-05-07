package entity

import (
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/paymentintent"
)

type PaymentIntentService interface {
	New(params *stripe.PaymentIntentParams) (*stripe.PaymentIntent, error)
}

type StripePaymentIntentService struct{}

func (s StripePaymentIntentService) New(params *stripe.PaymentIntentParams) (*stripe.PaymentIntent, error) {
	return paymentintent.New(params)
}

func CreatePaymentIntent(service PaymentIntentService, amount int64, currency string, paymentMethodID string) (*stripe.PaymentIntent, error) {
	params := &stripe.PaymentIntentParams{
		Amount:        stripe.Int64(amount),
		Currency:      stripe.String(currency),
		PaymentMethod: stripe.String(paymentMethodID),
		Confirm:       stripe.Bool(true),
	}

	return service.New(params)
}
