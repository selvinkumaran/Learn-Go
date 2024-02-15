package main

import "fmt"

func main() {

	s := struct {
		first     string
		friends   map[string]int
		favDrinks []string
	}{
		first: "selvin",
		friends: map[string]int{
			"nark":    2,
			"dknnd":   32,
			"qokqowd": 27,
		},
		favDrinks: []string{"pepsi", "bovonto"},
	}
	fmt.Println(s)

	for _, v := range s.favDrinks {
		fmt.Println(s.first, "fav drinks", v)
	}
	for _, v := range s.friends {
		fmt.Println(s.first, "have friends", v)
	}
}
