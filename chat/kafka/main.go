package main

import (
	"chat-kafka/network"
	"log"
)

func main() {
	n := network.NewServer()
	if err := n.StartServer(); err != nil {
		log.Fatalln("fail to start server")
	}

}
