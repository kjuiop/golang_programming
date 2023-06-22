package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	p := Person{Name: "John", Age: 30}

	m := make(map[string]*Person)
	m["00000001"] = &p

	fmt.Printf("before pass struct person : %p\n", &p)
	fmt.Printf("before pass map : %p\n", &m)

	passParameter(p, &p, m, &m)
}

func passParameter(p Person, p2 *Person, m map[string]*Person, m2 *map[string]*Person) {
	fmt.Printf("after pass struct person : %p\n", &p)
	fmt.Printf("after pass struct pointer person : %p\n", p2)
	fmt.Printf("after pass map : %p\n", m)
	fmt.Printf("after pass map pointer : %p\n", m2)
}
