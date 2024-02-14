package main

import "fmt"

func main() {

	m := map[string]int{

		"selvin": 22,
		"nandha": 23,
		"suman":  15,
	}

	// v, ok := m["selvin"]
	// if ok {
	// 	fmt.Println("The value is", v)
	// } else {
	// 	fmt.Println("The value does not exixts")

	// }

	if v, ok := m["selvin"]; ok {
		fmt.Println("The value is", v)
	} else {
		fmt.Println("The value does not exixts")

	}
	
	if v, ok := m["selvin"]; !ok {
		fmt.Println("The value does not exixts")
	} else {
		fmt.Println("The value is", v)

	}
}
