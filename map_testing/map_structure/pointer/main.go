package main

import (
	"fmt"
	"unsafe"
)

// 강제 메모리 할당하는 프로그램

type Person struct {
	Name  string // 16 bytes
	Age   int    // 8 bytes
	Phone string // 16 bytes
}

func main() {
	fmt.Println("============================================================================")
	var p Person
	fmt.Printf("p : %v, pointer p : %p, size : %v\n", p, &p, int(unsafe.Sizeof(p)))

	var pp *Person
	fmt.Printf("pp : %p, pointer pp : %p, size : %v\n", pp, &pp, int(unsafe.Sizeof(pp)))
	pp = &Person{Name: "park"}
	fmt.Printf("pp : %p, pointer pp : %p, size : %v\n", pp, &pp, int(unsafe.Sizeof(*pp)))

	fmt.Println("============================================================================")
	var m map[string]Person
	fmt.Printf("m : %p, &m : %p, size : %v\n", m, &m, int(unsafe.Sizeof(m)))
	m = make(map[string]Person)
	m["ryu"] = Person{Name: "ryu", Age: 123, Phone: "010-1234-5678"}
	fmt.Printf("&m : %p, size : %v\n", &m, int(unsafe.Sizeof(&m)))
	fmt.Printf("m[ryu] : %v, size : %v\n", m["ryu"], int(unsafe.Sizeof(m["ryu"])))

	fmt.Println("============================================================================")
	mp := map[string]*Person{}
	mp["ryu"] = &Person{Name: "ryu"}
	// 변수의 위치는 다르지만 포인터는 같다.
	fmt.Printf("mp[ryu] : %p, pointer mp : %p, size : %v\n", mp, &mp, int(unsafe.Sizeof(mp["ryu"])))
	fmt.Printf("mp[ryu] : %p, pointer mp : %p, size : %v\n", mp["ryu"], &mp, int(unsafe.Sizeof(*mp["ryu"])))
}
