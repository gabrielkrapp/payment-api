package main

import (
	"log"
	"os"
	"os/signal"
	"payment-api/config"
	"payment-api/pkg/kafka"
)

func main() {

	consumerGroup := config.NewKafkaConsumer("payment-worker-group")
	defer consumerGroup.Close()

	go kafka.ConsumePayments(consumerGroup, "payment-intents")

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	<-sig
	log.Println("Worker shutting down.")
}
