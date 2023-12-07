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

	go printRoutine(context.Background(), time.Second*1)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		log.Println("receive cancel signal")
		//cancel()
		// cancel 은 신호 전파용
	}()

	log.Println("receive sigint signal")
	<-ctx.Done()
	log.Println("ctx done")
}

func printRoutine(ctx context.Context, period time.Duration) {
	i := 0
	for {
		select {
		case <-time.After(period):
			log.Printf("print Routine : %d\n", i)
			i++
		case <-ctx.Done():
			log.Println("printRoutine close")
			return
		}
	}
}
