package main

import "fmt"

func main() {
	fmt.Println("result : ", Concat("a", "bcd"))
}

func Concat(a, b string) string {
	return a + b
}
