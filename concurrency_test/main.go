package main

import (
	"fmt"
	"sync"
)

var mutex sync.Mutex

func main() {

	wg := sync.WaitGroup{}

	fmt.Println("map writes start!")
	m := map[int]int{}

	wg.Add(1)
	go writeMap(m, 10000, &wg)

	wg.Add(1)
	go readMap(m, 1000, &wg)

	wg.Wait()
	fmt.Println("map writes end!")
}

func writeMap(m map[int]int, num int, wg *sync.WaitGroup) {

	mutex.Lock()

	defer func() {
		mutex.Unlock()
		wg.Done()
	}()

	for i := 0; i < num; i++ {
		//mutex.Lock()
		m[i] = i
		fmt.Println(m[i])
		//mutex.Unlock()
	}

}

func readMap(m map[int]int, num int, wg *sync.WaitGroup) {

	defer func() {
		wg.Done()
		fmt.Println("readMap stop")
	}()

	fmt.Println("readMap start")

	readMapWg := sync.WaitGroup{}

	readMapWg.Add(1)
	go printValue(m, num, &readMapWg)

	readMapWg.Wait()
}

func printValue(m map[int]int, num int, readMapWg *sync.WaitGroup) {

	mutex.Lock()

	defer func() {
		mutex.Unlock()
		readMapWg.Done()
	}()

	for i := 0; i < num; i++ {
		//mutex.Lock()
		fmt.Println("read Map Value : ", m[i])
		//mutex.Unlock()
	}

}
