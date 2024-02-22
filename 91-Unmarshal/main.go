package main

import (
	"encoding/json"
	"fmt"
)

// website go to json

type user struct {
	Name string `json:"Name"`
	Age  int    `json:"Age"`
}

func main() {

	s := `[{"Name":"selvin","Age":22},{"Name":"narayan","Age":44},{"Name":"new","Age":25}]`

	fmt.Println("JSON Object : ", s)
	var users []user
	err := json.Unmarshal([]byte(s), &users)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("GO Object : ", users)

	for i, v := range users {
		fmt.Println("index of user", i)
		fmt.Println("The name is", v.Name, "and age is", v.Age)
	}
}
