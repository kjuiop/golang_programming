package main

import (
	"fmt"
	"log"
)

// 유클리드 호제법으로 최대공약수 계산
func gcd(a, b int) int {
	if b == 0 {
		return a
	}

	return gcd(b, a%b)
}

// 최소공배수 계산
func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

func main() {
	var num1, num2 int

	// 사용자로부터 두 수 입력 받기
	fmt.Print("첫 번째 수를 입력하세요: ")
	if _, err := fmt.Scan(&num1); err != nil {
		log.Fatalf("정수를 입력해주세요.")
	}

	fmt.Print("두 번째 수를 입력하세요: ")
	if _, err := fmt.Scan(&num2); err != nil {
		log.Fatalf("정수를 입력해주세요.")
	}

	// 최대공약수 계산 및 출력
	gcdResult := gcd(num1, num2)
	fmt.Printf("두 수의 최대공약수(GCD): %d\n", gcdResult)

	// 최소공배수 계산 및 출력
	lcmResult := lcm(num1, num2)
	fmt.Printf("두 수의 최소공배수(LCM): %d\n", lcmResult)
}
