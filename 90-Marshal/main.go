package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	Name string
	Age  int
}

func main() {

	users := []user{
		{"selvin", 22}, {"narayan", 44}, {"new", 25},
	}
	jm, err := json.Marshal(users)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(jm))
}
