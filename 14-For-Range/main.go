package main

import (
	"fmt"
	"math/rand"
)

func main() {

	x := []int{43, 32, 538, 49, 35, 56, 77, 88}

	for i, x := range x {
		fmt.Printf("Indes %v and value %v \n", i, x)
	}

	fmt.Println("**********************")

	m := map[string]int{
		"selvin": 20,
		"oadwod": 34,
		"sojwiw": 44,
	}

	fmt.Println("**********************")

	for k, v := range m {
		fmt.Printf("the key is %v and the value is %v \n", k, v)
	}

	fmt.Println("**********************")

	age1 := m["selvin"]
	age2 := m["oadwod"]
	age3 := m["sojwiw"]
	age4 := m["nnn"]

	fmt.Println(age1)
	fmt.Println(age2)
	fmt.Println(age3)
	fmt.Println(age4)

	fmt.Println("**********************")

	if v, ok := m["selvin"]; ok {
		fmt.Println("the value", v)
	}

	fmt.Println("**********************")

	for i := 0; i <= 100; i++ {
		if x := rand.Intn(5); x == 3 {
			fmt.Printf("Iteration %v \t x is %v \n", i, x)
		}
	}
}
