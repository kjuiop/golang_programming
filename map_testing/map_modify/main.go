package main

import "fmt"

func modifyMap(m2 map[string]int) {
	m2["key"] = 42
	fmt.Printf("[02] m1 address : %p,m1 pointer : %p m1 value : %v\n", &m2, m2, m2["key"])
}

func main() {
	m1 := make(map[string]int)
	m1["key"] = 10

	fmt.Printf("[01] m1 address : %p,m1 pointer : %p m1 value : %v\n", &m1, m1, m1["key"])
	modifyMap(m1)

	fmt.Printf("[03] m1 address : %p,m1 pointer : %p m1 value : %v\n", &m1, m1, m1["key"])
}
