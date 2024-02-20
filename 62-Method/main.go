package main

import "fmt"

type Person struct {
	first string
	age   int
}

func main() {

	p := Person{
		first: "selvin",
		age:   22,
	}
	p.speak()

}

func (p Person) speak() {
	fmt.Println(p.first, "age is", p.age)
}
