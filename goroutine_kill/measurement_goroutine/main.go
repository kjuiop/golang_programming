package main

import (
	"fmt"
	"os"
	"runtime/trace"
	"time"
)

func main() {
	// trace 파일 생성
	traceFile, err := os.Create("trace.out")
	if err != nil {
		panic(err)
	}
	defer traceFile.Close()

	// trace 시작
	trace.Start(traceFile)
	defer trace.Stop()

	// 작업 시작
	for i := 0; i < 5; i++ {
		worker(i)
	}
}

func worker(id int) {
	fmt.Printf("Worker %d started\n", id)
	defer fmt.Printf("Worker %d done\n", id)

	// 간단한 작업 수행
	for j := 0; j < 50; j++ {
		fmt.Printf("Worker %d working on task %d\n", id, j)
		time.Sleep(time.Second * 1)
	}
}
