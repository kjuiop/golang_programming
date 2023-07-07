package main

import (
	"fmt"
	"strconv"
	"sync"
)

type Person struct {
	Name  string
	Age   int
	Phone string
}

func main() {

	testMap := make(map[int]Person)
	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 1; i <= 4000; i++ {
			testMap[i] = Person{
				Name:  "jake_" + strconv.Itoa(i),
				Age:   i * 10,
				Phone: "010-3182-0825",
			}

			fmt.Printf("test 1번 고루틴 : %d\n", i)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 1; i <= 4000; i++ {
			printGoroutineNumber(testMap, i)
		}
	}()

	wg.Wait()
}

func printGoroutineNumber(testMap map[int]Person, i int) {
	if _, ok := testMap[i]; ok {
		fmt.Printf("exist test 2번 고루틴 : %d\n", i)
	} else {
		fmt.Printf("not exist test 2번 고루틴 : %d\n", i)
	}
}
