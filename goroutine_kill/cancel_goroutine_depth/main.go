package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	ctx, cancelFn := context.WithCancel(context.Background())
	var wg sync.WaitGroup
	wg.Add(1)
	go func(ctx context.Context) {
		defer wg.Done()
		i := 0
		for {
			select {
			case <-ctx.Done():
				fmt.Println("first goroutine close")
				return
			default:
				i++
				time.Sleep(time.Second * 20)
				do(ctx, i)
				select {
				case <-ctx.Done():
					fmt.Println("second goroutine close")
					return
				default:
					fmt.Printf("print goroutine : %d\n", i)
					time.Sleep(time.Second * 5)
				}
			}
		}
	}(ctx)

	<-exitSignal()
	cancelFn()
	wg.Wait()
	fmt.Println("main goroutine stop")
}

func exitSignal() <-chan os.Signal {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	return sig
}

func do(ctx context.Context, i int) {
	fmt.Println("exec do")

	select {
	case <-ctx.Done():
		fmt.Println("three goroutine close")
		time.Sleep(time.Second * 1)
		return
	default:
		fmt.Printf("three goroutine : %d\n", i)
		twoDo(ctx, i)
	}
}

func twoDo(ctx context.Context, i int) {

	select {
	case <-ctx.Done():
		fmt.Println("four goroutine close")
		time.Sleep(time.Second * 1)
		return
	default:
		fmt.Printf("four goroutine : %d\n", i)
		threeDo(ctx, i)
	}
}

func threeDo(ctx context.Context, i int) {

	select {
	case <-ctx.Done():
		fmt.Println("five goroutine close")
		time.Sleep(time.Second * 20)
		return
	default:
		fmt.Printf("five goroutine : %d\n", i)
	}
}
