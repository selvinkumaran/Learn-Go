package main

import "fmt"

type engine struct {
	electric bool
}
type Vehicle struct {
	engine
	make   string
	model  string
	doors  int
	colors string
}

func main() {

	v1 := Vehicle{
		engine: engine{
			electric: true,
		},
		make:   "bmw",
		model:  "basic",
		doors:  4,
		colors: "red",
	}
	v2 := Vehicle{
		engine: engine{
			electric: false,
		},
		make:   "audi",
		model:  "advanced",
		doors:  4,
		colors: "blue",
	}

	fmt.Println(v1)
	fmt.Println(v2)

	fmt.Println(v1.electric, v1.make, v1.colors)
	fmt.Println(v2.electric, v2.make, v1.colors)

	fmt.Println(v1.engine.electric, v1.make, v1.colors)
	fmt.Println(v2.engine.electric, v2.make, v1.colors)
}
