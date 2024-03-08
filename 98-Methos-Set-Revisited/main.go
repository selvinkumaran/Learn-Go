package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}

type Shape interface {
	Area() float64
}

func (c *circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func Info(s Shape) {
	fmt.Println(s.Area())
}

func main() {

	c := circle{
		4,
	}
	// Info(c) //if use pass value it wont work ,You have pass address(pointer) of the value
	fmt.Println(c.Area())
}
