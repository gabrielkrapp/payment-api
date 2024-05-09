package config

import (
	"log"

	"github.com/IBM/sarama"
)

var brokers = []string{"localhost:9092"}

func NewKafkaProducer() sarama.SyncProducer {
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

func NewKafkaConsumer(groupID string) sarama.ConsumerGroup {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	consumer, err := sarama.NewConsumerGroup(brokers, groupID, config)
	if err != nil {
		log.Fatalf("Error creating Kafka consumer group: %v", err)
	}
	return consumer
}
