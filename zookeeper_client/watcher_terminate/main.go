package main

import (
	"golang_programming/zookeeper_client/watcher_terminate/config"
	"golang_programming/zookeeper_client/watcher_terminate/zookeeper"
	"log"
)

func main() {

	cfg, err := config.ReadConfig()
	if err != nil {
		log.Fatalf("fail to read environments: %v", err)
		return
	}

	zkCon, err := zookeeper.NewZkClient(cfg)
	if err != nil {
		log.Printf("[NewHandler] failed zookeeper Connection error : %v\n", err)
		return
	}
	defer zkCon.EndZookeeper()

}
