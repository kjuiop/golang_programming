package main

import (
	"fmt"
	"github.com/shamaton/msgpack/v2"
	"log"
)

type Person struct {
	Name string
}

func main() {

	v := Person{Name: "msgpack"}

	d, err := msgpack.Marshal(v)
	if err != nil {
		log.Println("error : ", err.Error())
	}

	r := Person{}
	if err := msgpack.Unmarshal(d, &r); err != nil {
		log.Println("error : ", err.Error())
	}

	fmt.Println("Original:", v)
	fmt.Println("Unmarshaled:", r)
}
