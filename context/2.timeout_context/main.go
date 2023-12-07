package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	go printRoutine(context.Background())

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	<-ctx.Done()
	log.Println("receive sigint signal")
}

func printRoutine(ctx context.Context) {
	i := 0
	isComplete := false

	defer func() {
		log.Printf("is complete print Routine : %v\n", isComplete)
	}()

	for {
		select {
		case <-ctx.Done():
			log.Println("ctx done printRoutine close")
			return
		default:
			i++
			log.Printf("print Routine : %d\n", i)
			isComplete = false
			time.Sleep(1 * time.Second)
			isComplete = true
		}
	}
}
