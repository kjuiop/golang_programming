package main

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

type RabbitMq struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	queue   amqp.Queue
}

func NewRabbitMqClient() (*RabbitMq, error) {

	r := &RabbitMq{}

	host := "amqp://guest:guest@localhost:5672/"
	if err := r.connect(host); err != nil {
		return nil, err
	}

	return r, nil
}

func (r *RabbitMq) connect(host string) error {
	conn, err := amqp.Dial(host)
	if err != nil {
		return fmt.Errorf("fail connected host : %s, err : %s\n", host, err.Error())
	}
	r.conn = conn
	return nil
}

func (r *RabbitMq) getChannel() error {
	ch, err := r.conn.Channel()
	if err != nil {
		return err
	}

	r.channel = ch
	return nil
}

func (r *RabbitMq) queueDeclare(name string, durable bool, autoDelete bool, exclusive bool, noWait bool, args amqp.Table) error {

	q, err := r.channel.QueueDeclare(name, durable, autoDelete, exclusive, noWait, args)
	if err != nil {
		return err
	}

	r.queue = q
	return nil
}

func (r *RabbitMq) publishMq(qName, exchange string, mandatory, immediate bool, payload []byte) error {

	err := r.channel.Publish(
		exchange,
		qName,
		mandatory,
		immediate,
		amqp.Publishing{
			Headers:         amqp.Table{"Content-Type": "application/json"},
			ContentType:     "application/json",
			ContentEncoding: "UTF-8",
			Body:            payload,
			DeliveryMode:    amqp.Persistent, // 1=non-persistent, 2=persistent
			Priority:        0,               // 0-9
		},
	)

	return err
}

func (r *RabbitMq) Close() {
	if err := r.channel.Close(); err != nil {
		log.Printf("rabbitmq channel connection close, err: %s\n", err.Error())
	}

	if err := r.conn.Close(); err != nil {
		log.Printf("rabbitmq connection close, err: %s\n", err.Error())
	}
}
