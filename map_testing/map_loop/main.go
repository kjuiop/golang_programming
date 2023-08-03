package main

import "fmt"

type Person struct {
	Name     string
	Age      int
	Children []Child
}

type Child struct {
	Name   string
	Age    int
	Parent string
}

func main() {

	myMap := make(map[string]Person)
	assembleTestData(myMap)

	idArrays := [...]string{"jake", "kan", "jake", "john"}
	for idx, id := range idArrays {
		obj := myMap[id]
		if idx == 0 {
			obj.Age = 2
			obj.Children[1].Age = 99
			continue
		}

		obj.Age = 10
		obj.Children[0].Age = 10
		myMap[id] = obj
	}

	for _, obj := range myMap {
		fmt.Printf("result : obj parsing name: %s, age: %d, children: %v\n", obj.Name, obj.Age, obj.Children)
	}
}

func assembleTestData(data map[string]Person) {
	data["jake"] = Person{
		Name: "jake",
		Age:  1,
		Children: []Child{
			{
				Name:   "jake_1",
				Age:    2,
				Parent: "jake",
			},
			{
				Name:   "jake_2",
				Age:    2,
				Parent: "jake",
			},
			{
				Name:   "jake_3",
				Age:    2,
				Parent: "jake",
			},
		},
	}
	data["kan"] = Person{
		Name: "kan",
		Age:  2,
		Children: []Child{
			{
				Name:   "jake_1",
				Age:    2,
				Parent: "jake",
			},
		},
	}
	data["john"] = Person{
		Name: "john",
		Age:  3,
		Children: []Child{
			{
				Name:   "jake_1",
				Age:    2,
				Parent: "jake",
			},
		},
	}
}
