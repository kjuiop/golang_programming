package main

import (
	"fmt"
	cmap "github.com/orcaman/concurrent-map"
)

func main() {
	// generic type
	// m := cmap.New[any]()
	m := cmap.New()
	m.Set("a1", 1234)
	m.Set("a2", "morning")
	m.Set("a3", []int{2, 3, 4, 5})
	if v, ok := m.Get("a2"); ok {
		fmt.Println(v)
	}
	// 이부분을 고루틴으로 사용하여 쓰기 지연 기능을 구현
	m.IterCb(func(key string, val any) {
		fmt.Printf("%s: %#v\n", key, val)
	})
}
