package main

import "fmt"

func main() {

	type person struct {
		first_name string
		last_name  string
		age        int
	}

	p1 := person{
		first_name: "sambath",
		last_name:  "kumar",
		age:        23,
	}
	p2 := person{	
		first_name: "mani",
		last_name:  "kumar",
		age:        43,
	}

	fmt.Println(p1)
	fmt.Println(p2)

	fmt.Printf("%T \t %v \n", p1, p1)

	fmt.Println(p1.first_name, p1.last_name, p1.age)

}
