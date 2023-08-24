package main

import (
	"fmt"
	"github.com/samuel/go-zookeeper/zk"
	"golang_programming/zookeeper_client/watcher_terminate/config"
	"golang_programming/zookeeper_client/watcher_terminate/zookeeper"
	"log"
	"sync"
)

func main() {

	wg := sync.WaitGroup{}
	wg.Add(1)

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

	path := fmt.Sprintf("%s/%s", cfg.ZookeeperInfo.RootNode, "transcoders/kollus")
	if err := watcherChildrenNode(path, zkCon); err != nil {
		log.Printf("watcher children node err : %v\n", err)
	}

	wg.Wait()
}

func watcherChildrenNode(path string, zkCon *zookeeper.ZKCon) error {

	nodes, err := zkCon.WatcherChildrenNodeToMap(path, func(event *zk.Event) {
		switch event.Type {
		case zk.EventNodeChildrenChanged:
			if err := watcherChildrenNode(path, zkCon); err != nil {
				return
			}
		}
	})

	if err != nil {
		return err
	}

	log.Printf("wathcer nodes : %v\n", nodes)
	return nil
}
