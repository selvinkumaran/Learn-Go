package main

import (
	"fmt"
	"math"
)

type Square struct {
	length float64
	width  float64
}

type Circle struct {
	radius float64
}

type Shape interface {
	Area() float64
}

func main() {

	s := Square{
		5, 8,
	}
	c := Circle{4}
	fmt.Println(info(s))
	fmt.Println(info(c))
}

func info(s Shape) float64 {
	return s.Area()
}
func (s Square) Area() float64 {
	return s.length * s.width
}

func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}
