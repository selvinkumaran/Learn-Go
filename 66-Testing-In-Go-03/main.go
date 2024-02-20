package main

import "fmt"

func main() {
	fmt.Println(Paradise("Australia"))
}

func Paradise(place string) string {
	return fmt.Sprint("My place is ", place)
}
