package main

import (
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"sync"
	"syscall"
)

func main() {

	wg := sync.WaitGroup{}

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()

		cmd := exec.Command("sh", "time.sh")
		if err := cmd.Start(); err != nil {
			fmt.Println("start error:", err)
			return
		}
		fmt.Printf("waiting process %d ....\n", cmd.Process.Pid)

		sig := <-sigs
		fmt.Println("--- --- --- receive signal:", sig)
		if err := cmd.Process.Kill(); err != nil {
			fmt.Println("kill sub process error:", err)
		}

		fmt.Println("MAIN END")

	}(&wg)

	wg.Wait()
}
