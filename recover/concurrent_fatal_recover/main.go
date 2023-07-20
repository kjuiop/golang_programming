package main

import (
	"fmt"
	"log"
	"strconv"
	"sync"
)

type Person struct {
	Name  string
	Age   int
	Phone string
}

var mutex = &sync.Mutex{}

// panic error 는 런타임 패닉을 발생시킬 때 Recover 로 잡을 수는 있음
// fatal 은 에로 로그를 출력하고 그대로 프로그램을 종료한다. Recover 로 되살릴 수 없음
func main() {

	defer func() {
		if r := recover(); r != nil {
			log.Printf("   >>> main recoverd: %v\n", r)
		}
	}()

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
			log.Printf("test 1번 고루틴 %p: %d\n", testMap[1], i)
		}
	}()

	go printPersonAge(testMap)

	wg.Wait()

	fmt.Println("end recover")
}

func printPersonAge(testMap map[int]*Person) {

	defer Recover()

	for {
		for _, person := range testMap {
			log.Printf("test 2번 고루틴 Person Age : %d\n", person.Age)
		}
	}
}

func Recover() {
	if r := recover(); r != nil {
		log.Printf("   >>> recoverd: %v\n", r)
	}
}

func (p *Person) printPersonName() {
	log.Printf("test 2번 고루틴 Person Name : %s\n", p.Name)
}

func NewPerson(i int) *Person {
	return &Person{
		Name:  "jake_" + strconv.Itoa(i),
		Age:   i * 10,
		Phone: "010-3182-0825",
	}
}
