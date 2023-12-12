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
	ctx, cancel := context.WithCancel(context.Background())

	wg.Add(1)
	go signalChan(&wg)

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

func signalChan(wg *sync.WaitGroup) {
	defer wg.Done()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("signalChan receive sigterm signal")
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
			time.Sleep(8 * time.Second)
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
