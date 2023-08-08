package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	var term bool = false
	sigs := make(chan os.Signal, 1)
	defer close(sigs)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGKILL)

	// interrupt 시그널을 수신하면 프로그램을 종료한다.
	go func() {
		sig := <-sigs
		term = true
		fmt.Println("=======================")
		fmt.Println("Receive signal: ", sig)
		fmt.Println("=======================")
	}()
	go func() {
		traceFile, err := os.Create("trace.log")
		if err != nil {
			panic(err)
		}
		defer traceFile.Close()

		var i int
		fmt.Println("i : ", &i)
		for i = 0; i < 10; i++ {
			if i == 9 {
				i = 0
			}

			_, err = traceFile.WriteString(fmt.Sprintf("i : %d\n", i))
			if err != nil {
				fmt.Println("Error writing to file:", err)
				return
			}

			time.Sleep(time.Second)
		}
	}()
	for {
		if term {
			fmt.Println("term")
			break
		} else {
			fmt.Println("wait")
			time.Sleep(time.Second)
		}
	}
}
