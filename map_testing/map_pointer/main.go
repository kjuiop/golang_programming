package main

import (
	"fmt"
	"unsafe"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	p := Person{Name: "John", Age: 30}

	// 변수를 선언만 해도, 위치 (&m) , 사이즈는 8만큼 할당을 받은 상태 64bit 일 때 8, 32bit 일 때 4
	// 변수의 주소값 과 포인터는 다르다. 변수의 주소값은 map 을 가리키는 변수 자체의 주소, 포인터 값은 map 을 가리키는 주소
	var m map[string]Person
	fmt.Printf("m address : %p, m pointer : %p size : %v\n", &m, m, unsafe.Sizeof(m))

	m = make(map[string]Person)
	fmt.Printf("m address : %p, m pointer : %p size : %v\n", &m, m, unsafe.Sizeof(m))

	m["00000001"] = p

	fmt.Printf("before pass struct person : %p\n", &p)
	passParameterStruct(p, &p)

	fmt.Printf("before pass map : %p\n", m)
	passParameterStructMap(m, &m)

}

func passParameterStruct(
	p Person, p2 *Person,
) {
	// 파라미터로 넘겼을 때, 포인터로 넘기지 않고,
	// 넘기기전의 변수와 넘긴 후의 변수가 같은지를 알아보고 싶다.
	fmt.Printf("after pass struct person : %p\n", &p)
	fmt.Printf("after pass struct pointer person : %p\n", p2)
}

func passParameterStructMap(m1 map[string]Person, m2 *map[string]Person) {
	fmt.Printf("after pass map : %p\n", m1)
	fmt.Printf("after pass map pointer : %p\n", m2)
}
