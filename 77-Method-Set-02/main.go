package main

import "fmt"

type dog struct {
	name string
}

type youngin interface {
	walk()
	run()
}

func youngRun(y youngin) {
	y.run()
	y.walk()
}

func main() {
	d1 := dog{
		"hendry",
	}
	d1.walk()
	d1.run()
	// youngRun(d1)

	d2 := &dog{
		"buddy",
	}
	d2.walk()
	d2.run()
	youngRun(d2)
}

func (d dog) walk() {
	fmt.Println("my name is ", d.name, "I am walking")
}
func (d *dog) run() {
	d.name = "new changed"
	fmt.Println("my name is ", d.name, "I am walking")
}

/*

  Receiver           Values
   (t T)	    	T and *T
   (t *T)	     	*T

*/
