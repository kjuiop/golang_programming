package main

import "fmt"

var m map[string]map[string]string

// 최초 m 을 선언후 생성하였을 때 우리는 하나의 map (map[string]map[string]string) 에 대한 공간을 받음
// 이후 최초의 m 을 다시 make 로 재할당 받지 않는 이상 원래 m 은 보존됨
// child map 은 재할당할 때마다 pointer 가 바뀌지만, 기존 m 이 유지됨으로 재할당되어도 m 에서 접근 가능함
func main() {

	fmt.Printf("parent map 할당 전 : m address : %p, m pointer : %p\n", &m, m)

	m = make(map[string]map[string]string)

	fmt.Printf("parent map 할당 후 : m address : %p, m pointer : %p\n", &m, m)

	roomId := "0000000001"
	userIds := []string{"spring", "jake", "golang"}

	m[roomId] = make(map[string]string)

	for _, userId := range userIds {
		m[roomId][userId] = "data"
	}

	fmt.Printf("child map 할당 후 : m address : %p, m pointer : %p, children pointer : %p children value : %v\n", &m, m, m[roomId], m[roomId])

	m[roomId] = make(map[string]string)

	for _, userId := range userIds {
		m[roomId][userId] = "ss"
	}

	fmt.Printf("child map 할당 후 : m address : %p, m pointer : %p, children pointer : %p children value : %v\n", &m, m, m[roomId], m[roomId])
}
