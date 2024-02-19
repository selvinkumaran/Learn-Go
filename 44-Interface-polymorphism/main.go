package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}
type secrectAgent struct {
	person
	ltk bool
}

func (p person) speak() {
	fmt.Println(p.firstName, p.lastName)
}

func (sa secrectAgent) speak() {
	fmt.Println("I m a secret agent", sa.person.firstName, sa.person.lastName, sa.ltk)
}

func main() {
	sa := secrectAgent{
		person: person{
			firstName: "mani",
			lastName:  "kumar",
		},
		ltk: true,
	}
	p2 := person{
		firstName: "sanjay",
		lastName:  "kumar",
	}

	sa.speak()
	p2.speak()

}
