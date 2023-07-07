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

var mutex = &sync.Mutex{}

func main() {

	testMap := make(map[int]*Person)
	wg := sync.WaitGroup{}
	testMap[1] = NewPerson(1)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 1; i <= 20000; i++ {
			mutex.Lock()
			testMap[i] = NewPerson(i)
			mutex.Unlock()
			fmt.Printf("test 1번 고루틴 %p: %d\n", testMap[1], i)
		}
	}()

	go printPersonAge(testMap)

	wg.Wait()
}

func printPersonAge(testMap map[int]*Person) {

	for {
		mutex.Lock()
		for _, person := range testMap {
			fmt.Printf("test 2번 고루틴 Person Age : %d\n", person.Age)
		}
		mutex.Unlock()

		mutex.Lock()
		printGoroutineNumber(testMap, 1)
		mutex.Unlock()

		mutex.Lock()
		person := testMap[1]
		mutex.Unlock()

		person.Age = 1
		person.printPersonName()

	}
}

func printGoroutineNumber(testMap map[int]*Person, i int) {
	if _, ok := testMap[i]; ok {
		fmt.Printf("exist test 2번 고루틴 : %d\n", i)
	} else {
		fmt.Printf("not exist test 2번 고루틴 : %d\n", i)
	}
}

func (p *Person) printPersonName() {
	fmt.Printf("test 2번 고루틴 Person Name : %s\n", p.Name)
}

func NewPerson(i int) *Person {
	return &Person{
		Name:  "jake_" + strconv.Itoa(i),
		Age:   i * 10,
		Phone: "010-3182-0825",
	}
}
