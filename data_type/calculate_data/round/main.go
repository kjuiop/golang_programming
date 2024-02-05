package main

import (
	"fmt"
	"math"
)

func main() {
	inputNumber := 10.56789

	fmt.Printf("Original Number %.5f\n", roundTwoScale(inputNumber))
	fmt.Printf("Original Number %.5f\n", roundThreeScale(inputNumber))
}

func roundTwoScale(val float64) float64 {
	return math.Round(val*100) / 100
}

func roundThreeScale(val float64) float64 {
	return math.Round(val*1000) / 1000
}
