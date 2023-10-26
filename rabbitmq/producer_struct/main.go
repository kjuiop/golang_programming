package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("rabbitmq producer application start")
	http.HandleFunc("/producer", ProduceRabbitmqMsg)

	if err := http.ListenAndServe(":3002", nil); err != nil {
		log.Printf("http server listen err : %s\n", err.Error())
	}
}
