package main

import "fmt"

func main() {

	m := map[string]int{

		"selvin": 22,
		"nandha": 23,
		"suman":  15,
	}

	fmt.Println(m)
	fmt.Println(len(m))

	fmt.Println("***********************")

	delete(m, "nandha")

	//WONT GIVE ERROR
	fmt.Println("**********START*************")
	delete(m, "nandDKDha")
	fmt.Println(m["selvon"]) //If ther no key it returs 0 as a value
	fmt.Println("***********END************")

	fmt.Println(m)
	fmt.Println(len(m))

}
