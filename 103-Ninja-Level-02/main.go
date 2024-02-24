package main

import "fmt"

type person struct {
	first string
	age   int
}

func (p *person) speak() {
	fmt.Println(p.first, "has age", p.age)
}

type human interface {
	speak()
}

func saySomething(h human) {
	h.speak()
}

func main() {

	p := person{
		"selvin", 22,
	}
	//It wont work
	// saySomething(p)

	saySomething(&p)

	p.speak()

}
