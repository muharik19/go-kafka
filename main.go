package main

import (
	"fmt"
	"os"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type TransactionPlacer struct {
	producer  *kafka.Producer
	topic     string
	deliverch chan kafka.Event
}

func NewTransactionPlacer(p *kafka.Producer, topic string) *TransactionPlacer {
	return &TransactionPlacer{
		producer:  p,
		topic:     topic,
		deliverch: make(chan kafka.Event, 10000),
	}
}

func (r *TransactionPlacer) placeTransaction(orderType string, size int) error {
	var (
		format  = fmt.Sprintf("%s - %d", orderType, size)
		payload = []byte(format)
	)

	err := r.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{
			Topic:     &r.topic,
			Partition: kafka.PartitionAny,
		},
		Value: payload,
	}, r.deliverch)
	if err != nil {
		fmt.Printf("Err: %v\n", err)
	}

	<-r.deliverch
	fmt.Printf("placed transaction on the queue %s\n", format)
	return nil
}

func main() {
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"client.id":         "blackboard",
		"acks":              "all",
	})
	if err != nil {
		fmt.Printf("Failed to create producer: %v\n", err)
		os.Exit(1)
	}

	const topic = "TRANSACTION"
	t := NewTransactionPlacer(p, topic)
	for i := 0; i < 100; i++ {
		if err := t.placeTransaction("transaction", i+1); err != nil {
			fmt.Printf("Err: %v\n", err)
		}
		time.Sleep(time.Second * 1)
	}
}
