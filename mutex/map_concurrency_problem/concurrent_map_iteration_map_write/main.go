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

	testMap := make(map[int]*Person)
	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 1; i <= 50000; i++ {
			testMap[i] = NewPerson(i)
			fmt.Printf("test 1번 고루틴 %p: %d\n", testMap[1], i)
		}
	}()

	go printPersonAge(testMap)

	wg.Wait()
}

func printPersonAge(testMap map[int]*Person) {

	for {
		for _, person := range testMap {
			fmt.Printf("test 2번 고루틴 Person Age : %d\n", person.Age)
		}

		//time.Sleep(1 * time.Millisecond)
	}
}

func NewPerson(i int) *Person {
	return &Person{
		Name:  "jake_" + strconv.Itoa(i),
		Age:   i * 10,
		Phone: "010-3182-0825",
	}
}
