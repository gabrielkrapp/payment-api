package kafka

import (
	"context"
	"fmt"
	"log"

	"github.com/IBM/sarama"
)

type PaymentConsumerHandler struct{}

func (PaymentConsumerHandler) Setup(sarama.ConsumerGroupSession) error {
	return nil
}

func (PaymentConsumerHandler) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

func (h PaymentConsumerHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for message := range claim.Messages() {
		fmt.Printf("Message claimed: value = %s\n", string(message.Value))
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
