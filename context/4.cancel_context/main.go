package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {

	wg := sync.WaitGroup{}
	parentContext := context.Background()

	ctx, cancel := context.WithCancel(parentContext)

	wg.Add(1)
	go printRoutine(ctx, &wg)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	log.Println("receive sigint signal")

	cancel()
	wg.Wait()
	log.Println("sub ctx receive sigterm signal")
}

func printRoutine(ctx context.Context, wg *sync.WaitGroup) {
	i := 0
	isComplete := false

	defer func() {
		log.Printf("is complete print Routine : %v\n", isComplete)
	}()

	for {
		select {
		case <-ctx.Done():
			log.Println("ctx done printRoutine close")
			wg.Done()
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
