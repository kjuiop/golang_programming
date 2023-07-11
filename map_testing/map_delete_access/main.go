package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	myMap := make(map[string]*Person)
	myMap["1"] = &Person{
		Name: "ddd",
		Age:  1,
	}

	fmt.Println(myMap["1"])

	delete(myMap, "1")

	fmt.Println(myMap["1"].Age)
}
