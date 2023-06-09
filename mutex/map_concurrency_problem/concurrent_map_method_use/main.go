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
	testMap[1] = NewPerson(1)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 1; i <= 50000; i++ {
			person := testMap[1]
			fmt.Printf("test 1번 고루틴 %p: %d\n", person, i)
		}
	}()

	go printPersonAge(testMap)

	wg.Wait()
}

func printPersonAge(testMap map[int]*Person) {

	for {
		for _, person := range testMap {
			person.printPersonName()
		}
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
