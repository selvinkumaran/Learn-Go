package main

import "fmt"

type dog struct {
	name string
}

func main() {
	d1 := dog{
		"hendry",
	}
	d1.walk()
	d1.run()

	d2 := &dog{
		"buddy",
	}
	d2.walk()
	d2.run()
}

func (d dog) walk() {
	fmt.Println("my name is ", d.name, "I am walking")
}
func (d *dog) run() {
	d.name = "new changed"
	fmt.Println("my name is ", d.name, "I am walking")
}
