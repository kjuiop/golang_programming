package main

import (
	"chat-kafka/config"
	"chat-kafka/network"
	"flag"
	"log"
)

var pathFlag = flag.String("config", "./config.toml", "config set")
var port = flag.String("port", ":1010", "port set")

func main() {
	flag.Parse()

	_ = config.NewConfig(*pathFlag)

	n := network.NewServer()
	if err := n.StartServer(); err != nil {
		log.Fatalln("fail to start server")
	}

}
