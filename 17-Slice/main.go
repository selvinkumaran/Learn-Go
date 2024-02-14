package main

import "fmt"

func main() {

	xs := []string{"selvin", "newwww", "selvqkdnwidiwin", "we2oejihqwiof"}

	fmt.Println(xs)
	fmt.Println(len(xs))
	fmt.Printf("the type xs is %T", xs)

	for i, v := range xs {
		fmt.Printf("\nThe index is %v and the value is %v ", i, v)
	}

	fmt.Println("***********************")

	for _, v := range xs {
		fmt.Printf("\nThe value is %v \n", v)
	}

	fmt.Println("***********************")

	fmt.Println(xs[0])
	fmt.Println(xs[1])
	fmt.Println(xs[2])
	fmt.Println(xs[3])

	fmt.Println("***********************")

	for i := 0; i < len(xs); i++ {
		fmt.Println(xs[i])
	}

}
