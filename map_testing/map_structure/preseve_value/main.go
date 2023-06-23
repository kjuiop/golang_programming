package main

import (
	"fmt"
)

type Person struct {
	Name  string // 16 bytes
	Age   int    // 8 bytes
	Phone string // 16 bytes
}

func main() {
	p := Person{Name: "John", Age: 30}

	fmt.Println("============================================================================")
	m1 := make(map[string]Person)
	m1["00000001"] = p
	fmt.Printf("[01] m1 address : %p,m1 pointer : %p m1 value : %v\n", m1, &m1, m1)
	one(m1)
	fmt.Printf("[03] m1 address : %p,m1 pointer : %p m1 value : %v\n", m1, &m1, m1)
}

func one(m1 map[string]Person) {
	m1["00000001"] = Person{Name: "Jake", Age: 30}
	fmt.Printf("[02] m1 address : %p,m1 pointer : %p m1 value : %v\n", m1, &m1, m1)
}
