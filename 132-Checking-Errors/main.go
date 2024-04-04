package main

import (
	"fmt"
)

func main() {
	var name, favFood, favColor string

	fmt.Println("Enter your Name: ")
	_, err := fmt.Scanln(&name)
	if err != nil {
		panic(err)
	}

	fmt.Println("Enter your Favorite Food: ")
	_, err = fmt.Scanln(&favFood)
	if err != nil {
		panic(err)
	}

	fmt.Println("Enter your Favorite Color: ")
	_, err = fmt.Scanln(&favColor)
	if err != nil {
		panic(err)
	}

	fmt.Println("name:", name, "favfavFood:", favFood, "favColor:", favColor)
}
