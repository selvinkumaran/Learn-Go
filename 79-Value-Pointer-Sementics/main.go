package main

import "fmt"

type person struct {
	first string
}

func main() {
	p := person{
		"jenny",
	}
	fmt.Println(p)

	q := changeName(p, "jen")
	fmt.Println(q)

	changeNamePtr(&p, "hendry")
	fmt.Println(p)

}

func changeName(p person, s string) person {
	p.first = s
	return p
}

func changeNamePtr(p *person, s string) {
	p.first = s
}
