package infra

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/IBM/sarama"
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

type MockKafkaProducer struct {
	mock.Mock
}

func (m *MockKafkaProducer) AbortTxn() error {
	panic("unimplemented")
}

func (m *MockKafkaProducer) AddMessageToTxn(msg *sarama.ConsumerMessage, groupId string, metadata *string) error {
	panic("unimplemented")
}

func (m *MockKafkaProducer) AddOffsetsToTxn(offsets map[string][]*sarama.PartitionOffsetMetadata, groupId string) error {
	panic("unimplemented")
}

func (m *MockKafkaProducer) BeginTxn() error {
	panic("unimplemented")
}

func (m *MockKafkaProducer) Close() error {
	panic("unimplemented")
}

func (m *MockKafkaProducer) CommitTxn() error {
	panic("unimplemented")
}

func (m *MockKafkaProducer) IsTransactional() bool {
	panic("unimplemented")
}

func (m *MockKafkaProducer) SendMessages(msgs []*sarama.ProducerMessage) error {
	panic("unimplemented")
}

func (m *MockKafkaProducer) TxnStatus() sarama.ProducerTxnStatusFlag {
	panic("unimplemented")
}

func (m *MockKafkaProducer) SendMessage(msg *sarama.ProducerMessage) (partition int32, offset int64, err error) {
	args := m.Called(msg)
	return args.Get(0).(int32), args.Get(1).(int64), args.Error(2)
}

func TestPaymentHandler(t *testing.T) {
	mockService := new(MockPaymentIntentService)
	mockKafkaProducer := new(MockKafkaProducer)
	topic := "payment-intents"
	handler := MakePaymentHandler(mockService, mockKafkaProducer, topic)

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

	mockKafkaProducer.On("SendMessage", mock.Anything).Return(int32(0), int64(0), nil)

	handler(responseRecorder, request)

	assert.Equal(t, http.StatusOK, responseRecorder.Code)
	responsePi := &stripe.PaymentIntent{}
	err := json.NewDecoder(responseRecorder.Body).Decode(responsePi)
	assert.NoError(t, err)
	assert.Equal(t, expectedPi.ID, responsePi.ID)

	mockService.AssertExpectations(t)
	mockKafkaProducer.AssertExpectations(t)
}
