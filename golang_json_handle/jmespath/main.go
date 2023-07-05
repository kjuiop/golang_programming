package main

import (
	"encoding/json"
	"fmt"
	"github.com/jmespath/go-jmespath"
)

const JsonString = `
{
	"store" : {
		"book": [
			{
				"category": "reference",
				"author": "Nigel Rees",
				"title": "Saying of the Century",
				"price": 8.95
			},
			{
				"category": "fiction",
				"author": "Evelyn Waugh",
				"title": "Sword of Honour",
				"price": 12.99
			},
			{
				"category": "fiction",
				"author": "Herman Malvilie",
				"title": "Moby Dick",
				"isbn": "0-553-21311-3",
				"price": 8.99
			}
		],
		"bicycle": {
			"color": "red",
			"price": 19.95
		}
	},
	"expensive": 12.43
}

`

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

	var jsonData interface{}
	if err := json.Unmarshal([]byte(JsonString), &jsonData); err != nil {
		panic(err)
	}
	var lookupStrings = []string{
		"expensive",
		"store.book[0].price",
		"store.book[-1].isbn",
		"store.book[1:2].price",
		"store.book[?isbn].price",
		"store.book[?price <`10`].price",
	}

	for _, str := range lookupStrings {
		res, err := lookup(jsonData, str)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%s => %s\n", str, res)
	}
}

func lookup(data interface{}, expression string) (string, error) {
	res, err := jmespath.Search(expression, data)
	if err != nil {
		panic(err)
	}
	encoded, err := json.Marshal(res)
	if err != nil {
		panic(err)
	}
	return string(encoded), nil
}
