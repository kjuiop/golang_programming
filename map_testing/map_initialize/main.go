package main

import "fmt"

func main() {
	var test map[string]int
	test = make(map[string]int)
	fmt.Printf("1-1 : %p\n", test)
	one(test)
	fmt.Printf("1-2 : %p\n", test)
	fmt.Printf("test : %v\n", test)
	two(test)
	fmt.Printf("1-3 : %p\n", test)
	fmt.Printf("test : %v\n", test)
}
func one(test map[string]int) {
	fmt.Printf("2-1 : %p\n", test)
	fmt.Printf("2-2 : %p\n", &test)
	test = make(map[string]int)
	test["a"] = 1
	fmt.Printf("2-3 : %p\n", test)
	fmt.Printf("2-4 : %p\n", &test)
}
func two(test map[string]int) {
	fmt.Printf("3-1 : %p\n", test)
	test["b"] = 2
	fmt.Printf("3-2 : %p\n", test)
}

//==========================
//결과
//1-1 : 0xc000088150 map 의 주소
//2-1 : 0xc000088150 map 의 주소 (test 자체가 map 을 가리키는 포인터 값이므로 함수로 넘겨도 값이 보전)
//2-2 : 0xc00000e030 변수 test 의 주소
//2-3 : 0xc000088180 새로 할당된 map 의 주소
//2-4 : 0xc00000e030 변수 test 의 주소
//1-2 : 0xc000088150 기존 map 의 주소
//test : map[] 기존 map 의 값
//3-1 : 0xc000088150 기존 map 의 주소
//3-2 : 0xc000088150 기존 map 의 주소
//1-3 : 0xc000088150 기존 map 의 주소
//test : map[b:2] 기존 map 의 값
