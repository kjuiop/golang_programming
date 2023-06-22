package main

import "fmt"

func modifyMap(m map[string]int) {
	m["key"] = 42
}

func main() {
	myMap := make(map[string]int)
	myMap["key"] = 10

	modifyMap(myMap)

	fmt.Println(myMap["key"]) // 출력: 42
}
