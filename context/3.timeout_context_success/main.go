package main

import (
	"context"
	"log"
	"time"
)

func main() {

	parentContext := context.Background()

	ctx, cancel := context.WithTimeout(parentContext, 5*time.Second)

	go printRoutine(ctx)

	<-ctx.Done()
	log.Println("sub ctx receive sigterm signal")
	cancel()
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
