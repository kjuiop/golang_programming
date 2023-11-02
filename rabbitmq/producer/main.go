package main

import (
	"github.com/streadway/amqp"
	"log"
)

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("rabbitmq connection fail, error : %s\n", err.Error())
	}
	defer func(conn *amqp.Connection) {
		if err := conn.Close(); err != nil {
			log.Printf("rabbitmq connection close, err: %s\n", err.Error())
		}
	}(conn)

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("fail connection rabbitmq channel, err: %s\n", err.Error())
	}
	defer func(ch *amqp.Channel) {
		if err := ch.Close(); err != nil {
			log.Printf("rabbitmq channel connection close, err: %s\n", err.Error())
		}
	}(ch)

	queueName := "go-test-request"
	q, err := ch.QueueDeclare(queueName, false, false, false, false, nil)
	if err != nil {
		log.Printf("Failed to declare a queue, err : %s", err.Error())
	}

	body := "Hello rabbitMq Consumer"
	if err := ch.Publish("", q.Name, false, false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		}); err != nil {
		log.Printf("Failed to publish a queue, err : %s", err.Error())
	}

	log.Printf("Success Sent : %s\n", body)
}
