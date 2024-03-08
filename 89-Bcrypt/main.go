package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	pass := "password123"

	bpass, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(pass)
	fmt.Println(string(bpass))

	check := "password1232"

	//To compare password with HashPassword

	err = bcrypt.CompareHashAndPassword(bpass, []byte(check))
	if err != nil {
		fmt.Println("You are not logged In")
		return
	}
	fmt.Println("you are logged In")
}
