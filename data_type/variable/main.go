package main

import "fmt"

func main() {

	var age int
	var name string

	// 초기값과 함께 변수 선언
	var score float64 = 92.5
	var isActive bool = true

	// 짧은 변수 선언 (타입 생략)
	height := 180
	country := "Korea"

	// 여러 변수 선언
	var (
		x int    = 5
		y string = "Hello"
		z bool   = true
	)

	// 사용된 변수들을 출력
	fmt.Println("age:", age)
	fmt.Println("name:", name)
	fmt.Println("score:", score)
	fmt.Println("isActive:", isActive)
	fmt.Println("height:", height)
	fmt.Println("country:", country)
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("z:", z)

	m, d, j := 5, 3, 5
	m, d, j = d, j, m

	fmt.Println("m:", m)
	fmt.Println("d:", d)
	fmt.Println("j:", j)
}
