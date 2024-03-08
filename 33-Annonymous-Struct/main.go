package main

import "fmt"

func main() {

	type person struct {
		firstName string
		lastName  string
		age       int
	}

	//Annoymous struct
	//If you dont wanna use many place go to annonymous struct
	p1 := struct {
		firstName string
		lastName  string
		age       int
	}{
		firstName: "sambath",
		lastName:  "kumar",
		age:       23,
	}
	p2 := person{
		firstName: "mani",
		lastName:  "kumar",
		age:       43,
	}

	fmt.Println(p1)
	fmt.Println(p2)

	fmt.Printf("%T \t %v \n", p1, p1)
	fmt.Printf("%T \t %v \n", p2, p2)

	fmt.Println(p1.firstName, p1.lastName, p1.age)

}
