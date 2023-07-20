package main

import "fmt"

var person = &Person{}

func main() {

	if person != nil {
		fmt.Println("is exist ", person)
	}
}

type Person struct {
	Name  string // 16 bytes
	Age   int    // 8 bytes
	Phone string // 16 bytes
}
