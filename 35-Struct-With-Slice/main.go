package main

import "fmt"

type person struct {
	first_name string
	last_name  string
	ice_cream  []string
}

func main() {

	p1 := person{
		first_name: "sambath",
		last_name:  "kumar",
		ice_cream:  []string{"chocalate", "vanila"},
	}
	p2 := person{		first_name: "mani",
		last_name:  "kumar",
		ice_cream:  []string{"straberry", "black current"},
	}

	fmt.Println(p1)
	fmt.Println(p2)

	fmt.Println(p1.ice_cream)
	fmt.Println(p2.ice_cream)

	for _, v := range p1.ice_cream {
		fmt.Println(p1.first_name, "favorite Ice Cream is", v)
	}

	for _, v := range p2.ice_cream {
		fmt.Println(p2.first_name, "favorite Ice Cream is", v)
	}
}
