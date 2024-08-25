package kafka

import (
	"chat_controller/config"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type Kafka struct {
	cfg *config.Config

	producer *kafka.Producer
}

func NewKafka(cfg *config.Config) (*Kafka, error) {
	k := &Kafka{cfg: cfg}
	var err error

	// Kafka 프로듀서 초기화
	k.producer, err = kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": cfg.Kafka.URL,
		"client.id":         cfg.Kafka.ClientID,
		"acks":              "all",
	})

	if err != nil {
		return nil, err
	}

	return k, nil
}
