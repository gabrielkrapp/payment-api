package kafka

import (
	"log"
	"time"

	"github.com/IBM/sarama"
)

func KafkaProducer(producer sarama.SyncProducer, topic string, message []byte) {
	kafkaMsg := &sarama.ProducerMessage{
		Topic:     topic,
		Value:     sarama.ByteEncoder(message),
		Timestamp: time.Now(),
	}

	_, _, err := producer.SendMessage(kafkaMsg)
	if err != nil {
		log.Printf("Failed to publish message to Kafka: %v", err)
	}
}
