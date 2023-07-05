package main

import (
	"encoding/json"
	"fmt"
	"github.com/jmespath/go-jmespath"
)

func main() {
	var jsondata = []byte(`{"foo": {"bar": {"baz" : [0,1,2,3,4]}}}`)
	var data interface{}
	err := json.Unmarshal(jsondata, &data)
	if err != nil {
		panic(err)
	}
	result, err := jmespath.Search("foo.bar.baz[2]", data)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}
