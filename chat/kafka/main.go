package main

import (
	"chat-kafka/config"
	"chat-kafka/network"
	"chat-kafka/repository"
	"flag"
	"fmt"
	"log"
)

var pathFlag = flag.String("config", "./config.toml", "config set")
var port = flag.String("port", ":1010", "port set")

func main() {
	flag.Parse()

	cfg := config.NewConfig(*pathFlag)
	mysqlRepo, err := repository.NewRepository(cfg)
	if err != nil {
		log.Fatalf("failed connect mysql db, err : %v\n", err)
	}

	fmt.Println(mysqlRepo)

	n := network.NewServer()
	if err := n.StartServer(); err != nil {
		log.Fatalln("fail to start server")
	}

}
