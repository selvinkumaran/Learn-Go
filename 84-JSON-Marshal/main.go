package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string
	Last  string
	Age   int
}

func main() {

	p1 := person{
		"mani", "kumar", 23,
	}
	p2 := person{
		"karan", "kumar", 24,
	}
	p3 := person{
		"sunnil", "kumar", 25,
	}

	people := []person{p1, p2, p3}
	fmt.Println(people)

	jm, err := json.Marshal(people)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(jm))
	fmt.Printf("%T", jm)
	fmt.Printf("%T", people)

}
