package main

import (
	"github.com/streadway/amqp"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Fatalf("rabbitmq connection fail, error : %s\n", err.Error())
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Printf("channel error : %s", err.Error())
	}
	defer ch.Close()
	ch.Qos(1, 0, false)

	numConsumers := 3

	for i := 1; i <= numConsumers; i++ {
		go consumeMessages(conn, ch, i)
	}

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-sigterm:
		log.Println("Received SIGINT/SIGTERM, exiting gracefully...")
	}

}

func consumeMessages(conn *amqp.Connection, ch *amqp.Channel, consumerID int) {
	//ch, err := conn.Channel()
	//if err != nil {
	//	log.Printf("channel error : %s", err.Error())
	//}
	//defer ch.Close()
	//ch.Qos(1, 0, false)

	notify := conn.NotifyClose(make(chan *amqp.Error)) //error channel

	q, err := ch.QueueDeclare(
		"go-test-request",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Printf("queue declare error : %s", err.Error())
	}

	msgCh, err := ch.Consume(
		q.Name,
		"",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Printf("consume error : %s", err.Error())
	}

	for {
		select {
		case err = <-notify:
			log.Printf("channel error : %s", err.Error())
			time.Sleep(1 * time.Second)
			break
		case d := <-msgCh:
			log.Println("message : ", string(d.Body))
			time.Sleep(3 * time.Second)
			d.Ack(false)
		}
	}
}
