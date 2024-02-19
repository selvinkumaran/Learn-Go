package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

//method
func (p person) /*receiver*/ speak() {
	fmt.Println(p.firstName, p.lastName)
}
func main() {
	p1 := person{
		firstName: "mani",
		lastName:  "kumar",
	}
	p2 := person{
		firstName: "sanjay",
		lastName:  "kumar",
	}

	p1.speak()
	p2.speak()

}
