package main

import (
	"context"
	"log"
	"sync"
	"time"
)

func main() {

	wg := sync.WaitGroup{}
	parentContext := context.Background()

	ctx, cancel := context.WithTimeout(parentContext, 5*time.Second)

	wg.Add(1)
	go printRoutine(ctx, &wg)

	<-ctx.Done()
	log.Println("sub ctx receive sigterm signal")
	cancel()
	log.Println("receive sigint signal")
	wg.Wait()
	log.Println("server shutdown complete")
}

func printRoutine(ctx context.Context, wg *sync.WaitGroup) {
	i := 0
	isComplete := false

	defer func() {
		log.Printf("is complete print Routine : %v\n", isComplete)
		wg.Done()
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
