package main

import (
	"fmt"
)

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
	p2 := person{
		first_name: "mani",
		last_name:  "kumar",
		ice_cream:  []string{"straberry", "black current"},
	}

	m := map[string]person{
		p1.first_name: p1,
		p2.first_name: p2,
	}

	for _, v := range m {
		fmt.Println(v)
		for _, v2 := range v.ice_cream {
			fmt.Println(v.first_name, "Likes ice cream", v2)
		}
	}
}
