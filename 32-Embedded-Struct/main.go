package main

import "fmt"

func main() {

	type person struct {
		firstName string
		lastName  string
		age       int
	}

	type student struct {
		person
		firstName string
		graduated bool
	}

	s1 := student{
		person: person{
			firstName: "sambath",
			lastName:  "kumar",
			age:       23,
		},
		firstName: "New Name",
		graduated: false,
	}

	s2 := student{
		person: person{
			firstName: "mani",
			lastName:  "kumar",
			age:       43,
		},
		firstName: "old Name",
		graduated: true,
	}

	fmt.Println(s1)
	fmt.Println(s2)

	fmt.Printf("%T \t %#v \n", s1, s1)

	fmt.Println(s1.firstName, s1.lastName, s1.age, s1.person.firstName)

}
