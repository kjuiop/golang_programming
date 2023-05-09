package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rect struct {
	width  float64
	height float64
}

type Circle struct {
	radius float64
}

func (r Rect) Area() float64 {
	return r.width * r.height
}

func (r Rect) Perimeter() float64 {
	return 2 * (r.width + r.height)
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func main() {

	var s Shape = Rect{10, 3}
	fmt.Printf("type of s is %T\n", s)
	fmt.Printf("value of s is %v\n", s)
	fmt.Printf("value of s is %0.2f\n\n", s.Area())

	var d Shape = Circle{10}
	fmt.Printf("type of s is %T\n", d)
	fmt.Printf("value of s is %v\n", d)
	fmt.Printf("value of s is %0.2f\n\n", d.Area())

}
