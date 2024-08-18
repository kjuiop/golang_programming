package main

import (
	"chat-kafka/config"
	"chat-kafka/network"
	"chat-kafka/repository"
	"chat-kafka/service"
	"flag"
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

	roomService := service.NewService(mysqlRepo)

	n := network.NewServer(cfg, roomService, mysqlRepo)
	if err := n.StartServer(); err != nil {
		log.Fatalln("fail to start server")
	}

}
