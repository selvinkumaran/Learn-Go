package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type user struct {
	Name string
	Age  int
}

func main() {

	users := []user{
		{"selvin", 22}, {"narayan", 44}, {"new", 25},
	}
	fmt.Println(users)
	err := json.NewEncoder(os.Stdout).Encode(users)
	if err != nil {
		fmt.Println(err)
	}

}
