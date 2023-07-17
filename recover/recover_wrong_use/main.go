package main

import "fmt"

// defer 는 함수가 올바르게 종료되도, panic 으로 인해 중단되도 마지막에 호출되는 구문이다.
// 따라서 recover 는 defer 함수 내에서 선언되어야 하며, defer 구문으로 함수가 직접 호출이 되어야 한다.
// 만일 아래와 같은 예시로 recover 함수가 호출이 된다면 (defer -> Recover() -> recover()) recover 는 적용되지 않는다.
// 또한, recover 이 된다고 해서 함수 내에서 err 가 발생하는 지점부터 살아나는 것이 아니라, 그 함수를 종료시키고 다음 함수를 실행시킨다.

func main() {

	occurPanic()
	recoverPrint()
}

func occurPanic() {

	defer func() {
		Recover()
	}()

	/**
	recover 를 직접 호출하는 경우 정상 동작
	defer Recover()
	*/

	var m map[string]*Person
	// panic occur
	fmt.Printf("person name : %s", m["jake"].Name)
}

func recoverPrint() {
	fmt.Println("recover statement worked")
}

func Recover() {
	if r := recover(); r != nil {
		fmt.Printf("   >>> recoverd: %v\n", r)
	}
}

type Person struct {
	Name  string // 16 bytes
	Age   int    // 8 bytes
	Phone string // 16 bytes
}
