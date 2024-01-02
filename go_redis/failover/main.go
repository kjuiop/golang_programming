package main

import (
	"context"
	"github.com/go-redis/redis"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	wg := sync.WaitGroup{}

	options := &redis.Options{
		Addr: "localhost:6379",
	}
	client := redis.NewClient(options)

	wg.Add(1)
	go redisPing(ctx, &wg, client)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("receive sigint signal")

	cancel()
	wg.Wait()
	log.Println("sub ctx receive sigterm signal")
}

func redisPing(ctx context.Context, wg *sync.WaitGroup, client *redis.Client) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			log.Println("ctx done printRoutine close")
			return
		default:
			time.Sleep(1 * time.Second)
			pingResult, err := client.Ping().Result()
			if err != nil {
				log.Printf("Error: %v", err)
				continue
			}
			log.Printf("Ping Result: %v", pingResult)
		}
	}
}
