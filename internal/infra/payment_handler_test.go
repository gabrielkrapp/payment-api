package infra

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
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

func TestPaymentHandler(t *testing.T) {
	mockService := new(MockPaymentIntentService)
	handler := MakePaymentHandler(mockService)

	requestBody := map[string]interface{}{
		"amount":          1000,
		"currency":        "usd",
		"paymentMethodId": "pm_123456789",
	}
	bodyBytes, _ := json.Marshal(requestBody)
	request := httptest.NewRequest(http.MethodPost, "/payment", bytes.NewBuffer(bodyBytes))
	request.Header.Set("Content-Type", "application/json")
	responseRecorder := httptest.NewRecorder()

	expectedPi := &stripe.PaymentIntent{
		ID:       "pi_123456789",
		Amount:   1000,
		Currency: "usd",
		Status:   stripe.PaymentIntentStatusSucceeded,
	}
	mockService.On("New", mock.Anything).Return(expectedPi, nil)

	handler(responseRecorder, request)

	assert.Equal(t, http.StatusOK, responseRecorder.Code)
	responsePi := &stripe.PaymentIntent{}
	err := json.NewDecoder(responseRecorder.Body).Decode(responsePi)
	assert.NoError(t, err)
	assert.Equal(t, expectedPi.ID, responsePi.ID)

	mockService.AssertExpectations(t)
}
