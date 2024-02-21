package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string `json:"First"`
	Last  string `json:"Last"`
	Age   int    `json:"Age"`
}

func main() {
	s := `[{"First":"mani","Last":"kumar","Age":23},{"First":"karan","Last":"kumar","Age":24},{"First":"sunnil","Last":"kumar","Age":25}]`

	bs := []byte(s)

	fmt.Printf("%T\n", s)
	fmt.Printf("%T\n", bs)

	// people:=[]person{} //use any one this
	var people []person

	err := json.Unmarshal(bs, &people)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("print all the data", people)

	for i, v := range people {
		fmt.Println("******************************")
		fmt.Println("Index of ", i)
		fmt.Println(v.First, v.Last, v.Age)
	}
}
