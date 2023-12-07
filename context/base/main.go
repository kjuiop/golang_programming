package main

import (
	"log"
	"time"
)

func main() {
	printRoutine()
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
			log.Printf("print Routine : %d\n", i)
			isComplete = false
			time.Sleep(3 * time.Second)
			isComplete = true
		}
	}
}
