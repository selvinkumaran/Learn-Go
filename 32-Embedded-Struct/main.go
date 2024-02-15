package main

import "fmt"

func main() {

	type person struct {
		first_name string
		last_name  string
		age        int
	}

	type student struct {
		person
		first_name string
		graduated  bool
	}

	s1 := student{
		person: person{
			first_name: "sambath",
			last_name:  "kumar",
			age:        23,
		},
		first_name: "New Name",
		graduated:  false,
	}

	s2 := student{
		person: person{
			first_name: "mani",
			last_name:  "kumar",
			age:        43,
		},
		first_name: "old Name",
		graduated:  true,
	}

	fmt.Println(s1)
	fmt.Println(s2)

	fmt.Printf("%T \t %#v \n", s1, s1)

	fmt.Println(s1.first_name, s1.last_name, s1.age, s1.person.first_name)

}
