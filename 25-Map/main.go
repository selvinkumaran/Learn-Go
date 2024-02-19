package main

import "fmt"

func main() {

	m := map[string]int{

		"selvin": 22,
		"nandha": 23,
		"suman":  15,
	}

	for k, v := range m {
		fmt.Println(k, v)
	}

	fmt.Println(m["selvin"])

	mm := make(map[string]int)

	mm["shyam"] = 55
	mm["daniel"] = 66
	mm["sam"] = 33

	fmt.Println(mm)
	fmt.Printf("%#v \n", mm)

	fmt.Println(len(mm))

	//This will print the key
	for v := range m {
		fmt.Println(v)
	}

	//This will print the value
	for _, v := range m {
		fmt.Println(v)
	}

	//This will print the key
	for v := range m {
		fmt.Println(v)
	}

	//For range with slice

	xi := []int{2, 3, 5, 6, 77, 8}

	fmt.Println("***************************")

	for i, v := range xi {
		fmt.Println(i, v)
	}

	fmt.Println("***************************")

	for _, v := range xi {
		fmt.Println(v)
	}

	fmt.Println("***************************")

	for v := range xi {
		fmt.Println(v)
	}

}
