package kafka

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

var producer *kafka.Producer

func InitKafka() error {
	var err error
	producer, err = kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"acks":              "all",
		"retries":           5,
	})
	return err
}

func SendMessage(topic string, value string) error {
	deliveryChan := make(chan kafka.Event)
	err := producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte(value),
	}, deliveryChan)

	e := <-deliveryChan
	m := e.(*kafka.Message)

	close(deliveryChan)

	if m.TopicPartition.Error != nil {
		return m.TopicPartition.Error
	}
	return err
}

func CloseKafka() {
	producer.Close()
}
