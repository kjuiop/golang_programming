package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"runtime"
	"sync"
	"syscall"
	"time"
)

func main() {

	wg := sync.WaitGroup{}
	ctx, cancel := context.WithCancel(context.Background())
	sigs := make(chan os.Signal, 1)
	defer close(sigs)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGKILL)

	wg.Add(1)
	go CloseWithContext(sigs, &wg, ctx, cancel)

	wg.Wait()
}

func CloseWithContext(sigs chan os.Signal, wg *sync.WaitGroup, ctx context.Context, cancel context.CancelFunc) {

	defer wg.Done()

	i := 0
	for {
		select {
		case <-sigs:
			log.Printf("[Module 1] Receive exit signal sigs: %v\n", sigs)
			cancel()
		case <-ctx.Done():
			log.Println("[Module 1] CloseWithContext Close Goroutine")
			return
		default:
			numGoroutines := runtime.NumGoroutine()
			log.Printf("[Module 1] i numbering %d / Number of goroutines: %d\n", i, numGoroutines)
			i++
			time.Sleep(time.Second * 1)
		}
	}
}
