package main

import "fmt"

func main() {
	fmt.Println("result : ", calculateDivide(3, 2))
}

func calculateDivide(num1, num2 int) int {
	if num2 <= 0 {
		return -1
	}

	result := float64(num1) / float64(num2) * 1000
	return int(result)
}
