package main

import "fmt"

func main() {
	fmt.Println("2 + 3 = ", mySum(2, 3))
	fmt.Println("4 + 7 = ", mySum(4, 7))
}

func mySum(xi ...int) int {
	sum := 0
	for _, val := range xi {
		sum += val
	}
	return sum
}
