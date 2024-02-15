package main

import "fmt"

type Engine struct {
	//Engine fileds
}

type Car struct {
	Engine //Embedding the engine type
	name   string
	//Car fields
}

func main() {
	car := Car{
		name: "selvin",
	}
	car.Start()
}

func (e Car) Start() {
	fmt.Println("Started")
	fmt.Println(e.name)
}
