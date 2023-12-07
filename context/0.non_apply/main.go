package main

import (
	"log"
	"time"
)

func main() {
	go printRoutine()
}

func printRoutine() {
	i := 0
	isComplete := false

	defer func() {
		log.Printf("is complete print Routine : %v\n", isComplete)
	}()

	for {
		select {
		default:
			i++
			log.Printf("print Routine : %d\n", i)
			isComplete = false
			time.Sleep(3 * time.Second)
			isComplete = true
		}
	}
}
