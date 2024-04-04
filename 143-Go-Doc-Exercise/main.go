package main

import (
	"Module/143-Go-Doc-Exercise/dog"
	"fmt"
)

type lab struct {
	name string
	age  int
}

func main() {
	l1 := lab{
		"vilo",
		dog.Years(10),
	}
	fmt.Println(l1)
} 
