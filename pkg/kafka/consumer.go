package kafka

import (
	"context"
	"encoding/json"
	"log"

	"payment-api/internal/worker"

	"github.com/IBM/sarama"
)

type PaymentMessage struct {
	Email     string `json:"email"`
	PaymentID string `json:"payment_id"`
	Amount    string `json:"amount"`
}

type PaymentConsumerHandler struct{}

func (PaymentConsumerHandler) Setup(sarama.ConsumerGroupSession) error {
	return nil
}

func (PaymentConsumerHandler) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

func (h PaymentConsumerHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for message := range claim.Messages() {
		var payment PaymentMessage
		if err := json.Unmarshal(message.Value, &payment); err != nil {
			log.Printf("Error decoding message: %v", err)
			continue
		}

		subject := "Confirmação de Pagamento"
		body := "Seu pagamento foi processado com sucesso. Detalhes do pagamento: ID " + payment.PaymentID + ", valor: " + payment.Amount + "."
		if err := worker.SendPaymentEmail(payment.Email, subject, body); err != nil {
			log.Printf("Error sending email: %v", err)
		}

		session.MarkMessage(message, "")
	}
	return nil
}

func ConsumePayments(consumerGroup sarama.ConsumerGroup, topic string) {
	handler := PaymentConsumerHandler{}
	for {
		if err := consumerGroup.Consume(context.TODO(), []string{topic}, handler); err != nil {
			log.Fatalf("Error consuming messages: %v", err)
		}
	}
}
