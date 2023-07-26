package main

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
)

func main() {
	val := []byte(`{"ID":1, "Name": "Reds", "Colors":["Crimson", "Red", "Ruby"]}`)
	fmt.Println(jsoniter.Get(val, "ID").ToInt())
	fmt.Println(jsoniter.Get(val, "Colors", 0).ToString())
}
