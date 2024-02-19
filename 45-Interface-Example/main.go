package main

import (
	"fmt"
)

// Shape is an interface that defines a method Area()
type Shape interface {
	Area() float64
}

// Circle struct
type Circle struct {
	Radius float64
}

// Area method for Circle
func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

// Rectangle struct
type Rectangle struct {
	Width  float64
	Height float64
}

// Area method for Rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func main() {

	c := Circle{Radius: 5}
	r := Rectangle{Width: 3, Height: 4}

	// Define a slice of shapes
	shapes := []Shape{c, r}

	// Calculate and print the area of each shape
	for _, shape := range shapes {
		fmt.Println("Area:", shape.Area())
	}
}
