package entity

import (
	"testing"

	"github.com/stretchr/testify/mock"
	"github.com/stripe/stripe-go"
)

type MockPaymentIntentService struct {
	mock.Mock
}

func (m *MockPaymentIntentService) New(params *stripe.PaymentIntentParams) (*stripe.PaymentIntent, error) {
	args := m.Called(params)
	return args.Get(0).(*stripe.PaymentIntent), args.Error(1)
}

func TestCreatePaymentIntent(t *testing.T) {
	mockService := new(MockPaymentIntentService)
	expectedPaymentIntent := &stripe.PaymentIntent{
		ID:       "pi_123456789",
		Amount:   1000,
		Currency: "usd",
		Status:   stripe.PaymentIntentStatusSucceeded,
	}

	mockService.On("New", mock.Anything).Return(expectedPaymentIntent, nil)

	paymentIntent, err := CreatePaymentIntent(mockService, 1000, "usd", "pm_123456789")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if paymentIntent != nil {
		if paymentIntent.ID != expectedPaymentIntent.ID {
			t.Errorf("Expected PaymentIntent ID to be %s, got %s", expectedPaymentIntent.ID, paymentIntent.ID)
		}
	} else {
		t.Errorf("Expected a valid paymentIntent, got nil")
	}

	mockService.AssertExpectations(t)
}
