package main

import "fmt"

func main() {

	type person struct {
		first_name string
		last_name  string
		age        int
	}

	//Annoymous struct
	//If you dont wanna use many place go to  annonymous struct
	p1 := struct {
		first_name string
		last_name  string
		age        int
	}{
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
	fmt.Printf("%T \t %v \n", p2, p2)

	fmt.Println(p1.first_name, p1.last_name, p1.age)

}
