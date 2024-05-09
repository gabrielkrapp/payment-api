package config

import (
	"log"

	"github.com/IBM/sarama"
)

func NewKafkaProducer(brokers []string) sarama.SyncProducer {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5
	config.Producer.Return.Successes = true

	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		log.Fatalf("Failed to create Kafka producer: %v", err)
	}
	return producer
}

func NewKafkaConsumer(brokers []string, groupID string, config *sarama.Config) (sarama.ConsumerGroup, error) {
	consumer, err := sarama.NewConsumerGroup(brokers, groupID, config)
	if err != nil {
		log.Fatalf("Error creating Kafka consumer group: %v", err)
	}
	return consumer, nil
}
