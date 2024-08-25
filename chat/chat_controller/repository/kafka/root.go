package kafka

import (
	"chat_controller/config"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type Kafka struct {
	cfg *config.Config

	consumer *kafka.Consumer
}

func NewKafka(cfg *config.Config) (*Kafka, error) {
	k := &Kafka{cfg: cfg}
	var err error

	// Kafka 프로듀서 초기화
	k.consumer, err = kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": cfg.Kafka.URL,
		"group.id":          cfg.Kafka.GroupID,
		"auto.offset.reset": "latest",
	})

	if err != nil {
		return nil, err
	}

	return k, nil
}

func (k *Kafka) RegisterSubTopic(topic string) error {
	if err := k.consumer.Subscribe(topic, nil); err != nil {
		return err
	} else {
		return nil
	}
}

func (k *Kafka) Pool(timeoutMs int) kafka.Event {
	return k.consumer.Poll(timeoutMs)
}
